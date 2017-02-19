[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/zorkian/go-datadog-api)
[![Build
status](https://travis-ci.org/zorkian/go-datadog-api.svg)](https://travis-ci.org/zorkian/go-datadog-api)

# Datadog API in Go

A Go wrapper for the Datadog API. Use this library if you need to interact with the Datadog system. You can post metrics
with it if you want, but this library is probably mostly used for automating dashboards/alerting and retrieving data
(events, etc).

The master branch contains the v1 version of this library. If you are going to use this library for the first time you
should switch to the v2 branch. If you are a v1 user, please upgrade to the v2 branch. This version is currently soaking,
and should you find any issues now is the time to raise an issue or PR.

The source API documentation is here: <http://docs.datadoghq.com/api/>

## Installation
### v1 (deprecated)
 To use the default branch, include it in your code like so:
```go
    import "github.com/zorkian/go-datadog-api"
```
To be able to use v1 after the default branch switched to version 2; import using [gopkg.in](http://labix.org/gopkg.in):
```go
    import "gopkg.in/zorkian/go-datadog-api.v1"
```

Using `go get`:
```bash
go get gopkg.in/zorkian/go-datadog-api.v1
```
### v2
Import:
```go
    import "gopkg.in/zorkian/go-datadog-api.v2"
```

Or `go get`:
```bash
go get gopkg.in/zorkian/go-datadog-api.v2
```
## USAGE
Using the client:
```go
    client := datadog.NewClient("api key", "application key")

    dash, err := client.GetDashboard(10880)
    if err != nil {
        log.Fatalf("fatal: %s\n", err)
    }

    log.Printf("dashboard %d: %s\n", dash.Id, dash.Title)
```

Check out the Godoc link for the available API methods and, if you can't find the one you need,
let us know (or patches welcome)!

## DOCUMENTATION

Please see: <http://godoc.org/github.com/zorkian/go-datadog-api>

## BUGS/PROBLEMS/CONTRIBUTING

There are certainly some, but presently no known major bugs. If you do
find something that doesn't work as expected, please file an issue on
Github:

<https://github.com/zorkian/go-datadog-api/issues>

Thanks in advance! And, as always, patches welcome!

## DEVELOPMENT
* Run tests tests with `make test`.
* Integration tests can be run with `make testacc`. Run specific integration tests with `make testacc TESTARGS='-run=TestCreateAndDeleteMonitor'`

The acceptance tests require _DATADOG_API_KEY_ and _DATADOG_APP_KEY_ to be available
in your environment variables.

*Warning: the integrations tests will create and remove real resources in your Datadog account.*

## COPYRIGHT AND LICENSE

Please see the LICENSE file for the included license information.

Copyright 2013-2017 by authors and contributors.