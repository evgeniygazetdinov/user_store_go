package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// const PORT = ":8080"

type User struct {
	Name                string
	Age                 uint16
	Money               int16
	AvgGrades, Happines float64
	Hobbies             []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s He is %d and he"+
		"has money: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(page http.ResponseWriter, r *http.Request) {
	bob := User{"boh", 25, -50, 3.0, 5.0, []string{"football", "dancer"}}
	bob.setNewName("Alex")
	// fmt.Fprintf(page, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(page, bob)
}

func contactsPage(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "contacts page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)

	http.ListenAndServe(PORT, nil)
}
