# Etapa 1: Compilar el binario
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compilar con el nombre del servicio
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deleteProduct-cart .

# Etapa 2: Imagen final liviana
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/deleteProduct-cart .

# Documentar el puerto expuesto
EXPOSE 3038

CMD ["./deleteProduct-cart"]
