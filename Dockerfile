FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" . 

FROM scratch

COPY --from=builder /build/rh /usr/local/bin/rh

ENTRYPOINT ["/usr/local/bin/rh"]

EXPOSE 8080 8081


