package api

import (
	"awesomeProject1/adminrights/adminauth"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ShowWaitingList(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/showWaitingList", nil)
	if err != nil {
		t.Fatal(err)
	}
	adminauth.ShowWaitingList(response, request)
	assert.Equal(t, http.StatusFound, response.Code)
	assert.Contains(t, response.Body.String(), "_id")
}
