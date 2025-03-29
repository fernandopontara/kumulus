# Etapa de build
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o server main.go

# Imagem final
FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/server /app/server
COPY .env /app/.env

EXPOSE 8080

ENTRYPOINT ["/app/server"]
