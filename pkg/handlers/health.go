package handlers

import (
	"net/http"
	"sync/atomic"
)

// HealthHandler manages the health endpoint
type HealthHandler struct {
	isReady uint32
}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{
		isReady: 0, // Initially not ready
	}
}

// SetReady marks the service as ready (or not ready)
func (h *HealthHandler) SetReady(ready bool) {
	if ready {
		atomic.StoreUint32(&h.isReady, 1)
	} else {
		atomic.StoreUint32(&h.isReady, 0)
	}
}

// IsReady checks if the service is ready
func (h *HealthHandler) IsReady() bool {
	return atomic.LoadUint32(&h.isReady) == 1
}

// ServeHTTP implements the http.Handler interface
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !h.IsReady() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service is initializing"))
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy"))
}