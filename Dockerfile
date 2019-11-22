#FROM alpine:latest
FROM golang:1.6

WORKDIR /

RUN mkdir /app

ADD go-204 /app/go-204

EXPOSE 80

ENTRYPOINT ["/app/go-204"]


