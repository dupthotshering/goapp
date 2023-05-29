package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdmLogin(t *testing.T) {
	url := "http://localhost:8080/login"

	var jsonStr = []byte(`{"email":"12220090.gcit@rub.edu.bt", "password":"@Tshermax10"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expResp := `{"message":"login success"}`

	assert.JSONEq(t, expResp, string(body))
}
