package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Executando API")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":8000", r))

}
