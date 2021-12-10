// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package honeycomb

import (
	"fmt"
	"path/filepath"

	"github.com/antifuchs/pulumi-honeycomb/provider/pkg/version"
	"github.com/kvrhdn/terraform-provider-honeycombio/honeycombio"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "honeycomb"
	// modules:
	mainMod = "index" // the xyz module
	docBase = ""
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(honeycombio.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "honeycombio",
		Description: "A Pulumi package for creating and managing honeycomb.io resources.",
		Keywords:    []string{"pulumi", "honeycomb", "tracing", "o11y", "monitoring", "otel"},
		License:     "Apache-2.0",
		GitHubOrg:   "kvrhdn",
		Homepage:    "https://github.com/antifuchs/pulumi-honeycomb",
		Repository:  "https://github.com/kvrhdn/honeycombio",
		Config: map[string]*tfbridge.SchemaInfo{
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"HONEYCOMBIO_APIKEY"},
				},
			},
			"api_url": {
				Default: &tfbridge.DefaultInfo{
					Value: "https://api.honeycomb.io",
				},
			},
			"debug": {
				Default: &tfbridge.DefaultInfo{
					Value: false,
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"honeycombio_board":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Board")},
			"honeycombio_column":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Column")},
			"honeycombio_dataset":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Dataset")},
			"honeycombio_derived_column":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DerivedColumn")},
			"honeycombio_marker":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Marker")},
			"honeycombio_query":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Query")},
			"honeycombio_query_annotation": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "QueryAnnotation")},
			"honeycombio_trigger":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Trigger")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"honeycombio_datasets":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetDatasets")},
			"honeycombio_query":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetQuery")},
			"honeycombio_trigger_recipient": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetTriggerRecipient")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
