package cryptoguess

import "crypto/x509"

func init() {
	AvailableExperiments = append(AvailableExperiments, NewX509PKIXPublicKey)
}

type X509PKIXPublicKey struct{ *baseExperiment }

func runX509PKIXPublicKey(exp Experiment) []Result {
	results := []Result{}
	// FIXME: input should be already decoded by pem
	ret, err := x509.ParsePKIXPublicKey(exp.Input())
	result := &baseResult{
		exp:  exp,
		data: ret,
		err:  err,
		// FIXME: name: RSA/ECDSA/...
	}
	results = append(results, result)
	return results
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
