package cryptoguess

import (
	"encoding/base64"
)

func init() {
	AvailableExperiments = append(AvailableExperiments, NewBASE64Block)
}

type BASE64Block struct{ *baseExperiment }

func runBASE64Block(exp Experiment) []Result {
	results := []Result{}
	result := &baseResult{exp: exp}
	result.data, result.err = base64.StdEncoding.DecodeString(string(exp.Input()))
	results = append(results, result)
	// FIXME: recursively call other parsers with prefix=base64-encoded (like multiaddr)
	return results
}

func NewBASE64Block(input []byte) Experiment {
	return &BASE64Block{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "BASE64 encoded data",
			run:   runBASE64Block,
		},
	}
}
