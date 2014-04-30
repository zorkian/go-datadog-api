# Datadog API in Go

Hi!

This is a very early stage project. Someone mentioned there was no Go
implementation of a Datadog API, so I took a crack at it. It is not
fully featured yet.

The source API documentation is here: <http://docs.datadoghq.com/api/>

## USAGE

To use this project, include it in your code like:

    import "github.com/zorkian/go-datadog-api"

Then, you can work with it:

    client := datadog.NewClient("api key", "application key")
    
    dash, err := client.GetDashboard(10880)
    if err != nil {
        log.Fatalf("fatal: %s\n", err)
    }
    log.Printf("dashboard %d: %s\n", dash.Id, dash.Title)

That's all; it's pretty easy to use.

## DOCUMENTATION

Please see: <http://godoc.org/github.com/zorkian/go-datadog-api>

## BUGS/PROBLEMS/CONTRIBUTING

There are certainly some, but presently no known major bugs. If you do
find something that doesn't work as expected, please file an issue on
Github:

<https://github.com/zorkian/go-datadog-api/issues>

Thanks in advance! And, as always, patches welcome!

## COPYRIGHT AND LICENSE

Please see the LICENSE file for the included license information.

Copyright 2013 by authors and contributors.
