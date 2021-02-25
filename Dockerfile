FROM golang:1.15.6

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/main.go" --command=./main
