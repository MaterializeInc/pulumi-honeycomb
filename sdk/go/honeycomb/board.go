// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: Board
//
// Creates a board. For more information about boards, check out [Collaborate with Boards](https://docs.honeycomb.io/working-with-your-data/collaborating/boards/#docs-sidebar).
//
// ## Example Usage
// ### Simple Board
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			queryGetQuerySpecification, err := honeycomb.GetQuerySpecification(ctx, &GetQuerySpecificationArgs{
//				Calculations: []GetQuerySpecificationCalculation{
//					GetQuerySpecificationCalculation{
//						Op:     "P99",
//						Column: pulumi.StringRef("duration_ms"),
//					},
//				},
//				Filters: []GetQuerySpecificationFilter{
//					GetQuerySpecificationFilter{
//						Column: "trace.parent_id",
//						Op:     "does-not-exist",
//					},
//				},
//				Breakdowns: []string{
//					"app.tenant",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			queryQuery, err := honeycomb.NewQuery(ctx, "queryQuery", &honeycomb.QueryArgs{
//				Dataset:   pulumi.Any(_var.Dataset),
//				QueryJson: pulumi.String(queryGetQuerySpecification.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = honeycomb.NewBoard(ctx, "board", &honeycomb.BoardArgs{
//				Queries: BoardQueryArray{
//					&BoardQueryArgs{
//						Dataset: pulumi.Any(_var.Dataset),
//						QueryId: queryQuery.ID(),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Annotated Board
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			latencyByUseridGetQuerySpecification, err := honeycomb.GetQuerySpecification(ctx, &GetQuerySpecificationArgs{
//				TimeRange: pulumi.IntRef(86400),
//				Breakdowns: []string{
//					"app.user_id",
//				},
//				Calculations: []GetQuerySpecificationCalculation{
//					GetQuerySpecificationCalculation{
//						Op:     "HEATMAP",
//						Column: pulumi.StringRef("duration_ms"),
//					},
//					GetQuerySpecificationCalculation{
//						Op:     "P99",
//						Column: pulumi.StringRef("duration_ms"),
//					},
//				},
//				Filters: []GetQuerySpecificationFilter{
//					GetQuerySpecificationFilter{
//						Column: "trace.parent_id",
//						Op:     "does-not-exist",
//					},
//				},
//				Orders: []GetQuerySpecificationOrder{
//					GetQuerySpecificationOrder{
//						Column: pulumi.StringRef("duration_ms"),
//						Op:     pulumi.StringRef("P99"),
//						Order:  pulumi.StringRef("descending"),
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			latencyByUseridQuery, err := honeycomb.NewQuery(ctx, "latencyByUseridQuery", &honeycomb.QueryArgs{
//				Dataset:   pulumi.Any(_var.Dataset),
//				QueryJson: pulumi.String(latencyByUseridGetQuerySpecification.Json),
//			})
//			if err != nil {
//				return err
//			}
//			latencyByUseridQueryAnnotation, err := honeycomb.NewQueryAnnotation(ctx, "latencyByUseridQueryAnnotation", &honeycomb.QueryAnnotationArgs{
//				Dataset:     pulumi.Any(_var.Dataset),
//				QueryId:     latencyByUseridQuery.ID(),
//				Description: pulumi.String("A breakdown of trace latency by User over the last 24 hours"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = honeycomb.NewBoard(ctx, "overview", &honeycomb.BoardArgs{
//				Style: pulumi.String("visual"),
//				Queries: BoardQueryArray{
//					&BoardQueryArgs{
//						Caption:           pulumi.String("Latency by User"),
//						QueryId:           latencyByUseridQuery.ID(),
//						QueryAnnotationId: latencyByUseridQueryAnnotation.ID(),
//						GraphSettings: &BoardQueryGraphSettingsArgs{
//							UtcXaxis: pulumi.Bool(true),
//						},
//					},
//				},
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
// Boards can be imported using their ID, e.g.
//
// ```sh
//
//	$ pulumi import honeycomb:index/board:Board my_board AobW9oAZX71
//
// ```
//
//	You can find the ID in the URL bar when visiting the board from the UI.
type Board struct {
	pulumi.CustomResourceState

	// the number of columns to layout on the board, either `multi` (the default) or `single`. Only `visual` style boards (see below) have a column layout.
	ColumnLayout pulumi.StringOutput `pulumi:"columnLayout"`
	// Description of the board. Supports markdown.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the board.
	Name pulumi.StringOutput `pulumi:"name"`
	// Zero or more configurations blocks (described below) with the queries of the board.
	Queries BoardQueryArrayOutput `pulumi:"queries"`
	// How the board should be displayed in the UI, either `list` (the default) or `visual`.
	Style pulumi.StringPtrOutput `pulumi:"style"`
}

// NewBoard registers a new resource with the given unique name, arguments, and options.
func NewBoard(ctx *pulumi.Context,
	name string, args *BoardArgs, opts ...pulumi.ResourceOption) (*Board, error) {
	if args == nil {
		args = &BoardArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Board
	err := ctx.RegisterResource("honeycomb:index/board:Board", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBoard gets an existing Board resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBoard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BoardState, opts ...pulumi.ResourceOption) (*Board, error) {
	var resource Board
	err := ctx.ReadResource("honeycomb:index/board:Board", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Board resources.
type boardState struct {
	// the number of columns to layout on the board, either `multi` (the default) or `single`. Only `visual` style boards (see below) have a column layout.
	ColumnLayout *string `pulumi:"columnLayout"`
	// Description of the board. Supports markdown.
	Description *string `pulumi:"description"`
	// Name of the board.
	Name *string `pulumi:"name"`
	// Zero or more configurations blocks (described below) with the queries of the board.
	Queries []BoardQuery `pulumi:"queries"`
	// How the board should be displayed in the UI, either `list` (the default) or `visual`.
	Style *string `pulumi:"style"`
}

type BoardState struct {
	// the number of columns to layout on the board, either `multi` (the default) or `single`. Only `visual` style boards (see below) have a column layout.
	ColumnLayout pulumi.StringPtrInput
	// Description of the board. Supports markdown.
	Description pulumi.StringPtrInput
	// Name of the board.
	Name pulumi.StringPtrInput
	// Zero or more configurations blocks (described below) with the queries of the board.
	Queries BoardQueryArrayInput
	// How the board should be displayed in the UI, either `list` (the default) or `visual`.
	Style pulumi.StringPtrInput
}

func (BoardState) ElementType() reflect.Type {
	return reflect.TypeOf((*boardState)(nil)).Elem()
}

type boardArgs struct {
	// the number of columns to layout on the board, either `multi` (the default) or `single`. Only `visual` style boards (see below) have a column layout.
	ColumnLayout *string `pulumi:"columnLayout"`
	// Description of the board. Supports markdown.
	Description *string `pulumi:"description"`
	// Name of the board.
	Name *string `pulumi:"name"`
	// Zero or more configurations blocks (described below) with the queries of the board.
	Queries []BoardQuery `pulumi:"queries"`
	// How the board should be displayed in the UI, either `list` (the default) or `visual`.
	Style *string `pulumi:"style"`
}

// The set of arguments for constructing a Board resource.
type BoardArgs struct {
	// the number of columns to layout on the board, either `multi` (the default) or `single`. Only `visual` style boards (see below) have a column layout.
	ColumnLayout pulumi.StringPtrInput
	// Description of the board. Supports markdown.
	Description pulumi.StringPtrInput
	// Name of the board.
	Name pulumi.StringPtrInput
	// Zero or more configurations blocks (described below) with the queries of the board.
	Queries BoardQueryArrayInput
	// How the board should be displayed in the UI, either `list` (the default) or `visual`.
	Style pulumi.StringPtrInput
}

func (BoardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*boardArgs)(nil)).Elem()
}

type BoardInput interface {
	pulumi.Input

	ToBoardOutput() BoardOutput
	ToBoardOutputWithContext(ctx context.Context) BoardOutput
}

func (*Board) ElementType() reflect.Type {
	return reflect.TypeOf((**Board)(nil)).Elem()
}

func (i *Board) ToBoardOutput() BoardOutput {
	return i.ToBoardOutputWithContext(context.Background())
}

func (i *Board) ToBoardOutputWithContext(ctx context.Context) BoardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoardOutput)
}

// BoardArrayInput is an input type that accepts BoardArray and BoardArrayOutput values.
// You can construct a concrete instance of `BoardArrayInput` via:
//
//	BoardArray{ BoardArgs{...} }
type BoardArrayInput interface {
	pulumi.Input

	ToBoardArrayOutput() BoardArrayOutput
	ToBoardArrayOutputWithContext(context.Context) BoardArrayOutput
}

type BoardArray []BoardInput

func (BoardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Board)(nil)).Elem()
}

func (i BoardArray) ToBoardArrayOutput() BoardArrayOutput {
	return i.ToBoardArrayOutputWithContext(context.Background())
}

func (i BoardArray) ToBoardArrayOutputWithContext(ctx context.Context) BoardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoardArrayOutput)
}

// BoardMapInput is an input type that accepts BoardMap and BoardMapOutput values.
// You can construct a concrete instance of `BoardMapInput` via:
//
//	BoardMap{ "key": BoardArgs{...} }
type BoardMapInput interface {
	pulumi.Input

	ToBoardMapOutput() BoardMapOutput
	ToBoardMapOutputWithContext(context.Context) BoardMapOutput
}

type BoardMap map[string]BoardInput

func (BoardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Board)(nil)).Elem()
}

