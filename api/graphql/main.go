package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// Schema GraphQL
	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"listOrders": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// Lógica para listar pedidos aqui
						return "Listagem de Pedidos via GraphQL", nil
					},
				},
			},
		}),
	})

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	fmt.Println("Serviço GraphQL executando na porta 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}