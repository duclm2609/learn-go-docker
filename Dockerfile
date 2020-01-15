FROM golang:alpine AS builder
WORKDIR /app
COPY ./ /app
RUN go mod download
RUN go build -o server command/server.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
ENTRYPOINT ./server