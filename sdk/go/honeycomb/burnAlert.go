// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: BurnAlert
//
// Creates a burn alert. For more information about burn alerts, check out [Define Burn Alerts](https://docs.honeycomb.io/working-with-your-data/slos/slo-process/#define-burn-alerts).
//
// ## Example Usage
// ### Basic Example
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		dataset := cfg.Require("dataset")
// 		sloId := cfg.Require("sloId")
// 		_, err := honeycomb.NewBurnAlert(ctx, "exampleAlert", &honeycomb.BurnAlertArgs{
// 			Dataset:           pulumi.String(dataset),
// 			SloId:             pulumi.String(sloId),
// 			ExhaustionMinutes: pulumi.Int(480),
// 			Recipients: BurnAlertRecipientArray{
// 				&BurnAlertRecipientArgs{
// 					Type:   pulumi.String("email"),
// 					Target: pulumi.String("hello@example.com"),
// 				},
// 				&BurnAlertRecipientArgs{
// 					Type:   pulumi.String("slack"),
// 					Target: pulumi.String("#example-channel"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Example with PagerDuty Recipient and Severity
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		dataset := cfg.Require("dataset")
// 		sloId := cfg.Require("sloId")
// 		pd_prod, err := honeycomb.GetRecipient(ctx, &GetRecipientArgs{
// 			Type: "pagerduty",
// 			DetailFilter: GetRecipientDetailFilter{
// 				Name:  "integration_name",
// 				Value: pulumi.StringRef("Prod On-Call"),
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = honeycomb.NewBurnAlert(ctx, "exampleAlert", &honeycomb.BurnAlertArgs{
// 			Dataset:           pulumi.String(dataset),
// 			SloId:             pulumi.String(sloId),
// 			ExhaustionMinutes: pulumi.Int(60),
// 			Recipients: BurnAlertRecipientArray{
// 				&BurnAlertRecipientArgs{
// 					Id: pulumi.String(pd_prod.Id),
// 					NotificationDetails: &BurnAlertRecipientNotificationDetailsArgs{
// 						PagerdutySeverity: pulumi.String("critical"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Burn Alerts can be imported using a combination of the dataset name and their ID, e.g.
//
// ```sh
//  $ pulumi import honeycomb:index/burnAlert:BurnAlert my_alert my-dataset/bj9BwOb1uKz
// ```
type BurnAlert struct {
	pulumi.CustomResourceState

	// The dataset this burn alert is associated with.
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// The amount of time, in minutes, remaining before the SLO's error budget will be exhausted and the alert will fire.
	ExhaustionMinutes pulumi.IntOutput `pulumi:"exhaustionMinutes"`
	// Zero or more configuration blocks (described below) with the recipients to notify when the alert fires.
	Recipients BurnAlertRecipientArrayOutput `pulumi:"recipients"`
	// ID of the SLO this burn alert is associated with.
	SloId pulumi.StringOutput `pulumi:"sloId"`
}

// NewBurnAlert registers a new resource with the given unique name, arguments, and options.
func NewBurnAlert(ctx *pulumi.Context,
	name string, args *BurnAlertArgs, opts ...pulumi.ResourceOption) (*BurnAlert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.ExhaustionMinutes == nil {
		return nil, errors.New("invalid value for required argument 'ExhaustionMinutes'")
	}
	if args.Recipients == nil {
		return nil, errors.New("invalid value for required argument 'Recipients'")
	}
	if args.SloId == nil {
		return nil, errors.New("invalid value for required argument 'SloId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BurnAlert
	err := ctx.RegisterResource("honeycomb:index/burnAlert:BurnAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBurnAlert gets an existing BurnAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBurnAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BurnAlertState, opts ...pulumi.ResourceOption) (*BurnAlert, error) {
	var resource BurnAlert
	err := ctx.ReadResource("honeycomb:index/burnAlert:BurnAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BurnAlert resources.
type burnAlertState struct {
	// The dataset this burn alert is associated with.
	Dataset *string `pulumi:"dataset"`
	// The amount of time, in minutes, remaining before the SLO's error budget will be exhausted and the alert will fire.
	ExhaustionMinutes *int `pulumi:"exhaustionMinutes"`
	// Zero or more configuration blocks (described below) with the recipients to notify when the alert fires.
	Recipients []BurnAlertRecipient `pulumi:"recipients"`
	// ID of the SLO this burn alert is associated with.
	SloId *string `pulumi:"sloId"`
}

type BurnAlertState struct {
	// The dataset this burn alert is associated with.
	Dataset pulumi.StringPtrInput
	// The amount of time, in minutes, remaining before the SLO's error budget will be exhausted and the alert will fire.
	ExhaustionMinutes pulumi.IntPtrInput
	// Zero or more configuration blocks (described below) with the recipients to notify when the alert fires.
	Recipients BurnAlertRecipientArrayInput
	// ID of the SLO this burn alert is associated with.
	SloId pulumi.StringPtrInput
}

func (BurnAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*burnAlertState)(nil)).Elem()
}

type burnAlertArgs struct {
	// The dataset this burn alert is associated with.
	Dataset string `pulumi:"dataset"`
	// The amount of time, in minutes, remaining before the SLO's error budget will be exhausted and the alert will fire.
	ExhaustionMinutes int `pulumi:"exhaustionMinutes"`
	// Zero or more configuration blocks (described below) with the recipients to notify when the alert fires.
	Recipients []BurnAlertRecipient `pulumi:"recipients"`
	// ID of the SLO this burn alert is associated with.
	SloId string `pulumi:"sloId"`
}

// The set of arguments for constructing a BurnAlert resource.
type BurnAlertArgs struct {
	// The dataset this burn alert is associated with.
	Dataset pulumi.StringInput
	// The amount of time, in minutes, remaining before the SLO's error budget will be exhausted and the alert will fire.
	ExhaustionMinutes pulumi.IntInput
	// Zero or more configuration blocks (described below) with the recipients to notify when the alert fires.
	Recipients BurnAlertRecipientArrayInput
	// ID of the SLO this burn alert is associated with.
	SloId pulumi.StringInput
}

func (BurnAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*burnAlertArgs)(nil)).Elem()
}

type BurnAlertInput interface {
	pulumi.Input

	ToBurnAlertOutput() BurnAlertOutput
	ToBurnAlertOutputWithContext(ctx context.Context) BurnAlertOutput
}

func (*BurnAlert) ElementType() reflect.Type {
	return reflect.TypeOf((**BurnAlert)(nil)).Elem()
}

func (i *BurnAlert) ToBurnAlertOutput() BurnAlertOutput {
	return i.ToBurnAlertOutputWithContext(context.Background())
}

func (i *BurnAlert) ToBurnAlertOutputWithContext(ctx context.Context) BurnAlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BurnAlertOutput)
}

// BurnAlertArrayInput is an input type that accepts BurnAlertArray and BurnAlertArrayOutput values.
// You can construct a concrete instance of `BurnAlertArrayInput` via:
//
//          BurnAlertArray{ BurnAlertArgs{...} }
type BurnAlertArrayInput interface {
	pulumi.Input

	ToBurnAlertArrayOutput() BurnAlertArrayOutput
	ToBurnAlertArrayOutputWithContext(context.Context) BurnAlertArrayOutput
}

type BurnAlertArray []BurnAlertInput

func (BurnAlertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BurnAlert)(nil)).Elem()
}

