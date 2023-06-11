package usecase_test

import (
	"2/model"
	"2/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) Create(task *model.Task) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}

func (m *TaskRepositoryMock) Read(id int) (*model.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) Update(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *TaskRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestTaskUsecase(t *testing.T) {
	mockRepo := new(TaskRepositoryMock)
	taskUsecase := usecase.NewTaskUsecase(mockRepo)

	task := &model.Task{Title: "title"}

	mockRepo.On("Create", task).Return(1, nil)

	id, err := taskUsecase.CreateTask(task.Title)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}
