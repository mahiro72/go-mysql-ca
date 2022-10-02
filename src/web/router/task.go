package router

import (
	"github.com/mahiro72/go-mysql-ca/infrastructure/database"
	"github.com/mahiro72/go-mysql-ca/infrastructure/persistance"
	"github.com/mahiro72/go-mysql-ca/usecase"
	"github.com/mahiro72/go-mysql-ca/web/handler"
)

// NewTaskRouterはRouterへのハンドラ登録やミドルウェアの設定をします
func (r Router) NewTaskRouter(conn *database.Conn) {
	taskRepository := persistance.NewTaskRepository(conn)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	g := r.engine.Group("/task")
	g.GET("/all", taskHandler.GetAllTask)
	g.POST("/create", taskHandler.CreateTask)
	g.PUT("/status", taskHandler.ChangeTaskStatus)
}
