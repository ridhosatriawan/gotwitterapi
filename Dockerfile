FROM golang:1.17-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN go mod tidy
RUN go build -o main

CMD ["/app/main"]
