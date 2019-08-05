package cryptoguess

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type X509PKIXPublicKey struct{ *base }

func NewX509PKIXPublicKey(input []byte) Experiment {
	return &X509PKIXPublicKey{
		base: &base{
			input: input,
			name:  "x509: DER encoded public key",
			run: func(input []byte) (interface{}, error) {
				block, _ := pem.Decode(input)
				if block == nil {
					return nil, fmt.Errorf("no PEM formatted block found")
				}
				return x509.ParsePKIXPublicKey(block.Bytes)
			},
		},
	}
}
