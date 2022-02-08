build:
	go mod tidy
	go get ./...
	go build -o goblog ./cmd/main.go

test:
	go mod tidy
	go get ./...
	go test ./...

dev:
	go mod tidy
	go get ./...
	go run ./cmd/main.go

image:
	docker build -f ./deploy/Dockerfile -t goblog .

