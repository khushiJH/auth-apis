FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /auth-api

EXPOSE 8000

CMD ["/auth-api"]