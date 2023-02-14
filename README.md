Notification Service
====================

This service is designed to do 2 things: 

* Act as a fully-fledged example of how my REST API course principles can be applied to larger, more monolithic-styled applications
* Hopefully act as a replacement for TutorialEdge's current email campaign system which costs almost $1,000 a year as it stands.

## Running the Service

To run the service, you will need `docker` up and running on your machine as well as the `Taskfile` tool:

```bash
$ export MAILGUN_DOMAIN=yourdomain.com
$ export MAILGUN_API_KEY=mailgun-api-key

$ task run
```

This will start up a Postgres Docker container as well as the application itself.

## Testing the Service

```bash
# Running the unit tests
$ task test

# Running the acceptance tests - note: this requires the app to be running locally
$ task acceptance-test