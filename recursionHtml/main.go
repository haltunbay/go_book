package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	check(err)
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElenentNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "a" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func check(e error) {
	if e != nil {
		fmt.Printf("unexpected error %v\n", e)
		os.Exit(-1)
	}
}
