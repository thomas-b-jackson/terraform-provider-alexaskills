module github.com/scg/terraform-provider-alexaskills

go 1.15

require (
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/hashicorp/terraform-plugin-docs v0.5.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.20.0
	github.com/posener/complete v1.2.1 // indirect
	github.com/scg/va/smapi_client v0.0.0
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
)

replace github.com/scg/va/smapi_client => ./smapi_client
