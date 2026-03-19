package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"projeto_go/models"
)

// Mock em memória (simples pra exemplo)
var usuarios = []models.Usuario{
	{ID: 1, Nome: "João", Email: "joao@email.com"},
	{ID: 2, Nome: "Maria", Email: "maria@email.com"},
}

var nextID = 3

// GET /usuarios
func GetUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, usuarios)
}

// GET /usuarios/:id
func GetUsuarioByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for _, u := range usuarios {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
}

// POST /usuarios
func CreateUsuario(c *gin.Context) {
	var novo models.Usuario

	if err := c.ShouldBindJSON(&novo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	novo.ID = nextID
	nextID++

	usuarios = append(usuarios, novo)

	c.JSON(http.StatusCreated, novo)
}

// PUT /usuarios/:id
func UpdateUsuario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var atualizado models.Usuario

	if err := c.ShouldBindJSON(&atualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	for i, u := range usuarios {
		if u.ID == id {
			atualizado.ID = id
			usuarios[i] = atualizado
			c.JSON(http.StatusOK, atualizado)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
}

// DELETE /usuarios/:id
func DeleteUsuario(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for i, u := range usuarios {
		if u.ID == id {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Usuário removido"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
}