# webstub

A Go package for stubbing HTTP requests.

Can be useful when you are writing tests for your golang code that talks to HTTP servers.

## Installation

	$ go get github.com/arvida/webstub

## Example

```go
package main

import (
	"fmt"
	"github.com/arvida/webstub"
	"net/http"
)

func main() {
	// Inject webstub
	webstub.Enable()

	// Setup a stubbed response for GET requests to http://example.com/my-endpoint
	p := webstub.Request{
		method:   "GET",
		url:      "http://example.com/my-endpoint",
		response: "Hello from the example!",
	}
	webstub.Register(p)

	// Make a request
	resp, err := http.Get("http://example.com/my-endpoint")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
}

```

### Response from file

Save the response to a text file. It is easy to save a live response with `curl`:

	$ curl -is www.example.com > stubbed_response.txt

Set the `webstub.Request` `file` parameter to specify the response file to use:

```go
p := webstub.Request{
	method: "GET",
    url: "http://example.com/",
    file: "/stubbed_response.txt",
}
webstub.Register(p)
```

## Contribute

Please do. Create a issue or pull request.

## Copyright

Copyright (c) 2013 Arvid Andersson. See LICENSE for details.
