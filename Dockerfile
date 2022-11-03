FROM golang:1.17-alpine

ENV CGO_ENABLED 0
RUN mksdir /app
WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
COPY vendor ./vendor

RUN go build -o /gotwitterapi

EXPOSE 8090

CMD ["/gotwitterapi"]
