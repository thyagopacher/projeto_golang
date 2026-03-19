package controllers

import (
	"github.com/gin-gonic/gin"
	"projeto_go/internal/services"
)

type HomeController struct {
	service *services.HomeService
}

func NewHomeController(service *services.HomeService) *HomeController {
	return &HomeController{
		service: service,
	}
}

func (h *HomeController) GetHome(c *gin.Context) {

    status := h.service.GetHome()

    c.JSON(200, status)
}
