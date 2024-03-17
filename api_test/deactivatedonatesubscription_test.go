package api

import (
	"awesomeProject1/adminrights/donationsubscribe"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_DeactivateDonateSubscription(t *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/deactivateDonateSubscription?_id=65e1717668e3bc279918be6a", nil)
	if err != nil {
		t.Fatal(err)
	}
	donationsubscribe.DeactivateDonateSubscription(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, response.Body.String(), "\"Successfully deactivated\"\n")
}
