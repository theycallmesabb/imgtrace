package main

import (
	"imgtrace/db"
	"imgtrace/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()

	r.POST("/upload", routes.UploadImage) //  Pehle route define karo
	r.Run()                               // Phir server run karo

}
