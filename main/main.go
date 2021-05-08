package main

import (
	"HTML_Link_parser/parse"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "message.html", "The parsing of HTML file ")
	flag.Parse()
	s, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(s)
	if err != nil {
		panic(err)
	}

	for _, i := range links {
		fmt.Println("Href: ", i.Href)
		fmt.Println("Text: ", i.Text)
	}
	return
}
