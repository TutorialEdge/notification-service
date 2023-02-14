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
```

## Flow

1. Create a List
2. Create a subscriber
3. Create a notification
4. Send a notification with an email

## TODO

* The concept of a list needs to be fleshed out better. Currently subscribers aren't associated
with a list - subscribers should belong to 1 or more different lists.
* Notifications can then be sent to entire lists which could fetch all subscribers under a list
and send all emails.
* Notifications can also be sent to individual subscribers so we'll need that distinction.