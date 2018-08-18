package main

import (
	"fmt"
	"net/http"

	"github.com/esoros/gophercises/urlshortner"
)

func main() {
	mux := http.NewServeMux()
	u := urlshortner.Urlshortner{
		Paths: make(map[string]string, 0),
	}
	mux.HandleFunc("/", u.Handler)

	fmt.Println("listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
