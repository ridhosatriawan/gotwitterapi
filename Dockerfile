FROM golang:1.17-alpine

RUN apk update && apk --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
COPY vendor ./vendor

RUN go mod tidy
RUN go build -o /gotwitterapi

EXPOSE 8090

CMD ["/gotwitterapi"]
