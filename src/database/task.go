package database

import (
	"database/sql"
	"fmt"

	"github.com/mahiro72/go-mysql-ca/domain/entity"
	"github.com/mahiro72/go-mysql-ca/domain/repository"
)

// TaskRepositoryがdomainのinterfaceを満たしているか確認します
var _ repository.TaskRepository = &TaskRepository{}

// TaskRepositoryはrepository.TaskRepositoryを満たす構造体です
type TaskRepository struct {
	db *sql.DB
}

// NewTaskRepositoryはTaskRepositoryのポインタを生成する関数です
func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

// GetAllTaskはすべてのTaskを返します
func (r *TaskRepository) GetAllTask() ([]*entity.Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("TaskRepositoy.GetAllTask Query Error : %w", err)
	}
	defer rows.Close()

	var tasks []*entity.Task

	for rows.Next() {
		t := &entity.Task{}
		if err := rows.Scan(&t.Id, &t.Name, &t.Done); err != nil {
			return nil, fmt.Errorf("TaskRepositoy.GetAllTask Scan Error : %w", err)
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("TaskRepositoy.GetAllTask rows Error : %w", err)
	}

	return tasks, nil
}
