package main

import (
	"log"

	"github.com/mahiro72/go-mysql-ca/infrastructure/database"
	"github.com/mahiro72/go-mysql-ca/web/router"
)

func main() {
	conn, err := database.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Routerの初期化
	r := router.NewRouter()
	r.SetMiddleware()

	r.Health()
	r.NewTaskRouter(conn)

	// Routerの起動
	r.Serve()
}
