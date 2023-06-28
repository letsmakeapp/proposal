package domain

import (
	"proposal/internal/pkgs/maybe"
	"time"
)

type TodoID struct {
	value string
}

func UnsafeCreateTodoID(value string) TodoID {
	return TodoID{
		value: value,
	}
}

func (id TodoID) Value() string {
	return id.value
}

type Todo struct {
	id        TodoID
	title     string
	dueDate   maybe.T[time.Time]
	createdAt time.Time
	updatedAt time.Time
}

func UnsafeCreateTodo(
	id TodoID,
	title string,
	dueDate maybe.T[time.Time],
	createdAt time.Time,
	updatedAt time.Time,
) Todo {
	return Todo{
		id:        id,
		title:     title,
		dueDate:   dueDate,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func CreateNewTodo(
	id TodoID,
	title string,
	dueDate maybe.T[time.Time],
	now time.Time,
) Todo {
	return UnsafeCreateTodo(
		id,
		title,
		dueDate,
		now,
		now,
	)
}

func (t Todo) ID() TodoID                  { return t.id }
func (t Todo) Title() string               { return t.title }
func (t Todo) DueDate() maybe.T[time.Time] { return t.dueDate }
func (t Todo) CreatedAt() time.Time        { return t.createdAt }
func (t Todo) UpdatedAt() time.Time        { return t.createdAt }

func (t Todo) IsDueDateExceeded(now time.Time) bool {
	if dd, ok := t.dueDate.TryGetValue(); ok {
		return now.Equal(dd) || now.After(dd)
	}
	return false
}
