package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/pion/dtls/v2"
)

// NewDefaultConfig creates a new default configuration.  This requires you to have
// a .devcli/certs directory under your home directory where the cert.crt and key.pem
// are stored.  This function is a bit ugly since it terminates if anything goes wrong,
// so it is only useful in clients where this behavior is acceptable.  If you want
// to be able to do proper error handling please see the NewConfig function.
func NewDefaultConfig() Config {
	userHomedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("can't figure out user home directory: %v", err)
	}
	certsDir := path.Join(userHomedir, certsDirFragment)

	certBytes, err := os.ReadFile(certsDir + "/" + certFile)
	if err != nil {
		log.Fatalf("Error reading cert file: %v", err)
	}

	keyBytes, err := os.ReadFile(certsDir + "/" + keyFile)
	if err != nil {
		log.Fatalf("Error reading key file: %v", err)
	}

	cfg, err := NewConfig(certBytes, keyBytes)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

// NewConfig creates a new configuration from certBytes and keyBytes. Note that both
// certBytes and keyBytes are expected to be PEM-encoded.
func NewConfig(certBytes []byte, keyBytes []byte) (Config, error) {
	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return Config{}, fmt.Errorf("%w: %v", ErrCannotReadKeyPair, err)
	}

	intermediates, roots, err := loadCertPool(certBytes)
	if err != nil {
		return Config{}, fmt.Errorf("%w: %v", ErrCannotLoadCertPool, err)
	}

	return Config{
		DTLSConfig: &dtls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: false,
			RootCAs:            roots,
			ClientCAs:          intermediates,
			ConnectContextMaker: func() (context.Context, func()) {
				return context.WithTimeout(context.Background(), defaultDTLSTimeout)
			},
		},
		SpanAddr: defaultSpanAddr,
	}, nil
}
