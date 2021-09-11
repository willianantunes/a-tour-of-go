package main

import (
	"github.com/willianantunes/a-tour-of-go/tutorials/writing_web_applications/internal/handlers"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", makeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", makeHandler(handlers.SaveHandler))

	// Fatal is equivalent to Print() followed by a call to os.Exit(1).
	log.Fatal(http.ListenAndServe(":8080", nil))
}
