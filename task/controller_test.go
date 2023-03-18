package task

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"paulboony/go-rest-api/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) FindAll() []Task {
	args := m.Called()
	return args.Get(0).([]Task)
}

func (m *MockTaskService) FindById(id string) (*Task, error) {
	args := m.Called(id)
	return args.Get(0).(*Task), args.Error(1)
}

func (m *MockTaskService) Create(task Task) (*Task, error) {
	args := m.Called(task)
	return args.Get(0).(*Task), args.Error(1)
}

func (m *MockTaskService) Update(task Task) (*Task, error) {
	args := m.Called(task)
	return args.Get(0).(*Task), args.Error(1)
}

func (m *MockTaskService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestFindAll(t *testing.T) {
	r := gin.Default()
	tasks := []Task{{ID: "123", Title: "Buy milk"}}
	mockTaskService := new(MockTaskService)
	mockTaskService.On("FindAll").Return(tasks)
	controller := NewController(mockTaskService)
	Route(r, controller)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assertTasks(t, tasks, res.Body.String())
}

func TestFindById(t *testing.T) {
	r := gin.Default()
	id := "cafe"
	task := Task{ID: id, Title: "Buy milk"}
	mockTaskService := new(MockTaskService)
	mockTaskService.On("FindById", id).Return(&task, nil)
	controller := NewController(mockTaskService)
	Route(r, controller)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/"+id, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assertTask(t, task, res.Body.String())
}

func TestCreate(t *testing.T) {
	r := gin.Default()
	mockTaskService := new(MockTaskService)
	id := "cafe"
	task := Task{ID: id, Title: "Buy milk"}
	mockTaskService.On("Create", task).Return(&task, nil)
	controller := NewController(mockTaskService)
	Route(r, controller)

	b, _ := json.Marshal(task)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(b))
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
	assertTask(t, task, res.Body.String())
}

func TestUpdate(t *testing.T) {
	r := gin.Default()
	mockTaskService := new(MockTaskService)
	id := "cafe"
	task := Task{ID: id, Title: "Buy milk"}
	mockTaskService.On("FindById", id).Return(&task, nil)
	mockTaskService.On("Update", task).Return(&task, nil)
	controller := NewController(mockTaskService)
	Route(r, controller)

	b, _ := json.Marshal(task)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/tasks/"+id, bytes.NewBuffer(b))
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assertTask(t, task, res.Body.String())
}

func TestDelete(t *testing.T) {
	r := gin.Default()
	id := "cafe"
	task := Task{ID: id, Title: "Buy milk"}
	mockTaskService := new(MockTaskService)
	mockTaskService.On("FindById", id).Return(&task, nil)
	mockTaskService.On("Delete", id).Return(nil)
	controller := NewController(mockTaskService)
	Route(r, controller)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/"+id, nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusNoContent, res.Code)
	assert.Equal(t, "", res.Body.String())
}

func assertTasks(t *testing.T, tasks []Task, actualJson string) {
	expectedBytes, _ := json.Marshal(util.ResourcePayload(tasks))
	assert.Equal(t, string(expectedBytes), actualJson)
}

func assertTask(t *testing.T, task Task, actualJson string) {
	expectedBytes, _ := json.Marshal(util.ResourcePayload(task))
	assert.Equal(t, string(expectedBytes), actualJson)
}
