package usecase

import (
	"fmt"
	"testing"

	"github.com/mahiro72/go-mysql-ca/domain/entity"
)

// TestTaskUsecase_GetAllTaskはusecaseのGetAllTaskをテストします
func TestTaskUsecase_GetAllTask(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		task    *entity.Task
		wantErr error
	}{
		{
			name:    "正常に動いている場合",
			task:    &entity.Task{},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskUC := NewTaskUseCase(&taskRepositoryMock{})
			if _, err := taskUC.GetAllTask(); err != tt.wantErr {
				t.Errorf("TestTaskUsecase_GetAllTask Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

// TestTaskUsecase_CreateTaskはusecaseのCreateTaskをテストします
func TestTaskUsecase_CreateTask(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		task    *entity.Task
		wantErr error
	}{
		{
			name: "正常に動いている場合",
			task: &entity.Task{
				Name: "hoge",
			},
			wantErr: nil,
		},
		{
			name:    "必要なフィールド(name)がなかった場合",
			task:    &entity.Task{},
			wantErr: fmt.Errorf("TaskUseCase.CreateTask Name Error : task name required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskUC := NewTaskUseCase(&taskRepositoryMock{})
			_, err := taskUC.CreateTask(tt.task)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestTaskUsecase_GetAllTask Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

// TestTaskUsecase_ChangeTaskStatusはusecaseのChangeTaskStatusをテストします
func TestTaskUsecase_ChangeTaskStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		task    *entity.Task
		wantErr error
	}{
		{
			name: "正常に動いている場合",
			task: &entity.Task{
				Id:     1,
				Status: "done",
			},
			wantErr: nil,
		},
		{
			name:    "必要なフィールド(id)がなかった場合",
			task:    &entity.Task{},
			wantErr: fmt.Errorf("TaskUseCase.ChangeTaskStatus Id Error : task id required"),
		},
		{
			name: "必要なフィールド(status)がなかった場合",
			task: &entity.Task{
				Id: 1,
			},
			wantErr: fmt.Errorf("TaskUseCase.ChangeTaskStatus Status Error : task status required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskUC := NewTaskUseCase(&taskRepositoryMock{})
			_, err := taskUC.ChangeTaskStatus(tt.task)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("TestTaskUsecase_ChangeTaskStatus Error : want %v, but got %v", tt.wantErr, err)
			}
		})
	}
}

// taskRepositoryMockはtaskRepositoryのmockの構造体です
type taskRepositoryMock struct{}

func (r *taskRepositoryMock) GetAllTask() (entity.Tasks, error) {
	return entity.Tasks{}, nil
}

func (r *taskRepositoryMock) CreateTask(task *entity.Task) (*entity.Task, error) {
	return task, nil
}

func (r *taskRepositoryMock) ChangeTaskStatus(task *entity.Task) (*entity.Task, error) {
	return task, nil
}
