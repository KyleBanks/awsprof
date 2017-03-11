package credentials

import (
	"os"
)

const (
	// AccessKeyEnv is the name of the AWS access key environment variable.
	AccessKeyEnv = "AWS_ACCESS_KEY_ID"

	// SecretKeyEnv is the name of the AWS secret key environment variable.
	SecretKeyEnv = "AWS_SECRET_ACCESS_KEY"
)

// Env returns the current access and secret key evironment variables.
func Env() (string, string) {
	return os.Getenv(AccessKeyEnv), os.Getenv(SecretKeyEnv)
}
