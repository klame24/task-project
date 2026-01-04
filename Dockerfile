FROM golang:1.24.11-bookworm

WORKDIR /app

COPY . .

RUN go mod download

CMD ["make", "service-run"]