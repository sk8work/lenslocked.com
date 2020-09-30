package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)

	// http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", mux)

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my super duper site</h1>")
	} else if r.URL.Path == "/contact" || r.URL.Path == "/contact/" {
		fmt.Fprint(w, "To get in touch? please send an email to <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1><p>Please email us if uoy keep being sent to invalid page <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a></p>")
	}
}

// video 9
