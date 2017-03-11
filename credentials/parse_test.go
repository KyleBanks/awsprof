package credentials

import (
	"testing"
)

func Test_parse(t *testing.T) {
	sample := `
[default]
aws_access_key_id = ACCESS
aws_secret_access_key = SECRET
[profile2]
aws_access_key_id = ACCESS2
aws_secret_access_key = SECRET2
`
	p, err := parse(sample)
	if err != nil {
		t.Fatal(err)
	}

	if len(p) != 2 {
		t.Fatalf("Unexpected number of Profiles returned, expected=2, got=%v", len(p))
	}

	if p[0].Name != "default" || p[0].AccessKey != "ACCESS" || p[0].SecretKey != "SECRET" {
		t.Fatalf("Unexpected Profile at index 0, expected=%v, got=%v", "default ACCESS SECRET", p[0].Name+" "+p[0].AccessKey+" "+p[0].SecretKey)
	} else if p[1].Name != "profile2" || p[1].AccessKey != "ACCESS2" || p[1].SecretKey != "SECRET2" {
		t.Fatalf("Unexpected Profile at index 0, expected=%v, got=%v", "profile2 ACCESS2 SECRET2", p[1].Name+" "+p[1].AccessKey+" "+p[1].SecretKey)
	}
}
