# GOTIFY - Golang notification application for digIT internal services

Sends notifications trough calls to rest interface.

Currently supports the following notification types:
* Mail

## Usage
How to use the running application

All request must inclue a header with the preshared key.

`Authorization`: `pre-shared: your...key`

### Mail
POST `/mail`

Json Request:
```json5
{
    "to": "....",
    "from": "....",
    "subject": "....",
    "body": "....",
    "attachments": [
        {
            "name": "....",        // File name
            "data": "....",        // Base64 encoded file
            "content_type": "...." // MIME type
        }
    ]
}
```

## Setup
Steps to run the application.
this include configuration and key files at the moment

### Config
The application can be configured through a config file or environment variables. Environment variables take precedence.

#### config.toml
config.toml can reside in your working directory, `/etc/gotify/` or `$HOME/.gotify/`

```
port = "8080"
pre-shared-key = "......"
debug-mode = false
mock-mode = false
max-mail-size = 20000000

[google-mail]
    keyfile = "gapps.json"
    admin-mail = "admin@example.ex"
```
See [Environment Variables](#environment-variables) for config explanation

#### Environment Variables
* `GOTIFY_PORT`: Port for the web service, defaults to `8080` (string)
* `GOTIFY_PRE-SHARED-KEY`*: Random string used by other apps to authenticate
* `GOTIFY_DEBUG-MODE`: Bool indicating debug mode defaults to `false`
* `GOTIFY_MAX-MAIL-SIZE`: The maximum size of an email in bytes, defaults to `20000000` (int)
* `GOTIFY_GOOGLE-MAIL.KEYFILE`: the file described in [Google config file](#google-config-file) defaults to `gapps.json`
* `GOTIFY_GOOGLE-MAIL.ADMIN-MAIL`*: The google administrator email.

### Google config file
This file (gapps.json by default config) should be placed in the working directory

(digIT can find this file on their wiki)


Go to [Google developer console](https://console.developers.google.com) to retrieve this file

* go to credentials
* create new service account för this app
* use the downloaded file


You must also allow mail api calls:

* go to security > advanced settings > Manage API client access
* use the `client_id` from the credentials file previously retrieved
* use api scope `https://www.googleapis.com/auth/gmail.send`

## Development
You can either set this project up manually or with a simple docker compose setup. The manual setup is recommended if you'll be doing extensive development.

Please referer to the software design document before starting development: `DESIGN.md`

See issues for suggested features.
### Manual
Make sure you have golang installed and you `$GOPATH` setup.
1. Follow the steps in [Setup](#setup) and enable debug mode.
2. Grab all dependencies by standing in the project root and run `go get -d ./...`
3. You find the main file in `cmd/main.go`
4. Go to http://localhost:8080

Use gin for hot reloading.
1. Grab it with `go get github.com/codegangsta/gin`
2. Run gotify with `gin -d cmd -a 8080 run main.go`
3. Go to http://localhost:3000

### Docker Compose
1. Get a [Google key file](#google-config-file).
2. Run `docker-compose up --build`
3. Go to http://localhost:8080

You can install additional dependencies without restarting the container by running `docker exec gotify_web_1 go get ...`, gotify_web_1 is the name of the container and ... is the dependency.

### As mock
1. Set the `pre-shared-key` config/environment variable.
2. Set the `mock-mode` config/environment variable to true
3. Enjoy

Example docker-compose entry for mock service:
```
services:
    ...
    gotify:
        image: cthit/gotify:latest
        environment:
          GOTIFY_PRE-SHARED-KEY: "123abc"
          GOTIFY_MOCK-MODE: "true"

```

Other services would then be able to reach this service on `http://gotify:8080/...` with `123abc` as the preshared key