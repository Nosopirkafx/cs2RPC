package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
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

	stateCh := make(chan rpc.State, 1)
	go rpc.Run(clientID, stateCh)

	store := &stateStore{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", dashboardHandler)
	mux.HandleFunc("/api/state", makeStateHandler(stateCh, store))
	mux.HandleFunc("/api/status", store.statusHandler)

	addr := "127.0.0.1:" + port
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	shutdown := func() {
		log.Println("shutting down")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}
	mux.HandleFunc("/api/stop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || !validHost(r) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		go shutdown()
	})

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		shutdown()
	}()

	log.Printf("listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

type stateStore struct {
	mu    sync.RWMutex
	state rpc.State
	seen  time.Time
}

func (s *stateStore) set(state rpc.State) {
	s.mu.Lock()
	s.state = state
	s.seen = time.Now()
	s.mu.Unlock()
}

func (s *stateStore) statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || !validHost(r) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	s.mu.RLock()
	response := struct {
		Running bool      `json:"running"`
		Seen    time.Time `json:"seen,omitempty"`
		State   rpc.State `json:"state"`
	}{Running: !s.seen.IsZero(), Seen: s.seen, State: s.state}
	s.mu.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(response)
}

func makeStateHandler(stateCh chan rpc.State, store *stateStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method == http.MethodGet {
			store.statusHandler(w, r)
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
		store.set(s)

		select {
		case stateCh <- s:
		default:
			select {
			case <-stateCh:
			default:
			}
			select {
			case stateCh <- s:
			default:
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func validHost(r *http.Request) bool {
	h := strings.ToLower(r.Host)
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
