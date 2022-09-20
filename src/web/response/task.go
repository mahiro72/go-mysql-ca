package response

import "github.com/mahiro72/go-mysql-ca/domain/entity"

type TaskResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// NewTaskResponseはentity.TaskをTaskResponse型に変換します
func NewTaskResponse(t *entity.Task) *TaskResponse {
	return &TaskResponse{
		Id:   t.Id,
		Name: t.Name,
		Done: t.Done,
	}
}

// NewTasksResponseはentity.TaskのリストをTaskResponseのリストに変換します
func NewTasksResponse(tasks []*entity.Task) []*TaskResponse {
	var tasksResponse []*TaskResponse
	for _, t := range tasks {
		tasksResponse = append(tasksResponse, NewTaskResponse(t))
	}
	return tasksResponse
}
