# Dockerfile for protobuf generation
FROM znly/protoc:0.4.0 AS protocGenerator
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

# create dir
RUN mkdir /app
WORKDIR /app

ENTRYPOINT ["/bin/sh"]

FROM protocGenerator AS protocGen

COPY . /app

RUN ./scripts/protoc-gen.sh

# Dockerfile for gotify production
FROM golang:alpine AS buildStage
MAINTAINER digIT <digit@chalmers.it>

# Install git
RUN apk update
RUN apk upgrade
RUN apk add --update git

# Copy sources
RUN mkdir /app
COPY --from=protocGen /app /app
WORKDIR /app/cmd/gotify

# Grab dependencies
RUN go mod download

# build binary
RUN go build

##########################
#    PRODUCTION STAGE    #
##########################
FROM alpine
MAINTAINER digIT <digit@chalmers.it>

# Add standard certificates
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Set user
RUN addgroup -S app
RUN adduser -S -G app -s /bin/bash app
USER app:app

# Copy binary
COPY --from=buildStage /app/cmd/gotify/gotify /app/gotify

# Set good defaults
WORKDIR /app
ENTRYPOINT /app/gotify
