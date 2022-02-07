build:
	go build -o goblog ./...

test:
	go test ./...

dev:
	go run ./cmd/main.go

image:
	docker build -f ./infra/Dockerfile -t goblog .

