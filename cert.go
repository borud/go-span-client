package client

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
)

// loadCertPool parses and loads certificates from a byte buffer into two pools;
// one for intermediates and one for root certificates.
func loadCertPool(pemBytes []byte) (*x509.CertPool, *x509.CertPool, error) {
	var certs []*x509.Certificate

	block, remain := pem.Decode(pemBytes)
	for block != nil {
		if block.Type == "CERTIFICATE" {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err == nil {
				certs = append(certs, cert)
			}
		}
		block, remain = pem.Decode(remain)
	}
	roots := x509.NewCertPool()
	intermediates := x509.NewCertPool()

	for _, crt := range certs {
		if isIntermediateCertificate(crt) {
			intermediates.AddCert(crt)
			continue
		}
		if isRootCA(crt) {
			roots.AddCert(crt)
			continue
		}
	}
	return intermediates, roots, nil
}

// Intermediate certificates can sign new certificates (the IsCA flag is set) but the
// issuer is different from the certificate
func isIntermediateCertificate(c *x509.Certificate) bool {
	return c.IsCA && !bytes.Equal(c.RawIssuer, c.RawSubject)
}

// Root CAs have the same issuer and subject fields
func isRootCA(c *x509.Certificate) bool {
	return c.IsCA && bytes.Equal(c.RawIssuer, c.RawSubject)
}
