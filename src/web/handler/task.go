package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahiro72/go-mysql-ca/usecase"
	"github.com/mahiro72/go-mysql-ca/web/response"
)

type TaskHandler struct {
	taskUC *usecase.TaskUseCase
}

func NewTaskHandler(u *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUC: u,
	}
}

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

	tasksRes := response.NewTasksResponse(tasks)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data": tasksRes,
		},
	)
}
