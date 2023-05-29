package Model

import "myapp/datastore/postgres"

type Student struct {
	StdId     int64  `json:"stdid"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
}

const queryInsertUser = "INSERT INTO student(stdid, firstname, lastname, email) VALUES($1, $2, $3, $4);"
const queryGetUser = "SELECT stdid, firstname, lastname, email from student WHERE stdid=$1;"
const queryUpdateUser = "UPDATE student SET stdid=$1, firstname=$2, lastname=$3, email=$4 WHERE stdid=$5;"
const queryDeleteUser = "DELETE from student WHERE stdid=$1;"

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}

func (s *Student) Read() error {
	return postgres.Db.QueryRow(queryGetUser, s.StdId).Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
}

func (s *Student) Update(oldID int64) error {
	_, error := postgres.Db.Exec(queryUpdateUser, s.StdId, s.FirstName, s.Email, oldID)
	return error
}

func (s *Student) Delete() error {
	if _, err := postgres.Db.Exec(queryDeleteUser, s.StdId); err != nil {
		return err
	}
	return nil
}

func GetAllStudents() ([]Student, error) {
	rows, getErr := postgres.Db.Query("SELECT * from student;")
	if getErr != nil {
		return nil, getErr
	}
	//Create a slice type student
	students := []Student{}

	for rows.Next() {
		var s Student
		dbErr := rows.Scan(&s.StdId, &s.FirstName, &s.LastName, &s.Email)
		if dbErr != nil {
			return nil, dbErr
		}
		students = append(students, s)
	}
	rows.Close()
	return students, nil
}

// signup
type Admin struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

const queryInsertAdmin = "INSERT into admin(firstname, lastname, email, password) VALUES ($1, $2, $3, $4);"

func (adm *Admin) Create() error {
	_, err := postgres.Db.Exec(queryInsertAdmin, adm.Firstname, adm.Lastname, adm.Email, adm.Password)
	return err
}

// login
type Login struct {
	Email    string
	Password string
}

const queryGetAdmin = "SELECT email, password FROM admin WHERE email = $1 and password = $2;"

func (adm *Login) Get() error {
	return postgres.Db.QueryRow(queryGetAdmin, adm.Email, adm.Password).Scan(&adm.Email, &adm.Password)
}

// course
type Course struct {
	Courseid   string `json:"courseid"`
	CourseName string `json:"coursename"`
}

const courseinsert = "INSERT INTO course(courseid, coursename)VALUES($1,$2);"

func (c *Course) Create() error {
	_, err := postgres.Db.Exec(courseinsert, c.Courseid, c.CourseName)
	return err
}

var queryGetCourse = "SELECT courseid, coursename FROM course WHERE courseid=$1;"

func (c *Course) Read() error {
	return postgres.Db.QueryRow(queryGetCourse, c.Courseid).Scan(&c.Courseid, &c.CourseName)
}

var queryUpdateCourse = "UPDATE course SET courseid = $1, coursename = $2 WHERE courseid = $3;"

func (c *Course) Update(oldID string) error {
	_, err := postgres.Db.Exec(queryUpdateCourse, c.Courseid, c.CourseName, oldID)
	return err
}

var queryDeleteCourse = "DELETE FROM course WHERE courseid=$1;"

func (c *Course) Delete() error {
	if _, err := postgres.Db.Exec(queryDeleteCourse, c.Courseid); err != nil {
		return err
	}
	return nil
}

func GetAllCourses() ([]Course, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM course;")
	if getErr != nil {
		return nil, getErr
	}
	courses := []Course{}

	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.Courseid, &c.CourseName)
		if dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}

//enroll
type Enroll struct {
	StdId         int64  `json:"stdid"`
	CourseID      string `json:"courseid"`
	Date_Enrolled string `json:"date"`
}

const queryEnrollStd = "INSERT INTO enroll(std_id, course_id, date_enrolled) VALUES($1, $2, $3);"
const queryGetEnroll = "SELECT std_Id, course_id, date_enrolled FROM enroll WHERE std_id=$1 and course_id=$2;"
const queryDeleteEnroll = "DELETE FROM enroll WHERE std_id=$1 and course_id=$2;"

func (e *Enroll) EnrollStud() error {
	if _, err := postgres.Db.Exec(queryEnrollStd, e.StdId, e.CourseID, e.Date_Enrolled); err != nil {
		return err
	}
	return nil
}

func (e *Enroll) Get() error {
	return postgres.Db.QueryRow(queryGetEnroll,
		e.StdId, e.CourseID).Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
}

func GetAllEnrolls() ([]Enroll, error) {
	rows, getErr := postgres.Db.Query("SELECT std_id, course_id, date_enrolled FROM enroll;")
	if getErr != nil {
		return nil, getErr
	}
	enrolls := []Enroll{}

	for rows.Next() {
		var e Enroll
		dbErr := rows.Scan(&e.StdId, &e.CourseID, &e.Date_Enrolled)
		if dbErr != nil {
			return nil, dbErr
		}
		enrolls = append(enrolls, e)
	}
	rows.Close()
	return enrolls, nil
}

func (e *Enroll) Delete() error {
	if _, err := postgres.Db.Exec(queryDeleteEnroll, e.StdId, e.CourseID); err != nil {
		return err
	}
	return nil
}
