package controllers

import (
	"mmf"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceController struct {
	Device *mmf.Device
}

func (dc *DeviceController) SetDevice(c *gin.Context) {
	on := c.Query("on") == "true"
	dc.Device.SetOn(on)
}

func (dc *DeviceController) ResetDevice(c *gin.Context) {
	dc.Device.ResetValue()
}

func (dc *DeviceController) GetDeviceValue(c *gin.Context) {
	c.JSON(http.StatusOK, dc.Device.SomeValue)
}

func (dc *DeviceController) GetDeviceJson(c *gin.Context) {
	c.JSON(http.StatusOK, dc.Device)
}
