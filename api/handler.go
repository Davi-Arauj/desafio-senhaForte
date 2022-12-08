package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RotaPrincipal(c *gin.Context) {

	var senha Senha

	if err := c.ShouldBindJSON(&senha); err != nil {
		panic(err)
	}

	erro := ValidaSenha(senha)
	if erro != nil {
		c.JSON(400, false)
	}

}

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/", RotaPrincipal)
}

func ValidaSenha(senha Senha) error {

	var erroCausas []string
	for _, rule := range senha.Rules {
		if rule.Name == "minSize" {
			//então o nome tem que ter no minimo o valor passado de caracters
			if len(senha.Password) < rule.Value {
				erroCausas = append(erroCausas, rule.Name)
			}
		}
	}

	if len(erroCausas) > 0 {
		for _, erro := range erroCausas {
			return fmt.Errorf("%v", erro)
		}
	}
	return nil
}
