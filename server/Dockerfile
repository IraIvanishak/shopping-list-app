FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

EXPOSE 8080

CMD goose -dir=/app/migrations postgres "postgres:password@db/shopping-list" up && ./main