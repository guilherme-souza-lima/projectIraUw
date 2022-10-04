package response

type ListAll struct {
	ID         int    `json:"id"`
	Canal      string `json:"canal"`
	Arithmetic string `json:"arithmetic"`
	Quantidade string `json:"quantidade"`
	Descricao  string `json:"descricao"`
	Pessoa     string `json:"pessoa"`
	Data       string `json:"data"`
}
