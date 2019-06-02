package anchor

import "golang.org/x/net/html"

// Anchor contains the key attributes of an anchor tag
type Anchor struct {
	href string
	text string
}

// FindAnchors traverses an HTML doc to return an array of anchor tags
func FindAnchors(n *html.Node) []Anchor {
	var anchors []Anchor

	if n.Type == html.ElementNode && n.Data == "a" {
		text := extractText(n)
		href := extractAttribute(n, "href")
		anchor := Anchor{href, text}
		anchors = append(anchors, anchor)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		anchors = append(anchors, FindAnchors(c)...)
	}

	return anchors
}

func extractAttribute(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func extractText(n *html.Node) string {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			return c.Data
		}
	}
	return ""
}
