// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stripe

import (
	"context"
	"reflect"

	"github.com/georgegebbett/pulumi-stripe/sdk/go/stripe/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// With this resource, you can create a coupon - [Stripe API coupon documentation](https://stripe.com/docs/api/coupons).
//
// A coupon contains information about a percent-off or amount-off discount you might want to apply to a customer.
//
// A coupon has either a `percentOff` or an `amountOff` and `currency`. If you set an `amountOff`, that amount will be subtracted from any invoice’s subtotal.
//
// For example, an invoice with a subtotal of $100 will have a final total of $0 if a coupon with an amountOff of 20000 is applied to it and an invoice with a subtotal of $300 will have a final total of $100 if a coupon with an amountOff of 20000 is applied to it.
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
//			// coupon for the amount off discount
//			_, err := stripe.NewCoupon(ctx, "couponCoupon", &stripe.CouponArgs{
//				AmountOff:      pulumi.Int(1000),
//				Currency:       pulumi.String("aud"),
//				Duration:       pulumi.String("once"),
//				MaxRedemptions: pulumi.Int(10),
//			})
//			if err != nil {
//				return err
//			}
//			// coupon for the percentage off discount
//			_, err = stripe.NewCoupon(ctx, "couponIndex/couponCoupon", &stripe.CouponArgs{
//				PercentOff: pulumi.Float64(33.3),
//				Duration:   pulumi.String("forever"),
//			})
//			if err != nil {
//				return err
//			}
//			// coupon with limitation to a date and the product only
//			_, err = stripe.NewCoupon(ctx, "couponStripeIndex/couponCoupon", &stripe.CouponArgs{
//				AmountOff: pulumi.Int(2000),
//				Duration:  pulumi.String("once"),
//				RedeemBy:  pulumi.String("2025-07-23T03:27:06+00:00"),
//				AppliesTos: pulumi.StringArray{
//					stripe_product.Product.Id,
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
type Coupon struct {
	pulumi.CustomResourceState

	// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff pulumi.IntPtrOutput `pulumi:"amountOff"`
	// List(String). A list of product IDs this coupon applies to.
	AppliesTos pulumi.StringArrayOutput `pulumi:"appliesTos"`
	// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
	CouponId pulumi.StringOutput `pulumi:"couponId"`
	// String. Required if `amountOff` has been set, the three-letter ISO code for the currency of the amount to take off.
	Currency pulumi.StringPtrOutput `pulumi:"currency"`
	// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
	Duration pulumi.StringPtrOutput `pulumi:"duration"`
	// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
	DurationInMonths pulumi.IntPtrOutput `pulumi:"durationInMonths"`
	// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptions pulumi.IntPtrOutput `pulumi:"maxRedemptions"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// String. Name of the coupon displayed to customers on for instance invoices or receipts.
	Name pulumi.StringOutput `pulumi:"name"`
	// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percentOff of 50 will make a $100 invoice $50 instead.
	PercentOff pulumi.Float64PtrOutput `pulumi:"percentOff"`
	// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
	RedeemBy pulumi.StringPtrOutput `pulumi:"redeemBy"`
	// Int. Number of times this coupon has been applied to a customer.
	TimesRedeemed pulumi.IntOutput `pulumi:"timesRedeemed"`
	// Bool. Taking account of the above properties, whether this coupon can still be applied to a customer.
	Valid pulumi.BoolOutput `pulumi:"valid"`
}

// NewCoupon registers a new resource with the given unique name, arguments, and options.
func NewCoupon(ctx *pulumi.Context,
	name string, args *CouponArgs, opts ...pulumi.ResourceOption) (*Coupon, error) {
	if args == nil {
		args = &CouponArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Coupon
	err := ctx.RegisterResource("stripe:index/coupon:Coupon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCoupon gets an existing Coupon resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCoupon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CouponState, opts ...pulumi.ResourceOption) (*Coupon, error) {
	var resource Coupon
	err := ctx.ReadResource("stripe:index/coupon:Coupon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Coupon resources.
type couponState struct {
	// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff *int `pulumi:"amountOff"`
	// List(String). A list of product IDs this coupon applies to.
	AppliesTos []string `pulumi:"appliesTos"`
	// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
	CouponId *string `pulumi:"couponId"`
	// String. Required if `amountOff` has been set, the three-letter ISO code for the currency of the amount to take off.
	Currency *string `pulumi:"currency"`
	// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
	Duration *string `pulumi:"duration"`
	// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
	DurationInMonths *int `pulumi:"durationInMonths"`
	// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptions *int `pulumi:"maxRedemptions"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// String. Name of the coupon displayed to customers on for instance invoices or receipts.
	Name *string `pulumi:"name"`
	// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percentOff of 50 will make a $100 invoice $50 instead.
	PercentOff *float64 `pulumi:"percentOff"`
	// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
	RedeemBy *string `pulumi:"redeemBy"`
	// Int. Number of times this coupon has been applied to a customer.
	TimesRedeemed *int `pulumi:"timesRedeemed"`
	// Bool. Taking account of the above properties, whether this coupon can still be applied to a customer.
	Valid *bool `pulumi:"valid"`
}

type CouponState struct {
	// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff pulumi.IntPtrInput
	// List(String). A list of product IDs this coupon applies to.
	AppliesTos pulumi.StringArrayInput
	// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
	CouponId pulumi.StringPtrInput
	// String. Required if `amountOff` has been set, the three-letter ISO code for the currency of the amount to take off.
	Currency pulumi.StringPtrInput
	// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
	Duration pulumi.StringPtrInput
	// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
	DurationInMonths pulumi.IntPtrInput
	// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptions pulumi.IntPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// String. Name of the coupon displayed to customers on for instance invoices or receipts.
	Name pulumi.StringPtrInput
	// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percentOff of 50 will make a $100 invoice $50 instead.
	PercentOff pulumi.Float64PtrInput
	// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
	RedeemBy pulumi.StringPtrInput
	// Int. Number of times this coupon has been applied to a customer.
	TimesRedeemed pulumi.IntPtrInput
	// Bool. Taking account of the above properties, whether this coupon can still be applied to a customer.
	Valid pulumi.BoolPtrInput
}

func (CouponState) ElementType() reflect.Type {
	return reflect.TypeOf((*couponState)(nil)).Elem()
}

type couponArgs struct {
	// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff *int `pulumi:"amountOff"`
	// List(String). A list of product IDs this coupon applies to.
	AppliesTos []string `pulumi:"appliesTos"`
	// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
	CouponId *string `pulumi:"couponId"`
	// String. Required if `amountOff` has been set, the three-letter ISO code for the currency of the amount to take off.
	Currency *string `pulumi:"currency"`
	// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
	Duration *string `pulumi:"duration"`
	// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
	DurationInMonths *int `pulumi:"durationInMonths"`
	// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptions *int `pulumi:"maxRedemptions"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// String. Name of the coupon displayed to customers on for instance invoices or receipts.
	Name *string `pulumi:"name"`
	// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percentOff of 50 will make a $100 invoice $50 instead.
	PercentOff *float64 `pulumi:"percentOff"`
	// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
	RedeemBy *string `pulumi:"redeemBy"`
}

// The set of arguments for constructing a Coupon resource.
type CouponArgs struct {
	// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
	AmountOff pulumi.IntPtrInput
	// List(String). A list of product IDs this coupon applies to.
	AppliesTos pulumi.StringArrayInput
	// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
	CouponId pulumi.StringPtrInput
	// String. Required if `amountOff` has been set, the three-letter ISO code for the currency of the amount to take off.
	Currency pulumi.StringPtrInput
	// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
	Duration pulumi.StringPtrInput
	// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
	DurationInMonths pulumi.IntPtrInput
	// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
	MaxRedemptions pulumi.IntPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// String. Name of the coupon displayed to customers on for instance invoices or receipts.
	Name pulumi.StringPtrInput
	// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percentOff of 50 will make a $100 invoice $50 instead.
	PercentOff pulumi.Float64PtrInput
	// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
	RedeemBy pulumi.StringPtrInput
}

func (CouponArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*couponArgs)(nil)).Elem()
}

type CouponInput interface {
	pulumi.Input

	ToCouponOutput() CouponOutput
	ToCouponOutputWithContext(ctx context.Context) CouponOutput
}

func (*Coupon) ElementType() reflect.Type {
	return reflect.TypeOf((**Coupon)(nil)).Elem()
}

func (i *Coupon) ToCouponOutput() CouponOutput {
	return i.ToCouponOutputWithContext(context.Background())
}

func (i *Coupon) ToCouponOutputWithContext(ctx context.Context) CouponOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CouponOutput)
}

// CouponArrayInput is an input type that accepts CouponArray and CouponArrayOutput values.
// You can construct a concrete instance of `CouponArrayInput` via:
//
//	CouponArray{ CouponArgs{...} }
type CouponArrayInput interface {
	pulumi.Input

	ToCouponArrayOutput() CouponArrayOutput
	ToCouponArrayOutputWithContext(context.Context) CouponArrayOutput
}

type CouponArray []CouponInput

func (CouponArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Coupon)(nil)).Elem()
}

