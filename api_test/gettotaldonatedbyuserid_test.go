package api

import (
	"awesomeProject1/donation"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetTotalDonatedByUserIdAndPeriod(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"id":"65c1da85b9683cf7113767cf",
    	"from":"2024-01-01T00:00:00Z",
    	"to":"2024-12-31T00:00:00Z"
	}`)
	request, err := http.NewRequest("POST", "/getTotalDonatedByUserIdAndPeriod", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	donation.GetTotalDonatedByUserIdAndPeriod(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	//assert.Contains(t, response.Body.String(), "816")
}
