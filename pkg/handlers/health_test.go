package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_NotReady(t *testing.T) {
	// Create a new health handler (defaults to not ready)
	handler := NewHealthHandler()
	
	// Create a test request
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	
	// Serve the request
	handler.ServeHTTP(w, req)
	
	// Check the response
	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("Expected status code %d, got %d", http.StatusServiceUnavailable, w.Code)
	}
	
	expectedBody := "Service is initializing"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, w.Body.String())
	}
}

func TestHealthHandler_Ready(t *testing.T) {
	// Create a new health handler and mark it as ready
	handler := NewHealthHandler()
	handler.SetReady(true)
	
	// Create a test request
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	
	// Serve the request
	handler.ServeHTTP(w, req)
	
	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
	
	expectedBody := "Service is healthy"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, w.Body.String())
	}
}

func TestHealthHandler_SetReadyToggle(t *testing.T) {
	// Create a new health handler
	handler := NewHealthHandler()
	
	// Initial state should be not ready
	if handler.IsReady() {
		t.Error("Expected handler to not be ready initially")
	}
	
	// Set to ready
	handler.SetReady(true)
	if !handler.IsReady() {
		t.Error("Expected handler to be ready after SetReady(true)")
	}
	
	// Set back to not ready
	handler.SetReady(false)
	if handler.IsReady() {
		t.Error("Expected handler to not be ready after SetReady(false)")
	}
}