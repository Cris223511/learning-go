# Multi-stage build: el stage "builder" compila el binario y el stage final
# solo copia el ejecutable. Resultado: imagen liviana sin el compilador de Go.

# Stage 1: compilar
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Copiamos primero el go.mod del repo raíz porque este servicio es parte del módulo.
# El build context de docker-compose apunta a la raíz del repo.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO_ENABLED=0 produce un binario estático que funciona en cualquier Linux,
# incluso en imágenes mínimas como alpine o scratch.
RUN CGO_ENABLED=0 go build -o /polizas-svc ./13-microservicios/polizas-svc/

# Stage 2: imagen final liviana
FROM alpine:3.20
# ca-certificates permite hacer HTTPS desde el contenedor.
RUN apk add --no-cache ca-certificates

COPY --from=builder /polizas-svc /polizas-svc

EXPOSE 8081
CMD ["/polizas-svc"]
