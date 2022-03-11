---
title: Honeycomb.io Setup
meta_desc: Information on how to install the Honeycomb.io provider.
layout: installation
---

## Installation

The Pulumi Honeycomb.io provider is available as a package in all Pulumi languages:

// add links to the packages

## Setup

To provision resources with the Pulumi Honeycomb.io provider, you need to have Honeycomb.io credentials.

### Set environment variables

Once you have provisioned these credentials, you can set environment variables to provision resources in Scaleway:

{{< chooser os "linux,macos,windows" >}}
{{% choosable os linux %}}

```bash
$ export HONEYCOMBIO_APIKEY=<HONEYCOMBIO_APIKEY>
```

{{% /choosable %}}

{{% choosable os macos %}}

```bash
$ export HONEYCOMBIO_APIKEY=<HONEYCOMBIO_APIKEY>
```

{{% /choosable %}}

{{% choosable os windows %}}

```powershell
> $env:HONEYCOMBIO_APIKEY = "<HONEYCOMBIO_APIKEY>"
```

{{% /choosable %}}
{{< /chooser >}}

## Configuration Options

Use `pulumi config set honeycomb:<option>` or pass options to the [constructor of `new honeycomb.Provider`]({{< relref "/registry/packages/honeycomb/api-docs/provider" >}}).

| Option   | Required/Optional | Description                                                                                          |
|----------|-------------------|------------------------------------------------------------------------------------------------------|
| `apiKey` | Required          | The Honeycomb API key to use. It can also be set using the `HONEYCOMBIO_APIKEY` environment variable |
| `apiUrl` | Optional          | Override the url of the Honeycomb.io API. Defaults to `https://api.honeycomb.io`.                    |
| `debug`  | Optional          | Enable to log additional debug information. To view the logs, set `TF_LOG` to at least debug         |
