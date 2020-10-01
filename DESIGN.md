# Design rationale for gotify

## Project structure
See [Project structure in go](https://github.com/golang-standards/project-layout) for further explanation.

## API Structure
One api endpoint for every notification type.
See readme for existing api endpoints

A post request to an endpoint with the matching json notification type should send a notification and on
success return the sent notification in json.