package main

import (
	"os"

	"github.com/KyleBanks/awsprof/credentials"
)

func Example_list() {
	os.Setenv(credentials.AccessKeyEnv, "ACTIVE ACCESS")
	os.Setenv(credentials.SecretKeyEnv, "ACTIVE SECRET")

	p := []credentials.Profile{
		credentials.Profile{Name: "default"},
		credentials.Profile{Name: "website", AccessKey: "ACTIVE ACCESS", SecretKey: "ACTIVE SECRET"},
		credentials.Profile{Name: "production"},
		credentials.Profile{Name: "devops"},
	}

	list(p)
	// Output:
	//    default
	//    devops
	//    production
	//  * website
}

func Example_set() {
	p := []credentials.Profile{
		credentials.Profile{Name: "default"},
		credentials.Profile{Name: "website"},
		credentials.Profile{Name: "production"},
		credentials.Profile{Name: "devops", AccessKey: "SAMPLE ACCESS", SecretKey: "SAMPLE SECRET"},
	}

	set(p, "devops")
	// Output:
	// export AWS_ACCESS_KEY_ID='SAMPLE ACCESS'
	// export AWS_SECRET_ACCESS_KEY='SAMPLE SECRET'
}
