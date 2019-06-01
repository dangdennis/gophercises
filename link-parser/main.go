package main

import (
	"bufio"
	"os"

	"golang.org/x/net/html"
)

// Anchor contains the key attributes of an anchor tag
type Anchor struct {
	href string
	text string
}

func main() {

	f, err := os.Open("ex1.html")
	check(err)

	r := bufio.NewReader(f)

	doc, err := html.Parse(r)
	check(err)

	findAnchors(doc)

}

func findAnchors(doc *html.Node) []Anchor {
	anchors := make([]Anchor, 50)

	anchors = append(anchors, Anchor{"hello", "world"})

	return anchors
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
