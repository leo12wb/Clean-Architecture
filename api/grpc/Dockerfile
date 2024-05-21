# Use the official Golang image from Docker Hub
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy all project files to the container
COPY . .

# Command to run the service
CMD ["go", "run", "main.go"]