# HTML_Link_parser

In this project my  goal is to create a package that makes it easy to parse an HTML file and extract all of the links (<a href="">...</a> tags). For each extracted link it will return a data structure that includes both the href, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Using the x/net/html package for this task, which I will need to go get. It is provided by the Go team, but isn't included in the standard library. This makes it a little easier to parse HTML files. I will ignore any links nested inside of another link. In this project I end up using a DFS, which is a graph theory algorithm.
