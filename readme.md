# Desafio em Go

Este projeto é parte de um desafio para criar um sistema de listagem de pedidos usando Go, com endpoints REST, serviço gRPC e consulta GraphQL.

## Passos para Execução

1. Clone este repositório.
2. Execute `docker-compose up` para iniciar os serviços.
3. Acesse os seguintes endpoints:

   - Serviço REST: [http://localhost:8080/order](http://localhost:8080/order)
   - Serviço GraphQL: [http://localhost:8081/graphql](http://localhost:8081/graphql)
   - Serviço gRPC: Porta 50051

## Estrutura do Projeto

- `api/`: Contém os códigos fonte dos serviços REST, gRPC e GraphQL.
- `docker-compose.yaml`: Arquivo de configuração do Docker Compose.
