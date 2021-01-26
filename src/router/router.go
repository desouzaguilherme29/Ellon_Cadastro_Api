package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar a nossa rota
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
