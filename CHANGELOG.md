CHANGELOG
=========

## 0.0.14

* Upgraded the terraform provider to 0.11.1.
* New resource `MarkerSetting`.

## 0.0.13

* Add missing resource `SlackRecipient` and data source `GetRecipients`.

## 0.0.12

* Upgraded the honeycomb terraform provider to 0.10.0. Thanks, [markfickett](https://github.com/markfickett)!
* New resources `EmailRecipient`, `PagerdutyRecipient`, `WebhookRecipient`
* Improved the README and added an example. Thanks, [markfickett](https://github.com/markfickett)!

## 0.0.11

* Upgraded the terraform provider to 0.7.0
* New data source `Recipient`
* New resources `SLO` and `BurnAlert`
* Breaking change: the `QueryResult` data source now takes the query
  specification JSON directly.
* The provider should now identify itself to the honeycomb API as
  "pulumi-honeycomb/<version>" in the User-Agent header.

## 0.0.10

* Fixed the resource download URL.

* Hopefully fixed the nodejs release package name.

## 0.0.9

* Upgraded terraform-provider-honeycomb to
  [0.6.0](https://github.com/honeycombio/terraform-provider-honeycombio/releases/tag/v0.6.0).

* Added builds for the go, .NET and nodejs SDKs - thanks
  [@stack72](https://github.com/stack72)!

## 0.0.8

* Upgraded terraform-provider-honeycomb to
  [0.3.1](https://github.com/honeycombio/terraform-provider-honeycombio/releases/tag/v0.3.1),
  skipping
  [0.3.0](https://github.com/honeycombio/terraform-provider-honeycombio/releases/tag/v0.3.0),
  which deprecates `value_boolean`, `value_string` and other filter
  values in favor of `value`.

## 0.0.7

* Corrected setting configuration settings in the provider configuration.

## 0.0.6

* Corrected the plugin download URL to allow actual installation of the python package.

## 0.0.5

* Bumped terraform-provider-honeycomb to 0.2.0, the first version released by honeycomb.
* Added corresponding type for `honeycombio_query_specification`

---
