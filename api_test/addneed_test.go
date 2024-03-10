package api

import (
	"awesomeProject1/needs"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_AddNeed(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"amount":1,
    	"categoryofdonate":"1",
    	"sizeofclothes":"1",
    	"typeofcount":"1",
    	"typeofdonate":"1",
    	"orphanageid":"1",
    	"isactive":true
	}`)
	request, err := http.NewRequest("POST", "/addNeed", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	needs.AddNeeds(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Contains(t, response.Body.String(), "\"Successfully added\"")
}
