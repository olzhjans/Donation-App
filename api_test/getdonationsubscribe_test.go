package api

import (
	"awesomeProject1/donationsubscribe"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetDonationSubscribeByUserId(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/getDonationSubscribeByUserId?userid=65c1da85b9683cf7113767cf", nil)
	if err != nil {
		t.Fatal(err)
	}
	donationsubscribe.GetDonationSubscribeByUserId(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Contains(t, response.Body.String(), "\"_id\":\"65e18ce8b2c0a7eabd1af1c4\",")
}
