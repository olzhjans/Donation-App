package api

import (
	"awesomeProject1/auth"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_UserLogIn(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"phone":"+77787787878",
    	"password":"12345678"
	}`)
	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	auth.UserLogin(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, response.Body.String(), "{\"_id\":\"65c1da85b9683cf7113767cf\",\"donated\":72,\"email\":\"nurbeksila@gmail.com\",\"firstname\":\"Nurbek\",\"lastname\":\"Cvetmet\",\"password\":\"12345678\",\"phone\":\"+77787787878\",\"region\":\"Shymkent\",\"signupdate\":\"2024-02-05T19:00:00Z\"}\n")
}
