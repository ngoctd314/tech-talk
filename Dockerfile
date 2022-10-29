FROM golang:1.17-alpine

WORKDIR /app

ENV GOMAXPROCS 2
# 1 job - 1s
ENV JOBS 2
# 0: sequential, 1: concurrent
ENV VER 1

COPY main.go .

RUN go build -o run main.go

CMD ["./run"]