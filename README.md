# Terraform Provider Alexa Skills

Provider for building alexa skills via terraform

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.15
- [ask cli]

## Building The Provider

1. Clone the repository
2. Run the `install` target as: 
```sh
$ make install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

See [./examples](./examples)

## Developing the Provider

Create a token as:

`ask util generate-lwa-tokens --no-browser`

Save the token to an environment variable as:

`export SMAPI_TOKEN=<token>`

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (per [Requirements](#requirements) above).

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources (and then clean them up after). So you'll need credentials to https://developer.amazon.com/ in order to run them.

```sh
$ make test
```

## Release

Create a release using GoReleaser. 

**Note:** steps are adapted from [these instructions](https://www.terraform.io/docs/registry/providers/publishing.html#using-goreleaser-locally)

Setup Steps:
* Install GoReleaser
* Obtain fingerprint of GPG private key for signing (key currently controlled by Tom Jackson)
  * fingerprint is 40 chars and is obtained by running this command:
    `gpg --list-secret-keys --keyid-format=long`
* Obtain Personal Access Token for repo (token currently controlled by Tom Jackson)

Release Steps:
* Commit changes locally
* Set GITHUB_TOKEN to a Personal Access Token
* Set GPG_FINGERPRINT to fingerprint
* Tag your release commit to match version in GNUmakefile, e.g.:
  `git tag 0.2.0`
* Build, sign, and upload your release with:
  `goreleaser release --rm-dist`
* Re-run terraform init against the release in the registry (to make sure it has sync'd from github)
* Test the released provider in a pipeline
* Push commit and do pull request