func (i BurnAlertArray) ToBurnAlertArrayOutput() BurnAlertArrayOutput {
	return i.ToBurnAlertArrayOutputWithContext(context.Background())
}

func (i BurnAlertArray) ToBurnAlertArrayOutputWithContext(ctx context.Context) BurnAlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BurnAlertArrayOutput)
}

// BurnAlertMapInput is an input type that accepts BurnAlertMap and BurnAlertMapOutput values.
// You can construct a concrete instance of `BurnAlertMapInput` via:
//
//          BurnAlertMap{ "key": BurnAlertArgs{...} }
type BurnAlertMapInput interface {
	pulumi.Input

	ToBurnAlertMapOutput() BurnAlertMapOutput
	ToBurnAlertMapOutputWithContext(context.Context) BurnAlertMapOutput
}

type BurnAlertMap map[string]BurnAlertInput

func (BurnAlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BurnAlert)(nil)).Elem()
}

func (i BurnAlertMap) ToBurnAlertMapOutput() BurnAlertMapOutput {
	return i.ToBurnAlertMapOutputWithContext(context.Background())
}

func (i BurnAlertMap) ToBurnAlertMapOutputWithContext(ctx context.Context) BurnAlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BurnAlertMapOutput)
}

type BurnAlertOutput struct{ *pulumi.OutputState }

