FROM golang:1.17-alpine

WORKDIR /go/src
ADD . /go/src
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on &&\
        go env -w GOPROXY=https://goproxy.cn,direct &&\
        GO111MODULE=on GOPROXY=https://goproxy.cn,direct
RUN go get -d -v ./... \
    && go install -v ./...
RUN go build -o server ./cmd/main.go

EXPOSE 8080
CMD ["./server"]