#!/bin/sh
FROM golang:alpine
MAINTAINER MKP Mobile Production <mkpproduction@gmail.com>
RUN apk add --no-cache git
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
COPY ./aws /root/.aws
RUN go build -o main .
CMD ["/app/main"]