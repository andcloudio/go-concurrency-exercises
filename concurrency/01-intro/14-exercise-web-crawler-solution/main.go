package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

var fetched map[string]bool

type result struct {
	url   string
	urls  []string
	err   error
	depth int
}

// Crawl uses findLinks to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int) {
	results := make(chan *result)

	fetch := func(url string, depth int) {
		urls, err := findLinks(url)
		results <- &result{url, urls, err, depth}
	}

	go fetch(url, depth)
	fetched[url] = true

	for fetching := 1; fetching > 0; fetching-- {
		res := <-results
		if res.err != nil {
			// fmt.Println(res.err)
			continue
		}

		fmt.Printf("found: %s\n", res.url)
		if res.depth > 0 {
			for _, u := range res.urls {
				if !fetched[u] {
					fetching++
					go fetch(u, res.depth-1)
					fetched[u] = true
				}
			}
		}
	}
	close(results)
}

func main() {
	fetched = make(map[string]bool)
	now := time.Now()
	Crawl("http://andcloud.io", 2)
	fmt.Println("time taken:", time.Since(now))
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
