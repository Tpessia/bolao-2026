init:
	go install github.com/swaggo/swag/cmd/swag@latest
	go mod download

generate:
	go generate ./...

swagger:
	swag init -g ./internal/server/handler.go

build: swagger
	go build -o bolao-2026 cmd/api/main.go

run: swagger
	go run cmd/api/main.go

test:
	go test ./... -v

docker_build:
	docker build -f Dockerfile -t bolao-2026 .

docker_run: docker_build
	docker run -it --rm -p 8080:8080 bolao-2026

docker_exec:
	docker exec -it bolao-2026 /bin/bash