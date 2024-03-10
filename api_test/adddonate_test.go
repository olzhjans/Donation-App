package api

import (
	"awesomeProject1/donation"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Donate(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"bankdetails-id":"65bbbe64a29af2768a9009cb",
    	"orphanage-id":["65ba7800ab0265f4fa9d4b60","65c07ee1dfda391f6fe449a6"],
    	"sum":18
	}`)
	request, err := http.NewRequest("POST", "/addDonate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	donation.AddDonate(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Contains(t, response.Body.String(), "\"Success\"")
}
