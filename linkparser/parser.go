package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type anchor struct {
	Href string `json:"href"`
	Text string `json:"text"`
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
	anchors := make([]anchor, 0)
	insideAnchor := false

	//refactor this out to use classes instead.. (or whatever makes sense in go)
	for {
		tokenType := tokenizer.Next()
		switch {

		case tokenType == html.ErrorToken:
			bytes, _ := json.Marshal(anchors)
			fmt.Println(fmt.Sprintf("%s", bytes))
			return

		case tokenType == html.StartTagToken:

			token := tokenizer.Token()
			if token.Data == "a" && !insideAnchor {
				currentAnchor.Href = getAttribute("href", token)
				insideAnchor = true
			}

		case tokenType == html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "a" && insideAnchor {
				insideAnchor = false
				anchors = append(anchors, currentAnchor)
				fmt.Printf("Anchor close %s\n", currentAnchor.Text)
			}

		case tokenType == html.TextToken:
			token := tokenizer.Token()
			if insideAnchor {
				currentAnchor.Text = fmt.Sprintf("%s%s ", currentAnchor.Text, strings.TrimSpace(token.String()))
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
