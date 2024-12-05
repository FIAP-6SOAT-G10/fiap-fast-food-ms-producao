FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# COPY .env /app/.env

RUN go build .

EXPOSE 8080

ENV ENV_FILE_LOCATION=/app/.env

CMD ["./fiap-fast-food-ms-producao"]
