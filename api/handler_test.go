package api

import "testing"

func TestValidaSenha(t *testing.T) {

	t.Run("teste Falha minSize", func(t *testing.T) {
		var senha Senha
		var rule Rule

		senha.Password = "teste"

		rule.Name = "minSize"
		rule.Value = 6

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify != false {
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

		if verify != true {
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

		if verify != false {
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

		if verify != true {
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

		if verify != false {
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

		if verify != true {
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

		if verify != false {
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

		if verify != true {
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

		if verify != false {
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

		if verify != true {
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
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify != false {
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
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify != false {
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
		rule.Value = 2

		senha.Rules = append(senha.Rules, rule)

		verify, rules, _ := ValidaSenha(senha)

		if verify != true {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", len(rules))
		}
	})

}

func TestValidaGeral(t *testing.T) {
	t.Run("teste Completo", func(t *testing.T) {
		var senha Senha
		var rule Rule
		var rule2 Rule
		var rule3 Rule
		var rule4 Rule
		var rule5 Rule
		var rule6 Rule

		senha.Password = "tE1#TT!"

		rule.Name = "minSize"
		rule.Value = 6
		rule2.Name = "minUppercase"
		rule2.Value = 2
		rule3.Name = "minLowercase"
		rule3.Value = 2
		rule4.Name = "minDigit"
		rule4.Value = 2
		rule5.Name = "minSpecialChars"
		rule5.Value = 2
		rule6.Name = "noRepeted"
		//rule6.Value = 0

		senha.Rules = append(senha.Rules, rule)
		senha.Rules = append(senha.Rules, rule2)
		senha.Rules = append(senha.Rules, rule3)
		senha.Rules = append(senha.Rules, rule4)
		senha.Rules = append(senha.Rules, rule5)
		senha.Rules = append(senha.Rules, rule6)

		verify, rules, _ := ValidaSenha(senha)

		if verify != true {
			t.Errorf("Esperava um TRUE e veio um %v", verify)
		}
		if len(rules) > 0 {
			t.Errorf("Esperava Nenhuma regra e veio %v", rules)
		}
	})
}
