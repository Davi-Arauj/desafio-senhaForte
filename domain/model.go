package domain

type Senha struct {
	Password string `json:"password"`
	Rules    []Rule
}

type Rule struct {
	Name  string `json:"rule"`
	Value int    `json:"value"`
}
