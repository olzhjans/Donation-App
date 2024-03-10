package api

import (
	"awesomeProject1/donation"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetTotalDonatedByOrphanageIdAndPeriod(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"id":"65c07ee1dfda391f6fe449a6",
    	"from":"2024-01-01T00:00:00Z",
    	"to":"2024-12-31T00:00:00Z"
	}`)
	request, err := http.NewRequest("POST", "/getTotalDonatedByOrphanageIdAndPeriod", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	donation.GetTotalDonatedByOrphanageIdAndPeriod(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Contains(t, response.Body.String(), "173")
}
