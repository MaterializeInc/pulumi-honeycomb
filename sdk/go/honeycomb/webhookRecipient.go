// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: WebhookRecipient
//
// `WebhookRecipient` allows you to define and manage a Webhook recipient that can be used by Triggers or BurnAlerts notifications.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := honeycomb.NewWebhookRecipient(ctx, "prod", &honeycomb.WebhookRecipientArgs{
// 			Secret: pulumi.String("a63dab148496ecbe04a1a802ca9b95b8"),
// 			Url:    pulumi.String("https://my.url.corp.net"),
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
// Webhook Recipients can be imported by their ID, e.g.
//
// ```sh
//  $ pulumi import honeycomb:index/webhookRecipient:WebhookRecipient my_recipient nx2zsegA0dZ
// ```
type WebhookRecipient struct {
	pulumi.CustomResourceState

	// The name of the Webhook Integration to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The secret to include when sending the notification to the webhook.
	Secret pulumi.StringOutput `pulumi:"secret"`
	// The URL of the endpoint to send the notification to.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewWebhookRecipient registers a new resource with the given unique name, arguments, and options.
func NewWebhookRecipient(ctx *pulumi.Context,
	name string, args *WebhookRecipientArgs, opts ...pulumi.ResourceOption) (*WebhookRecipient, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Secret == nil {
		return nil, errors.New("invalid value for required argument 'Secret'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource WebhookRecipient
	err := ctx.RegisterResource("honeycomb:index/webhookRecipient:WebhookRecipient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhookRecipient gets an existing WebhookRecipient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhookRecipient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookRecipientState, opts ...pulumi.ResourceOption) (*WebhookRecipient, error) {
	var resource WebhookRecipient
	err := ctx.ReadResource("honeycomb:index/webhookRecipient:WebhookRecipient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebhookRecipient resources.
type webhookRecipientState struct {
	// The name of the Webhook Integration to create.
	Name *string `pulumi:"name"`
	// The secret to include when sending the notification to the webhook.
	Secret *string `pulumi:"secret"`
	// The URL of the endpoint to send the notification to.
	Url *string `pulumi:"url"`
}

type WebhookRecipientState struct {
	// The name of the Webhook Integration to create.
	Name pulumi.StringPtrInput
	// The secret to include when sending the notification to the webhook.
	Secret pulumi.StringPtrInput
	// The URL of the endpoint to send the notification to.
	Url pulumi.StringPtrInput
}

func (WebhookRecipientState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookRecipientState)(nil)).Elem()
}

type webhookRecipientArgs struct {
	// The name of the Webhook Integration to create.
	Name *string `pulumi:"name"`
	// The secret to include when sending the notification to the webhook.
	Secret string `pulumi:"secret"`
	// The URL of the endpoint to send the notification to.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a WebhookRecipient resource.
type WebhookRecipientArgs struct {
	// The name of the Webhook Integration to create.
	Name pulumi.StringPtrInput
	// The secret to include when sending the notification to the webhook.
	Secret pulumi.StringInput
	// The URL of the endpoint to send the notification to.
	Url pulumi.StringInput
}

func (WebhookRecipientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookRecipientArgs)(nil)).Elem()
}

type WebhookRecipientInput interface {
	pulumi.Input

	ToWebhookRecipientOutput() WebhookRecipientOutput
	ToWebhookRecipientOutputWithContext(ctx context.Context) WebhookRecipientOutput
}

func (*WebhookRecipient) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookRecipient)(nil)).Elem()
}

func (i *WebhookRecipient) ToWebhookRecipientOutput() WebhookRecipientOutput {
	return i.ToWebhookRecipientOutputWithContext(context.Background())
}

func (i *WebhookRecipient) ToWebhookRecipientOutputWithContext(ctx context.Context) WebhookRecipientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookRecipientOutput)
}

// WebhookRecipientArrayInput is an input type that accepts WebhookRecipientArray and WebhookRecipientArrayOutput values.
// You can construct a concrete instance of `WebhookRecipientArrayInput` via:
//
//          WebhookRecipientArray{ WebhookRecipientArgs{...} }
type WebhookRecipientArrayInput interface {
	pulumi.Input

	ToWebhookRecipientArrayOutput() WebhookRecipientArrayOutput
	ToWebhookRecipientArrayOutputWithContext(context.Context) WebhookRecipientArrayOutput
}

