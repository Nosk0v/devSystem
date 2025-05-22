FROM --platform=linux/amd64 golang:1.22-bullseye AS builder

WORKDIR /app
COPY . .
COPY ./config/config.json /app/config/config.json

RUN go mod download && go mod tidy
RUN go build -o main ./cmd/main.go

FROM debian:bullseye

WORKDIR /app


RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .
COPY .env .env
COPY --from=builder /app/config/config.json ./config/config.json
COPY --from=builder /app/db/migrations ./db/migrations

EXPOSE 25502
CMD ["./main"]
