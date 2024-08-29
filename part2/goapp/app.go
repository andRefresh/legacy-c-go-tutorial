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

	// we can call more readable methods
	device.ResetValue()
	device.SetOn(true)
	device.SetOn(false)

	// if we want we can also use here some C exported structs,
	// because of the public alias
	d2 := mmf.Device{}
	d2.SetOn(true) // does nothing, just showing purposes

	// we can also do now more advanced tasks
	controller := &controllers.DeviceController{Device: device}

	r := gin.Default()

	r.GET("/reset", controller.ResetDevice)
	r.GET("/value", controller.GetDeviceValue)
	r.GET("/set_on", controller.SetDevice)
	r.GET("/device_json", controller.GetDeviceJson)

	r.Run()
}
