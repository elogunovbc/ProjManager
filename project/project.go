package project

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Project struct {
	Id    uuid.UUID
	Name  string
	Tasks map[uuid.UUID]Task
}

var (
	ErrEmptyProjectName  = errors.New("project name cannot be empty")
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrTaskDoesNotExist  = errors.New("task does not exist")
)

func New(id uuid.UUID, name string) (*Project, error) {
	if len(name) == 0 {
		return nil, ErrEmptyProjectName
	}

	return &Project{
		Id:    id,
		Name:  name,
		Tasks: make(map[uuid.UUID]Task),
	}, nil
}

func (p Project) TaskExists(task Task) bool {
	_, exists := p.Tasks[task.Id]
	return exists
}

func (p *Project) AddTask(task Task) error {
	if p.TaskExists(task) {
		return ErrTaskAlreadyExists
	}

	p.Tasks[task.Id] = task
	return nil
}

func (p *Project) UpdateTask(task Task) error {
	if !p.TaskExists(task) {
		return ErrTaskDoesNotExist
	}

	p.Tasks[task.Id] = task
	return nil
}

func (p Project) PrintInfo() {
	//fmt.Printf("%+v\n", p)

	var builder strings.Builder

	builder.WriteString("Project:\n")
	builder.WriteString(fmt.Sprintf("\tId: %s\n", p.Id))
	builder.WriteString(fmt.Sprintf("\tName: %s\n", p.Name))
	builder.WriteString("\tTasks:\n")
	for _, task := range p.Tasks {
		builder.WriteString("\t\tTask:\n")
		builder.WriteString(fmt.Sprintf("\t\t\tId: %s\n", task.Id))
		builder.WriteString(fmt.Sprintf("\t\t\tTitle: %s\n", task.Title))
		builder.WriteString(fmt.Sprintf("\t\t\tDescription: %s\n", task.Description))
		builder.WriteString(fmt.Sprintf("\t\t\tStatus: %s\n", task.Status))
	}

	fmt.Print(builder.String())
}

func (p Project) FilterTasksByStatus(status TaskStatus) []Task {
	result := []Task{}
	for _, task := range p.Tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}
	return result
}
