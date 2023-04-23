package handler

import "github.com/gin-gonic/gin"

const (
	DownloadPath = "/tmp/:filename"
)

func DownloadHandler(c *gin.Context) {
	filename := c.Param("filename")
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("/tmp/" + filename)
}
