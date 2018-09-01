# Dockerfile for gotify development
FROM golang:alpine
MAINTAINER digIT <digit@chalmers.it>

# Install git
RUN apk update
RUN apk upgrade
RUN apk add --update git

RUN mkdir -p $GOPATH/bin && \
    go get github.com/codegangsta/gin

# Add standard certificates
RUN apk add ca-certificates && rm -rf /var/cache/apk/*

# create dir
RUN mkdir -p /go/src/github.com/cthit/gotify
WORKDIR $GOPATH/src/github.com/cthit/gotify

CMD go get -d -v ./... && gin -d cmd -a 8080 run main.go
