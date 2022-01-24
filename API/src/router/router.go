package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

//Generate vai retornar um router com as routas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configure(r)
}
