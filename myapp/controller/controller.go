package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func CourseHandler(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)

	course := p["course"]
	_, err := w.Write([]byte("Hello world. \n This is course is " + course))
	if err != nil {
		fmt.Println("error: ", err)
	}
}
