package api

import (
	"awesomeProject1/needs"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetNeedsByRegionAndType(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"region":"Almaty",
    	"category-of-donate":"Clothes"
	}`)
	request, err := http.NewRequest("POST", "/getNeedsByRegionAndType", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	needs.GetNeedsByRegionAndType(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Contains(t, response.Body.String(), "\"_id\":\"65ba75f0ab0265f4fa9d4b5f\",")
}
