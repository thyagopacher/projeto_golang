package controllers

import (
	"strconv"
	"projeto_go/internal/models"
	"github.com/gin-gonic/gin"
	"projeto_go/internal/services"
)

type UsuarioController struct {
	service *services.UsuarioService
}

func NewUsuarioController(service *services.UsuarioService) *UsuarioController {
	return &UsuarioController{
		service: service,
	}
}

/**
* GET /usuarios
*/
func (uc *UsuarioController) GetUsuarios(c *gin.Context) {
    usuarios, err := uc.service.GetUsuarios()
    if err != nil {
        c.JSON(500, gin.H{
            "error": "Erro ao buscar usuários",
            "details": err.Error(),
        })
        return
    }
    
    c.JSON(200, usuarios)
}

/**
* GET /usuarios/:id
*/
func (uc *UsuarioController) GetUsuarioByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	usuario, err := uc.service.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, usuario)
}

/**
* POST /usuarios
*/
func (uc *UsuarioController) CreateUsuario(c *gin.Context) {
	var input models.Usuario

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	usuario, err := uc.service.CreateUsuario(input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, usuario)
}

/**
* PUT /usuarios/:id
*/
func (uc *UsuarioController) UpdateUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Usuario
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	usuario, err := uc.service.UpdateUsuario(id, input)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, usuario)
}

/**
* DELETE /usuarios/:id
*/
func (uc *UsuarioController) DeleteUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	status, err := uc.service.DeleteUsuario(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Usuário removido", "success": status})
}