func (i BoardMap) ToBoardMapOutput() BoardMapOutput {
	return i.ToBoardMapOutputWithContext(context.Background())
}

func (i BoardMap) ToBoardMapOutputWithContext(ctx context.Context) BoardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BoardMapOutput)
}

type BoardOutput struct{ *pulumi.OutputState }

func (BoardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Board)(nil)).Elem()
}

func (o BoardOutput) ToBoardOutput() BoardOutput {
	return o
}

func (o BoardOutput) ToBoardOutputWithContext(ctx context.Context) BoardOutput {
	return o
}

// the number of columns to layout on the board, either `multi` (the default) or `single`. Only `visual` style boards (see below) have a column layout.
func (o BoardOutput) ColumnLayout() pulumi.StringOutput {
	return o.ApplyT(func(v *Board) pulumi.StringOutput { return v.ColumnLayout }).(pulumi.StringOutput)
}

// Description of the board. Supports markdown.
func (o BoardOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Board) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the board.
func (o BoardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Board) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Zero or more configurations blocks (described below) with the queries of the board.
func (o BoardOutput) Queries() BoardQueryArrayOutput {
	return o.ApplyT(func(v *Board) BoardQueryArrayOutput { return v.Queries }).(BoardQueryArrayOutput)
}

// How the board should be displayed in the UI, either `list` (the default) or `visual`.
func (o BoardOutput) Style() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Board) pulumi.StringPtrOutput { return v.Style }).(pulumi.StringPtrOutput)
}

