package api

import (
	"awesomeProject1/adminrights/orphanage"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_EditOrphanage(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"_id":"65ebfa792a10e66ce0a6a8f5",
    	"name":"Umit",
    	"region":"Almaty",
    	"address":"Abay st. 1",
    	"description":"Description of orphanage",
    	"childs-count":40,
    	"working-hours":"8AM - 6PM",
    	"photos":["link_to_photo_1","link_to_photo_2"],
    	"bill":0
	}`)
	request, err := http.NewRequest("POST", "/editOrphanage", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	orphanage.EditOrphanage(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	expectedBody := "\"65ebfa792a10e66ce0a6a8f5 successfully edited\"\n"
	assert.Equal(t, expectedBody, response.Body.String())
}
