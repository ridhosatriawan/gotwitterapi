FROM golang:1.17-alpine

RUN mkdir /app
WORKDIR /app

RUN export GO111MODULE=on
RUN go get github.com/ridhosatriawan/gotwitterapi
RUN cd /app && git clone https://github.com/ridhosatriawan/gotwitterapi.git

RUN cd /app/gotwitterapi && go build

EXPOSE 8090

CMD ["/app/gotwitterapi"]
