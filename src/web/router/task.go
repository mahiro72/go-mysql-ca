package router

import (
	"database/sql"

	"github.com/mahiro72/go-mysql-ca/database"
	"github.com/mahiro72/go-mysql-ca/usecase"
	"github.com/mahiro72/go-mysql-ca/web/handler"
)

// NewTaskRouterはRouterへのハンドラ登録やミドルウェアの設定をします
func (r Router) NewTaskRouter(db *sql.DB) {
	taskRepository := database.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	g := r.engine.Group("/task")
	g.GET("/all", taskHandler.GetAllTask)
}
