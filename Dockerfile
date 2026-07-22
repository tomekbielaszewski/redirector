FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor/ ./vendor/
COPY main.go ./

RUN go build -mod=vendor -o redirector .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/redirector /redirector

EXPOSE 8080

ENTRYPOINT ["/redirector"]
