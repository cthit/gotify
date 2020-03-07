# Dockerfile for gotify development
FROM golang:alpine as dev
MAINTAINER digIT <digit@chalmers.it>

# Install git
RUN apk update
RUN apk upgrade
RUN apk add --update git

RUN go get -u github.com/cespare/reflex

# Add standard certificates
RUN apk add ca-certificates && rm -rf /var/cache/apk/*

# create dir
RUN mkdir /app
WORKDIR /app

RUN which reflex

CMD reflex -r '\.go$' -s -- go run ./cmd/gotify

# Dockerfile for protobuf generation
FROM znly/protoc:0.4.0 as dev_gen
MAINTAINER digIT <digit@chalmers.it>

RUN apk update
RUN apk upgrade
RUN apk add --update git

# Add standard certificates
RUN apk add ca-certificates && rm -rf /var/cache/apk/*

# Add proto imports
RUN mkdir -p /src/github.com/grpc-ecosystem
WORKDIR /src/github.com/grpc-ecosystem
RUN git clone https://github.com/grpc-ecosystem/grpc-gateway.git

COPY --from=dev /go/bin/reflex /bin/reflex

# create dir
RUN mkdir /app
WORKDIR /app


ENTRYPOINT ["/bin/sh"]

