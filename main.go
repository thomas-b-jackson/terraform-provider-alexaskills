package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/scg/terraform-provider-alexaskills/internal/provider"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.Provider(version)
		},
	})
}
