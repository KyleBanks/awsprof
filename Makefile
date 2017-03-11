VERSION = 1.0.0

RELEASE_PKG = ./cmd/awsprof
RELEASE_PLATFORMS = darwin/386 darwin/amd64 linux/386 linux/amd64 linux/arm
INSTALL_PKG = $(RELEASE_PKG)

# Remote includes require 'mmake' 
# github.com/tj/mmake
include github.com/KyleBanks/make/go/install
include github.com/KyleBanks/make/go/sanity
include github.com/KyleBanks/make/go/release
include github.com/KyleBanks/make/git/precommit

# Example demonstrates listing profiles, and setting one.
example: | install
	@awsprof 
	@awsprof default
	@eval $$(awsprof default)
.PHONY: example
