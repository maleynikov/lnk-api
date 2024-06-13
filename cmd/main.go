package main

import (
	"fmt"
	"net/http"
	"os"

	"gitlab.maleynikov.me/url-short/api/pkg/app"
)

func main() {
	if err := runServer(); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}

func runServer() error {
	// Load config
	cfg, err := app.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Start server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: routes(),
	}

	fmt.Printf("Starting server on :%d\n", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
