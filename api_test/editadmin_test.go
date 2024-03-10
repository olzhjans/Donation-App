package api

import (
	"awesomeProject1/edituser"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_EditAdmin(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"_id":"65e71c90cbb401656adb55a6",
    	"phone":"+77766787878"
	}`)
	request, err := http.NewRequest("POST", "/editAdmin", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	edituser.EditAdmin(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	expectedBody := "\"65e71c90cbb401656adb55a6 successfully edited\"\n"
	assert.Equal(t, expectedBody, response.Body.String())
}
