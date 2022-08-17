// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: DerivedColumn
//
// Creates a derived column. For more information about derived columns, check out [Calculate with derived columns](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			dataset := cfg.Require("dataset")
//			_, err := honeycomb.NewDerivedColumn(ctx, "durationMsLog", &honeycomb.DerivedColumnArgs{
//				Alias:       pulumi.String("duration_ms_log10"),
//				Expression:  pulumi.String(fmt.Sprintf("LOG10($duration_ms)")),
//				Description: pulumi.String("LOG10 of duration_ms"),
//				Dataset:     pulumi.String(dataset),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Derived columns can be imported using a combination of the dataset name and their alias, e.g.
//
// ```sh
//
//	$ pulumi import honeycomb:index/derivedColumn:DerivedColumn my_column my-dataset/duration_ms_log10
//
// ```
type DerivedColumn struct {
	pulumi.CustomResourceState

	// The name of the derived column. Must be unique per dataset.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The dataset this derived column is added to.
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// A description that is shown in the UI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The function of the derived column. See [Derived Column Syntax](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/#derived-column-syntax).
	Expression pulumi.StringOutput `pulumi:"expression"`
}

// NewDerivedColumn registers a new resource with the given unique name, arguments, and options.
func NewDerivedColumn(ctx *pulumi.Context,
	name string, args *DerivedColumnArgs, opts ...pulumi.ResourceOption) (*DerivedColumn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.Expression == nil {
		return nil, errors.New("invalid value for required argument 'Expression'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DerivedColumn
	err := ctx.RegisterResource("honeycomb:index/derivedColumn:DerivedColumn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDerivedColumn gets an existing DerivedColumn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDerivedColumn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DerivedColumnState, opts ...pulumi.ResourceOption) (*DerivedColumn, error) {
	var resource DerivedColumn
	err := ctx.ReadResource("honeycomb:index/derivedColumn:DerivedColumn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DerivedColumn resources.
type derivedColumnState struct {
	// The name of the derived column. Must be unique per dataset.
	Alias *string `pulumi:"alias"`
	// The dataset this derived column is added to.
	Dataset *string `pulumi:"dataset"`
	// A description that is shown in the UI.
	Description *string `pulumi:"description"`
	// The function of the derived column. See [Derived Column Syntax](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/#derived-column-syntax).
	Expression *string `pulumi:"expression"`
}

type DerivedColumnState struct {
	// The name of the derived column. Must be unique per dataset.
	Alias pulumi.StringPtrInput
	// The dataset this derived column is added to.
	Dataset pulumi.StringPtrInput
	// A description that is shown in the UI.
	Description pulumi.StringPtrInput
	// The function of the derived column. See [Derived Column Syntax](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/#derived-column-syntax).
	Expression pulumi.StringPtrInput
}

func (DerivedColumnState) ElementType() reflect.Type {
	return reflect.TypeOf((*derivedColumnState)(nil)).Elem()
}

type derivedColumnArgs struct {
	// The name of the derived column. Must be unique per dataset.
	Alias string `pulumi:"alias"`
	// The dataset this derived column is added to.
	Dataset string `pulumi:"dataset"`
	// A description that is shown in the UI.
	Description *string `pulumi:"description"`
	// The function of the derived column. See [Derived Column Syntax](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/#derived-column-syntax).
	Expression string `pulumi:"expression"`
}

// The set of arguments for constructing a DerivedColumn resource.
type DerivedColumnArgs struct {
	// The name of the derived column. Must be unique per dataset.
	Alias pulumi.StringInput
	// The dataset this derived column is added to.
	Dataset pulumi.StringInput
	// A description that is shown in the UI.
	Description pulumi.StringPtrInput
	// The function of the derived column. See [Derived Column Syntax](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/#derived-column-syntax).
	Expression pulumi.StringInput
}

func (DerivedColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*derivedColumnArgs)(nil)).Elem()
}

type DerivedColumnInput interface {
	pulumi.Input

	ToDerivedColumnOutput() DerivedColumnOutput
	ToDerivedColumnOutputWithContext(ctx context.Context) DerivedColumnOutput
}

func (*DerivedColumn) ElementType() reflect.Type {
	return reflect.TypeOf((**DerivedColumn)(nil)).Elem()
}

func (i *DerivedColumn) ToDerivedColumnOutput() DerivedColumnOutput {
	return i.ToDerivedColumnOutputWithContext(context.Background())
}

func (i *DerivedColumn) ToDerivedColumnOutputWithContext(ctx context.Context) DerivedColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DerivedColumnOutput)
}

// DerivedColumnArrayInput is an input type that accepts DerivedColumnArray and DerivedColumnArrayOutput values.
// You can construct a concrete instance of `DerivedColumnArrayInput` via:
//
//	DerivedColumnArray{ DerivedColumnArgs{...} }
type DerivedColumnArrayInput interface {
	pulumi.Input

	ToDerivedColumnArrayOutput() DerivedColumnArrayOutput
	ToDerivedColumnArrayOutputWithContext(context.Context) DerivedColumnArrayOutput
}

type DerivedColumnArray []DerivedColumnInput

func (DerivedColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DerivedColumn)(nil)).Elem()
}

