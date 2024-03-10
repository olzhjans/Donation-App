package api

import (
	"awesomeProject1/edituser"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_EditUser(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"_id":"65e719fd9c39339929ae5b5d",
    	"phone":"+77766787878"
	}`)
	request, err := http.NewRequest("POST", "/editUser", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	edituser.EditUser(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	expectedBody := "\"65e719fd9c39339929ae5b5d successfully edited\"\n"
	assert.Equal(t, expectedBody, response.Body.String())
}
