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

// With this resource, you can create a promotion code - [Stripe API promotion code documentation](https://stripe.com/docs/api/promotion_codes).
//
// A Promotion Code represents a customer-redeemable code for a coupon. It can be used to create multiple codes for a single coupon.
//
// > Removal of the promotion code isn't supported through the Stripe SDK.
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
//			// promotion code for the coupon
//			_, err := stripe.NewPromotionCode(ctx, "codePromotionCode", &stripe.PromotionCodeArgs{
//				Coupon: pulumi.Any(stripe_coupon.Coupon.Id),
//				Code:   pulumi.String("FREE"),
//			})
//			if err != nil {
//				return err
//			}
//			// promotion code for the coupon with limitations
//			_, err = stripe.NewPromotionCode(ctx, "codeIndex/promotionCodePromotionCode", &stripe.PromotionCodeArgs{
//				Coupon:         pulumi.Any(stripe_coupon.Coupon.Id),
//				Code:           pulumi.String("FREE"),
//				MaxRedemptions: pulumi.Int(5),
//				ExpiresAt:      pulumi.String("2025-08-03T08:37:18+00:00"),
//			})
//			if err != nil {
//				return err
//			}
//			// promotion code for the coupon to customer
//			_, err = stripe.NewPromotionCode(ctx, "codeStripeIndex/promotionCodePromotionCode", &stripe.PromotionCodeArgs{
//				Coupon:   pulumi.Any(stripe_coupon.Coupon.Id),
//				Code:     pulumi.String("FREE"),
//				Customer: pulumi.String("cus..."),
//			})
//			if err != nil {
//				return err
//			}
//			// promotion code for the coupon with restrictions
//			_, err = stripe.NewPromotionCode(ctx, "codeStripeIndex/promotionCodePromotionCode1", &stripe.PromotionCodeArgs{
//				Coupon: pulumi.Any(stripe_coupon.Coupon.Id),
//				Code:   pulumi.String("FREE"),
//				Restrictions: &stripe.PromotionCodeRestrictionsArgs{
//					FirstTimeTransaction:  pulumi.Bool(true),
//					MinimumAmount:         pulumi.Int(100),
//					MinimumAmountCurrency: pulumi.String("aud"),
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
type PromotionCode struct {
	pulumi.CustomResourceState

	// Bool. Whether the promotion code is currently active. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
	Code pulumi.StringPtrOutput `pulumi:"code"`
	// String. The coupon for this promotion code.
	Coupon pulumi.StringOutput `pulumi:"coupon"`
	// String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
	Customer pulumi.StringPtrOutput `pulumi:"customer"`
	// String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
	MaxRedemptions pulumi.IntPtrOutput `pulumi:"maxRedemptions"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
	Restrictions PromotionCodeRestrictionsPtrOutput `pulumi:"restrictions"`
}

// NewPromotionCode registers a new resource with the given unique name, arguments, and options.
func NewPromotionCode(ctx *pulumi.Context,
	name string, args *PromotionCodeArgs, opts ...pulumi.ResourceOption) (*PromotionCode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Coupon == nil {
		return nil, errors.New("invalid value for required argument 'Coupon'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PromotionCode
	err := ctx.RegisterResource("stripe:index/promotionCode:PromotionCode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPromotionCode gets an existing PromotionCode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPromotionCode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PromotionCodeState, opts ...pulumi.ResourceOption) (*PromotionCode, error) {
	var resource PromotionCode
	err := ctx.ReadResource("stripe:index/promotionCode:PromotionCode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PromotionCode resources.
type promotionCodeState struct {
	// Bool. Whether the promotion code is currently active. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
	Code *string `pulumi:"code"`
	// String. The coupon for this promotion code.
	Coupon *string `pulumi:"coupon"`
	// String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
	Customer *string `pulumi:"customer"`
	// String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
	ExpiresAt *string `pulumi:"expiresAt"`
	// Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
	MaxRedemptions *int `pulumi:"maxRedemptions"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
	Restrictions *PromotionCodeRestrictions `pulumi:"restrictions"`
}

