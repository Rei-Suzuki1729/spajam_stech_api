FROM golang:1.19.0-alpine3.16

RUN apk add --update &&  apk add git curl

ADD ./app /go/src/app

WORKDIR /go/src/app

ENV GO111MODULE=on


#RUN go mod download

EXPOSE $PORT
