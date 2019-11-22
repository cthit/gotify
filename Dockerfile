# Dockerfile for gotify production
FROM golang:alpine AS buildStage
MAINTAINER digIT <digit@chalmers.it>

# Install git
RUN apk update
RUN apk upgrade
RUN apk add --update git

# Copy sources
RUN mkdir /app
COPY . /app
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
