FROM golang:alpine

RUN mkdir /server

WORKDIR /server

COPY . .

RUN go build -o /go/bin/main

EXPOSE 8001

ENTRYPOINT ["/go/bin/main"]