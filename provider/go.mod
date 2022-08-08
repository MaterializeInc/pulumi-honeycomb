module github.com/MaterializeInc/pulumi-honeycomb/provider

go 1.16

require github.com/honeycombio/terraform-provider-honeycombio v0.8.0

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-exec => github.com/hashicorp/terraform-exec v0.17.2
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20220725190814-23001ad6ec03
)

require github.com/pulumi/pulumi-terraform-bridge/v3 v3.26.1
