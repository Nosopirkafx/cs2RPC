package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
	statusRes := httptest.NewRecorder()
	store.statusHandler(statusRes, statusReq)
	var response struct {
		Running bool      `json:"running"`
		State   rpc.State `json:"state"`
	}
	if err := json.NewDecoder(statusRes.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	if !response.Running || response.State.Status != "queue" {
		t.Fatalf("unexpected status response: %+v", response)
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
