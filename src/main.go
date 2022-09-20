package main

import (
	"log"

	"github.com/mahiro72/go-mysql-ca/database"
	"github.com/mahiro72/go-mysql-ca/web/router"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Routerの初期化
	r := router.NewRouter()
	r.Health()
	r.NewTaskRouter(db)

	// Routerの起動
	r.Serve()
}
