package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppController struct{}

func NewAppController() (ac *AppController) {
	ac = &AppController{}
	return
}

func (ac *AppController) Show(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
