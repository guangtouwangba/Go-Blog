build:
	go mod tidy
	go get ./...
	go build -o goblog ./cmd/main.go

test:
	go mod tidy
	go get ./...
	go test ./...

dev:
	-bash ./start.sh
	export GO111MODULE=on
	export GOPROXY=https://goproxy.io
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod tidy
	go get ./...
	go run ./cmd/main.go

image:
	docker build -f ./deploy/Dockerfile -t goblog .

push:
	docker build -f ./deploy/Dockerfile -t 819110182/goblog:1.0 .
	docker push 819110182/goblog:1.0

