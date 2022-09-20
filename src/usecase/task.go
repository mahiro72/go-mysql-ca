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

func NewTaskUseCase(r repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: r,
	}
}

func (u *TaskUseCase) GetAllTask() ([]*entity.Task, error) {
	tasks, err := u.taskRepo.GetAllTask()
	if err != nil {
		return nil, fmt.Errorf("TaskUseCase.GetAllTask GetAllTask Error : %w", err)
	}
	return tasks, nil
}
