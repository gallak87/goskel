package main

import (
	"fmt"
	"sync"
)

var (
	mux sync.Mutex
	complete map[string]string
)

func main() {
	complete = make(map[string]string)
	Crawl("https://golang.org/", 4, fetcher)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	mux.Lock()
	if _, ok := complete[url]; ok {
		mux.Unlock()
		return
	}

	body, urls, err := fetcher.Fetch(url)
	complete[url] = body
	mux.Unlock()
	if err != nil {
		fmt.Println(err)
		return
	}
	done := make(chan bool)
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go func(u string) {
			Crawl(u, depth-1, fetcher)
			done <- true
		}(u)
	}
	for _, _ = range urls {
		<-done
	}
	return
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}