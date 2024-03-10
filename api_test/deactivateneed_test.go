package api

import (
	"awesomeProject1/needs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_DeactivateNeed(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/activateNeed?needid=65ba75f0ab0265f4fa9d4b5f", nil)
	if err != nil {
		t.Fatal(err)
	}
	needs.DeactivateNeedByNeedId(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, response.Body.String(), "\"Successfully deactivated\"\n")
}
