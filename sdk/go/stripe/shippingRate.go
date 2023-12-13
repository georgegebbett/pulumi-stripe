// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stripe

import (
	"context"
	"reflect"

	"errors"
	"github.com/georgegebbett/pulumi-stripe/sdk/go/stripe/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// With this resource, you can create a shipping rate - [Stripe API shipping rate documentation](https://stripe.com/docs/api/shipping_rates).
//
// Shipping rates let you display various shipping options—like standard, express, and overnight—with more accurate delivery estimates.
// Charge your customer for shipping using different Stripe products, some of which require coding.
//
// > Removal of the shipping rate isn't supported through the Stripe SDK. The best practice, which this provider follows,
// is to archive the shipping rate by marking it as inactive on destroy, which indicates that the shipping rate is no longer
// available.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/georgegebbett/pulumi-stripe/sdk/go/stripe"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := stripe.NewShippingRate(ctx, "shippingRate", &stripe.ShippingRateArgs{
//				DeliveryEstimates: stripe.ShippingRateDeliveryEstimateArray{
//					&stripe.ShippingRateDeliveryEstimateArgs{
//						Maximum: &stripe.ShippingRateDeliveryEstimateMaximumArgs{
//							Unit:  pulumi.String("day"),
//							Value: pulumi.Int(4),
//						},
//						Minimum: &stripe.ShippingRateDeliveryEstimateMinimumArgs{
//							Unit:  pulumi.String("hour"),
//							Value: pulumi.Int(24),
//						},
//					},
//				},
//				DisplayName: pulumi.String("shipping rate"),
//				FixedAmount: &stripe.ShippingRateFixedAmountArgs{
//					Amount:   pulumi.Int(1000),
//					Currency: pulumi.String("aud"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = stripe.NewShippingRate(ctx, "shipping", &stripe.ShippingRateArgs{
//				DisplayName: pulumi.String("shipping rate"),
//				FixedAmount: &stripe.ShippingRateFixedAmountArgs{
//					Amount:   pulumi.Int(1000),
//					Currency: pulumi.String("aud"),
//					CurrencyOptions: stripe.ShippingRateFixedAmountCurrencyOptionArray{
//						&stripe.ShippingRateFixedAmountCurrencyOptionArgs{
//							Amount:   pulumi.Int(350),
//							Currency: pulumi.String("eur"),
//						},
//						&stripe.ShippingRateFixedAmountCurrencyOptionArgs{
//							Amount:   pulumi.Int(500),
//							Currency: pulumi.String("usd"),
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
type ShippingRate struct {
	pulumi.CustomResourceState

	// Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// List(Resource). The estimated range for how long shipping will take,
	// meant to be displayable to the customer. This will appear on CheckoutSessions.
	// For details please see Delivery Estimate.
	DeliveryEstimates ShippingRateDeliveryEstimateArrayOutput `pulumi:"deliveryEstimates"`
	// String. The name of the shipping rate, meant to be displayable to the customer.
	// This will appear on CheckoutSessions.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// List(Resource). Describes a fixed amount to charge for shipping.
	// Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
	FixedAmount ShippingRateFixedAmountOutput `pulumi:"fixedAmount"`
	// Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode pulumi.BoolOutput `pulumi:"livemode"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
	// unspecified.
	TaxBehavior pulumi.StringPtrOutput `pulumi:"taxBehavior"`
	// String. A tax code ID. The Shipping tax code is `txcd92010001`.
	TaxCode pulumi.StringPtrOutput `pulumi:"taxCode"`
	// String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewShippingRate registers a new resource with the given unique name, arguments, and options.
func NewShippingRate(ctx *pulumi.Context,
	name string, args *ShippingRateArgs, opts ...pulumi.ResourceOption) (*ShippingRate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.FixedAmount == nil {
		return nil, errors.New("invalid value for required argument 'FixedAmount'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ShippingRate
	err := ctx.RegisterResource("stripe:index/shippingRate:ShippingRate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetShippingRate gets an existing ShippingRate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetShippingRate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ShippingRateState, opts ...pulumi.ResourceOption) (*ShippingRate, error) {
	var resource ShippingRate
	err := ctx.ReadResource("stripe:index/shippingRate:ShippingRate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ShippingRate resources.
type shippingRateState struct {
	// Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
	Active *bool `pulumi:"active"`
	// List(Resource). The estimated range for how long shipping will take,
	// meant to be displayable to the customer. This will appear on CheckoutSessions.
	// For details please see Delivery Estimate.
	DeliveryEstimates []ShippingRateDeliveryEstimate `pulumi:"deliveryEstimates"`
	// String. The name of the shipping rate, meant to be displayable to the customer.
	// This will appear on CheckoutSessions.
	DisplayName *string `pulumi:"displayName"`
	// List(Resource). Describes a fixed amount to charge for shipping.
	// Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
	FixedAmount *ShippingRateFixedAmount `pulumi:"fixedAmount"`
	// Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode *bool `pulumi:"livemode"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
	// unspecified.
	TaxBehavior *string `pulumi:"taxBehavior"`
	// String. A tax code ID. The Shipping tax code is `txcd92010001`.
	TaxCode *string `pulumi:"taxCode"`
	// String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
	Type *string `pulumi:"type"`
}

type ShippingRateState struct {
	// Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
	Active pulumi.BoolPtrInput
	// List(Resource). The estimated range for how long shipping will take,
	// meant to be displayable to the customer. This will appear on CheckoutSessions.
	// For details please see Delivery Estimate.
	DeliveryEstimates ShippingRateDeliveryEstimateArrayInput
	// String. The name of the shipping rate, meant to be displayable to the customer.
	// This will appear on CheckoutSessions.
	DisplayName pulumi.StringPtrInput
	// List(Resource). Describes a fixed amount to charge for shipping.
	// Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
	FixedAmount ShippingRateFixedAmountPtrInput
	// Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode pulumi.BoolPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
	// unspecified.
	TaxBehavior pulumi.StringPtrInput
	// String. A tax code ID. The Shipping tax code is `txcd92010001`.
	TaxCode pulumi.StringPtrInput
	// String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
	Type pulumi.StringPtrInput
}

func (ShippingRateState) ElementType() reflect.Type {
	return reflect.TypeOf((*shippingRateState)(nil)).Elem()
}

type shippingRateArgs struct {
	// Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
	Active *bool `pulumi:"active"`
	// List(Resource). The estimated range for how long shipping will take,
	// meant to be displayable to the customer. This will appear on CheckoutSessions.
	// For details please see Delivery Estimate.
	DeliveryEstimates []ShippingRateDeliveryEstimate `pulumi:"deliveryEstimates"`
	// String. The name of the shipping rate, meant to be displayable to the customer.
	// This will appear on CheckoutSessions.
	DisplayName string `pulumi:"displayName"`
	// List(Resource). Describes a fixed amount to charge for shipping.
	// Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
	FixedAmount ShippingRateFixedAmount `pulumi:"fixedAmount"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
	// unspecified.
	TaxBehavior *string `pulumi:"taxBehavior"`
	// String. A tax code ID. The Shipping tax code is `txcd92010001`.
	TaxCode *string `pulumi:"taxCode"`
	// String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ShippingRate resource.
type ShippingRateArgs struct {
	// Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
	Active pulumi.BoolPtrInput
	// List(Resource). The estimated range for how long shipping will take,
	// meant to be displayable to the customer. This will appear on CheckoutSessions.
	// For details please see Delivery Estimate.
	DeliveryEstimates ShippingRateDeliveryEstimateArrayInput
	// String. The name of the shipping rate, meant to be displayable to the customer.
	// This will appear on CheckoutSessions.
	DisplayName pulumi.StringInput
	// List(Resource). Describes a fixed amount to charge for shipping.
	// Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
	FixedAmount ShippingRateFixedAmountInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
	// unspecified.
	TaxBehavior pulumi.StringPtrInput
	// String. A tax code ID. The Shipping tax code is `txcd92010001`.
	TaxCode pulumi.StringPtrInput
	// String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
	Type pulumi.StringPtrInput
}

func (ShippingRateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*shippingRateArgs)(nil)).Elem()
}

type ShippingRateInput interface {
	pulumi.Input

	ToShippingRateOutput() ShippingRateOutput
	ToShippingRateOutputWithContext(ctx context.Context) ShippingRateOutput
}

func (*ShippingRate) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingRate)(nil)).Elem()
}

func (i *ShippingRate) ToShippingRateOutput() ShippingRateOutput {
	return i.ToShippingRateOutputWithContext(context.Background())
}

func (i *ShippingRate) ToShippingRateOutputWithContext(ctx context.Context) ShippingRateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingRateOutput)
}

// ShippingRateArrayInput is an input type that accepts ShippingRateArray and ShippingRateArrayOutput values.
// You can construct a concrete instance of `ShippingRateArrayInput` via:
//
//	ShippingRateArray{ ShippingRateArgs{...} }
type ShippingRateArrayInput interface {
	pulumi.Input

	ToShippingRateArrayOutput() ShippingRateArrayOutput
	ToShippingRateArrayOutputWithContext(context.Context) ShippingRateArrayOutput
}

type ShippingRateArray []ShippingRateInput

func (ShippingRateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShippingRate)(nil)).Elem()
}

func (i ShippingRateArray) ToShippingRateArrayOutput() ShippingRateArrayOutput {
	return i.ToShippingRateArrayOutputWithContext(context.Background())
}

func (i ShippingRateArray) ToShippingRateArrayOutputWithContext(ctx context.Context) ShippingRateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingRateArrayOutput)
}

// ShippingRateMapInput is an input type that accepts ShippingRateMap and ShippingRateMapOutput values.
// You can construct a concrete instance of `ShippingRateMapInput` via:
//
//	ShippingRateMap{ "key": ShippingRateArgs{...} }
type ShippingRateMapInput interface {
	pulumi.Input

	ToShippingRateMapOutput() ShippingRateMapOutput
	ToShippingRateMapOutputWithContext(context.Context) ShippingRateMapOutput
}

type ShippingRateMap map[string]ShippingRateInput

func (ShippingRateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShippingRate)(nil)).Elem()
}

func (i ShippingRateMap) ToShippingRateMapOutput() ShippingRateMapOutput {
	return i.ToShippingRateMapOutputWithContext(context.Background())
}

func (i ShippingRateMap) ToShippingRateMapOutputWithContext(ctx context.Context) ShippingRateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShippingRateMapOutput)
}

type ShippingRateOutput struct{ *pulumi.OutputState }

func (ShippingRateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ShippingRate)(nil)).Elem()
}

func (o ShippingRateOutput) ToShippingRateOutput() ShippingRateOutput {
	return o
}

func (o ShippingRateOutput) ToShippingRateOutputWithContext(ctx context.Context) ShippingRateOutput {
	return o
}

// Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
func (o ShippingRateOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ShippingRate) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// List(Resource). The estimated range for how long shipping will take,
// meant to be displayable to the customer. This will appear on CheckoutSessions.
// For details please see Delivery Estimate.
func (o ShippingRateOutput) DeliveryEstimates() ShippingRateDeliveryEstimateArrayOutput {
	return o.ApplyT(func(v *ShippingRate) ShippingRateDeliveryEstimateArrayOutput { return v.DeliveryEstimates }).(ShippingRateDeliveryEstimateArrayOutput)
}

// String. The name of the shipping rate, meant to be displayable to the customer.
// This will appear on CheckoutSessions.
func (o ShippingRateOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ShippingRate) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// List(Resource). Describes a fixed amount to charge for shipping.
// Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
func (o ShippingRateOutput) FixedAmount() ShippingRateFixedAmountOutput {
	return o.ApplyT(func(v *ShippingRate) ShippingRateFixedAmountOutput { return v.FixedAmount }).(ShippingRateFixedAmountOutput)
}

// Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
func (o ShippingRateOutput) Livemode() pulumi.BoolOutput {
	return o.ApplyT(func(v *ShippingRate) pulumi.BoolOutput { return v.Livemode }).(pulumi.BoolOutput)
}

// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
// storing additional information about the object in a structured format.
func (o ShippingRateOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ShippingRate) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
// unspecified.
func (o ShippingRateOutput) TaxBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingRate) pulumi.StringPtrOutput { return v.TaxBehavior }).(pulumi.StringPtrOutput)
}

// String. A tax code ID. The Shipping tax code is `txcd92010001`.
func (o ShippingRateOutput) TaxCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingRate) pulumi.StringPtrOutput { return v.TaxCode }).(pulumi.StringPtrOutput)
}

// String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
func (o ShippingRateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ShippingRate) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type ShippingRateArrayOutput struct{ *pulumi.OutputState }

func (ShippingRateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ShippingRate)(nil)).Elem()
}

func (o ShippingRateArrayOutput) ToShippingRateArrayOutput() ShippingRateArrayOutput {
	return o
}

func (o ShippingRateArrayOutput) ToShippingRateArrayOutputWithContext(ctx context.Context) ShippingRateArrayOutput {
	return o
}

func (o ShippingRateArrayOutput) Index(i pulumi.IntInput) ShippingRateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ShippingRate {
		return vs[0].([]*ShippingRate)[vs[1].(int)]
	}).(ShippingRateOutput)
}

type ShippingRateMapOutput struct{ *pulumi.OutputState }

func (ShippingRateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ShippingRate)(nil)).Elem()
}

func (o ShippingRateMapOutput) ToShippingRateMapOutput() ShippingRateMapOutput {
	return o
}

func (o ShippingRateMapOutput) ToShippingRateMapOutputWithContext(ctx context.Context) ShippingRateMapOutput {
	return o
}

func (o ShippingRateMapOutput) MapIndex(k pulumi.StringInput) ShippingRateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ShippingRate {
		return vs[0].(map[string]*ShippingRate)[vs[1].(string)]
	}).(ShippingRateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingRateInput)(nil)).Elem(), &ShippingRate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingRateArrayInput)(nil)).Elem(), ShippingRateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShippingRateMapInput)(nil)).Elem(), ShippingRateMap{})
	pulumi.RegisterOutputType(ShippingRateOutput{})
	pulumi.RegisterOutputType(ShippingRateArrayOutput{})
	pulumi.RegisterOutputType(ShippingRateMapOutput{})
}