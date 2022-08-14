module github.com/MaterializeInc/pulumi-honeycomb/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20220505215311-795430389fa7
)

require (
	github.com/honeycombio/terraform-provider-honeycombio v0.9.0
	github.com/kr/pretty v0.3.0 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.24.1
)
