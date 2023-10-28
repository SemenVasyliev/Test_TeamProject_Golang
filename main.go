package main

import (
	"fmt"
	"net/http"
)

func main() {
	// call a home page and start of func "home_page"
	http.HandleFunc("/", home_page)
	// Added a local server with parametrs (nil = null = none)
	http.ListenAndServe(":8080", nil)
}
func home_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Go is super easy!")
}
