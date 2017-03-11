// Package credentials provides the ability to parse and represent an AWS credentials file.
package credentials

import (
	"fmt"
	"strings"
)

// Profile represents an AWS profile.
type Profile struct {
	Name      string
	AccessKey string
	SecretKey string
}

// IsActive returns true of this profile matches the current access/secret key environment variables.
func (p *Profile) IsActive() bool {
	a, s := Env()

	return p.AccessKey == a && p.SecretKey == s
}

func (p *Profile) validate() error {
	if !strings.HasPrefix(p.Name, namePrefix) || !strings.HasSuffix(p.Name, nameSuffix) {
		return fmt.Errorf("expected profile name, got %v", p.Name)
	}
	p.Name = p.Name[1 : len(p.Name)-1]

	if len(p.AccessKey) == 0 {
		return fmt.Errorf("failed to parse access key for %v", p.Name)
	}
	if len(p.SecretKey) == 0 {
		return fmt.Errorf("failed to parse secret key for %v", p.Name)
	}

	return nil
}

func (p *Profile) String() string {
	return p.Name
}
