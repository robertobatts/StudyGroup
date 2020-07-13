package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func run() {
	url := "https://golang.org"
	doc, err := getHTMLFromURL(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(visit(nil, doc))
	outline(nil, doc)
	fmt.Println(countWords(doc))
	forEachNode(doc, startElement, endElement)
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s='%s'", attr.Key, attr.Val)
		}
		fmt.Println(">")
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func getHTMLFromURL(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, nil
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

func countWords(n *html.Node) int {
	nWords := 0
	if n == nil {
		return nWords
	}
	if n.Type == html.TextNode {
		words := strings.Split(n.Data, " ")
		nWords = len(words)
	}
	nWords += countWords(n.FirstChild)
	nWords += countWords(n.NextSibling)
	return nWords
}
