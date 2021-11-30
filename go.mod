module github.com/scg/terraform-provider-alexaskills

go 1.15

require (
	github.com/hashicorp/terraform-plugin-docs v0.5.0
	github.com/hashicorp/terraform-plugin-sdk v1.17.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/scg/va/ask_client v0.0.0
)

replace github.com/scg/va/ask_client => ./ask_client
