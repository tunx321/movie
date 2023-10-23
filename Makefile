build:
	go build -o app cmd/server/main.go

test:
	go test -v ./...


lint:
	golangci-lint run

run:
	docker-compose up --build

export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_TABLE=postgres
export DB_HOST=localhost
export DB_PORT=5432
export DB_DB=postgres
export SSL_MODE=disable
integration-test: 

	docker-compose up -d db
	go test -tags=integration -v ./...
	