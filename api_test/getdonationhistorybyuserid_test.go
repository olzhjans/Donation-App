package api

import (
	"awesomeProject1/donation"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetDonationHistoryByUserId(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/getDonationSubscribeByUserId?userid=65c1da85b9683cf7113767cf", nil)
	if err != nil {
		t.Fatal(err)
	}
	donation.GetDonationHistoryByUserId(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Contains(t, response.Body.String(), "\"_id\":\"65e026c0a2585a9512507cd0\",")
}
