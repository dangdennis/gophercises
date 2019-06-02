package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/dangdennis/gophercises/link-parser/parser"
	"golang.org/x/net/html"
)

func main() {
	flag.Parse()
	f, err := os.Open(flag.Arg(0))
	check(err)

	r := bufio.NewReader(f)

	doc, err := html.Parse(r)
	check(err)

	fmt.Printf("%v", parser.FindAnchors(doc))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
