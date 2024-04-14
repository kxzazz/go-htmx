package main

import (
	"html/template"
	"log"
	"net/http"
)


type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	IsCompleted bool `json:"IsCompleted"`
}

var todo = []Todo{
	{Id: 1, Name: "Learn Go", IsCompleted: false},
	{Id: 2, Name: "Learn Alpine", IsCompleted: false},
	{Id: 3, Name: "Go to the gym", IsCompleted: true},
}

var templates map[string]*template.Template

func init(){
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index.html"] = template.Must(template.ParseFiles("index.html"))
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	tmpl := templates["index.html"]
	tmpl.ExecuteTemplate(w,"index.html",nil)
}

func main(){
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8000",nil))
}