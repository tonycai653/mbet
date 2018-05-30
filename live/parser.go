package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

func parse(content io.Reader) {
	rootNode, err := html.Parse(content)
	if err != nil {

	}
}
