package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	id       uint
	Login    string
	Password string
	Name     string
	Age      uint16
	Articles []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d."+
		"And his login is %s", u.Name, u.Age, u.Login)
}

// * - using an object like a link (using a link on the object)
func (u *User) setNewName(newName string) {
	u.Name = newName
}

func main() {
	handleRequest()
}

func handleRequest() {
	// call a homeа page and start of func "home_page"
	http.HandleFunc("/", home_page)
	http.HandleFunc("/editor-page/", editor_page)
	// Added a local server with parametrs (nil = null = none)
	http.ListenAndServe(":8080", nil)
}

func editor_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Editor page")
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{1, "ddd", "222", "BOB", 1, []string{"A1", "A2", "A3"}}
	// bob.setNewName("Alex")
	// fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}
