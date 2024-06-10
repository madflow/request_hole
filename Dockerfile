FROM golang:1-alpine as builder

WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go generate ./...
RUN GOOS=linux GOARCH=amd64 go build -o rh

EXPOSE 8080 8081

CMD ["/app/rh", "--web", "-p", "8000", "--web_port", "8080", "-a", "0.0.0.0", "--web_address", "0.0.0.0", "http"]
