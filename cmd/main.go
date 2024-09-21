package main

import (
	"github.com/gin-gonic/gin"
	"templ-todo/internals"
)

func main() {
	router := gin.Default()

	//initialize config
	app := internals.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8080")
}
