package controllers

import (
	"api/src/baseDedados"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := basededados.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)

	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID Inserido: %d", usuarioID)))
}

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}

func GetUsario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario"))
}

func PutUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuario"))
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Apagando Usuario"))
}
