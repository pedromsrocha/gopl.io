// See ex 5.2
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "populate: %v\n", err)
		os.Exit(1)
	}
	elementsCount := populate(map[string]int{}, doc)
	fmt.Printf("element\tcount\n")
	for element, count := range elementsCount {
		fmt.Printf("%s: %d\n", element, count)
	}
}
func populate(elementsCount map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elementsCount[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		populate(elementsCount, c)
	}
	return elementsCount
}
