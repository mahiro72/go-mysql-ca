package repository

import (
	"github.com/mahiro72/go-mysql-ca/domain/entity"
)

type TaskRepository interface {
	GetAllTask() ([]*entity.Task, error)
}
