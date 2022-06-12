// Package client implements a DTLS based client for Span
//
// Sample usage:
//     package main
//
//     import (
//         "log"
//         "time"
//         client "github.com/borud/go-span-client"
//     )
//
//     func main() {
//         client, err := client.Connect(client.NewDefaultConfig())
//         if err != nil {
//             log.Fatal(err)
//         }
//         defer client.Close()
//
//         n, err := client.Write([]byte("this is a test"), time.Second)
//         if err != nil {
//             log.Fatal(err)
//         }
//         log.Printf("wrote %d bytes", n)
//
//         buffer := make([]byte, 1024)
//
//         n, err = client.Read(buffer, time.Second*5)
//         if err != nil {
//             log.Fatal(err)
//         }
//         log.Printf("read %d bytes: [%s]", n, string(buffer))
//     }
///
package client

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/pion/dtls/v2"
)

// Client is a Span client that connects as a device.
type Client struct {
	conn   *dtls.Conn
	config Config
}

// Errors for Client
var (
	ErrCannotResolveSpanAddress = errors.New("cannot resolve Span address")
	ErrCannotConnect            = errors.New("cannot connect to Span")
	ErrCannotReadKeyPair        = errors.New("cannot read keypair")
	ErrCannotLoadCertPool       = errors.New("cannot load certpool")
)

// Connect client to Span using DTLS
func Connect(config Config) (*Client, error) {
	addr, err := net.ResolveUDPAddr("udp", config.SpanUDPAddr)
	if err != nil {
		return nil, fmt.Errorf("%w [%s]: %v", ErrCannotResolveSpanAddress, config.SpanUDPAddr, err)
	}

	dtlsConn, err := dtls.Dial("udp", addr, config.DTLSConfig)
	if err != nil {
		return nil, fmt.Errorf("%w [%s]: %v", ErrCannotConnect, config.SpanUDPAddr, err)
	}

	return &Client{
		conn:   dtlsConn,
		config: config,
	}, nil
}

// Write data with optional deadline.
func (c *Client) Write(data []byte, deadline ...time.Duration) (int, error) {
	if len(deadline) != 0 {
		c.conn.SetWriteDeadline(time.Now().Add(deadline[0]))
	}
	return c.conn.Write(data)
}

// Read data with optional deadline.
func (c *Client) Read(buffer []byte, deadline ...time.Duration) (int, error) {
	if len(deadline) != 0 {
		c.conn.SetReadDeadline(time.Now().Add(deadline[0]))
	}
	return c.conn.Read(buffer)
}

// Close the client connection.
func (c *Client) Close() error {
	return c.conn.Close()
}
