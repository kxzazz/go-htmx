package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)


type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	IsCompleted bool `json:"isCompleted"`
}

var todos = []Todo{
	{Id: 1, Name: "Learn Go", IsCompleted: false},
	{Id: 2, Name: "Learn Alpine", IsCompleted: false},
	{Id: 3, Name: "Go to the gym", IsCompleted: true},
}

var templates map[string]*template.Template

func init(){
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	// fmt.Println("init execd")
	templates["index.html"] = template.Must(template.ParseFiles("index.html"))
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	json, err:=json.Marshal(todos)

	if err != nil{
		log.Fatal(err)
	}

	// fmt.Println("handler")
	tmpl := templates["index.html"]
	tmpl.ExecuteTemplate(w,"index.html",map[string]template.JS{"Todos":template.JS(json)})
}

func submitTodoHandler(w http.ResponseWriter, r *http.Request){
	name := r.PostFormValue("name")
	
	todo := Todo{ Id : 4, Name: name, IsCompleted: false}
	fmt.Println(todo)
}

func main(){
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit-todo/", submitTodoHandler)
	log.Fatal(http.ListenAndServe(":8000",nil))
}