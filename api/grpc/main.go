package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type orderService struct{}

func (*orderService) ListOrders() {
	// Lógica para listar pedidos aqui
	fmt.Println("Listando Pedidos via gRPC")
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	RegisterOrderServiceServer(server, &orderService{})

	fmt.Println("Serviço gRPC executando na porta 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}