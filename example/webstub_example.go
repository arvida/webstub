package main

import (
	"github.com/arvida/webstub"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Inject webstub
	webstub.Enable()

	// Setup a stubbed response for GET requests to http://example.com/my-endpoint
	r := webstub.Request{
		Method:   "GET",
		Url:      "http://example.com/my-endpoint",
		Response: "Hello from the example!",
	}
	webstub.Register(r)

	// Make a request
	resp, err := http.Get("http://example.com/my-endpoint")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
