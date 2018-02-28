package main

import (
	"github.com/gin-gonic/gin"

	controllers "Helio/Tugas1.3/controllers"
)

func main() {
	route := gin.Default()

	route.GET("/books", controllers.GetBook)

	route.Run(":8080")

}
