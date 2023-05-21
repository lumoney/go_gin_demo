package util

import (
	"fmt"
	"net/http"
	"strconv"
	"work/response"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		fmt.Println(file.Filename)
		// 上传文件至指定目录
		dst := "./static/" + file.Filename
		c.SaveUploadedFile(file, dst)
	}
	response.Response(c, http.StatusOK, 200, nil, strconv.Itoa(len(files))+"文件上传成功")
}
