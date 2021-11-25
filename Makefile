dev:
	go build -v ./cmd/apiserver
	./apiserver
	rm apiserver
build:
	go build -v ./cmd/apiserver
