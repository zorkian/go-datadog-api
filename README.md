datadog-api -- a Go implementation

Hi!

This is a very early stage project. Someone mentioned there was no Go
implementation of a Datadog API, so I took a crack at it. It is not
fully featured yet.

The source API documentation is here: <http://docs.datadoghq.com/api/>

To use this project, include it in your code like:

    import "github.com/xb95/datadog-api"

Then, you can work with it:

    client := datadog.NewClient("api key", "application key")
    
    dash, err := client.GetDashboard(10880)
    if err != nil {
        log.Fatalf("fatal: %s\n", err)
    }
    log.Printf("dashboard %d: %s\n", dash.Id, dash.Title)

Supported methods:

    PostMetrics([]Metric)       upload timeseries data
    
    PostEvent(Event)            create a new event
    GetEvent(id)                return a single event
    GetEvents(see code)         return a list of events matching a query
    
    GetTags(...)                get all tags
    GetHostTags(...)            get all host tags
    GetHostTagsBySource(...)    get host tags, with source information
    AddTagsToHost(...)          add new tags to a host/source
    UpdateHostTags(...)         overwrite all tags for this host/source
    RemoveHostTags(...)         remove all tags that match this host/source
    
    CreateDashboard(Dashboard)  creates a new dashboard
    DeleteDashboard(id)         deletes one dashboard
    GetDashboard(id)            returns one (with details, graphs, etc)
    GetDashboards()             returns all (only title/desc/url/id)
    UpdateDashboard(Dashboard)  persists changes to a dashboard

More information to come. Please see the LICENSE file for the included
license information.

Copyright 2013 by authors and contributors.
