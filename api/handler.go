package api

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func RotaPrincipal(c *gin.Context) {
	type Resposta struct {
		verify  bool
		noMatch []string
	}
	var senha Senha

	if err := c.ShouldBindJSON(&senha); err != nil {
		panic(err)
	}

	verify, rules, erro := ValidaSenha(senha)
	if erro != nil {
		c.JSON(400, false)
	}

	var res Resposta
	res.verify = verify
	res.noMatch = rules

	c.JSON(http.StatusOK, res)

}

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/", RotaPrincipal)
}

func ValidaSenha(senha Senha) (bool, []string, error) {

	var ruleCausas []string
	for _, rule := range senha.Rules {
		switch rule.Name {
		case "minSize":
			if len(senha.Password) < rule.Value {
				ruleCausas = append(ruleCausas, rule.Name)
			}
		case "minUppercase":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.*[A-Z]){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				ruleCausas = append(ruleCausas, rule.Name)
			}
		case "minLowercase":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.*[a-z]){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				ruleCausas = append(ruleCausas, rule.Name)
			}
		case "minDigit":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.*[0-9]){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				ruleCausas = append(ruleCausas, rule.Name)
			}
		case "minSpecialChars":
			valorString := strconv.Itoa(rule.Value)
			match := regexp.MustCompile(`(.*[!@#$%^&*()-+{}/]){` + valorString + `}`)
			if !match.MatchString(senha.Password) {
				ruleCausas = append(ruleCausas, rule.Name)
			}
		case "noRepeted":
			splitPalavra := strings.Split(senha.Password, "")
			for indice, letra := range splitPalavra {
				if indice != len(splitPalavra)-1 {
					if letra == splitPalavra[indice+1] {
						ruleCausas = append(ruleCausas, rule.Name)
						break
					}
				}
			}
		}

	}

	if len(ruleCausas) > 0 {
		return false, ruleCausas, nil
	}
	return true, nil, nil
}
