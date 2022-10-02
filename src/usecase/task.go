package usecase

import (
	"fmt"

	"github.com/mahiro72/go-mysql-ca/domain/entity"
	"github.com/mahiro72/go-mysql-ca/domain/repository"
)

// TaskUseCaseはTaskに関するユースケースです
type TaskUseCase struct {
	taskRepo repository.TaskRepository
}

// NewTaskUseCaseはTaskUseCaseのオブジェクトのポインタを生成する関数です
func NewTaskUseCase(r repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: r,
	}
}

// GetAllTaskはrepositoryの関数を呼び出し、すべてのタスクを返します
func (u *TaskUseCase) GetAllTask() ([]*entity.Task, error) {
	tasks, err := u.taskRepo.GetAllTask()
	if err != nil {
		return nil, fmt.Errorf("TaskUseCase.GetAllTask GetAllTask Error : %w", err)
	}
	return tasks, nil
}

// CreateTaskはrepositoryの関数を呼び出し、新しいtaskを保存します
func (u *TaskUseCase) CreateTask(task *entity.Task) (*entity.Task, error) {
	task, err := u.taskRepo.CreateTask(task)
	if err != nil {
		return nil, fmt.Errorf("TaskUseCase.CreateTask CreateTask Error : %w", err)
	}
	return task, nil
}

func (u *TaskUseCase) ChangeTaskStatus(task *entity.Task) (*entity.Task, error) {
	task, err := u.taskRepo.ChangeTaskStatus(task)
	if err != nil {
		return nil, fmt.Errorf("TaskUseCase.ChangeTaskStatus ChangeTaskStatus Error : %w", err)
	}
	return task, nil
}
