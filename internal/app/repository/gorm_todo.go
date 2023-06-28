package repository

import (
	"context"
	"proposal/internal/app/domain"
	"proposal/internal/pkgs/maybe"
	"proposal/internal/pkgs/sliceutil"
	"time"

	"gorm.io/gorm"
)

type gormTodoModel struct {
	ID        string     `gorm:"primaryKey;column:id"`
	Title     string     `gorm:"column:title"`
	DueDate   *time.Time `gorm:"column:due_date"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
}

func (m gormTodoModel) ToDomain() domain.Todo {
	return domain.UnsafeCreateTodo(
		domain.UnsafeCreateTodoID(m.ID),
		m.Title,
		maybe.FromPointer(m.DueDate),
		m.CreatedAt,
		m.UpdatedAt,
	)
}

func createGormTodoModelFromDomain(t domain.Todo) gormTodoModel {
	return gormTodoModel{
		ID:        t.ID().Value(),
		Title:     t.Title(),
		DueDate:   t.DueDate().ToPointer(),
		CreatedAt: t.CreatedAt(),
		UpdatedAt: t.UpdatedAt(),
	}
}

type GormTodoRepository struct {
	db *gorm.DB
}

var _ TodoRepository = (*GormTodoRepository)(nil)

func NewGormTodoRepository(db *gorm.DB) GormTodoRepository {
	return GormTodoRepository{
		db: db,
	}
}

func (r *GormTodoRepository) FindAll(ctx context.Context) ([]domain.Todo, error) {
	ctx, span := tracer.Start(ctx, "GormTodoRepository.FindAll")
	defer span.End()

	var models []gormTodoModel
	res := r.db.WithContext(ctx).Find(&models)
	if err := res.Error; err != nil {
		return nil, err
	}

	return sliceutil.Map(
		models,
		func(m gormTodoModel) domain.Todo {
			return m.ToDomain()
		},
	), nil
}

func (r *GormTodoRepository) FindByID(ctx context.Context, id domain.TodoID) (maybe.T[domain.Todo], error) {
	ctx, span := tracer.Start(ctx, "GormTodoRepository.FindByID")
	defer span.End()

	var model gormTodoModel
	res := r.db.WithContext(ctx).Where("id = ?", id.Value()).Take(&model)
	if err := res.Error; err != nil {
		if isErrorGorm(err, errRecordNotFound) {
			return maybe.None[domain.Todo](), ErrTodoNotFound
		}
		return maybe.None[domain.Todo](), err
	}
	return maybe.Some[domain.Todo](model.ToDomain()), nil
}

func (r *GormTodoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	ctx, span := tracer.Start(ctx, "GormTodoRepository.Create")
	defer span.End()

	model := createGormTodoModelFromDomain(*todo)
	res := r.db.WithContext(ctx).Create(&model)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r *GormTodoRepository) Update(ctx context.Context, todo *domain.Todo) error {
	ctx, span := tracer.Start(ctx, "GormTodoRepository.Update")
	defer span.End()

	model := createGormTodoModelFromDomain(*todo)
	res := r.db.WithContext(ctx).Save(&model)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
