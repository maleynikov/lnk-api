package main

import (
	"fmt"
	"net/http"

	"gitlab.maleynikov.me/url-short/api/pkg/app"
	"gitlab.maleynikov.me/url-short/api/pkg/app/storage"
)

type Server struct {
	storage *storage.Storage
}

func (s *Server) Run() error {
	// Load config
	cfg, err := app.LoadConfig()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	s.storage, err = storage.NewStorage(cfg)
	if err != nil {
		return fmt.Errorf("storage: %w", err)
	}
	// Start server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: s.routes(),
	}
	fmt.Printf("Starting server on :%d\n", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
