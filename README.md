# Terraform Provider Alexa Skills

Provider for building alexa skills via terraform

## Requirements

- wsl or mac
-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.15

## Using the provider

See [./examples](./examples)

## Developing the Provider

Create a token as:

1. log into the `Sempra Energy` account in developer.amazon.com
2. generate a token as:
   `ask util generate-lwa-tokens --no-browser`

    > The ask utility will output a json-encoded string to stdout containing `access_token` and `refresh_token` fields.

3. Save the `access_token` to an environment variable as:
    `export LWA_ACCESS_TOKEN=<token>`
4. make changes to provider sources
5. build and install the provider locally by running:
   `make install`
6. test the provider against examples in [./examples](./examples) as:
   1. reference the `localhost/va/awslex` version of the provider
   2. remove `.terraform.lock.hcl` between `make install` iterations

## Release

Steps for publishing a new version to the private registry in Terraform cloud:

2. develop and test changes to the provider per [Developing the Provider](#developing-the-provider)
3. verify all binaries can be built by running the `release` target as:
   `make release`
4. push changes to a feature branch
5. verify provider pipeline is green
6. submit pull request to review provider feature branch (but do NOT merge yet)
7. contact IaC and request they publish a new version. provide them:
   1. name of the feature branch
   2. the desired version tag
8.  update chatbot repo with new version tag on a feature branch
9.  verify chatbot pipeline is green
10. merge provider feature branch

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.