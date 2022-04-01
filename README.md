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