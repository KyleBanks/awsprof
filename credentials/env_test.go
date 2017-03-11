package credentials

import (
	"os"
	"testing"
)

func Test_Env(t *testing.T) {
	access := "test access"
	secret := "test secret"
	os.Setenv(AccessKeyEnv, access)
	os.Setenv(SecretKeyEnv, secret)

	outAccess, outSecret := Env()
	if outAccess != access {
		t.Fatalf("Unexpected Access Key returned, expected=%v, got=%v", access, outAccess)
	} else if outSecret != secret {
		t.Fatalf("Unexpected Secret Key returned, expected=%v, got=%v", secret, outSecret)
	}
}
