package api

import (
	"awesomeProject1/adminrights/needs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ActivateNeed(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/activateNeed?needid=65ba75f0ab0265f4fa9d4b5f", nil)
	if err != nil {
		t.Fatal(err)
	}
	needs.ActivateNeedByNeedId(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, response.Body.String(), "\"Successfully activated\"\n")
}
