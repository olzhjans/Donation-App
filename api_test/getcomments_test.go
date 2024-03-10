package api

import (
	"awesomeProject1/comments"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetComments(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"need-id":"65ba75f0ab0265f4fa9d4b5f",
    	"from":"2024-02-20T00:00:00Z",
    	"to":"2024-04-23T00:00:00Z"
	}`)
	request, err := http.NewRequest("POST", "/getComments", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	comments.GetComments(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Contains(t, response.Body.String(), "\"_id\":\"65deaa984fb9ab1f83c158bd\",")
}
