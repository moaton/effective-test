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

func (p *Postgres) GetUsers(ctx context.Context, params models.Params) ([]models.User, int64, error) {
	var total int64

	err := p.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE LOWER(name) LIKE $1 OR LOWER(surname) LIKE $1 OR LOWER(patronymic) LIKE $1", "%"+params.Query+"%").Scan(&total)
	if err != nil {
		return []models.User{}, 0, err
	}

	rows, err := p.db.QueryContext(ctx, "SELECT id, name, surname, patronymic, age, gender, nationality FROM users WHERE LOWER(name) LIKE $1 OR LOWER(surname) LIKE $1 OR LOWER(patronymic) LIKE $1 LIMIT $2 OFFSET $3", "%"+params.Query+"%", params.Limit, params.Offset)
	if err != nil {
		return []models.User{}, 0, err
	}
	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.Age, &user.Gender, &user.Nationality); err != nil {
			logger.Errorf("GetUsers rows.Scan err %v", err)
			continue
		}
		users = append(users, user)
	}

	return users, total, nil
}
func (p *Postgres) InsertUser(ctx context.Context, user *models.User) (int64, error) {
	var id int64
	err := p.db.QueryRowContext(ctx, "INSERT INTO users (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, user.Nationality).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := p.db.QueryContext(ctx, "UPDATE users SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6 WHERE id=$7", user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, user.Nationality, user.ID)
	return err
}

func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	_, err := p.db.QueryContext(ctx, "DELETE FROM users WHERE id=$1", id)
	return err
}
