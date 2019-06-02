package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const htmlMock1 = `
	<html>
	<body>
		<h1>Hello!</h1>
		<a href="/other-page">A link to another page</a>
	</body>
	</html>
`

func TestFindAnchor(t *testing.T) {
	r := strings.NewReader(htmlMock1)
	doc, err := html.Parse(r)
	anchors := main.FindAnchors(doc)
}
