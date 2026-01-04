include .env
export

service-run:
	@go run cmd/main.go

migrate-up:
	@migrate -path migrations -database "postgresql://postgres:root@localhost:5432/taskDB?sslmode=disable" up