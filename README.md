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
```
{
    "to": "....",
    "from": "....",
    "subject": "....",
    "body": "...."
}
```

## Setup
Steps to run the application.
this inklude cofiguration and key files at the moment

### Config
The application can be configured through a config file or environment variables. Environment variables take precedence.

#### config.toml
config.toml can reside in your working directory, `/etc/gotify/` or `$HOME/.gotify/`

```
port = "8080"
pre-shared-key = "......"
debug-mode = false

[google-mail]
    keyfile = "gapps.json"
    admin-mail = "admin@example.ex"
```
See [Environment Variables](#environment-variables) for config explanation

#### Environment Variables
* `GOTIFY_PORT`: Port for the web service, defaults to `8080` (string)
* `GOTIFY_PRE-SHARED-KEY`*: Random string used by other apps to authenticate
* `GOTIFY_DEBUG-MODE`: Bool indicating debug mode defaults to `false`
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
* úse api scope `https://www.googleapis.com/auth/gmail.send`

## Development
1. Follow the steps in [Setup](#setup) and enable debug mode.
2. Grab all dependencies by standing in the project root and run `go get -d ./...`
3. Run application, not hot-reload available