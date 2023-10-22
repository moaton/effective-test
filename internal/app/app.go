package app

import (
	"context"
	"effective-test/config"
	"effective-test/internal/db"
	"effective-test/internal/db/postgres"
	"effective-test/pkg/logger"
	"fmt"
)

func Run(ctx context.Context, cfg *config.Config) {
	addr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDB)
	repo, err := postgres.NewPostgres(addr)
	if err != nil {
		logger.Fatalf("postgres.NewPostgres err %v", err)
	}
	db.NewRepository(repo)
}
