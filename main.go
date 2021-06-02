package main

import (
	"fmt"
	"strings"

	"github.com/himanshu07070/HTML_Link_Parser/link"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
    My name is Himanshu Joshi
    <span> I am a Golang Developer  </span>
  </a>
  <a href="/page-two">This is a Html Link Parser</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
