package urlshortner

import (
	"fmt"
	"net/http"
)

//Urlshortner is
type Urlshortner struct {
	Paths map[string]string
}

//Handler is
func (u *Urlshortner) Handler(w http.ResponseWriter, r *http.Request) {

	if val, ok := u.Paths[r.URL.Path]; ok {
		http.Redirect(w, r, val, 302)
	} else {
		fmt.Fprintf(w, "Redirect not found %q", r.URL.Path)
	}
}
