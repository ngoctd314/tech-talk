FROM golang:latest-alpine

WORKDIR /app

COPY main.go .

<<<<<<< HEAD
ENV GOMAXPROCS 12
=======
ENV GOMAXPROCS 3
>>>>>>> e533d6b8910d14158c0abb1fbe98bcbde07b9408

RUN go build -o ./run main.go

CMD ["./run"]