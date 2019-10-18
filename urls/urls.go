package urls

import (
		"net/http"
		"SipToolKit/views"
       )

func LoadURLS(){
	http.Handle("/", http.FileServer(http.Dir("static")))
    http.HandleFunc("/testing", views.MainPage)
}