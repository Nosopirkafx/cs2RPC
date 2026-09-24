package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"faceit-discord-rpc/backend/rpc"
)

func TestStateHandlerStoresNewestState(t *testing.T) {
	states := make(chan rpc.State, 1)
	store := &stateStore{}
	handler := makeStateHandler(states, store)

	req := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:42157/api/state", strings.NewReader(`{"status":"queue"}`))
	req.Host = "127.0.0.1:42157"
	req.Header.Set("Origin", "chrome-extension://example")
	res := httptest.NewRecorder()
	handler(res, req)

	if res.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusNoContent)
	}
	if got := (<-states).Status; got != "queue" {
		t.Fatalf("queued status = %q, want queue", got)
	}

	statusReq := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:42157/api/status", nil)
	statusReq.Host = "127.0.0.1:42157"
	statusReq.Header.Set("Origin", "chrome-extension://example")
	statusRes := httptest.NewRecorder()
	store.statusHandler(statusRes, statusReq)
	var response struct {
		Connected bool      `json:"connected"`
		State     rpc.State `json:"state"`
	}
	if err := json.NewDecoder(statusRes.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if !response.Connected || response.State.Status != "queue" {
		t.Fatalf("unexpected status response: %+v", response)
	}
	if got := statusRes.Header().Get("Access-Control-Allow-Origin"); got != "chrome-extension://example" {
		t.Fatalf("allowed origin = %q, want extension origin", got)
	}
}

func TestStatusHandlerExpiresStaleMatch(t *testing.T) {
	store := &stateStore{
		state: rpc.State{Status: "match"},
		seen:  time.Now().Add(-31 * time.Second),
	}
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:42157/api/status", nil)
	req.Host = "127.0.0.1:42157"
	res := httptest.NewRecorder()
	store.statusHandler(res, req)
	var response struct {
		Connected bool      `json:"connected"`
		State     rpc.State `json:"state"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if response.Connected || response.State.Status != "idle" {
		t.Fatalf("stale match should be inactive: %+v", response)
	}
}

func TestStateHandlerRejectsWebOrigins(t *testing.T) {
	handler := makeStateHandler(make(chan rpc.State, 1), &stateStore{})
	req := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:42157/api/state", strings.NewReader(`{"status":"idle"}`))
	req.Host = "127.0.0.1:42157"
	req.Header.Set("Origin", "https://example.com")
	res := httptest.NewRecorder()
	handler(res, req)
	if res.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusForbidden)
	}
}
