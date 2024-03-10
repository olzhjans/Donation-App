package api

import (
	"awesomeProject1/donationsubscribe"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_SubscribeToDonation(t *testing.T) {
	response := httptest.NewRecorder()
	requestBody := []byte(`{
    	"bank-details":{
        	"name":"LOL LOL",
        	"expiring":"01.28",
        	"cvv":"777",
        	"cardnumber":"1234123412341234",
        	"userid":"65c1da85b9683cf7113767cf",
        	"bill":5000
    	},
    	"orphanageid":["65ba7800ab0265f4fa9d4b60"],
    	"amount":5000,
    	"whichday":32,
    	"isactive":false
	}`)
	request, err := http.NewRequest("POST", "/addDonationSubscribe", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	donationsubscribe.SubscribeToDonation(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)
	expectedBody := "\"Success\"\n"
	assert.Equal(t, expectedBody, response.Body.String())
}
