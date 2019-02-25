package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"time"
)

func FileUpload(c *gin.Context)  {
	file, _ := c.FormFile("file")
	date := time.Now().Format("2006-01-02")

	// Upload the file to specific dst.

	//os.Mkdir("/upload",766)
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
