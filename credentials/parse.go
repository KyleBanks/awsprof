package credentials

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	credentialsFile = ".aws/credentials"

	accessKeyPrefix = "aws_access_key_id="
	secretKeyPrefix = "aws_secret_access_key="
	namePrefix      = "["
	nameSuffix      = "]"
)

// Parse reads and parses the default aws credentials file, and returns the profiles it contains.
func Parse() ([]Profile, error) {
	c, err := readCredentials()
	if err != nil {
		return nil, err
	}

	return parse(c)
}

func readCredentials() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadFile(filepath.Join(usr.HomeDir, credentialsFile))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func parse(s string) ([]Profile, error) {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	var profiles []Profile

	for i := 0; i < len(lines)-2; i += 3 {
		var p Profile

		p.Name = lines[i]
		toKey(lines[i+1], &p)
		toKey(lines[i+2], &p)

		if err := p.validate(); err != nil {
			return nil, err
		}

		profiles = append(profiles, p)
	}

	return profiles, nil
}

func toKey(s string, p *Profile) {
	s = strings.Replace(s, " ", "", -1)

	if strings.HasPrefix(s, accessKeyPrefix) {
		p.AccessKey = strings.Replace(s, accessKeyPrefix, "", 1)
	} else if strings.HasPrefix(s, secretKeyPrefix) {
		p.SecretKey = strings.Replace(s, secretKeyPrefix, "", 1)
	}
}
