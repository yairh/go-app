package main

import (
	"./models"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func Landing(w http.ResponseWriter, r *http.Request) {
	me := models.Person{Name: "Yair", Age: 28}
	templatePath := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		err := tmpl.Execute(w, me)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	http.HandleFunc("/", Landing)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error running server")
	} else {
		fmt.Println("Server running at localhost:3000")
	}
}
