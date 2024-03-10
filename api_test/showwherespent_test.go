package api

import (
	"awesomeProject1/wherespent"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ShowWhereSpent(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"orphanage-id":"65c07ee1dfda391f6fe449a6",
    	"from":"2024-01-23T00:00:00Z",
    	"to":"2024-02-24T00:00:00Z"
	}`)
	request, err := http.NewRequest("POST", "/showWhereSpent", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	wherespent.ShowWhereSpent(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Contains(t, response.Body.String(), "\"_id\":\"65c0849ddfda391f6fe449ad\",")
}
