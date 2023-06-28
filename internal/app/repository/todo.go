package repository

import (
	"context"
	"errors"
	"proposal/internal/app/domain"
	"proposal/internal/pkgs/maybe"
)

var (
	ErrTodoNotFound      = errors.New("todo not found")
	ErrTodoAlreadyExists = errors.New("todo already exists")
)

type TodoRepository interface {
	FindByID(ctx context.Context, id domain.TodoID) (maybe.T[domain.Todo], error)
	FindAll(ctx context.Context) ([]domain.Todo, error)

	Create(ctx context.Context, todo *domain.Todo) error
	Update(ctx context.Context, todo *domain.Todo) error
}
