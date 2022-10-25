FROM golang:1.17-alpine

WORKDIR /app

COPY main.go .

ENV GOMAXPROCS 3

RUN go build -o ./run main.go

CMD ["./run"]