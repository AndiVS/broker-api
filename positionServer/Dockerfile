# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN apk add build-base
WORKDIR /

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go /
COPY /internal /internal
COPY /positionProtocol /positionProtocol

RUN go build -tags musl ./..

CMD [ "/positionServer" ]