package domain

import (
	"regexp"
	"strconv"
	"strings"
)

func ValidaSenha(senha Senha) (bool, []string, error) {

	ruleCausas := []string{}
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
			splitSenha := strings.Split(senha.Password, "")
			for _, letra := range splitSenha {
				match := regexp.MustCompile(letra + `{2}`)
				if match.MatchString(senha.Password) {
					ruleCausas = append(ruleCausas, rule.Name)
					break
				}
			}
		}

	}

	if len(ruleCausas) > 0 {
		return false, ruleCausas, nil
	}
	return true, ruleCausas, nil
}
