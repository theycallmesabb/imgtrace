package routes

import (
	"imgtrace/db"
	"imgtrace/hash"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SearchImage(c *gin.Context) {
	// Get uploaded file from form
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	// Save uploaded file temporarily
	tempPath := filepath.Join("static/uploads", file.Filename)
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Generate hash of uploaded image
	uploadedHash, err := hash.Generate(tempPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate hash"})
		return
	}

	// Get all images from DB
	var images []db.Image
	db.DB.Find(&images)

	// Compare uploaded hash with DB hashes
	for _, img := range images {
		distance, _ := hash.CompareHashes(uploadedHash, img.Hash)
		if distance <= 5 {
			c.JSON(http.StatusOK, gin.H{
				"message": "Image already exists",
				"url":     img.FilePath,
			})
			return
		}
	}

	// If not found, save new image info
	newImage := db.Image{
		FilePath: tempPath,
		Hash:     uploadedHash,
	}
	db.DB.Create(&newImage)

	c.JSON(http.StatusOK, gin.H{
		"message": "Image uploaded successfully and saved",
		"url":     tempPath,
	})
}
