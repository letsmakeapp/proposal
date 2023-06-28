package provider

import (
	"context"
	"proposal/internal/app/domain"

	"github.com/google/uuid"
)

type UUIDv4TodoIdProvider struct{}

var _ TodoIdProvider = (*UUIDv4TodoIdProvider)(nil)

func NewUUIDv4TodoIdProvider() UUIDv4TodoIdProvider {
	return UUIDv4TodoIdProvider{}
}

func (*UUIDv4TodoIdProvider) Next(ctx context.Context) (domain.TodoID, error) {
	_, span := tracer.Start(ctx, "UUIDv4TodoIdProvider.Next")
	defer span.End()

	id := uuid.New()
	return domain.UnsafeCreateTodoID(id.String()), nil
}
