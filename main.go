package main

import (
	"assignment-3/controllers"
	"assignment-3/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	go controllers.UpdateStatusPeriodically()

	r := gin.Default()

	r.PUT("/status/update", controllers.UpdateStatus)

	r.Run(":8080")
}