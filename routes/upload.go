package routes

import (
	"imgtrace/db"
	"imgtrace/hash"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UploadImage(c *gin.Context) {
	//  Image file fetch kar
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	// Unique filename bana (timestamp + original filename)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filepath.Base(file.Filename)
	filepath := "static/uploads/" + filename
	// Save the file first

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
		return
	}

	//  Hash generate kar saved file se
	imagehash, err := hash.Generate(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating image hash"})
		return
	}

	//  DB me check kar ki ye hash already hai ya nahi
	var existingImage db.Image
	if err := db.DB.Where("hash = ?", imagehash).First(&existingImage).Error; err == nil {
		// Image already exists, to new insert mat kar, existing record return kar
		c.JSON(http.StatusOK, gin.H{
			"message": "Image already exists",

			"file_path": existingImage.FilePath,
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		// Agar DB query me koi aur error aaya to bhi handle kar
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error", "details": err.Error()})
		return
	}

	//   Agar hash unique hai to DB me naya record insert kar
	imageRecord := db.Image{
		FilePath: filepath,
		Hash:     imagehash,
	}
	if err := db.DB.Create(&imageRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed"})
		return
	}

	// Success response bhej
	c.JSON(http.StatusOK, gin.H{
		"message": "Image uploaded successfully",
	})
}
