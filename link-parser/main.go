package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/dangdennis/gophercises/anchor"
	"golang.org/x/net/html"
)

func main() {
	flag.Parse()
	f, err := os.Open(flag.Arg(0))
	check(err)

	r := bufio.NewReader(f)

	doc, err := html.Parse(r)
	check(err)

	fmt.Printf("%v", anchor.FindAnchors(doc))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
