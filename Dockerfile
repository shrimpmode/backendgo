FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . /app

RUN go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

