FROM golang:alpine

RUN mkdir /test

WORKDIR /test

COPY . .

RUN go build -o /go/bin/main

EXPOSE 8001

ENTRYPOINT ["/go/bin/main"]