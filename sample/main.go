package main

import (
	"fmt"
	"gosoup"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("getting %s: %s", url, resp.Status)
	}
	soup, err := gosoup.Parse(resp.Body)
	if err != nil {
		log.Fatalf("parsing %s as HTML: %v", url, err)
	}
	fmt.Print(soup.GetText())
}
