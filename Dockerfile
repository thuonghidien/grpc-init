FROM golang:alpine

RUN apk update \
  && apk add git

COPY . /go/src/app

# Don't do this in production! Use vendoring instead.
RUN go get -v app/customer/server

RUN go install app/customer/server

ENTRYPOINT ["/go/bin/server"]
