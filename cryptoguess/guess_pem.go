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
	data, rest := pem.Decode(exp.Input())
	if data == nil {
		result := &baseResult{exp: exp, rest: rest, err: fmt.Errorf("no PEM data found")}
		results = append(results, result)
	} else {
		// FIXME: data.Type
		// FIXME: data.Headers
		result := &baseResult{exp: exp, rest: rest, data: data}
		for _, childExperiment := range New(data.Bytes).Experiments {
			for _, childResult := range childExperiment.Results() {
				results = append(
					results,
					&stackedResult{
						parent: result,
						child:  childResult,
					},
				)
			}
		}
		results = append(results, result)
	}
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
