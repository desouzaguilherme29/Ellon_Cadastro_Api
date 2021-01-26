package modelos

//Senha representa o formato da requisicao de alteravao de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