func (i DerivedColumnArray) ToDerivedColumnArrayOutput() DerivedColumnArrayOutput {
	return i.ToDerivedColumnArrayOutputWithContext(context.Background())
}

func (i DerivedColumnArray) ToDerivedColumnArrayOutputWithContext(ctx context.Context) DerivedColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DerivedColumnArrayOutput)
}

// DerivedColumnMapInput is an input type that accepts DerivedColumnMap and DerivedColumnMapOutput values.
// You can construct a concrete instance of `DerivedColumnMapInput` via:
//
//	DerivedColumnMap{ "key": DerivedColumnArgs{...} }
type DerivedColumnMapInput interface {
	pulumi.Input

	ToDerivedColumnMapOutput() DerivedColumnMapOutput
	ToDerivedColumnMapOutputWithContext(context.Context) DerivedColumnMapOutput
}

type DerivedColumnMap map[string]DerivedColumnInput

func (DerivedColumnMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DerivedColumn)(nil)).Elem()
}

func (i DerivedColumnMap) ToDerivedColumnMapOutput() DerivedColumnMapOutput {
	return i.ToDerivedColumnMapOutputWithContext(context.Background())
}

func (i DerivedColumnMap) ToDerivedColumnMapOutputWithContext(ctx context.Context) DerivedColumnMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DerivedColumnMapOutput)
}

type DerivedColumnOutput struct{ *pulumi.OutputState }

func (DerivedColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DerivedColumn)(nil)).Elem()
}

func (o DerivedColumnOutput) ToDerivedColumnOutput() DerivedColumnOutput {
	return o
}

func (o DerivedColumnOutput) ToDerivedColumnOutputWithContext(ctx context.Context) DerivedColumnOutput {
	return o
}

// The name of the derived column. Must be unique per dataset.
func (o DerivedColumnOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *DerivedColumn) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// The dataset this derived column is added to.
func (o DerivedColumnOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *DerivedColumn) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

// A description that is shown in the UI.
func (o DerivedColumnOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DerivedColumn) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The function of the derived column. See [Derived Column Syntax](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/#derived-column-syntax).
func (o DerivedColumnOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v *DerivedColumn) pulumi.StringOutput { return v.Expression }).(pulumi.StringOutput)
}

type DerivedColumnArrayOutput struct{ *pulumi.OutputState }

func (DerivedColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DerivedColumn)(nil)).Elem()
}

func (o DerivedColumnArrayOutput) ToDerivedColumnArrayOutput() DerivedColumnArrayOutput {
	return o
}

func (o DerivedColumnArrayOutput) ToDerivedColumnArrayOutputWithContext(ctx context.Context) DerivedColumnArrayOutput {
	return o
}

func (o DerivedColumnArrayOutput) Index(i pulumi.IntInput) DerivedColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DerivedColumn {
		return vs[0].([]*DerivedColumn)[vs[1].(int)]
	}).(DerivedColumnOutput)
}

type DerivedColumnMapOutput struct{ *pulumi.OutputState }

func (DerivedColumnMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DerivedColumn)(nil)).Elem()
}

func (o DerivedColumnMapOutput) ToDerivedColumnMapOutput() DerivedColumnMapOutput {
	return o
}

func (o DerivedColumnMapOutput) ToDerivedColumnMapOutputWithContext(ctx context.Context) DerivedColumnMapOutput {
	return o
}

func (o DerivedColumnMapOutput) MapIndex(k pulumi.StringInput) DerivedColumnOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DerivedColumn {
		return vs[0].(map[string]*DerivedColumn)[vs[1].(string)]
	}).(DerivedColumnOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DerivedColumnInput)(nil)).Elem(), &DerivedColumn{})
	pulumi.RegisterInputType(reflect.TypeOf((*DerivedColumnArrayInput)(nil)).Elem(), DerivedColumnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DerivedColumnMapInput)(nil)).Elem(), DerivedColumnMap{})
	pulumi.RegisterOutputType(DerivedColumnOutput{})
	pulumi.RegisterOutputType(DerivedColumnArrayOutput{})
	pulumi.RegisterOutputType(DerivedColumnMapOutput{})
}
