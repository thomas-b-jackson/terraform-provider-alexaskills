module github.com/scg/terraform-provider-alexaskills

go 1.15

require (
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/aws/aws-sdk-go v1.37.0 // indirect
	github.com/hashicorp/hcl/v2 v2.8.2 // indirect
	github.com/hashicorp/terraform-plugin-docs v0.5.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/posener/complete v1.2.1 // indirect
	github.com/scg/va/smapi_client v0.0.0
	golang.org/x/tools v0.0.0-20201028111035-eafbe7b904eb // indirect
	google.golang.org/api v0.34.0 // indirect
)

replace github.com/scg/va/smapi_client => ./smapi_client
