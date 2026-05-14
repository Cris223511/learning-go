# Mismo patrón multi-stage que polizas-svc: compilar en golang, correr en alpine.

FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /clientes-svc ./13-microservicios/clientes-svc/

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /clientes-svc /clientes-svc
EXPOSE 8082
CMD ["/clientes-svc"]
