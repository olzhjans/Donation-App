package api

import (
	"awesomeProject1/adminrights/orphanage"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_AddOrphanage(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"name":"321n",
		"region":"321r",
		"address":"321a",
    	"description":"321d",
		"childs-count":3,
		"working-hours":"321w",
		"photos":["kek.com","k"],
		"bill":0
	}`)
	request, err := http.NewRequest("POST", "/addOrphanage", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	orphanage.AddOrphanage(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	expectedBody := "\"Successfully added\"\n"
	assert.Equal(t, expectedBody, response.Body.String())
}
