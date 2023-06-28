package usecase

import (
	"context"
	"proposal/internal/app/domain"
	"proposal/internal/app/provider"
	"proposal/internal/app/repository"
	"proposal/internal/pkgs/maybe"
	"proposal/internal/pkgs/timeutil"
	"time"
)

type TodoUseCase interface {
	AddTodo(ctx context.Context, input AddTodoInput) (AddTodoOutput, error)
}

type AddTodoInput struct {
	Title   string
	DueDate maybe.T[time.Time]
}

type AddTodoOutput struct {
	TodoID string
}

type TodoUseCaseImpl struct {
	todoRepo repository.TodoRepository
	todoIdp  provider.TodoIdProvider
}

var _ TodoUseCase = (*TodoUseCaseImpl)(nil)

func NewTodoUseCaseImpl(
	todoRepo repository.TodoRepository,
	todoIdp provider.TodoIdProvider,
) *TodoUseCaseImpl {
	return &TodoUseCaseImpl{
		todoRepo: todoRepo,
		todoIdp:  todoIdp,
	}
}

func (u *TodoUseCaseImpl) AddTodo(
	ctx context.Context,
	input AddTodoInput,
) (AddTodoOutput, error) {
	ctx, span := tracer.Start(ctx, "TodoUseCase.AddTodo")
	defer span.End()

	id, err := u.todoIdp.Next(ctx)
	if err != nil {
		var zero AddTodoOutput
		return zero, err
	}

	todo := domain.CreateNewTodo(
		id,
		input.Title,
		input.DueDate,
		timeutil.Now(),
	)

	err = u.todoRepo.Create(ctx, &todo)
	if err != nil {
		var zero AddTodoOutput
		return zero, err
	}

	return AddTodoOutput{
		TodoID: todo.ID().Value(),
	}, nil
}
