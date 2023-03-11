package task

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

type TaskService interface {
	FindAll() []Task
	FindById(id string) (*Task, error)
	Create(Task) (*Task, error)
	Update(Task) (*Task, error)
	Delete(id string) error
}

type taskService struct {
	tasks map[string]Task
}

func NewService() TaskService {
	return &taskService{
		tasks: map[string]Task{},
	}
}

func (t *taskService) FindAll() []Task {
	log.Println("FindAll")
	return maps.Values(t.tasks)
}

func (s *taskService) FindById(id string) (*Task, error) {
	log.Println("FindById", id)
	t, ok := s.tasks[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &t, nil
}

func (s *taskService) Create(t Task) (*Task, error) {
	log.Println("Create", t)
	if t.ID == "" {
		id := uuid.New().String()
		t.ID = id
		s.tasks[id] = t
	} else {
		s.tasks[t.ID] = t
	}
	return &t, nil
}

func (s *taskService) Update(t Task) (*Task, error) {
	log.Println("Update", t)
	if t.ID == "" {
		return nil, errors.New("ID is required")
	} else {
		s.tasks[t.ID] = t
	}
	return &t, nil
}

func (s *taskService) Delete(id string) error {
	log.Println("Delete", id)
	if _, err := s.FindById(id); err != nil {
		return err
	}
	delete(s.tasks, id)
	return nil
}
