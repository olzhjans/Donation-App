package api_test

import (
	"awesomeProject1/adminrights/adminauth"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_AdminSignUp(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"phone":"+77766787878",
    	"password":"12345678",
    	"firstname":"John",
    	"lastname":"Doe",
    	"email":"johndoe@gmail.com",
    	"region":"Almaty",
    	"id":"950101123123",
    	"who":"Moderator",
    	"orphanage-id":"65ba7800ab0265f4fa9d4b60"
	}`)
	request, err := http.NewRequest("POST", "/adminSignUp", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	adminauth.AdminSignUp(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	expectedBody := "\"Added successfully\"\n"
	assert.Equal(t, expectedBody, response.Body.String())
}
