FROM golang:latest-alpine

WORKDIR /app

COPY main.go .

ENV GOMAXPROCS 12

RUN go build -o ./run main.go

CMD ["./run"]