package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/Model"
	"myapp/utils/httpResp"
	"myapp/utils/httpResp/date"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "add student handler")
	var stud Model.Student
	_ = stud

	decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)

	if err := decoder.Decode(&stud); err != nil {
		// w.Write([]byte("Invalid json data"))
		response, _ := json.Marshal(map[string]string{"error": "Invalid Json Type"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	defer r.Body.Close()

	saveErr := stud.Create()
	fmt.Println(saveErr)
	if saveErr != nil {
		// w.Write([]byte("Database error"))
		response, _ := json.Marshal(map[string]string{"error": saveErr.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
	//no error
	// w.Write([]byte("response success"))
	response, _ := json.Marshal(map[string]string{"status": "student added"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func GetStud(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	fmt.Println(sid)
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := Model.Student{StdId: stdId}
	getErr := s.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		default:
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, s)
}

func getUserId(userIdParam string) (int64, error) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, userErr
	}
	return userId, nil
}

func UpdateStud(w http.ResponseWriter, r *http.Request) {
	old_sid := mux.Vars(r)["sid"]
	old_stdId, idErr := getUserId(old_sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	var stud Model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid jason body")
		return
	}
	defer r.Body.Close()

	err := stud.Update(old_stdId)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, stud)
}

func DeleteStud(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	stdId, idErr := getUserId(sid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	s := Model.Student{StdId: stdId}
	if err := s.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllStuds(w http.ResponseWriter, r *http.Request) {
	students, getErr := Model.GetAllStudents()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, students)
}

// signup
func Signup(w http.ResponseWriter, r *http.Request) {

	var admin Model.Admin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&admin)

	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid Json Typer")
		return
	}
	defer r.Body.Close()
	saveErr := admin.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}

// login
func Login(w http.ResponseWriter, r *http.Request) {

	// set cookie
	cookie := http.Cookie{
		Name:    "my-cookie",
		Value:   "my-value",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	http.SetCookie(w, &cookie)

	var admin Model.Login

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&admin)

	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
		return
	}
	defer r.Body.Close()

	getErr := admin.Get()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusUnauthorized, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "login success"})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	})
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "cookie deleted"})
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the "my-cookie" cookie from the request
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			// No cookie found, redirect to login page or return an error
			httpResp.RespondWithError(w, http.StatusSeeOther, "cookie no found")
			return false
		}
		//some other error occured
		httpResp.RespondWithError(w, http.StatusInternalServerError, "internal server error")
		return false
	}
	// Verify the cookie value
	if cookie.Value != "my-value" {
		// Invalid cookie value, redirect to login page or return an error
		httpResp.RespondWithError(w, http.StatusSeeOther, "cookie does not match")
		return false
	}
	return true
}

// course
func Addcourse(w http.ResponseWriter, r *http.Request) {
	var course Model.Course
	_ = course

	decoder := json.NewDecoder(r.Body)
	// fmt.Println(decoder)

	if err := decoder.Decode(&course); err != nil {
		// w.Write([]byte("Invalid json data"))
		response, _ := json.Marshal(map[string]string{"error": "Invalid Json Type"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	defer r.Body.Close()

	saveErr := course.Create()
	fmt.Println(saveErr)
	if saveErr != nil {
		// w.Write([]byte("Database error"))
		response, _ := json.Marshal(map[string]string{"error": saveErr.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
	//no error
	// w.Write([]byte("response success"))
	response, _ := json.Marshal(map[string]string{"status": "course added"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Getcour(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	fmt.Println(cid)
	courseid, idErr := getCourseId(cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := Model.Course{Courseid: courseid}
	getErr := c.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Course Not Found")
		default:
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func getCourseId(courseIdParam string) (string, error) {
	return courseIdParam, nil
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	old_cid := mux.Vars(r)["cid"]
	old_courseId, idErr := getCourseId(old_cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	var courses Model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&courses); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid jason body")
		return
	}
	defer r.Body.Close()

	err := courses.Update(old_courseId)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)

}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	courseid, idErr := getCourseId(cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := Model.Course{Courseid: courseid}
	if err := c.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, getErr := Model.GetAllCourses()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}

// enroll
func Enroll(w http.ResponseWriter, r *http.Request) {
	var e Model.Enroll
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&e); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	e.Date_Enrolled = date.GetDate()
	defer r.Body.Close()

	saveErr := e.EnrollStud()
	if saveErr != nil {
		if strings.Contains(saveErr.Error(), "duplicate key") {
			httpResp.RespondWithError(w, http.StatusForbidden, "Duplicate keys")
			return
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, saveErr.Error())
		}
	}
	// no error
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "enrolled"})
}

func GetEnroll(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]

	stdid, _ := strconv.ParseInt(sid, 10, 64)

	e := Model.Enroll{StdId: stdid, CourseID: cid}
	getErr := e.Get()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "No such Enrollments")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, e)
}

func GetEnrolls(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reaching here")
	enrolls, getErr := Model.GetAllEnrolls()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		fmt.Println("eeee", getErr)
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, enrolls)
	fmt.Println("result", enrolls)
}

func DeleteEnroll(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	cid := mux.Vars(r)["cid"]
	stdid, _ := strconv.ParseInt(sid, 10, 64)
	e := Model.Enroll{StdId: stdid, CourseID: cid}

	if err := e.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