type BoardArrayOutput struct{ *pulumi.OutputState }

func (BoardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Board)(nil)).Elem()
}

func (o BoardArrayOutput) ToBoardArrayOutput() BoardArrayOutput {
	return o
}

func (o BoardArrayOutput) ToBoardArrayOutputWithContext(ctx context.Context) BoardArrayOutput {
	return o
}

func (o BoardArrayOutput) Index(i pulumi.IntInput) BoardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Board {
		return vs[0].([]*Board)[vs[1].(int)]
	}).(BoardOutput)
}

type BoardMapOutput struct{ *pulumi.OutputState }

func (BoardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Board)(nil)).Elem()
}

func (o BoardMapOutput) ToBoardMapOutput() BoardMapOutput {
	return o
}

func (o BoardMapOutput) ToBoardMapOutputWithContext(ctx context.Context) BoardMapOutput {
	return o
}

func (o BoardMapOutput) MapIndex(k pulumi.StringInput) BoardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Board {
		return vs[0].(map[string]*Board)[vs[1].(string)]
	}).(BoardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BoardInput)(nil)).Elem(), &Board{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoardArrayInput)(nil)).Elem(), BoardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BoardMapInput)(nil)).Elem(), BoardMap{})
	pulumi.RegisterOutputType(BoardOutput{})
	pulumi.RegisterOutputType(BoardArrayOutput{})
	pulumi.RegisterOutputType(BoardMapOutput{})
}
