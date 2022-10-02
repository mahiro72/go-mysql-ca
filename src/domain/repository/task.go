package repository

import (
	"github.com/mahiro72/go-mysql-ca/domain/entity"
)

type ITaskRepository interface {
	GetAllTask() (entity.Tasks, error)
	CreateTask(task *entity.Task) (*entity.Task, error)
	ChangeTaskStatus(task *entity.Task) (*entity.Task, error)
}
