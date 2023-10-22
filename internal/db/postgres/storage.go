package postgres

import (
	"context"
	"database/sql"
	"effective-test/internal/models"
	"effective-test/pkg/client/postgresql"
	"effective-test/pkg/logger"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(url string) (*Postgres, error) {
	db, err := postgresql.NewClient(url)
	if err != nil {
		logger.Errorf("postgresql.NewClient err %v", err)
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) GetUsers(ctx context.Context, params models.Params) ([]models.User, error) {
	return []models.User{}, nil
}
func (p *Postgres) InsertUser(ctx context.Context, user models.User) (int64, error) {
	return 0, nil
}
func (p *Postgres) UpdateUser(ctx context.Context, user models.User) error {
	return nil
}
func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	return nil
}
