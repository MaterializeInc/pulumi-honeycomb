// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "honeycomb:index/board:Board":
		r = &Board{}
	case "honeycomb:index/burnAlert:BurnAlert":
		r = &BurnAlert{}
	case "honeycomb:index/column:Column":
		r = &Column{}
	case "honeycomb:index/dataset:Dataset":
		r = &Dataset{}
	case "honeycomb:index/derivedColumn:DerivedColumn":
		r = &DerivedColumn{}
	case "honeycomb:index/emailRecipient:EmailRecipient":
		r = &EmailRecipient{}
	case "honeycomb:index/marker:Marker":
		r = &Marker{}
	case "honeycomb:index/pagerdutyRecipient:PagerdutyRecipient":
		r = &PagerdutyRecipient{}
	case "honeycomb:index/query:Query":
		r = &Query{}
	case "honeycomb:index/queryAnnotation:QueryAnnotation":
		r = &QueryAnnotation{}
	case "honeycomb:index/sLO:SLO":
		r = &SLO{}
	case "honeycomb:index/trigger:Trigger":
		r = &Trigger{}
	case "honeycomb:index/webhookRecipient:WebhookRecipient":
		r = &WebhookRecipient{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:honeycomb" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/board",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/burnAlert",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/column",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/dataset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/derivedColumn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/emailRecipient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/marker",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/pagerdutyRecipient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/query",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/queryAnnotation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/sLO",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/trigger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"honeycomb",
		"index/webhookRecipient",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"honeycomb",
		&pkg{version},
	)
}
