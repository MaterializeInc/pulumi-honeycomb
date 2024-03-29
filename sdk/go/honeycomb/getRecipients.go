// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Data Source: GetRecipients
//
// `GetRecipients` data source provides recipient IDs of recipients matching a set of criteria.
//
// ## Example Usage
// ### Get all recipients
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
// 		_, err := honeycomb.GetRecipients(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Get all email recipients matching a specific domain
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
// 		_, err := honeycomb.GetRecipients(ctx, &GetRecipientsArgs{
// 			DetailFilter: GetRecipientsDetailFilter{
// 				Name:       "address",
// 				ValueRegex: pulumi.StringRef(".*@example.com"),
// 			},
// 			Type: pulumi.StringRef("email"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetRecipients(ctx *pulumi.Context, args *GetRecipientsArgs, opts ...pulumi.InvokeOption) (*GetRecipientsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRecipientsResult
	err := ctx.Invoke("honeycomb:index/getRecipients:GetRecipients", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetRecipients.
type GetRecipientsArgs struct {
	// a block to further filter recipients as described below. `type` must be set when providing a filter.
	DetailFilter *GetRecipientsDetailFilter `pulumi:"detailFilter"`
	// The type of recipient, allowed types are `email`, `pagerduty`, `slack` and `webhook`.
	Type *string `pulumi:"type"`
}

// A collection of values returned by GetRecipients.
type GetRecipientsResult struct {
	DetailFilter *GetRecipientsDetailFilter `pulumi:"detailFilter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of all the recipient IDs found.
	Ids  []string `pulumi:"ids"`
	Type *string  `pulumi:"type"`
}

func GetRecipientsOutput(ctx *pulumi.Context, args GetRecipientsOutputArgs, opts ...pulumi.InvokeOption) GetRecipientsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRecipientsResult, error) {
			args := v.(GetRecipientsArgs)
			r, err := GetRecipients(ctx, &args, opts...)
			var s GetRecipientsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRecipientsResultOutput)
}

// A collection of arguments for invoking GetRecipients.
type GetRecipientsOutputArgs struct {
	// a block to further filter recipients as described below. `type` must be set when providing a filter.
	DetailFilter GetRecipientsDetailFilterPtrInput `pulumi:"detailFilter"`
	// The type of recipient, allowed types are `email`, `pagerduty`, `slack` and `webhook`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetRecipientsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecipientsArgs)(nil)).Elem()
}

// A collection of values returned by GetRecipients.
type GetRecipientsResultOutput struct{ *pulumi.OutputState }

func (GetRecipientsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecipientsResult)(nil)).Elem()
}

func (o GetRecipientsResultOutput) ToGetRecipientsResultOutput() GetRecipientsResultOutput {
	return o
}

func (o GetRecipientsResultOutput) ToGetRecipientsResultOutputWithContext(ctx context.Context) GetRecipientsResultOutput {
	return o
}

func (o GetRecipientsResultOutput) DetailFilter() GetRecipientsDetailFilterPtrOutput {
	return o.ApplyT(func(v GetRecipientsResult) *GetRecipientsDetailFilter { return v.DetailFilter }).(GetRecipientsDetailFilterPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRecipientsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecipientsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of all the recipient IDs found.
func (o GetRecipientsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRecipientsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetRecipientsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecipientsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRecipientsResultOutput{})
}
