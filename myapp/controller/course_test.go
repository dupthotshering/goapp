package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	url := "http://localhost:8080/course"

	var jsonStr = []byte(`{"courseid":"CSC101", "coursename": "CSM101"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	// cookie, err := resp.Cookie("my-cookie")
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	expResp := `{"status":"course added"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestGetCourse(t *testing.T) {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/course/CSC101")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(t, http.StatusOK, r.StatusCode)
	expResp := `{"courseid":"CSC101", "coursename": "CSM101"}`
	assert.JSONEq(t, expResp, string(body))
}
func TestDeleteCourse(t *testing.T) {
	url := "http://localhost:8080/course/CSM101"
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expResp := `{"status":"deleted"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestCourseNotFound(t *testing.T) {
	assert := assert.New(t)
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/course/CD1002")
	body, _ := io.ReadAll(r.Body)
	assert.Equal(http.StatusNotFound, r.StatusCode)
	expResp := `{"error":"Course Not Found"}`
	assert.JSONEq(expResp, string(body))
}
