package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func dealWith(err interface{}, format string, args ...interface{}) {
	if (err != nil) {
		fmt.Fprintf(os.Stderr, format, args...)
		os.Exit(1)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	dealWith(err, "outline: %v\n", err)
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}