type PromotionCodeState struct {
	// Bool. Whether the promotion code is currently active. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
	Code pulumi.StringPtrInput
	// String. The coupon for this promotion code.
	Coupon pulumi.StringPtrInput
	// String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
	Customer pulumi.StringPtrInput
	// String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
	ExpiresAt pulumi.StringPtrInput
	// Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
	MaxRedemptions pulumi.IntPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
	Restrictions PromotionCodeRestrictionsPtrInput
}

func (PromotionCodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*promotionCodeState)(nil)).Elem()
}

type promotionCodeArgs struct {
	// Bool. Whether the promotion code is currently active. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
	Code *string `pulumi:"code"`
	// String. The coupon for this promotion code.
	Coupon string `pulumi:"coupon"`
	// String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
	Customer *string `pulumi:"customer"`
	// String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
	ExpiresAt *string `pulumi:"expiresAt"`
	// Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
	MaxRedemptions *int `pulumi:"maxRedemptions"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
	Restrictions *PromotionCodeRestrictions `pulumi:"restrictions"`
}

// The set of arguments for constructing a PromotionCode resource.
type PromotionCodeArgs struct {
	// Bool. Whether the promotion code is currently active. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
	Code pulumi.StringPtrInput
	// String. The coupon for this promotion code.
	Coupon pulumi.StringInput
	// String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
	Customer pulumi.StringPtrInput
	// String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
	ExpiresAt pulumi.StringPtrInput
	// Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
	MaxRedemptions pulumi.IntPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
	Restrictions PromotionCodeRestrictionsPtrInput
}

func (PromotionCodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*promotionCodeArgs)(nil)).Elem()
}

type PromotionCodeInput interface {
	pulumi.Input

	ToPromotionCodeOutput() PromotionCodeOutput
	ToPromotionCodeOutputWithContext(ctx context.Context) PromotionCodeOutput
}

func (*PromotionCode) ElementType() reflect.Type {
	return reflect.TypeOf((**PromotionCode)(nil)).Elem()
}

func (i *PromotionCode) ToPromotionCodeOutput() PromotionCodeOutput {
	return i.ToPromotionCodeOutputWithContext(context.Background())
}

func (i *PromotionCode) ToPromotionCodeOutputWithContext(ctx context.Context) PromotionCodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PromotionCodeOutput)
}

// PromotionCodeArrayInput is an input type that accepts PromotionCodeArray and PromotionCodeArrayOutput values.
// You can construct a concrete instance of `PromotionCodeArrayInput` via:
//
//	PromotionCodeArray{ PromotionCodeArgs{...} }
type PromotionCodeArrayInput interface {
	pulumi.Input

	ToPromotionCodeArrayOutput() PromotionCodeArrayOutput
	ToPromotionCodeArrayOutputWithContext(context.Context) PromotionCodeArrayOutput
}

type PromotionCodeArray []PromotionCodeInput

func (PromotionCodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PromotionCode)(nil)).Elem()
}

func (i PromotionCodeArray) ToPromotionCodeArrayOutput() PromotionCodeArrayOutput {
	return i.ToPromotionCodeArrayOutputWithContext(context.Background())
}

func (i PromotionCodeArray) ToPromotionCodeArrayOutputWithContext(ctx context.Context) PromotionCodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PromotionCodeArrayOutput)
}

// PromotionCodeMapInput is an input type that accepts PromotionCodeMap and PromotionCodeMapOutput values.
// You can construct a concrete instance of `PromotionCodeMapInput` via:
//
//	PromotionCodeMap{ "key": PromotionCodeArgs{...} }
type PromotionCodeMapInput interface {
	pulumi.Input

	ToPromotionCodeMapOutput() PromotionCodeMapOutput
	ToPromotionCodeMapOutputWithContext(context.Context) PromotionCodeMapOutput
}

type PromotionCodeMap map[string]PromotionCodeInput

func (PromotionCodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PromotionCode)(nil)).Elem()
}

func (i PromotionCodeMap) ToPromotionCodeMapOutput() PromotionCodeMapOutput {
	return i.ToPromotionCodeMapOutputWithContext(context.Background())
}

func (i PromotionCodeMap) ToPromotionCodeMapOutputWithContext(ctx context.Context) PromotionCodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PromotionCodeMapOutput)
}

type PromotionCodeOutput struct{ *pulumi.OutputState }

func (PromotionCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PromotionCode)(nil)).Elem()
}

func (o PromotionCodeOutput) ToPromotionCodeOutput() PromotionCodeOutput {
	return o
}

func (o PromotionCodeOutput) ToPromotionCodeOutputWithContext(ctx context.Context) PromotionCodeOutput {
	return o
}

// Bool. Whether the promotion code is currently active. Defaults to `true`.
func (o PromotionCodeOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PromotionCode) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
func (o PromotionCodeOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PromotionCode) pulumi.StringPtrOutput { return v.Code }).(pulumi.StringPtrOutput)
}

// String. The coupon for this promotion code.
func (o PromotionCodeOutput) Coupon() pulumi.StringOutput {
	return o.ApplyT(func(v *PromotionCode) pulumi.StringOutput { return v.Coupon }).(pulumi.StringOutput)
}

// String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
func (o PromotionCodeOutput) Customer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PromotionCode) pulumi.StringPtrOutput { return v.Customer }).(pulumi.StringPtrOutput)
}

// String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
func (o PromotionCodeOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PromotionCode) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
func (o PromotionCodeOutput) MaxRedemptions() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PromotionCode) pulumi.IntPtrOutput { return v.MaxRedemptions }).(pulumi.IntPtrOutput)
}

// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
func (o PromotionCodeOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PromotionCode) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
func (o PromotionCodeOutput) Restrictions() PromotionCodeRestrictionsPtrOutput {
	return o.ApplyT(func(v *PromotionCode) PromotionCodeRestrictionsPtrOutput { return v.Restrictions }).(PromotionCodeRestrictionsPtrOutput)
}

type PromotionCodeArrayOutput struct{ *pulumi.OutputState }

func (PromotionCodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PromotionCode)(nil)).Elem()
}

func (o PromotionCodeArrayOutput) ToPromotionCodeArrayOutput() PromotionCodeArrayOutput {
	return o
}

func (o PromotionCodeArrayOutput) ToPromotionCodeArrayOutputWithContext(ctx context.Context) PromotionCodeArrayOutput {
	return o
}

func (o PromotionCodeArrayOutput) Index(i pulumi.IntInput) PromotionCodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PromotionCode {
		return vs[0].([]*PromotionCode)[vs[1].(int)]
	}).(PromotionCodeOutput)
}

type PromotionCodeMapOutput struct{ *pulumi.OutputState }

func (PromotionCodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PromotionCode)(nil)).Elem()
}

func (o PromotionCodeMapOutput) ToPromotionCodeMapOutput() PromotionCodeMapOutput {
	return o
}

func (o PromotionCodeMapOutput) ToPromotionCodeMapOutputWithContext(ctx context.Context) PromotionCodeMapOutput {
	return o
}

func (o PromotionCodeMapOutput) MapIndex(k pulumi.StringInput) PromotionCodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PromotionCode {
		return vs[0].(map[string]*PromotionCode)[vs[1].(string)]
	}).(PromotionCodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PromotionCodeInput)(nil)).Elem(), &PromotionCode{})
	pulumi.RegisterInputType(reflect.TypeOf((*PromotionCodeArrayInput)(nil)).Elem(), PromotionCodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PromotionCodeMapInput)(nil)).Elem(), PromotionCodeMap{})
	pulumi.RegisterOutputType(PromotionCodeOutput{})
	pulumi.RegisterOutputType(PromotionCodeArrayOutput{})
	pulumi.RegisterOutputType(PromotionCodeMapOutput{})
}