func (i CouponArray) ToCouponArrayOutput() CouponArrayOutput {
	return i.ToCouponArrayOutputWithContext(context.Background())
}

func (i CouponArray) ToCouponArrayOutputWithContext(ctx context.Context) CouponArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CouponArrayOutput)
}

// CouponMapInput is an input type that accepts CouponMap and CouponMapOutput values.
// You can construct a concrete instance of `CouponMapInput` via:
//
//	CouponMap{ "key": CouponArgs{...} }
type CouponMapInput interface {
	pulumi.Input

	ToCouponMapOutput() CouponMapOutput
	ToCouponMapOutputWithContext(context.Context) CouponMapOutput
}

type CouponMap map[string]CouponInput

func (CouponMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Coupon)(nil)).Elem()
}

func (i CouponMap) ToCouponMapOutput() CouponMapOutput {
	return i.ToCouponMapOutputWithContext(context.Background())
}

func (i CouponMap) ToCouponMapOutputWithContext(ctx context.Context) CouponMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CouponMapOutput)
}

type CouponOutput struct{ *pulumi.OutputState }

func (CouponOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Coupon)(nil)).Elem()
}

func (o CouponOutput) ToCouponOutput() CouponOutput {
	return o
}

func (o CouponOutput) ToCouponOutputWithContext(ctx context.Context) CouponOutput {
	return o
}

// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
func (o CouponOutput) AmountOff() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Coupon) pulumi.IntPtrOutput { return v.AmountOff }).(pulumi.IntPtrOutput)
}

// List(String). A list of product IDs this coupon applies to.
func (o CouponOutput) AppliesTos() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Coupon) pulumi.StringArrayOutput { return v.AppliesTos }).(pulumi.StringArrayOutput)
}

// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
func (o CouponOutput) CouponId() pulumi.StringOutput {
	return o.ApplyT(func(v *Coupon) pulumi.StringOutput { return v.CouponId }).(pulumi.StringOutput)
}

// String. Required if `amountOff` has been set, the three-letter ISO code for the currency of the amount to take off.
func (o CouponOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Coupon) pulumi.StringPtrOutput { return v.Currency }).(pulumi.StringPtrOutput)
}

// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
func (o CouponOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Coupon) pulumi.StringPtrOutput { return v.Duration }).(pulumi.StringPtrOutput)
}

// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
func (o CouponOutput) DurationInMonths() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Coupon) pulumi.IntPtrOutput { return v.DurationInMonths }).(pulumi.IntPtrOutput)
}

// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
func (o CouponOutput) MaxRedemptions() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Coupon) pulumi.IntPtrOutput { return v.MaxRedemptions }).(pulumi.IntPtrOutput)
}

// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
func (o CouponOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Coupon) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// String. Name of the coupon displayed to customers on for instance invoices or receipts.
func (o CouponOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Coupon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percentOff of 50 will make a $100 invoice $50 instead.
func (o CouponOutput) PercentOff() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Coupon) pulumi.Float64PtrOutput { return v.PercentOff }).(pulumi.Float64PtrOutput)
}

// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
func (o CouponOutput) RedeemBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Coupon) pulumi.StringPtrOutput { return v.RedeemBy }).(pulumi.StringPtrOutput)
}

// Int. Number of times this coupon has been applied to a customer.
func (o CouponOutput) TimesRedeemed() pulumi.IntOutput {
	return o.ApplyT(func(v *Coupon) pulumi.IntOutput { return v.TimesRedeemed }).(pulumi.IntOutput)
}

// Bool. Taking account of the above properties, whether this coupon can still be applied to a customer.
func (o CouponOutput) Valid() pulumi.BoolOutput {
	return o.ApplyT(func(v *Coupon) pulumi.BoolOutput { return v.Valid }).(pulumi.BoolOutput)
}

type CouponArrayOutput struct{ *pulumi.OutputState }

func (CouponArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Coupon)(nil)).Elem()
}

func (o CouponArrayOutput) ToCouponArrayOutput() CouponArrayOutput {
	return o
}

func (o CouponArrayOutput) ToCouponArrayOutputWithContext(ctx context.Context) CouponArrayOutput {
	return o
}

func (o CouponArrayOutput) Index(i pulumi.IntInput) CouponOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Coupon {
		return vs[0].([]*Coupon)[vs[1].(int)]
	}).(CouponOutput)
}

type CouponMapOutput struct{ *pulumi.OutputState }

func (CouponMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Coupon)(nil)).Elem()
}

func (o CouponMapOutput) ToCouponMapOutput() CouponMapOutput {
	return o
}

func (o CouponMapOutput) ToCouponMapOutputWithContext(ctx context.Context) CouponMapOutput {
	return o
}

func (o CouponMapOutput) MapIndex(k pulumi.StringInput) CouponOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Coupon {
		return vs[0].(map[string]*Coupon)[vs[1].(string)]
	}).(CouponOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CouponInput)(nil)).Elem(), &Coupon{})
	pulumi.RegisterInputType(reflect.TypeOf((*CouponArrayInput)(nil)).Elem(), CouponArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CouponMapInput)(nil)).Elem(), CouponMap{})
	pulumi.RegisterOutputType(CouponOutput{})
	pulumi.RegisterOutputType(CouponArrayOutput{})
	pulumi.RegisterOutputType(CouponMapOutput{})
}
