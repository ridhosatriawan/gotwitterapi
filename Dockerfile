FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY vendor ./vendor

RUN go build -o /gotwitterapi

EXPOSE 8090

CMD ["/gotwitterapi"]
