// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"reflect"
// )

// type Todo struct {
// 	Title string
// 	Done  bool
// }

// type TodoPageData struct {
// 	PageTitle string
// 	Todos     []Todo
// }

// func main() {
// 	tmpl := template.Must(template.ParseFiles("layout.html"))
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		data := TodoPageData{
// 			PageTitle: "My todo title",
// 			Todos: []Todo{
// 				{Title: "Task 1", Done: false},
// 				{Title: "Task 2", Done: true},
// 				{Title: "Task 3", Done: false},
// 			},
// 		}
// 		fmt.Println(data.PageTitle)
// 		fmt.Println(reflect.TypeOf(data))
// 		tmpl.Execute(w, data)
// 	})
// 	http.ListenAndServe(":80", nil)
// }
