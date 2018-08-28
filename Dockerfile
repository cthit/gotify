# Dockerfile for goldapps production
FROM golang:alpine AS buildStage
MAINTAINER digIT <digit@chalmers.it>

# Install git
RUN apk update
RUN apk upgrade
RUN apk add --update git

# Copy sources
RUN mkdir -p $GOPATH/src/github.com/cthit/gotify
COPY . $GOPATH/src/github.com/cthit/gotify
WORKDIR $GOPATH/src/github.com/cthit/gotify/cmd

# Grab dependencies
RUN go get -d -v ./...

# build binary
RUN go install -v
RUN mkdir /app && mv $GOPATH/bin/cmd /app/gotify

##########################
#    PRODUCTION STAGE    #
##########################
FROM alpine
MAINTAINER digIT <digit@chalmers.it>

#Update
RUN apk update
RUN apk upgrade

# Set user
RUN addgroup -S app
RUN adduser -S -G app -s /bin/bash app
USER app:app

# Copy binary
COPY --from=buildStage /app/gotify /app/gotify

# Set good defaults
WORKDIR /app
ENTRYPOINT /app/gotify