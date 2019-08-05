package cryptoguess

import (
	"encoding/pem"
	"fmt"
)

func init() {
	AvailableExperiments = append(AvailableExperiments, NewPEMBlock)
}

type PEMBlock struct{ *base }

func NewPEMBlock(input []byte) Experiment {
	return &PEMBlock{
		base: &base{
			input: input,
			name:  "PEM encoded data",
			run: func(input []byte, base *base) error {
				block, rest := pem.Decode(input)
				base.rest = rest
				if block == nil {
					return fmt.Errorf("no PEM data found")
				}
				base.result = block
				// FIXME: recursively call other parsers with prefix=pem-encoded (like multiaddr)
				return nil
			},
		},
	}
}
