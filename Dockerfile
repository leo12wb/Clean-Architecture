FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

# Comandos de compilação e execução do serviço
CMD ["go", "run", "main.go"]