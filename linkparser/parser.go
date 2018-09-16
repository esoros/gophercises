package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

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
	for {
		tokenMetadata := tokenizer.Next()
		switch {
		case tokenMetadata == html.ErrorToken:
			return
		case tokenMetadata == html.StartTagToken:
			token := tokenizer.Token()
			fmt.Printf("token start %s\n", token.Data)
			//once we find the start of an <A> tag
		case tokenMetadata == html.EndTagToken:
			token := tokenizer.Token()
			fmt.Printf("token end %s\n", token.Data)
		case tokenMetadata == html.TextToken:
			token := tokenizer.Token()
			fmt.Printf("token text %s\n", strings.TrimSpace(token.String()))
		}
	}
}
