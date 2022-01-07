FROM golang:alpine

WORKDIR /appointment-scheduling

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./appointment-scheduling