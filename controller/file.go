package controller

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func FileUpload(c *gin.Context) {
	pwd, _ := os.Getwd()
	file, _ := c.FormFile("file")
	//检查图片大小，扩展名，类型
	//定义允许图片类型

	filetype := file.Header["Content-Type"]
	if filetype[0] == "image/jpeg" || filetype[0] == "image/png" {

	} else {
		ResJson(402, "图片格式错误仅支持JPG,PNG请检查", c)
		return
	}

	if file.Size > 4194304 {
		ResJson(402, "图片大小不能超过4M", c)
		return
	}
	ftime := time.Now().Format("20060102")
	filetime := pwd + "/upload/" + ftime
	if err := os.MkdirAll(filetime, 0755); err != nil {
		ResJson(402, "图片创建失败", c)
		return
	}
	if err := c.SaveUploadedFile(file, filetime+"/"+file.Filename); err != nil {
		ResJson(402, "图片上传失败", c)
		return
	}
	ResJson(200, "/upload/"+ftime+"/"+file.Filename, c)
	return
}
