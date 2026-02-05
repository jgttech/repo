package models

// SSHKey tracks an SSH key associated with a profile
type SSHKey struct {
	ID        int
	ProfileID int
	Host      string // "github.com", "gitlab.com"
	HostAlias string // "github-work", "github-personal"
	KeyPath   string // "~/.ssh/id_ed25519_work"
	KeyType   string // "ed25519", "rsa"
}
