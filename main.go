package main

import (
	"fmt"
	"net/http"
)

type User struct {
	id       uint
	login    string
	password string
	name     string
	age      uint16
}

func main() {
	handleRequest()
}

func handleRequest() {
	// call a home page and start of func "home_page"
	http.HandleFunc("/", home_page)
	http.HandleFunc("/editor-page/", editor_page)
	// Added a local server with parametrs (nil = null = none)
	http.ListenAndServe(":8080", nil)
}

func editor_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Editor page")
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{1, "ddd", "222", "BOB", 1}
	bob.name = "Alex"
	fmt.Fprintf(w, "User name is: "+bob.name)
}
