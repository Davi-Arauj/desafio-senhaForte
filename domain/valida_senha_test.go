package domain

import (
	"testing"
)

// TestRegrasValidaSenha - teste de cada regra da função ValidaSenha
func TestRegrasValidaSenha(t *testing.T) {

	t.Run("teste Falha minSize", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "teste"

		rule.Name = "minSize"
		rule.Value = 6

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify {
			t.Errorf("Esperava um FALSE e veio um %v", verify)
		}
		if len(rules) == 0 {
			t.Errorf("Esperava uma regra 'minSize' e veio %v", len(rules))
		}
	})

	t.Run("teste OK minSize", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "testes"

		rule.Name = "minSize"
		rule.Value = 6

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if !verify {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", len(rules))
		}
	})

	t.Run("teste Falha minUppercase", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "testE"

		rule.Name = "minUppercase"
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify {
			t.Errorf("Esperava um FALSE e veio um %v", verify)
		}
		if len(rules) == 0 {
			t.Errorf("Esperava uma regra 'minUppercase' e veio %v", len(rules))
		}
	})

	t.Run("teste OK minUppercase", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "tEstEs"

		rule.Name = "minUppercase"
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if !verify {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", len(rules))
		}
	})

	t.Run("teste Falha minLowercase", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "TEsTE"

		rule.Name = "minLowercase"
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify {
			t.Errorf("Esperava um FALSE e veio um %v", verify)
		}
		if len(rules) == 0 {
			t.Errorf("Esperava uma regra 'minLowercase' e veio %v", len(rules))
		}
	})

	t.Run("teste OK minLowercase", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "tEsTE"

		rule.Name = "minLowercase"
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if !verify {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", len(rules))
		}
	})

	t.Run("teste Falha minDigit", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "TEsTE1"

		rule.Name = "minDigit"
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify {
			t.Errorf("Esperava um FALSE e veio um %v", verify)
		}
		if len(rules) == 0 {
			t.Errorf("Esperava uma regra 'minDigit' e veio %v", len(rules))
		}
	})

	t.Run("teste OK minDigit", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "tE1sTE2"

		rule.Name = "minDigit"
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if !verify {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", len(rules))
		}
	})

	t.Run("teste Falha minSpecialChars", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "TEsTE1@"

		rule.Name = "minSpecialChars"
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify {
			t.Errorf("Esperava um FALSE e veio um %v", verify)
		}
		if len(rules) == 0 {
			t.Errorf("Esperava uma regra 'minSpecialChars' e veio %v", len(rules))
		}
	})

	t.Run("teste OK minSpecialChars", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "tE1#sTE!"

		rule.Name = "minSpecialChars"
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if !verify {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", len(rules))
		}
	})

	t.Run("teste Falha noRepeted Minuscula", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "TEssTE1@"

		rule.Name = "noRepeted"

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify {
			t.Errorf("Esperava um FALSE e veio um %v", verify)
		}
		if len(rules) == 0 {
			t.Errorf("Esperava uma regra 'noRepeted' e veio %v", len(rules))
		}
	})

	t.Run("teste Falha noRepeted letra Maiuscula", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "TETTE1@"

		rule.Name = "noRepeted"

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify {
			t.Errorf("Esperava um FALSE e veio um %v", verify)
		}
		if len(rules) == 0 {
			t.Errorf("Esperava uma regra 'noRepeted' e veio %v", len(rules))
		}
	})

	t.Run("teste OK noRepeted", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "tE1#sTE!"

		rule.Name = "noRepeted"

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if !verify {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", len(rules))
		}
	})

}

// TestGeralValidaSenha - teste completo da função
func TestGeralValidaSenha(t *testing.T) {
	t.Run("teste Completo Fail", func(t *testing.T) {
		var senha Senha
		var rule Rule
		senha.Password = "tEs21#T!"

		rules := []string{"noRepeted", "minSize", "minUppercase", "minLowercase", "minDigit", "minSpecialChars"}
		for i := 0; i < len(rules); i++ {
			rule.Name = rules[i]
			rule.Value = i
			senha.Rules = append(senha.Rules, rule)
		}

		verify, rules, _ := ValidaSenha(senha)

		if verify {
			t.Errorf("Esperava um FALSE e veio um %v", verify)
		}
		if len(rules) == 0 {
			t.Errorf("Esperava [minLowercase minDigit minSpecialChars] regra e veio %v", rules)
		}
	})

	t.Run("teste Completo OK", func(t *testing.T) {
		var senha Senha
		var rule Rule
		senha.Password = "tEsr21#T!@3%&85"
		// CARACTERES ESPECIAIS RESERVADOS -$,*,),(,+,?,/,\,^,]
		rules := []string{"noRepeted", "minSize", "minUppercase", "minLowercase", "minDigit", "minSpecialChars"}
		for i := 0; i < len(rules); i++ {
			rule.Name = rules[i]
			rule.Value = i
			senha.Rules = append(senha.Rules, rule)
		}

		verify, rules, _ := ValidaSenha(senha)

		if !verify {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", rules)
		}
	})
}
