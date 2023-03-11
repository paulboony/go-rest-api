package task

import (
	"reflect"
	"testing"
)

func TestCreateTaskSavedCorrectly(t *testing.T) {
	s := NewService()
	task := Task{ID: "t1", Title: "Buy milk"}

	s.Create(task)

	savedTask, _ := s.FindById("t1")

	if savedTask == nil {
		t.Errorf("Expected task, but got %v", *savedTask)
	}
	if !reflect.DeepEqual(task, *savedTask) {
		t.Errorf("Expected task to be %v, but got %v", task, *savedTask)
	}
}

func TestCreateTaskGeneratesUuidForId(t *testing.T) {
	s := NewService()

	s.Create(Task{Title: "Buy milk"})

	tasks := s.FindAll()
	if len(tasks) != 1 {
		t.Errorf("Expected task length of 1, but got %v", len(tasks))
	}
	if tasks[0].ID == "" {
		t.Errorf("Expected task to have ID, but it's empty")
	}
}

func TestUpdateTaskSavedCorrectly(t *testing.T) {
	s := NewService()
	expectedTask := Task{ID: "t1", Title: "Buy apple"}

	s.Create(Task{ID: "t1", Title: "Buy milk"})
	updatedTask, err := s.Update(expectedTask)

	if err != nil {
		t.Errorf("Expected Update() not to have error, but got %v", err)
	}
	if updatedTask == nil {
		t.Errorf("Expected Update() to return task, but got %v", *updatedTask)
	}
	if !reflect.DeepEqual(expectedTask, *updatedTask) {
		t.Errorf("Expected task to be %v, but got %v", expectedTask, *updatedTask)
	}
}

func TestUpdateTaskFailsWithoutId(t *testing.T) {
	s := NewService()
	expectedTask := Task{ID: "", Title: "Buy apple"}

	s.Create(Task{ID: "t1", Title: "Buy milk"})
	updatedTask, err := s.Update(expectedTask)

	if err == nil {
		t.Errorf("Expected Update() to have error, but got %v", err)
	}
	if updatedTask != nil {
		t.Errorf("Expected Update() not to return task, but got %v", *updatedTask)
	}
}

func TestDeleteTask(t *testing.T) {
	s := NewService()
	s.Create(Task{ID: "t1", Title: "Buy apple"})
	s.Create(Task{ID: "t2", Title: "Buy banana"})
	s.Create(Task{ID: "t3", Title: "Buy cherry"})

	s.Delete("t2")

	tasks := s.FindAll()
	if len(tasks) != 2 {
		t.Errorf("Expected task length of 2, but got %v", len(tasks))
	}
	if task, _ := s.FindById("t2"); task != nil {
		t.Errorf("Expected task removed, but it still exists")
	}
}

func TestDeleteTaskFailsWhenNotFound(t *testing.T) {
	s := NewService()

	err := s.Delete("t2")

	if err == nil {
		t.Errorf("Expected error, but got nothing")
	}
}
