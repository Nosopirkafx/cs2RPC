package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"faceit-discord-rpc/backend/rpc"
)

const clientID = "1540354848015388685"

func main() {
	port := os.Getenv("CS2RPC_PORT")
	if port == "" {
		port = "42157"
	}

	setupLogging()

	if alreadyRunning() {
		log.Println("another instance is already running, exiting")
		return
	}

	stateCh := make(chan rpc.State, 8)
	go rpc.Run(clientID, stateCh)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/state", makeStateHandler(stateCh))

	addr := "127.0.0.1:" + port
	srv := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		log.Println("shutting down")
		_ = srv.Close()
		os.Exit(0)
	}()

	log.Printf("listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func makeStateHandler(stateCh chan<- rpc.State) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if !validHost(r) {
			http.Error(w, "bad host", http.StatusForbidden)
			return
		}
		if !validOrigin(r) {
			http.Error(w, "bad origin", http.StatusForbidden)
			return
		}

		var s rpc.State
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		if err := s.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		select {
		case stateCh <- s:
		default:
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func validHost(r *http.Request) bool {
	h := r.Host
	return strings.HasPrefix(h, "127.0.0.1:") || strings.HasPrefix(h, "localhost:")
}

func validOrigin(r *http.Request) bool {
	o := r.Header.Get("Origin")
	if o == "" {
		return true
	}
	return strings.HasPrefix(o, "chrome-extension://") || strings.HasPrefix(o, "moz-extension://")
}

func setupLogging() {
	exe, err := os.Executable()
	if err != nil {
		return
	}
	dir := filepath.Dir(exe)
	f, err := os.OpenFile(filepath.Join(dir, "cs2rpc.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(f)
}
