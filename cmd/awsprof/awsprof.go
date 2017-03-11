package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/KyleBanks/awsprof/credentials"
)

const (
	activeIndicator = " * "
)

func main() {
	// Parse AWS Credentials
	profiles, err := credentials.Parse()
	if err != nil {
		fmt.Println("parsing credentials:", err)
		os.Exit(1)
	}

	// Check the command
	if len(os.Args) == 1 {
		list(profiles)
	} else {
		set(profiles, os.Args[1])
	}
}

func list(p []credentials.Profile) {
	sort.Slice(p, func(i, j int) bool {
		return p[i].Name < p[j].Name
	})

	for _, prof := range p {
		prefix := "   "
		if prof.IsActive() {
			prefix = activeIndicator
		}

		fmt.Printf("%v%v\n", prefix, prof.String())
	}
}

func set(p []credentials.Profile, name string) {
	for _, prof := range p {
		if prof.Name != name {
			continue
		}

		fmt.Printf("export %v='%v'\nexport %v='%v'\n", credentials.AccessKeyEnv, prof.AccessKey, credentials.SecretKeyEnv, prof.SecretKey)
		return
	}

	fmt.Println("profile not found:", name)
	os.Exit(1)
}
