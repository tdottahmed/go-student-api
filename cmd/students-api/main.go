package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/tdottahmed/students-api/internal"
)

func main() {
	// Load configuration
	cfg := config.MustLoadConfig()

	// Router setup
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students-api"))
	})

	// Server setup
	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	// Log server start
	slog.Info("Server starting", slog.String("address", cfg.HTTPServer.Addr))
	// Start server & listen for graceful shutdown
	// TODO: Need to study all the graceful shutdown options
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			// slog.Fatal(err)
			slog.Error("Server error: ", slog.String("error", err.Error()))
		}
	}()

	<-done
	slog.Info("Server shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown error: ", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown complete")
}
