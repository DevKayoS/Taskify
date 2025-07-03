package services

import (
	"testing"
	"time"

	"github.com/DevKayoS/taskify/internal/store"
	"github.com/stretchr/testify/assert"
)

type MockTaskStore struct{}

func (m *MockTaskStore) CreateTask(title, description string, priority int32) (store.Task, error) {
	return store.Task{
		Id:          1,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) GetTaskById(id int32) (store.Task, error) {
	return store.Task{
		Id:          id,
		Title:       "Mock Title",
		Description: "Mock Description",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) ListTasks() ([]store.Task, error) {
	return []store.Task{
		{
			Id:          1,
			Title:       "Mock Title",
			Description: "Mock Description",
			Priority:    9,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          2,
			Title:       "Mock Title 2",
			Description: "Mock Description 2",
			Priority:    3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}

func (m *MockTaskStore) UpdateTask(id int32, title, description string, priority int32) (store.Task, error) {
	return store.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) DeleteTask(id int32) error {
	return nil
}

func TestCreateTask(t *testing.T) {
	mockStore := MockTaskStore{}
	TaskService := NewTaskService(&mockStore)

	task, err := TaskService.Store.CreateTask("Mock Title", "Mock Description", 2)

	assert.NoError(t, err)
	assert.Equal(t, "Mock Title", task.Title)
	assert.Equal(t, "Mock Description", task.Description)
	assert.Equal(t, int32(2), task.Priority)
}

func TestGetTask(t *testing.T) {
	mockStore := MockTaskStore{}
	TaskService := NewTaskService(&mockStore)

	task, err := TaskService.Store.GetTaskById(1)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), task.Id)
	assert.Equal(t, "Mock Title", task.Title)
	assert.Equal(t, "Mock Description", task.Description)
	assert.Equal(t, int32(1), task.Priority)
}

func TestListTasks(t *testing.T) {
	mockStore := MockTaskStore{}
	TaskService := NewTaskService(&mockStore)

	tasks, err := TaskService.Store.ListTasks()

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, int32(1), tasks[0].Id)
	assert.Equal(t, "Mock Title", tasks[0].Title)
	assert.Equal(t, "Mock Description", tasks[0].Description)

	assert.Equal(t, int32(2), tasks[1].Id)
	assert.Equal(t, "Mock Title 2", tasks[1].Title)
	assert.Equal(t, "Mock Description 2", tasks[1].Description)
}
