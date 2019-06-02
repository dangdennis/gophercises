package parser

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

const htmlMock2 = `
<html>
  <body>
    <a href="/dog-cat"
      >dog cat
      <!-- commented text SHOULD NOT be included! --></a
    >
  </body>
</html>
`

func TestFindAnchors(t *testing.T) {
	r := strings.NewReader(htmlMock1)
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	anchors := FindAnchors(doc)
	expected := []Anchor{Anchor{"/other-page", "A link to another page"}}

	// Plain old anchors
	if anchors[0].href != expected[0].href {
		t.Errorf("FindAnchors(html)[0].href = %v, want %v", anchors[0].href, expected[0].href)
	}

	if anchors[0].text != expected[0].text {
		t.Errorf("FindAnchors(html)[0].text = %v, want %v", anchors[0].text, expected[0].text)
	}

	// Anchors with comments
	r = strings.NewReader(htmlMock2)
	doc, err = html.Parse(r)
	if err != nil {
		panic(err)
	}

	anchors = FindAnchors(doc)
	expected = []Anchor{Anchor{"/dog-cat", "dog cat"}}

	if anchors[0].href != expected[0].href {
		t.Errorf("FindAnchors(html)[0].href = %v, want %v", anchors[0].href, expected[0].href)
	}

	if anchors[0].text != expected[0].text {
		t.Errorf("FindAnchors(html)[0].text = %v, want %v", anchors[0].text, expected[0].text)
	}
}
