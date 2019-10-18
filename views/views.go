package views

import (
	    "fmt"
		"html/template"
		"net/http"
       )


func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w,nil)
}
