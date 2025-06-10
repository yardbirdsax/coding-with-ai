package server

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestServer_RegisterHandler(t *testing.T) {
	logger := log.New(os.Stdout, "test: ", log.LstdFlags)
	srv := New(":8080", logger)
	
	// Create a test handler
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test handler"))
	})
	
	// Register the handler
	srv.RegisterHandler("/test", testHandler)
	
	// Create a test request
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	
	// Serve the request using the mux directly
	mux := srv.server.Handler.(*http.ServeMux)
	mux.ServeHTTP(w, req)
	
	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
	
	expectedBody := "test handler"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, w.Body.String())
	}
}

type testLogger struct {
	entries []string
}

func (l *testLogger) Write(p []byte) (n int, err error) {
	l.entries = append(l.entries, string(p))
	return len(p), nil
}

func TestServer_New(t *testing.T) {
	// Simple test to verify that New creates a Server instance
	logger := log.New(os.Stdout, "test: ", log.LstdFlags)
	srv := New(":8080", logger)
	
	if srv == nil {
		t.Error("Expected New to return a non-nil Server")
	}
	
	if srv.server == nil {
		t.Error("Expected Server to have non-nil http.Server")
	}
	
	if srv.server.Addr != ":8080" {
		t.Errorf("Expected server address to be :8080, got %q", srv.server.Addr)
	}
}