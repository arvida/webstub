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
