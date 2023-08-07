FROM golang:1.20.7

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "main.go"]
