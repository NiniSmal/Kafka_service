FROM golang:1.21 AS builder
WORKDIR /app
COPY ../../go.mod go.sum ./
RUN go mod download
COPY ../.. .
RUN CGO_ENABLED=0 go build -o /app/service_2 .cmd/service_2/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/service_2 /app/service_2
ENTRYPOINT ["/app/service_2"]