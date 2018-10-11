package utils

import (
	"github.com/ninjadotorg/handshake-dispatcher/config"
)

// GetForwardingEndpoint ...
func GetForwardingEndpoint(t string) (string, bool) {
	conf := config.GetConfig()
	var endpoint string
	found := false

	for n, ep := range conf.GetStringMap("forwarding") {
		if n == t {
			endpoint = ep.(string)
			found = true
			break
		}
	}

	return endpoint, found
}

// GetServicesEndpoint ...
func GetServicesEndpoint(t string) (string, bool) {
	conf := config.GetConfig()
	var endpoint string
	found := false

	for n, ep := range conf.GetStringMap("services") {
		if n == t {
			endpoint = ep.(string)
			found = true
			break
		}
	}

	return endpoint, found
}
