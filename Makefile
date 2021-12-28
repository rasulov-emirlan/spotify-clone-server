POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/spotify-clone?sslmode=disable'

dev:
	go build -v ./cmd/apiserver
	./apiserver
	rm apiserver

build:
	go build -v ./cmd/apiserver

generate_swagger:
	swag init -g ./internal/server/server.go

migrate_up:
	migrate -database ${POSTGRESQL_URL} -path ./migrations/ up 

migrate_down:
	migrate -database ${POSTGRESQL_URL} -path ./migrations/ down

migrate_force_fix:
	migrate -path ./migrations/ -database ${POSTGRESQL_URL} force ${VERSION}

start_server:

setup:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	# migrate -database ${POSTGRESQL_URL} -path ./migrations/db/migrations up 
