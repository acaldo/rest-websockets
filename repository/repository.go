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
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPosts(ctx context.Context, page uint64) ([]*models.Post, error)
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

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implentation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post) error {
	return implentation.UpdatePost(ctx, post)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implentation.DeletePost(ctx, id, userId)
}

func ListPosts(ctx context.Context, page uint64) ([]*models.Post, error) {
	return implentation.ListPosts(ctx, page)
}
