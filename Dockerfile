FROM golang:1.17-alpine

WORKDIR /app

COPY main.go .

ENV GOMAXPROCS 12
# JOBS = number of threads is created
ENV JOBS 12

RUN go build -o ./run main.go

CMD ["./run"]