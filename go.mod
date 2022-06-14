module github.com/scg/terraform-provider-alexaskills

go 1.15

require (
	github.com/hashicorp/hcl/v2 v2.8.2 // indirect
	github.com/hashicorp/terraform-plugin-docs v0.10.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/scg/va/smapi_client v0.0.0
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
)

replace github.com/scg/va/smapi_client => ./smapi_client
