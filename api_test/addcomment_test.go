package api

import (
	"awesomeProject1/comments"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_AddComment(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"need-id":"65c0842fdfda391f6fe449ac",
    	"user-id":"65c3b0b072629755858ece76",
    	"text":"text"
	}`)
	request, err := http.NewRequest("POST", "/addComment", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	comments.AddComment(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Contains(t, response.Body.String(), "\"Added successfully\"")
}
