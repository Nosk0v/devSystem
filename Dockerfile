
FROM golang:1.22
WORKDIR /app

RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*

COPY . .
COPY ./config/config.json /app/config/config.json

RUN go mod download
RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]
