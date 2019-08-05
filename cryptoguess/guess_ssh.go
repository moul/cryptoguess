package cryptoguess

import "golang.org/x/crypto/ssh"

func init() {
	AvailableExperiments = append(AvailableExperiments, NewSSHAuthorizedKey)
}

type SSHAuthorizedKey struct{ *base }

type ParsedSSHAuthorizedKey struct {
	PublicKey ssh.PublicKey
	Comment   string
	Options   []string
}

func NewSSHAuthorizedKey(input []byte) Experiment {
	return &SSHAuthorizedKey{
		base: &base{
			input: input,
			name:  "ssh: authorized key",
			run: func(input []byte, base *base) error {
				pubkey, comment, options, _, err := ssh.ParseAuthorizedKey(input)
				if err != nil {
					return err
				}
				base.result = &ParsedSSHAuthorizedKey{
					PublicKey: pubkey,
					Comment:   comment,
					Options:   options,
				}
				// FIXME: return more info: RSA, length, etc
				return nil
			},
		},
	}
}
