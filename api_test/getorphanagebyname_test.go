package api

import (
	"awesomeProject1/adminrights/orphanage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetOrphanage(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/getOrphanage?name=Umit", nil)
	if err != nil {
		t.Fatal(err)
	}
	orphanage.GetOrphanageByName(response, request)
	assert.Equal(t, http.StatusFound, response.Code)
	assert.Contains(t, response.Body.String(), "\"_id\":\"65ebfa792a10e66ce0a6a8f5\"")
}
