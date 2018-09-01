# Design rationale for gotify

## Project structure
The root package/directory only contains the domain types and logic. This will probably be limited to notification service interfaces and notification types. This package may not depend on any other package in this repo.

cmd contains the main package and takes care of meta sutch as configuration and binding all other packages together, This package may depend on any other package

All other packages represent a functionality or dependency and may only depend on the root package as well as external packages.

See [Project structure in go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091) for further explanation.

## API Structure
One api endpoint for every notification type.
See readme for existing api endpoints

A post request to an endpoint with the matching jason notification type should send a notification and on success return the sent notification in json.

## Dependency injection

The web package has some weird dependency injection.

For now, take a look at it until you understand it. It looks like it does for a good reason.