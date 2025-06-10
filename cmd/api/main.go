package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joshuafeierman/repos/coding-with-ai/pkg/handlers"
	"github.com/joshuafeierman/repos/coding-with-ai/pkg/server"
)

func main() {
	// Set up logger
	logger := log.New(os.Stdout, "api: ", log.LstdFlags)
	
	// Create server instance
	srv := server.New(":8080", logger)
	
	// Register handlers
	healthHandler := handlers.NewHealthHandler()
	srv.RegisterHandler("/health", healthHandler)
	
	// Start server in a goroutine
	go func() {
		logger.Println("Starting server on :8080")
		if err := srv.Start(); err != http.ErrServerClosed {
			logger.Fatalf("Server error: %v", err)
		}
	}()
	
	// Wait for initialization to complete (simulated)
	logger.Println("Initializing service...")
	time.Sleep(2 * time.Second) // Simulate initialization
	
	// Mark service as ready
	logger.Println("Service initialized, marking as ready")
	healthHandler.SetReady(true)
	
	// Set up graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	
	// Wait for shutdown signal
	<-stop
	
	logger.Println("Shutting down server...")
	
	// Shutdown with a timeout
	shutdownTimeout := 5 * time.Second
	if err := srv.Stop(shutdownTimeout); err != nil {
		logger.Fatalf("Server shutdown failed: %v", err)
	}
	
	logger.Println("Server stopped")
}