package api

import (
	"net/http"

	"github.com/desafio-senhaForte/domain"
	"github.com/gin-gonic/gin"
)

func RotaPrincipal(c *gin.Context) {

	var senha domain.Senha

	if err := c.ShouldBindJSON(&senha); err != nil {
		panic(err)
	}

	verify, rules, erro := domain.ValidaSenha(senha)
	if erro != nil {
		c.JSON(500, "Dados inválidos")
	}

	c.JSON(http.StatusOK, gin.H{
		"verify":   verify,
		"noMatch:": rules,
	})

}

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/verify", RotaPrincipal)
}
