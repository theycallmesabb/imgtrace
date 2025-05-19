package main

import (
	"imgtrace/db"
	"imgtrace/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()

	r.POST("/upload", routes.UploadImage)
	r.POST("/search", routes.SearchImage)
	r.Static("/static", "./static")

	r.Run()

}
