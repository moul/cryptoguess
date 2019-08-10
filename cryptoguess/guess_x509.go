package cryptoguess

import "crypto/x509"

func init() {
	AvailableExperiments = append(AvailableExperiments, NewX509PKIXPublicKey)
}

type X509PKIXPublicKey struct{ *baseExperiment }

func runX509PKIXPublicKey(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParsePKIXPublicKey(exp.Input())
	// FIXME: name: RSA/ECDSA/...
	return []Result{result}
}

func NewX509PKIXPublicKey(input []byte) Experiment {
	return &X509PKIXPublicKey{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509 DER encoded public key",
			run:   runX509PKIXPublicKey,
		},
	}
}
