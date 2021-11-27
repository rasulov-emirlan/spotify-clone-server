dev:
	go build -v ./cmd/apiserver
	./apiserver
	rm apiserver
build:
	go build -v ./cmd/apiserver

generate_swagger:
	swag init -g ./internal/server/server.go

install_deps:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest