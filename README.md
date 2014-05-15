gotelemetry
===========

[Go](http://golang.org) bindings for the [Telemetry Visualization API](http://telemetryapp.com).

## Installing

    go get github.com/telemetryapp/gotelemetry

    import "github.com/telemetryapp/gotelemetry"

## Usage

Gotelemetry provides Go structures for all the [flow variants](https://telemetryapp.com/user/documentation/data) support by the Telemetry API. It allows you to create flows and submit them to the API servers:

```go
func ExampleFlow() {
    c := NewCredentials("test-api-token")

    g := Gauge{
        Value: 123,
    }

    f := NewFlow("test_gauge", &g)

    err := f.Publish(c)

    if err != nil {
        panic("Something went wrong…", err.Error)
    }
}
```

You can also submit flows in batches to limit network usage and latency:

```go
func ExampleBatch() {
    // Note that the test API token cannot be used for
    // batch submission.
    c, err := NewCredentials("myapitoken")

    So(err, ShouldBeNil)

    g := Gauge{Value: 10}
    fg := NewFlow("gauge_1", &g)

    v := Value{Value: 101.23}
    fv := NewFlow("gauge_2", &v)

    b := Batch{}

    b.SetFlow(fg)
    b.SetFlow(fv)

    err = b.Publish(c)

    if err != nil {
        panic("Something went wrong…", err.Error())
    }
}
```

Note that neither flow nor batch values are thread-safe. If you require thread safety, you must mediate access to your values through some kind of synchronization mechanism, like a mutex.

## Support and bug reports

Don't hesitate to open issues if you find a bug or require support.