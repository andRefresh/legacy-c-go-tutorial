package controllers

import (
	"go-app/service"
	"io"
	"mmf"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceViewModel struct {
	mmf.Device

	IsActivated bool
	IsRed       bool // just to show some colors
}

func NewDeviceViewModel(d mmf.Device) (vm DeviceViewModel) {
	vm.Device = d
	vm.IsActivated = d.IsActive == 1
	vm.IsRed = d.SomeValue%5 == 0

	return
}

type DeviceController struct {
	service.Watcher
	*mmf.Device
}

func NewDeviceController(d *mmf.Device) (dc *DeviceController) {
	dc = &DeviceController{Device: d}
	dc.Watcher = *service.NewWatcher()
	dc.StartWatch("device", dc.Device, 100)

	return
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

func (dc *DeviceController) Sse(c *gin.Context) {
	sub := dc.Subscribe(c)

	c.Stream(func(w io.Writer) bool {
		for {
			select {
			case msg := <-sub:
				msgWatch := msg.(service.WatchData)
				c.SSEvent(msgWatch.Id, msgWatch.Data)
				return true

			case <-c.Writer.CloseNotify():
				dc.Unsubscribe(c)
				return false

			}
		}
	})
}

func (dc *DeviceController) Show(c *gin.Context) {
	c.HTML(http.StatusOK, "device.html", gin.H{
		"viewModel": NewDeviceViewModel(*dc.Device),
	})
}
