package routes

import (
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {

	router := mux.NewRouter()

	router.HandleFunc("/home", controller.HomeHandler)
	router.HandleFunc("/home/{course}", controller.CourseHandler)
	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")
	router.HandleFunc("/students", controller.GetAllStuds)

	//course
	router.HandleFunc("/course", controller.Addcourse).Methods("POST")
	router.HandleFunc("/course/{cid}", controller.Getcour).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/courses", controller.GetAllCourses)

	//enroll
	router.HandleFunc("/enroll", controller.Enroll).Methods("POST")
	router.HandleFunc("/enroll/{sid}/{cid}", controller.GetEnroll).Methods("GET")
	router.HandleFunc("/enroll/{sid}/{cid}", controller.DeleteEnroll).Methods("DELETE")
	router.HandleFunc("/enrolls", controller.GetEnrolls)

	//signup
	router.HandleFunc("/signup", controller.Signup).Methods("POST")

	//login
	router.HandleFunc("/login", controller.Login).Methods("POST")

	//logout
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	fhandler := http.FileServer(http.Dir("./View"))
	router.PathPrefix("/").Handler(fhandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