func (BurnAlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BurnAlert)(nil)).Elem()
}

func (o BurnAlertOutput) ToBurnAlertOutput() BurnAlertOutput {
	return o
}

func (o BurnAlertOutput) ToBurnAlertOutputWithContext(ctx context.Context) BurnAlertOutput {
	return o
}

// The dataset this burn alert is associated with.
func (o BurnAlertOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *BurnAlert) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

// The amount of time, in minutes, remaining before the SLO's error budget will be exhausted and the alert will fire.
func (o BurnAlertOutput) ExhaustionMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v *BurnAlert) pulumi.IntOutput { return v.ExhaustionMinutes }).(pulumi.IntOutput)
}

// Zero or more configuration blocks (described below) with the recipients to notify when the alert fires.
func (o BurnAlertOutput) Recipients() BurnAlertRecipientArrayOutput {
	return o.ApplyT(func(v *BurnAlert) BurnAlertRecipientArrayOutput { return v.Recipients }).(BurnAlertRecipientArrayOutput)
}

// ID of the SLO this burn alert is associated with.
func (o BurnAlertOutput) SloId() pulumi.StringOutput {
	return o.ApplyT(func(v *BurnAlert) pulumi.StringOutput { return v.SloId }).(pulumi.StringOutput)
}

type BurnAlertArrayOutput struct{ *pulumi.OutputState }

func (BurnAlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BurnAlert)(nil)).Elem()
}

func (o BurnAlertArrayOutput) ToBurnAlertArrayOutput() BurnAlertArrayOutput {
	return o
}

func (o BurnAlertArrayOutput) ToBurnAlertArrayOutputWithContext(ctx context.Context) BurnAlertArrayOutput {
	return o
}

func (o BurnAlertArrayOutput) Index(i pulumi.IntInput) BurnAlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BurnAlert {
		return vs[0].([]*BurnAlert)[vs[1].(int)]
	}).(BurnAlertOutput)
}

type BurnAlertMapOutput struct{ *pulumi.OutputState }

func (BurnAlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BurnAlert)(nil)).Elem()
}

func (o BurnAlertMapOutput) ToBurnAlertMapOutput() BurnAlertMapOutput {
	return o
}

func (o BurnAlertMapOutput) ToBurnAlertMapOutputWithContext(ctx context.Context) BurnAlertMapOutput {
	return o
}

func (o BurnAlertMapOutput) MapIndex(k pulumi.StringInput) BurnAlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BurnAlert {
		return vs[0].(map[string]*BurnAlert)[vs[1].(string)]
	}).(BurnAlertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BurnAlertInput)(nil)).Elem(), &BurnAlert{})
	pulumi.RegisterInputType(reflect.TypeOf((*BurnAlertArrayInput)(nil)).Elem(), BurnAlertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BurnAlertMapInput)(nil)).Elem(), BurnAlertMap{})
	pulumi.RegisterOutputType(BurnAlertOutput{})
	pulumi.RegisterOutputType(BurnAlertArrayOutput{})
	pulumi.RegisterOutputType(BurnAlertMapOutput{})
}
