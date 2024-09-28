package main

import (
	"flag"
	"go-app/controllers"
	"mmf"

	"github.com/gin-gonic/gin"
)

var mmfFilePath string

func main() {
	flag.StringVar(&mmfFilePath, "mmf", "./memory_mapped_file.mmf", "memory mapped file path")
	flag.Parse()

	memoryFile := mmf.OpenMmf(mmfFilePath)
	device := memoryFile.GetDevice()

	deviceController := controllers.NewDeviceController(device)
	appController := controllers.NewAppController()

	r := gin.Default()
	r.LoadHTMLGlob("html/*.html")

	api := r.Group("/api")
	api.GET("/reset", deviceController.ResetDevice)
	api.GET("/value", deviceController.GetDeviceValue)
	api.GET("/set_on", deviceController.SetDevice)
	api.GET("/device_json", deviceController.GetDeviceJson)
	api.GET("/sse", deviceController.Sse)

	components := r.Group("/components")
	components.GET("/device", deviceController.Show)

	r.GET("/", appController.Show)

	r.Run()
}
