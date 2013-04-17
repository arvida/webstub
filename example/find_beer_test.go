package main

import (
	"github.com/arvida/webstub"
	"testing"
)

func TestFindBeers(t *testing.T) {
	webstub.Enable()

	r := webstub.Request{
		Method: "GET",
		Url:    "http://api.openbeerdatabase.com/v1/beers.json?query=mybeer",
		Response: `{
      "beers": [{
        "name": "Da Beer",
        "description": "Great",
        "abv": 3.5
      }]
    }`,
	}
	webstub.Register(r)

	result := findBeers("mybeer")

	if len(result.Beers) != 1 {
		t.Errorf("expected found beers to be 1, was %d", len(result.Beers))
	}

	if result.Beers[0].Name != "Da Beer" {
		t.Errorf("expected beer name to be ”Da Beer”, was %s", result.Beers[0].Name)
	}
}
