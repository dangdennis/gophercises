package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"os"
)

func main() {
	fmt.Println("heya")
	filename := flag.String("file", "gopher.json", "our JSON file for the story")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)

	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
}

// Chapter is a chapter
type Chapter struct {
	Intro struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"intro"`
}
