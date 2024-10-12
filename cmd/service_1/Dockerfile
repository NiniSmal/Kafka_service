FROM golang:1.21 AS builder
WORKDIR /app
COPY ../../go.mod go.sum ./
RUN go mod download
COPY ../.. .
RUN CGO_ENABLED=0 go build -o /app/service_1 .cmd/service_1/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/service_1 /app/service_1
EXPOSE 8083
ENTRYPOINT ["/app/service_1"]