package provider

import (
	"context"
	"proposal/internal/app/domain"
)

type TodoIdProvider interface {
	Next(ctx context.Context) (domain.TodoID, error)
}
