package cryptoguess

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func init() {
	AvailableExperiments = append(AvailableExperiments, NewX509PKIXPublicKey)
}

type X509PKIXPublicKey struct{ *base }

func NewX509PKIXPublicKey(input []byte) Experiment {
	return &X509PKIXPublicKey{
		base: &base{
			input: input,
			name:  "x509: DER encoded public key",
			run: func(input []byte, base *base) error {
				// FIXME: input should be already decoded by pem
				block, _ := pem.Decode(input)
				if block == nil {
					return fmt.Errorf("no PEM formatted block found")
				}
				ret, err := x509.ParsePKIXPublicKey(block.Bytes)
				if err != nil {
					return err
				}
				base.result = ret
				// FIXME: return specific type (RSA or equivalent)
				return nil
			},
		},
	}
}
