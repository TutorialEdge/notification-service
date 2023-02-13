ctxlog - the simple context logger
===========

ctxlog looks to simplify the way that we handle observability through the various layers of our Go applications.

The idea is that, when a request hits the left-most edge of our application, we can append information such as request ids and trace ids to the context. This will then be automagically added to all subsequent log messages for that particular request context. 

## Installation

```shell
$ go get github.com/TutorialEdge/ctxlog
```

## Example Code

A picture is worth a thousand words, a code snippet is likely worth a lot less than that, but it hopefully highlights how this simplifies our lives.

```go
ctx := context.Background()
log := ctxlog.New(
	ctxlog.WithJSONFormat(),
)
ctx = ctxlog.WithFields(ctx, ctxlog.Fields{
	"trace_id": "my-trace-id",
})

log.Info(ctx, "hello world")
// {"level":"info","msg":"hello world","time":"2022-07-23T12:01:43+01:00","trace_id":"my-trace-id"}
log.Error(ctx, "oh my goodness")
// {"level":"error","msg":"oh my goodness","time":"2022-07-23T12:01:43+01:00","trace_id":"my-trace-id"}
```

## To Do:

* Tests that capture the system log and perform appropriate validations.
* Benchmarking & Profiling.

## Contributing

This is a very early rendition of this package. I'll likely expand this out as I start to adopt it in TutorialEdge's services. 

If you'd like to accelerate the development, feel free to submit a Pull Request.

## Attribution

It should be noted this was heavily inspired by the usability of a library developed within CircleCI for our Go applications. I've distilled some of the core concepts down into this far smaller, focused library.