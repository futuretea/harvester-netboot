package main

import (
	"github.com/futuretea/harvester-netboot/pkg/config"
	"github.com/futuretea/harvester-netboot/pkg/handler"
	"github.com/futuretea/harvester-netboot/pkg/log"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	config.Load()

	// log
	log.Setup()

	// service
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET(handler.BootPath, handler.BootHandler)
	r.GET(handler.DownloadPath, handler.DownloadHandler)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
