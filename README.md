# GOTIFY - Golang notification application for digIT internal services

Sends notifications trough calls to rest interface.

Currently supports the following notification types:
* Mail

## Usage
How to use the running application

### Mail
POST `/mail`

Json Request:
```json5
{
    "to": "....",
    "from": "....", // (optional)
    "reply_to": "....", // (optional)
    "subject": "....",
    "body": "...."
}
```

## Setup
Steps to run the application.
this include configuration and key files at the moment

### Config
The application is configured through  environment variables.

#### Environment Variables
* `GOTIFY_WEB_PORT`: Port for the web service, defaults to `8080` (string)
* `GOTIFY_RPC_PORT`: Port for the rpc service, defaults to `8090` (string)
* `GOTIFY_DEBUG_MODE`: Bool indicating debug mode defaults to `false`
* `GOTIFY_GOOGLE_MAIL_KEYFILE`: the file described in [Google config file](#google-config-file) defaults 
to `gapps.json`
* `GOTIFY_MAIL_DEFAULT_FROM`: Default `from` address in the mail, defaults to `admin@chalmers.it`
* `GOTIFY_MAIL_DEFAULT_REPLY_TO`: Default `reply-to` address in the mail, defaults to `no-reply@chalmers.it`
* `GOTIFY_MOCK_MODE`: Enable mock mode, defaults to `false`

### Google config file
This file (gapps.json by default config) should be placed in the working directory

(digIT can find this file on their wiki)


Go to [Google developer console](https://console.developers.google.com) to retrieve this file

* go to credentials
* create a project for this app if you don't already have one
* create new service account för this app
* use the downloaded file


You must also allow mail api calls:

* go to security > advanced settings > Manage API client access
* use the `client_id` from the credentials file previously retrieved
* use api scope `https://www.googleapis.com/auth/gmail.send`

## Development
To start a dockerized development environment with hot-reloading:
```bash
$ make dev
```

To start a non-dockerized development environment:
```bash
$ make run
```

Please referer to the software design document before starting development: `DESIGN.md`

### As mock
1. Set the `mock-mode` config/environment variable to true
2. Enjoy
