FROM golang:1-alpine as builder

WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go generate ./...
RUN GOOS=linux GOARCH=amd64 go build -o rh

EXPOSE 8080 8081

ENTRYPOINT ["/app/rh"]
