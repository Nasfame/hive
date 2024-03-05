package utils_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CoopHive/hive/utils"
)

func TestGetPublicIP(t *testing.T) {
	// Create a mock HTTP server that always returns an error response
	/*	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Mock error", http.StatusInternalServerError)
		}))
		defer mockServer.Close()*/

	// Set the mock server URL as the endpoint for getPublicIP

	// Test if getPublicIP panics as expected
	/*	defer func() {
		if r := recover(); r != nil {
			t.Error("Expected getPublicIP to panic, but it did not")
		}
	}()*/
	ipString := utils.GetPublicIP()
	t.Logf("ip:%v", ipString)
	assert.NotEmpty(t, ipString, "ip is empty")

	ip := net.ParseIP(ipString)
	assert.NotNil(t, ip, "ip is nil")
	assert.Equal(t, ip.To4() != nil, true, "ip is not ipv4") // check if ipv4
	assert.Equal(t, ip.To16() != nil, true, "ip to ipv6 failed")

}
