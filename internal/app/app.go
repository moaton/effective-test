package app

import (
	"context"
	"effective-test/config"
	"effective-test/internal/db"
	"effective-test/internal/db/postgres"
	"effective-test/internal/handlers"
	"effective-test/internal/service"
	"effective-test/pkg/logger"
	"fmt"
	"net/http"
)

func Run(ctx context.Context, cfg *config.Config) {
	addr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDB)
	repo, err := postgres.NewPostgres(addr)
	if err != nil {
		logger.Fatalf("postgres.NewPostgres err %v", err)
	}
	db.NewRepository(repo)

	service := service.NewService(repo)

	handler := handlers.NewHandler(service)
	if err := http.ListenAndServe(":3000", handler); err != nil {
		logger.Errorf("http.ListenAndServe err %v", err)
	}
}
