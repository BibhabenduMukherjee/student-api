package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BibhabenduMukherjee/student-api/internal/config"
	"github.com/BibhabenduMukherjee/student-api/internal/http/handlers/students"
)

func main() {
	// load configuration
	cfg := config.MustLoad()
	// initialize storage

	// setup routers
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", students.New())

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}
	fmt.Printf("server listening on %s", cfg.HTTPServer.Address)

	// channels that takes signals come from os
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	<-done

	slog.Info("shutting down the server ")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("server forced to shutdown: %w", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown completed")
}
