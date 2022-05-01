package repository

import (
	"context"

	"github.com/acaldo/rest-ws/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	Close() error
}

var implentation Repository

func SetRepository(repository Repository) {
	implentation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implentation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implentation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implentation.GetUserByEmail(ctx, email)
}

func Close() error {
	return implentation.Close()
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implentation.InsertPost(ctx, post)
}
