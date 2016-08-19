package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := getlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "url get error: %v\n", err)
			continue
		}
		for _, l := range links {
			fmt.Println(l)
		}
	}
}

func getlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("%s status error", url)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	links := visit(nil, doc)
	return links, nil
}

func visit(stack []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, url := range n.Attr {
			if url.Key == "href" {
				stack = append(stack, url.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		stack = visit(stack, c)
	}
	return stack
}
