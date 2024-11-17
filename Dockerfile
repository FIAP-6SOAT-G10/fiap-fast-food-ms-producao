FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build .

EXPOSE 8080

CMD ["./fiap-fast-food-ms-producao"]
