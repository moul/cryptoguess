package cryptoguess

import "golang.org/x/crypto/ssh"

func init() {
	AvailableExperiments = append(AvailableExperiments, NewSSHAuthorizedKey)
}

type SSHAuthorizedKey struct{ *baseExperiment }

type ParsedSSHAuthorizedKey struct {
	PublicKey ssh.PublicKey
	Comment   string
	Options   []string
}

func runSSHAuthorizedKey(exp Experiment) []Result {
	results := []Result{}
	pubkey, comment, options, _, err := ssh.ParseAuthorizedKey(exp.Input())
	result := &baseResult{
		exp: exp,
		data: &ParsedSSHAuthorizedKey{
			PublicKey: pubkey,
			Comment:   comment,
			Options:   options,
		},
		err: err,
	}
	results = append(results, result)
	// FIXME: return more info: RSA, length, etc
	return results
}

func NewSSHAuthorizedKey(input []byte) Experiment {
	return &SSHAuthorizedKey{
		baseExperiment: &baseExperiment{
			input: input,
			name:  "SSH authorized key",
			run:   runSSHAuthorizedKey,
		},
	}
}
