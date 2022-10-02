package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahiro72/go-mysql-ca/domain/entity"
	"github.com/mahiro72/go-mysql-ca/usecase"
)

// TaskHandlerはhttpをルーティングするハンドラーに関する構造体です
type TaskHandler struct {
	taskUC *usecase.TaskUseCase
}

// NewTaskHandlerはTaskHandlerのオブジェクトのポインタを生成する関数です
func NewTaskHandler(u *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUC: u,
	}
}

// GetAllTaskはすべてのタスクを返すハンドラーです
func (h *TaskHandler) GetAllTask(ctx *gin.Context) {
	tasks, err := h.taskUC.GetAllTask()
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Errorf("TaskHandler.GetAllTask GetAllTask Error : %w", err).Error(),
			},
		)
		return
	}

	taskListResp := taskEntityListToJson(tasks)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data": taskListResp,
		},
	)
}

// CreateTaskは新しいタスクを保存するハンドラーです
func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var b taskJson
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Errorf("TaskHandler.CreateTask ShouldBindJSON Error : %w", err).Error(),
			},
		)
		return
	}
	task := taskJsonToEntity(&b)

	task, err := h.taskUC.CreateTask(task)
	log.Println("task2", task, b)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Errorf("TaskHandler.CreateTask CreateTask Error : %w", err).Error(),
			},
		)
		return
	}

	taskResp := taskEntityToJson(task)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data": taskResp,
		},
	)
}

// ChangeTaskStatusはtaskのStatusを更新するhandlerです
func (h *TaskHandler) ChangeTaskStatus(ctx *gin.Context) {
	var b taskJson
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Errorf("TaskHandler.ChangeTaskStatus ShouldBindJSON Error : %w", err).Error(),
			},
		)
		return
	}
	task := taskJsonToEntity(&b)

	task, err := h.taskUC.ChangeTaskStatus(task)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": fmt.Errorf("TaskHandler.ChangeTaskStatus ChangeTaskStatus Error : %w", err).Error(),
			},
		)
		return
	}

	taskResp := taskEntityToJson(task)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data": taskResp,
		},
	)
}

// taskJsonはtaskの情報をJSONにバインドするための構造体です
type taskJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Status string   `json:"status"`
}

// taskEntityToJsonはentity.TaskをtaskJson型に変換します
func taskEntityToJson(task *entity.Task) *taskJson {
	return &taskJson{
		Id:   task.Id,
		Name: task.Name,
		Status: task.Status,
	}
}

// taskEntityListToJsonはentity.TaskのリストをtaskJsonのリストに変換します
func taskEntityListToJson(tasks entity.Tasks) []*taskJson {
	var taskJsonList []*taskJson
	for _, t := range tasks {
		taskJsonList = append(taskJsonList, taskEntityToJson(t))
	}
	return taskJsonList
}

// taskJsonToEntityはtaskJson型のオブジェクトをentity.Taskに変換します
func taskJsonToEntity(taskJson *taskJson) *entity.Task {
	return &entity.Task{
		Id:   taskJson.Id,
		Name: taskJson.Name,
		Status: taskJson.Status,
	}
}
