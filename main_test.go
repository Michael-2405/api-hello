package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	tests := [] struct {
		name				string
		method 			string
		path 				string
		wantStatus 	int
		wantMessage string
	} {
		{"root", http.MethodGet, "/", http.StatusOK, "Hello from Go API"},
		{"health", http.MethodGet, "/health", http.StatusOK, "OK"},
		{"unknown path", http.MethodGet, "/nope", http.StatusNotFound, ""},
		{"wrong method", http.MethodPost, "/health", http.StatusMethodNotAllowed, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func (t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			newMux().ServeHTTP(rec, req)

			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			if tc.wantMessage == "" {
				return
			}
			var got Response
			if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if got.Message != tc.wantMessage {
				t.Fatalf("message = %q, want %q", got.Message, tc.wantMessage)
			}
		})
	}
}
