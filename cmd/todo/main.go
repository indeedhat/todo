package main

import (
	"github.com/gin-gonic/gin"
	"github.com/indeedhat/todo/internal/controllers"
	"github.com/indeedhat/todo/internal/env"
	"github.com/indeedhat/todo/internal/store"
)

func main() {
	db, err := store.Connect()
	if err != nil {
		panic(err)
	}

	if err = store.Migrate(db); err != nil {
		panic(err)
	}

	router := gin.Default()

	_ = controllers.NewLists(router, db)
	_ = controllers.NewEntries(router, db)

	router.Run(env.GetFallback(env.BindAddress, ":8080"))
}
