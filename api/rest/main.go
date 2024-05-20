package main

import (
	"fmt"
	"log"
	"net/http"
)

func listOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Lógica para listar pedidos aqui
	fmt.Fprintf(w, "Listagem de Pedidos")
}

func main() {
	http.HandleFunc("/order", listOrdersHandler)
	fmt.Println("Serviço REST executando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
