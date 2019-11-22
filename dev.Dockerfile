# Dockerfile for gotify development
FROM golang:alpine
MAINTAINER digIT <digit@chalmers.it>

# Install git
RUN apk update
RUN apk upgrade
RUN apk add --update git

RUN go get github.com/codegangsta/gin

# Add standard certificates
RUN apk add ca-certificates && rm -rf /var/cache/apk/*

# create dir
RUN mkdir /app
WORKDIR /app

CMD go mod download && gin -d cmd/gotify -a 8080 run main.go
