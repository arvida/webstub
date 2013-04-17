package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Beer struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Abv         float32 `json:"abv"`
}

type SearchResult struct {
	Beers []Beer
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage %s <beer to search for>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	result := findBeers(os.Args[1])

	beers_count := len(result.Beers)
	if beers_count == 0 {
		log.Fatal("No beers found")
	}

	fmt.Printf("Found %d beers\n", beers_count)
	for _, beer := range result.Beers {
		fmt.Println("---")
		fmt.Printf("%s, %g%%\n", beer.Name, beer.Abv)
		fmt.Println(beer.Description)
	}
}

func findBeers(query string) (result SearchResult) {
	resp, err := http.Get("http://api.openbeerdatabase.com/v1/beers.json?query=" + query)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
