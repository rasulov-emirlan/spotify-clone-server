POSTGRESQL_URL='postgres://aqqhwvfetzhtxs:895ab244936af93ecf0a9827b4ffa88df78f62340ab48bf2405e2f394d09c877@ec2-18-213-133-45.compute-1.amazonaws.com:5432/d2f40c70mg32vo'

dev:
	go build -v ./cmd/apiserver
	./apiserver
	rm apiserver

build:
	go build -v ./cmd/apiserver

generate_swagger:
	swag init -g ./internal/server/server.go

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path ./migrations/db/migrations up 

migrate_down:
	migrate -database ${POSTGRESQL_URL} -path ./migrations/db/migrations down


start_server:

setup:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -database ${POSTGRESQL_URL} -path ./migrations/db/migrations up 
	cd ../
	mkdir database
	cd database
	mkdir audio
	mkdir covers
