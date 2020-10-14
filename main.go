package main

import (
	"log"

	"github.com/SushanthGunjal/BooleanAsAService/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connector()
	r := gin.Default()
	r.POST("/", AddService)
	r.GET("/:id", GetService)
	r.PATCH("/:id", UpdateService)
	r.DELETE("/:id", DeleteService)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
