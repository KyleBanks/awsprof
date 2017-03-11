# awsprof

[![GoDoc](https://godoc.org/github.com/KyleBanks/awsprof?status.svg)](https://godoc.org/github.com/KyleBanks/awsprof)&nbsp; 
[![Build Status](https://travis-ci.org/KyleBanks/awsprof.svg?branch=master)](https://travis-ci.org/KyleBanks/awsprof)&nbsp;
[![Go Report Card](https://goreportcard.com/badge/github.com/KyleBanks/awsprof)](https://goreportcard.com/report/github.com/KyleBanks/awsprof)


`awsprof` is a little tool to quickly switch your AWS access and secret key environment variables using profile names. Many AWS tools and APIs support the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables, but they don't support the profile system - `awsprof` aims to help resolve that.

## Install

Download a binary from the [Releases](https://github.com/KyleBanks/awsprof/releases) page, or if you have a valid Go installation you can install from source:

```
$ go get github.com/KyleBanks/awsprof/cmd/awsprof
```

## Usage

`awsprof` can be run with either no arguments to view the names of all configured profiles. If one of the profiles matches the current `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables, it will be marked as the active profile with an asterisk:

```sh
$ awsprof
 * default
   devops
   production
   website
```

Alternatively, running `awsprof` with a profile name allows you to activate the pair of `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables associated with that profile. Unfortunately, due to the nature of environment variables, and the fact that a child process cannot set environment variables for the parent, `awsprof` can only output the export commands which you can then run:

```sh
$ awsprof website
export AWS_ACCESS_KEY_ID='EXAMPLE'
export AWS_SECRET_ACCESS_KEY='EXAMPLE'
```

You can either copy and paste the command to run it yourself, or more efficiently:

```sh
$ eval $(awsprof website)
```

For frequent usage, add an alias to your `~/.bash_profile`:

```sh
# ~/.bash_profile

alias awsprof-website="eval \$(awsprof website)"
```

And run it like so:

```sh
$ awsprof-website
```

## Author

`awsprof` was developed by [Kyle Banks](https://twitter.com/kylewbanks).

## License

`awsprof` is available under the [MIT](./LICENSE) license.