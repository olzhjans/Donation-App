package api

import (
	"awesomeProject1/auth"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_UserSignUp(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"phone":"+77072077737",
		"password":"12345678",
		"firstname":"John",
		"lastname":"Doe",
		"email":"johndoe@gmail.com",
		"region":"Almaty",
		"donated":0
	}`)
	request, err := http.NewRequest("POST", "/userSignUp", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	auth.UserSignUp(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	expectedBody := "\"Added successfully\"\n"
	assert.Equal(t, expectedBody, response.Body.String())
}
