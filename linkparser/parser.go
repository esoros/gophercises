package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type anchor struct {
	href string
	text string
}

func main() {

	file, fileErr := os.Open("input.html")
	if fileErr != nil {
		fmt.Println("Unable to load file")
		os.Exit(1)
	}

	tokenizer := html.NewTokenizer(file)
	process(tokenizer)
}

func process(tokenizer *html.Tokenizer) {

	currentAnchor := anchor{}
	insideAnchor := false

	//refactor this out to use classes instead..
	for {
		tokenType := tokenizer.Next()
		switch {

		case tokenType == html.ErrorToken:
			return

		case tokenType == html.StartTagToken:

			token := tokenizer.Token()
			if token.Data == "a" && !insideAnchor {
				currentAnchor.href = getAttribute("href", token)
				insideAnchor = true
			}

		case tokenType == html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "a" && insideAnchor {
				insideAnchor = false
				fmt.Printf("closing anchor %s\n", currentAnchor.href)
			}

		case tokenType == html.TextToken:
			token := tokenizer.Token()
			if insideAnchor {
				currentAnchor.text += strings.TrimSpace(token.String())
			}
		}
	}
}

func getAttribute(key string, token html.Token) string {
	for _, k := range token.Attr {
		if k.Key == key {
			return k.Val
		}
	}
	return ""
}
