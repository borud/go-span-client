package client

import (
	"github.com/plgd-dev/go-coap/v2/dtls"
	"github.com/plgd-dev/go-coap/v2/udp/client"
)

// COAPConnect creates a CoAP connection.
func COAPConnect(config Config) (*client.ClientConn, error) {
	return dtls.Dial(config.SpanCOAPAddr, config.DTLSConfig)
}
