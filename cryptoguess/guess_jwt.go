package cryptoguess

import jwt "gopkg.in/square/go-jose.v2/jwt"

func init() {
	AvailableExperiments = append(AvailableExperiments, NewJWTSignedToken)
}

type JWTSignedToken struct{ *baseExperiment }

type ParsedJWTSignedToken struct {
	Token  *jwt.JSONWebToken
	Claims map[string]interface{}
}

func runJWTSignedToken(exp Experiment) []Result {
	results := []Result{}
	parsed := &ParsedJWTSignedToken{}
	result := &baseResult{
		exp:  exp,
		data: parsed,
		err:  nil,
	}
	parsed.Token, result.err = jwt.ParseSigned(string(exp.Input()))
	if result.err == nil {
		result.err = parsed.Token.UnsafeClaimsWithoutVerification(&parsed.Claims)
	}
	results = append(results, result)
	return results
}

func NewJWTSignedToken(input []byte) Experiment {
	return &JWTSignedToken{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "JWT signed token",
			run:   runJWTSignedToken,
		},
	}
}
