FROM golang:alpine

RUN apk update \
  && apk add git

COPY . /go/src/app

RUN go install app/server

EXPOSE 50051

ENTRYPOINT ["/go/bin/server"]