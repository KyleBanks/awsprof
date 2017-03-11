package credentials

import (
	"os"
	"testing"
)

func TestProfile_IsActive(t *testing.T) {
	p := Profile{
		AccessKey: "accesskey",
		SecretKey: "secretkey",
	}

	os.Setenv(AccessKeyEnv, "incorrect")
	os.Setenv(SecretKeyEnv, "incorrect")

	if p.IsActive() {
		t.Fatal("Expected IsActive = false")
	}

	os.Setenv(AccessKeyEnv, p.AccessKey)
	os.Setenv(SecretKeyEnv, p.SecretKey)

	if !p.IsActive() {
		t.Fatal("Expected IsActive = true")
	}
}

// this is important!
// don't want to print the access/secret key when it is not necessary!
func TestProfile_String(t *testing.T) {
	p := Profile{
		Name:      "profile name",
		AccessKey: "access",
		SecretKey: "secret",
	}

	if p.String() != p.Name {
		t.Fatalf("Unexpected String(), expected=%v, got=%v", p.Name, p.String())
	}
}
