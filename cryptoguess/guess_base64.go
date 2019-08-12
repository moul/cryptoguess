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
	data, err := base64.StdEncoding.DecodeString(string(exp.Input()))
	if err != nil {
		result := &baseResult{exp: exp, err: err}
		results = append(results, result)
	} else {
		result := &baseResult{exp: exp, data: data}
		results = append(results, recursiveResults(result, data)...)
		results = append(results, result)
	}
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
