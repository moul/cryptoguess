package cryptoguess

import (
	"encoding/pem"
	"fmt"
)

func init() {
	AvailableExperiments = append(AvailableExperiments, NewPEMBlock)
}

type PEMBlock struct{ *baseExperiment }

func runPEMBlock(exp Experiment) []Result {
	results := []Result{}
	result := &baseResult{exp: exp}
	result.data, result.rest = pem.Decode(exp.Input())
	if result.data == nil {
		result.err = fmt.Errorf("no PEM data found")
	}
	results = append(results, result)
	// FIXME: recursively call other parsers with prefix=pem-encoded (like multiaddr)
	return results
}

func NewPEMBlock(input []byte) Experiment {
	return &PEMBlock{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "PEM encoded data",
			run:   runPEMBlock,
		},
	}
}
