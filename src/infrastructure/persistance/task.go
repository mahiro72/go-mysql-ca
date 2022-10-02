package persistance

import (
	"fmt"

	"github.com/mahiro72/go-mysql-ca/domain/entity"
	"github.com/mahiro72/go-mysql-ca/domain/repository"
	"github.com/mahiro72/go-mysql-ca/infrastructure/database"
)

// TaskRepositoryがdomainのinterfaceを満たしているか確認します
var _ repository.ITaskRepository = &TaskRepository{}

// TaskRepositoryはrepository.TaskRepositoryを満たす構造体です
type TaskRepository struct {
	conn *database.Conn
}

// NewTaskRepositoryはTaskRepositoryのポインタを生成する関数です
func NewTaskRepository(conn *database.Conn) *TaskRepository {
	return &TaskRepository{
		conn: conn,
	}
}

// GetAllTaskはすべてのTaskを返します
func (r *TaskRepository) GetAllTask() (entity.Tasks, error) {
	var dtos []taskDTO
	query := `SELECT * FROM tasks`

	if err := r.conn.DB.Select(&dtos, query); err != nil {
		return nil, fmt.Errorf("TaskRepositoy.GetAllTask Select Error : %w", err)
	}
	return taskDTOsToEntities(dtos), nil
}

// CreateTaskは受け取ったtaskをDBに保存します
func (r *TaskRepository) CreateTask(task *entity.Task) (*entity.Task, error) {
	query := `INSERT INTO tasks (name, status) VALUES (:name, "todo")`

	dto := taskEntityToDTO(task)
	_, err := r.conn.DB.NamedExec(query, dto)
	if err != nil {
		return nil, fmt.Errorf("TaskRepositoy.CreateTask NamedExec Error : %w", err)
	}
	return taskDTOtoEntity(dto), nil
}

// ChangeTaskStatusは指定したidをtasksで検索し、見つけた場合statusを更新します
func (r *TaskRepository) ChangeTaskStatus(task *entity.Task) (*entity.Task, error) {
	query := `UPDATE tasks SET status = :status where id = :id`

	dto := taskEntityToDTO(task)
	_, err := r.conn.DB.NamedExec(query, dto)
	if err != nil {
		return nil, fmt.Errorf("TaskRepositoy.ChangeTaskStatus NamedExec Error : %w", err)
	}
	return taskDTOtoEntity(dto), nil
}

// entityのtaskのデータをやり取りするオブジェクトです
// またDTOとは、Data Transfer Objectの略です
type taskDTO struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Status string `db:"status"`
}

// taskEntityToDTOはtaskのentityをdtoに変えます
func taskEntityToDTO(task *entity.Task) taskDTO {
	return taskDTO{
		Id:     task.Id,
		Name:   task.Name,
		Status: task.Status,
	}
}

// taskEntitiesToDTOsはtaskのdtoのリストをentityのリストをに変えます
func taskDTOsToEntities(dtos []taskDTO) entity.Tasks {
	tasks := entity.Tasks{}
	for _, dto := range dtos {
		tasks = append(tasks, &entity.Task{
			Id:     dto.Id,
			Name:   dto.Name,
			Status: dto.Status,
		})
	}
	return tasks
}

// taskDTOtoEntityはtaskのdtoをentityに変えます
func taskDTOtoEntity(dto taskDTO) *entity.Task {
	return &entity.Task{
		Id:     dto.Id,
		Name:   dto.Name,
		Status: dto.Status,
	}
}
