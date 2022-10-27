FROM golang:1.17-alpine

WORKDIR /app

ENV GOMAXPROCS 12
ENV JOBS 12

COPY main.go .

RUN go build -o run main.go

CMD ["./run"]