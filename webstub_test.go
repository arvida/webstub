package webstub

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func init() {
	Enable()
}

func TestStubResponseFromString(t *testing.T) {
	p := Request{
		Method:   "GET",
		Url:      "http://example.com/my-endpoint",
		Response: "yo",
	}
	Register(p)

	_, body := get("http://example.com/my-endpoint")

	if body != "yo" {
		t.Errorf("returned body is not ”yo”, was: ”%s”", body)
	}
}

func TestStubStatusCode(t *testing.T) {
	p := Request{
		Method:     "GET",
		Url:        "http://example.com/my-endpoint",
		StatusCode: 418,
	}
	Register(p)

	resp, _ := get("http://example.com/my-endpoint")

	if resp.StatusCode != 418 {
		t.Errorf("returned status code is not ”418”, was: ”%s”", resp.StatusCode)
	}
}

func TestStubHeaders(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	p := Request{
		Method:  "GET",
		Url:     "http://example.com/my-endpoint",
		Headers: headers,
	}
	Register(p)

	resp, _ := get("http://example.com/my-endpoint")

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("returned status code is not ”application/json”, was: ”%s”", resp.Header.Get("Content-Type"))
	}
}

func TestStubResponseFromFile(t *testing.T) {
	p := Request{
		Method: "GET",
		Url:    "http://example.com/my-endpoint",
		File:   "fixtures/hello_response.json",
	}
	Register(p)

	resp, body := get("http://example.com/my-endpoint")

	if body != "hello there" {
		t.Errorf("returned body is not ”hello there”, was: ”%s”", body)
	}

	if resp.Header.Get("Content-Type") != "application/text" {
		t.Errorf("returned content type is not ”application/text”, was: ”%s”", resp.Header.Get("Content-Type"))
	}
}

// Helpers
func get(url string) (*http.Response, string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return resp, string(body)
}
