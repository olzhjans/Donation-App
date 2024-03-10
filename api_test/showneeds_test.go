package api

import (
	"awesomeProject1/needs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ShowNeeds(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/showNeeds?orphanageid=65ba7800ab0265f4fa9d4b60", nil)
	if err != nil {
		t.Fatal(err)
	}
	needs.ShowNeeds(response, request)
	assert.Equal(t, http.StatusFound, response.Code)
	assert.Contains(t, response.Body.String(), "\"_id\":\"65ba75f0ab0265f4fa9d4b5f\",")
}
