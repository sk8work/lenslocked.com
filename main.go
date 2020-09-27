package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	// fmt.Fprint(w, r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my super duper site</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch? please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	}
	// fmt.Println("Someone visited our page")
	// fmt.Fprint(w, "<h1>Welcome to my super duper site</h1>")
	// fmt.Fprint(w, "To get in touch? please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
}

// video 7
