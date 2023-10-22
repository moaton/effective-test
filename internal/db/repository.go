package db

import (
	"context"
	"effective-test/internal/models"
)

type Repository interface {
	GetUsers(ctx context.Context, params models.Params) ([]models.User, error)
	InsertUser(ctx context.Context, user models.User) (int64, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id int64) error
}

var impl Repository

func NewRepository(repo Repository) {
	impl = repo
}

func GetUsers(ctx context.Context, params models.Params) ([]models.User, error) {
	return impl.GetUsers(ctx, params)
}

func InsertUser(ctx context.Context, user models.User) (int64, error) {
	return impl.InsertUser(ctx, user)
}

func UpdateUser(ctx context.Context, user models.User) error {
	return impl.UpdateUser(ctx, user)
}

func DeleteUser(ctx context.Context, id int64) error {
	return impl.DeleteUser(ctx, id)
}
