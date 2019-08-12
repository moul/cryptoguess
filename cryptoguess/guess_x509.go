package cryptoguess

import "crypto/x509"

func init() {
	AvailableExperiments = append(AvailableExperiments, NewX509PKIXPublicKey)
	AvailableExperiments = append(AvailableExperiments, NewX509PKCS8PrivateKey)
	AvailableExperiments = append(AvailableExperiments, NewX509PKCS1PublicKey)
	AvailableExperiments = append(AvailableExperiments, NewX509PKCS1PrivateKey)
	AvailableExperiments = append(AvailableExperiments, NewX509ECPrivateKey)
	AvailableExperiments = append(AvailableExperiments, NewX509DERCRL)
	AvailableExperiments = append(AvailableExperiments, NewX509Certificate)
	AvailableExperiments = append(AvailableExperiments, NewX509Certificates)
	AvailableExperiments = append(AvailableExperiments, NewX509CertificateRequest)
}

//
// PKIXPublicKey
//

type X509PKIXPublicKey struct{ *baseExperiment }

func runX509PKIXPublicKey(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParsePKIXPublicKey(exp.Input())
	// FIXME: name: RSA/DA/ECDSA/...
	return []Result{result}
}

func NewX509PKIXPublicKey(input []byte) Experiment {
	return &X509PKIXPublicKey{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: DER encoded public key",
			run:   runX509PKIXPublicKey,
		},
	}
}

//
// PKCS8PrivateKey
//

type X509PKCS8PrivateKey struct{ *baseExperiment }

func runX509PKCS8PrivateKey(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParsePKCS8PrivateKey(exp.Input())
	return []Result{result}
}

func NewX509PKCS8PrivateKey(input []byte) Experiment {
	return &X509PKCS8PrivateKey{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: unencrypted PKCS#8 private key",
			run:   runX509PKCS8PrivateKey,
		},
	}
}

//
// PKCS1PublicKey
//

type X509PKCS1PublicKey struct{ *baseExperiment }

func runX509PKCS1PublicKey(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParsePKCS1PublicKey(exp.Input())
	return []Result{result}
}

func NewX509PKCS1PublicKey(input []byte) Experiment {
	return &X509PKCS1PublicKey{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: PKCS#1 public key (RSA) in ASN.1 DER form",
			run:   runX509PKCS1PublicKey,
		},
	}
}

//
// PKCS1PrivateKey
//

type X509PKCS1PrivateKey struct{ *baseExperiment }

func runX509PKCS1PrivateKey(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParsePKCS1PrivateKey(exp.Input())
	return []Result{result}
}

func NewX509PKCS1PrivateKey(input []byte) Experiment {
	return &X509PKCS1PrivateKey{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: PKCS#1 private key (RSA) in ASN.1 DER form",
			run:   runX509PKCS1PrivateKey,
		},
	}
}

//
// ECPrivateKey
//

type X509ECPrivateKey struct{ *baseExperiment }

func runX509ECPrivateKey(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParseECPrivateKey(exp.Input())
	return []Result{result}
}

func NewX509ECPrivateKey(input []byte) Experiment {
	return &X509ECPrivateKey{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: ASN.1 Elliptic Curve private key",
			run:   runX509ECPrivateKey,
		},
	}
}

//
// DERCRL
//

type X509DERCRL struct{ *baseExperiment }

func runX509DERCRL(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParseDERCRL(exp.Input())
	return []Result{result}
}

func NewX509DERCRL(input []byte) Experiment {
	return &X509DERCRL{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: DER encoded CRL (pkix certificate list)",
			run:   runX509DERCRL,
		},
	}
}

//
// CRL
//

type X509CRL struct{ *baseExperiment }

func runX509CRL(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParseCRL(exp.Input())
	return []Result{result}
}

func NewX509CRL(input []byte) Experiment {
	return &X509CRL{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: CRL (pkix certificate list)",
			run:   runX509CRL,
		},
	}
}

//
// Certificate
//

type X509Certificate struct{ *baseExperiment }

func runX509Certificate(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParseCertificate(exp.Input())
	return []Result{result}
}

func NewX509Certificate(input []byte) Experiment {
	return &X509Certificate{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: ASN.1 DER certificate",
			run:   runX509Certificate,
		},
	}
}

//
// Certificates
//

type X509Certificates struct{ *baseExperiment }

func runX509Certificates(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParseCertificates(exp.Input())
	return []Result{result}
}

func NewX509Certificates(input []byte) Experiment {
	return &X509Certificates{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: ASN.1 DER certificate",
			run:   runX509Certificates,
		},
	}
}

//
// CertificateRequest
//

type X509CertificateRequest struct{ *baseExperiment }

func runX509CertificateRequest(exp Experiment) []Result {
	result := &baseResult{exp: exp}
	result.data, result.err = x509.ParseCertificateRequest(exp.Input())
	return []Result{result}
}

func NewX509CertificateRequest(input []byte) Experiment {
	return &X509CertificateRequest{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "x509: ASN.1 DER certificate request",
			run:   runX509CertificateRequest,
		},
	}
}
