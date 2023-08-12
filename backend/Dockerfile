FROM golang:1.20.7

WORKDIR /app

COPY . .

RUN go install github.com/cosmtrek/air@v1.44.0

RUN go mod download

CMD ["air"]
