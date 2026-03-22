package project

import (
	"errors"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Title       string
	Description string
	Status      TaskStatus
}

var (
	ErrEmptyTaskTitle       = errors.New("task title cannot be empty")
	ErrEmptyTaskDescription = errors.New("task description cannot be empty")
	ErrTaskAlreadyClosed    = errors.New("task already closed")
)

func NewTask(id uuid.UUID, title string, description string) (*Task, error) {
	if len(title) == 0 {
		return nil, ErrEmptyTaskTitle
	}
	if len(description) == 0 {
		return nil, ErrEmptyTaskDescription
	}

	return &Task{
		Id:          id,
		Title:       title,
		Description: description,
		Status:      StatusActive,
	}, nil
}

func (t *Task) Close() error {
	if t.Status == StatusClosed {
		return ErrTaskAlreadyClosed
	}

	t.Status = StatusClosed
	return nil
}

func (t *Task) UpdateDescription(description string) error {
	if t.Status == StatusClosed {
		return ErrTaskAlreadyClosed
	}
	if len(description) == 0 {
		return ErrEmptyTaskDescription
	}

	t.Description = description
	return nil
}