type WebhookRecipientArray []WebhookRecipientInput

func (WebhookRecipientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebhookRecipient)(nil)).Elem()
}

func (i WebhookRecipientArray) ToWebhookRecipientArrayOutput() WebhookRecipientArrayOutput {
	return i.ToWebhookRecipientArrayOutputWithContext(context.Background())
}

func (i WebhookRecipientArray) ToWebhookRecipientArrayOutputWithContext(ctx context.Context) WebhookRecipientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookRecipientArrayOutput)
}

// WebhookRecipientMapInput is an input type that accepts WebhookRecipientMap and WebhookRecipientMapOutput values.
// You can construct a concrete instance of `WebhookRecipientMapInput` via:
//
//          WebhookRecipientMap{ "key": WebhookRecipientArgs{...} }
type WebhookRecipientMapInput interface {
	pulumi.Input

	ToWebhookRecipientMapOutput() WebhookRecipientMapOutput
	ToWebhookRecipientMapOutputWithContext(context.Context) WebhookRecipientMapOutput
}

type WebhookRecipientMap map[string]WebhookRecipientInput

func (WebhookRecipientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebhookRecipient)(nil)).Elem()
}

func (i WebhookRecipientMap) ToWebhookRecipientMapOutput() WebhookRecipientMapOutput {
	return i.ToWebhookRecipientMapOutputWithContext(context.Background())
}

func (i WebhookRecipientMap) ToWebhookRecipientMapOutputWithContext(ctx context.Context) WebhookRecipientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookRecipientMapOutput)
}

type WebhookRecipientOutput struct{ *pulumi.OutputState }

func (WebhookRecipientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookRecipient)(nil)).Elem()
}

func (o WebhookRecipientOutput) ToWebhookRecipientOutput() WebhookRecipientOutput {
	return o
}

func (o WebhookRecipientOutput) ToWebhookRecipientOutputWithContext(ctx context.Context) WebhookRecipientOutput {
	return o
}

// The name of the Webhook Integration to create.
func (o WebhookRecipientOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebhookRecipient) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The secret to include when sending the notification to the webhook.
func (o WebhookRecipientOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *WebhookRecipient) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

// The URL of the endpoint to send the notification to.
func (o WebhookRecipientOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *WebhookRecipient) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type WebhookRecipientArrayOutput struct{ *pulumi.OutputState }

func (WebhookRecipientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebhookRecipient)(nil)).Elem()
}

func (o WebhookRecipientArrayOutput) ToWebhookRecipientArrayOutput() WebhookRecipientArrayOutput {
	return o
}

func (o WebhookRecipientArrayOutput) ToWebhookRecipientArrayOutputWithContext(ctx context.Context) WebhookRecipientArrayOutput {
	return o
}

func (o WebhookRecipientArrayOutput) Index(i pulumi.IntInput) WebhookRecipientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebhookRecipient {
		return vs[0].([]*WebhookRecipient)[vs[1].(int)]
	}).(WebhookRecipientOutput)
}

type WebhookRecipientMapOutput struct{ *pulumi.OutputState }

func (WebhookRecipientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebhookRecipient)(nil)).Elem()
}

func (o WebhookRecipientMapOutput) ToWebhookRecipientMapOutput() WebhookRecipientMapOutput {
	return o
}

func (o WebhookRecipientMapOutput) ToWebhookRecipientMapOutputWithContext(ctx context.Context) WebhookRecipientMapOutput {
	return o
}

func (o WebhookRecipientMapOutput) MapIndex(k pulumi.StringInput) WebhookRecipientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebhookRecipient {
		return vs[0].(map[string]*WebhookRecipient)[vs[1].(string)]
	}).(WebhookRecipientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookRecipientInput)(nil)).Elem(), &WebhookRecipient{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookRecipientArrayInput)(nil)).Elem(), WebhookRecipientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookRecipientMapInput)(nil)).Elem(), WebhookRecipientMap{})
	pulumi.RegisterOutputType(WebhookRecipientOutput{})
	pulumi.RegisterOutputType(WebhookRecipientArrayOutput{})
	pulumi.RegisterOutputType(WebhookRecipientMapOutput{})
}
