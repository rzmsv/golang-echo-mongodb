include .env

build:
	go build -o ${BINARY}.exe  ./cmd/api/main.go

start: 
	go run ./cmd/api/main.go