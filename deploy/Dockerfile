FROM golang:1.17-alpine

WORKDIR /go/src
ADD . /go/src
RUN go get -d -v ./... \
    && go install -v ./...
RUN go build -o server ./cmd/main.go

EXPOSE 8080
CMD ["./server"]