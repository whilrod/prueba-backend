# Etapa 1: build
FROM golang:1.21 AS builder
WORKDIR /app

# Copiar go.mod y go.sum primero para aprovechar cache
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar binario
RUN go build -o main .

# Etapa 2: runtime (más liviana)
FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/main .

# Puerto donde corre el backend
EXPOSE 8080

CMD ["./main"]