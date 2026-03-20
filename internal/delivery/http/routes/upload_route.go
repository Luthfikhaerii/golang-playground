package routes

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadRoute(r *gin.RouterGroup) {
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"message": "file not found"})
			return
		}

		// validasi ekstensi
		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".png" {
			c.JSON(400, gin.H{"error": "format harus jpg/png"})
			return
		}

		// rename
		filename := fmt.Sprintf("%d%s", time.Now().Unix(), ext)

		err = c.SaveUploadedFile(file, "./upload/"+filename)
		if err != nil {
			c.JSON(400, gin.H{"message": "failed save file"})
			return
		}

		c.JSON(200, gin.H{
			"message":  "upload berhasil",
			"filename": file.Filename,
		})
	})
}
