# A pulumi provider for Honeycomb.io

This provider is based on the [honeycomb terraform
provider](https://github.com/honeycombio/terraform-provider-honeycombio),
and is currently in an experimental state.

It is created using the [Pulumi Terraform
Bridge](https://github.com/pulumi/pulumi-terraform-bridge), and has standard
project layout:

*   `Makefile` builds binaries:
    *   `bin/pulumi-tfgen-honeycomb` is the design-time code generator, and
    *   `bin/pulumi-resource-honeycomb` is the runtime plugin linking to the
        Terraform provider.
*   `provider/resources.go` maps all the Terraform Provider metadata for use.
*   `sdk` is generated code of the provider; the `sdk/go` directory is
    generated but checked in.
*   `examples/` example usage of `pulumi_honeycomb`.

## Building

`make python-sdk VERSION=0.0.8` (yes, the old version; alternatively, any other SDK that you use)
`make bin/pulumi-resource-honeycomb`

Then in your pulumi code base:

install the SDK from the local directory (using `pip install -e path/to/pulumi-honeycomb/sdk/python/`)
`PATH=/path/to/bin/pulumi-resource-honeycomb:$PATH pulumi up`
