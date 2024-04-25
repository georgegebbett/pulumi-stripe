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

// With this resource, you can create a price - [Stripe API price documentation](https://stripe.com/docs/api/prices).
//
// Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of
// products. Products help you track inventory or provisioning, and prices help you track payment terms.
//
// Different physical goods or levels of service should be represented by products, and pricing options should be
// represented by prices. This approach lets you change prices without having to change your provisioning scheme.
//
// For example, you might have a single "gold" product that has prices for $10/month, $100/year, and €9 once.
//
// > Removal of the price isn't supported through the Stripe SDK. The best practice, which this provider follows,
// is to archive the price by marking it as inactive on destroy, which indicates that the price is no longer
// available for purchase.
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
//			// basic price for the product
//			_, err := stripe.NewPrice(ctx, "pricePrice", &stripe.PriceArgs{
//				Product:    pulumi.Any(stripe_product.Product.Id),
//				Currency:   pulumi.String("aud"),
//				UnitAmount: pulumi.Int(100),
//			})
//			if err != nil {
//				return err
//			}
//			// basic free price for the product
//			_, err = stripe.NewPrice(ctx, "priceIndex/pricePrice", &stripe.PriceArgs{
//				Product:    pulumi.Any(stripe_product.Product.Id),
//				Currency:   pulumi.String("aud"),
//				UnitAmount: -1,
//			})
//			if err != nil {
//				return err
//			}
//			// recurring price for the product
//			_, err = stripe.NewPrice(ctx, "priceStripeIndex/pricePrice", &stripe.PriceArgs{
//				Product:       pulumi.Any(stripe_product.Product.Id),
//				Currency:      pulumi.String("aud"),
//				BillingScheme: pulumi.String("per_unit"),
//				UnitAmount:    pulumi.Int(100),
//				Recurring: &stripe.PriceRecurringArgs{
//					Interval:      pulumi.String("week"),
//					IntervalCount: pulumi.Int(1),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// tiered price for the product
//			_, err = stripe.NewPrice(ctx, "priceStripeIndex/pricePrice1", &stripe.PriceArgs{
//				Product:       pulumi.Any(stripe_product.Product.Id),
//				Currency:      pulumi.String("aud"),
//				BillingScheme: pulumi.String("tiered"),
//				TiersMode:     pulumi.String("graduated"),
//				Tiers: stripe.PriceTierArray{
//					&stripe.PriceTierArgs{
//						UpTo:       pulumi.Int(10),
//						UnitAmount: pulumi.Int(0),
//					},
//					&stripe.PriceTierArgs{
//						UpTo:       pulumi.Int(100),
//						UnitAmount: pulumi.Int(300),
//					},
//					&stripe.PriceTierArgs{
//						UpTo:              -1,
//						UnitAmountDecimal: pulumi.Float64(100.5),
//					},
//				},
//				Recurring: &stripe.PriceRecurringArgs{
//					Interval:       pulumi.String("week"),
//					AggregateUsage: pulumi.String("sum"),
//					IntervalCount:  pulumi.Int(2),
//					UsageType:      pulumi.String("metered"),
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
// ## Note on updating prices
//
// Once created, you can update the `active`, `metadata`, `nickname`, `lookupKey`, `taxBehaviour` (only if unspecified)
// and `transferLookupKey` attributes.
//
// Other attribute edits will trigger a destroy action (archival) and creation of a new price entry.
type Price struct {
	pulumi.CustomResourceState

	// Bool. Whether the price can be used for new purchases. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// String. Describes how to compute the price per period. Either `perUnit` or `tiered`
	// . `perUnit` indicates that the fixed amount (specified in `unitAmount` or `unitAmountDecimal`) will be charged per
	// unit in quantity (for prices with `usage_type=licensed`), or per unit of total usage (for prices
	// with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as
	// defined using the `tiers` and `tiersMode` attributes.
	BillingScheme pulumi.StringOutput `pulumi:"billingScheme"`
	// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
	Currency pulumi.StringOutput `pulumi:"currency"`
	// List(Resource). Prices defined in each available currency option. For details
	// of individual arguments see Currency Options.
	CurrencyOptions PriceCurrencyOptionArrayOutput `pulumi:"currencyOptions"`
	// String. A lookup key used to retrieve prices dynamically from a static string.
	LookupKey pulumi.StringPtrOutput `pulumi:"lookupKey"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// String. A brief description of the price, hidden from customers.
	Nickname pulumi.StringPtrOutput `pulumi:"nickname"`
	// String. The ID of the product that this price will belong to.
	Product pulumi.StringOutput `pulumi:"product"`
	// List(Resource). The recurring components of a price such as `interval` and `usageType`. For
	// details of individual arguments see Recurring.
	Recurring PriceRecurringPtrOutput `pulumi:"recurring"`
	// String. Specifies whether the price is considered inclusive of taxes or exclusive of
	// taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it
	// cannot be changed, default is `unspecified`.
	TaxBehaviour pulumi.StringPtrOutput `pulumi:"taxBehaviour"`
	// List(Resource). Each element represents a pricing tier. This parameter requires `billingScheme`
	// to be set to `tiered`. See also the documentation for `billingScheme`. For details of individual arguments
	// see Tiers.
	Tiers PriceTierArrayOutput `pulumi:"tiers"`
	// String. Defines if the tiering price should be `graduated`
	// or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per-unit price,
	// in `graduated` tiering pricing can successively change as the quantity grows.
	TiersMode pulumi.StringPtrOutput `pulumi:"tiersMode"`
	// Bool. If set to `true`, will atomically remove the lookup key from the existing
	// price, and assign it to this price.
	TransferLookupKey pulumi.BoolPtrOutput `pulumi:"transferLookupKey"`
	// List(Resource). Apply a transformation to the reported usage or set quantity before
	// computing the billed price. Cannot be combined with `tiers`. For details of individual arguments
	// see Transform Quantity.
	TransformQuantity PriceTransformQuantityPtrOutput `pulumi:"transformQuantity"`
	// String. One of `oneTime` or `recurring` depending on whether the price is for a one-time purchase or a
	// recurring (subscription) purchase.
	Type pulumi.StringOutput `pulumi:"type"`
	// Int. A positive integer in cents (or `-1` for a free
	// price) representing how much to charge.
	UnitAmount pulumi.IntOutput `pulumi:"unitAmount"`
	// Float. Same as `unitAmount`, but accepts a decimal value in cents with at most 12
	// decimal places. Only one of `unitAmount` and `unitAmountDecimal` can be set.
	UnitAmountDecimal pulumi.Float64Output `pulumi:"unitAmountDecimal"`
}

// NewPrice registers a new resource with the given unique name, arguments, and options.
func NewPrice(ctx *pulumi.Context,
	name string, args *PriceArgs, opts ...pulumi.ResourceOption) (*Price, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Currency == nil {
		return nil, errors.New("invalid value for required argument 'Currency'")
	}
	if args.Product == nil {
		return nil, errors.New("invalid value for required argument 'Product'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Price
	err := ctx.RegisterResource("stripe:index/price:Price", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrice gets an existing Price resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PriceState, opts ...pulumi.ResourceOption) (*Price, error) {
	var resource Price
	err := ctx.ReadResource("stripe:index/price:Price", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Price resources.
type priceState struct {
	// Bool. Whether the price can be used for new purchases. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// String. Describes how to compute the price per period. Either `perUnit` or `tiered`
	// . `perUnit` indicates that the fixed amount (specified in `unitAmount` or `unitAmountDecimal`) will be charged per
	// unit in quantity (for prices with `usage_type=licensed`), or per unit of total usage (for prices
	// with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as
	// defined using the `tiers` and `tiersMode` attributes.
	BillingScheme *string `pulumi:"billingScheme"`
	// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
	Currency *string `pulumi:"currency"`
	// List(Resource). Prices defined in each available currency option. For details
	// of individual arguments see Currency Options.
	CurrencyOptions []PriceCurrencyOption `pulumi:"currencyOptions"`
	// String. A lookup key used to retrieve prices dynamically from a static string.
	LookupKey *string `pulumi:"lookupKey"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// String. A brief description of the price, hidden from customers.
	Nickname *string `pulumi:"nickname"`
	// String. The ID of the product that this price will belong to.
	Product *string `pulumi:"product"`
	// List(Resource). The recurring components of a price such as `interval` and `usageType`. For
	// details of individual arguments see Recurring.
	Recurring *PriceRecurring `pulumi:"recurring"`
	// String. Specifies whether the price is considered inclusive of taxes or exclusive of
	// taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it
	// cannot be changed, default is `unspecified`.
	TaxBehaviour *string `pulumi:"taxBehaviour"`
	// List(Resource). Each element represents a pricing tier. This parameter requires `billingScheme`
	// to be set to `tiered`. See also the documentation for `billingScheme`. For details of individual arguments
	// see Tiers.
	Tiers []PriceTier `pulumi:"tiers"`
	// String. Defines if the tiering price should be `graduated`
	// or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per-unit price,
	// in `graduated` tiering pricing can successively change as the quantity grows.
	TiersMode *string `pulumi:"tiersMode"`
	// Bool. If set to `true`, will atomically remove the lookup key from the existing
	// price, and assign it to this price.
	TransferLookupKey *bool `pulumi:"transferLookupKey"`
	// List(Resource). Apply a transformation to the reported usage or set quantity before
	// computing the billed price. Cannot be combined with `tiers`. For details of individual arguments
	// see Transform Quantity.
	TransformQuantity *PriceTransformQuantity `pulumi:"transformQuantity"`
	// String. One of `oneTime` or `recurring` depending on whether the price is for a one-time purchase or a
	// recurring (subscription) purchase.
	Type *string `pulumi:"type"`
	// Int. A positive integer in cents (or `-1` for a free
	// price) representing how much to charge.
	UnitAmount *int `pulumi:"unitAmount"`
	// Float. Same as `unitAmount`, but accepts a decimal value in cents with at most 12
	// decimal places. Only one of `unitAmount` and `unitAmountDecimal` can be set.
	UnitAmountDecimal *float64 `pulumi:"unitAmountDecimal"`
}

type PriceState struct {
	// Bool. Whether the price can be used for new purchases. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// String. Describes how to compute the price per period. Either `perUnit` or `tiered`
	// . `perUnit` indicates that the fixed amount (specified in `unitAmount` or `unitAmountDecimal`) will be charged per
	// unit in quantity (for prices with `usage_type=licensed`), or per unit of total usage (for prices
	// with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as
	// defined using the `tiers` and `tiersMode` attributes.
	BillingScheme pulumi.StringPtrInput
	// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
	Currency pulumi.StringPtrInput
	// List(Resource). Prices defined in each available currency option. For details
	// of individual arguments see Currency Options.
	CurrencyOptions PriceCurrencyOptionArrayInput
	// String. A lookup key used to retrieve prices dynamically from a static string.
	LookupKey pulumi.StringPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// String. A brief description of the price, hidden from customers.
	Nickname pulumi.StringPtrInput
	// String. The ID of the product that this price will belong to.
	Product pulumi.StringPtrInput
	// List(Resource). The recurring components of a price such as `interval` and `usageType`. For
	// details of individual arguments see Recurring.
	Recurring PriceRecurringPtrInput
	// String. Specifies whether the price is considered inclusive of taxes or exclusive of
	// taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it
	// cannot be changed, default is `unspecified`.
	TaxBehaviour pulumi.StringPtrInput
	// List(Resource). Each element represents a pricing tier. This parameter requires `billingScheme`
	// to be set to `tiered`. See also the documentation for `billingScheme`. For details of individual arguments
	// see Tiers.
	Tiers PriceTierArrayInput
	// String. Defines if the tiering price should be `graduated`
	// or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per-unit price,
	// in `graduated` tiering pricing can successively change as the quantity grows.
	TiersMode pulumi.StringPtrInput
	// Bool. If set to `true`, will atomically remove the lookup key from the existing
	// price, and assign it to this price.
	TransferLookupKey pulumi.BoolPtrInput
	// List(Resource). Apply a transformation to the reported usage or set quantity before
	// computing the billed price. Cannot be combined with `tiers`. For details of individual arguments
	// see Transform Quantity.
	TransformQuantity PriceTransformQuantityPtrInput
	// String. One of `oneTime` or `recurring` depending on whether the price is for a one-time purchase or a
	// recurring (subscription) purchase.
	Type pulumi.StringPtrInput
	// Int. A positive integer in cents (or `-1` for a free
	// price) representing how much to charge.
	UnitAmount pulumi.IntPtrInput
	// Float. Same as `unitAmount`, but accepts a decimal value in cents with at most 12
	// decimal places. Only one of `unitAmount` and `unitAmountDecimal` can be set.
	UnitAmountDecimal pulumi.Float64PtrInput
}

func (PriceState) ElementType() reflect.Type {
	return reflect.TypeOf((*priceState)(nil)).Elem()
}

type priceArgs struct {
	// Bool. Whether the price can be used for new purchases. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// String. Describes how to compute the price per period. Either `perUnit` or `tiered`
	// . `perUnit` indicates that the fixed amount (specified in `unitAmount` or `unitAmountDecimal`) will be charged per
	// unit in quantity (for prices with `usage_type=licensed`), or per unit of total usage (for prices
	// with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as
	// defined using the `tiers` and `tiersMode` attributes.
	BillingScheme *string `pulumi:"billingScheme"`
	// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
	Currency string `pulumi:"currency"`
	// List(Resource). Prices defined in each available currency option. For details
	// of individual arguments see Currency Options.
	CurrencyOptions []PriceCurrencyOption `pulumi:"currencyOptions"`
	// String. A lookup key used to retrieve prices dynamically from a static string.
	LookupKey *string `pulumi:"lookupKey"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// String. A brief description of the price, hidden from customers.
	Nickname *string `pulumi:"nickname"`
	// String. The ID of the product that this price will belong to.
	Product string `pulumi:"product"`
	// List(Resource). The recurring components of a price such as `interval` and `usageType`. For
	// details of individual arguments see Recurring.
	Recurring *PriceRecurring `pulumi:"recurring"`
	// String. Specifies whether the price is considered inclusive of taxes or exclusive of
	// taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it
	// cannot be changed, default is `unspecified`.
	TaxBehaviour *string `pulumi:"taxBehaviour"`
	// List(Resource). Each element represents a pricing tier. This parameter requires `billingScheme`
	// to be set to `tiered`. See also the documentation for `billingScheme`. For details of individual arguments
	// see Tiers.
	Tiers []PriceTier `pulumi:"tiers"`
	// String. Defines if the tiering price should be `graduated`
	// or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per-unit price,
	// in `graduated` tiering pricing can successively change as the quantity grows.
	TiersMode *string `pulumi:"tiersMode"`
	// Bool. If set to `true`, will atomically remove the lookup key from the existing
	// price, and assign it to this price.
	TransferLookupKey *bool `pulumi:"transferLookupKey"`
	// List(Resource). Apply a transformation to the reported usage or set quantity before
	// computing the billed price. Cannot be combined with `tiers`. For details of individual arguments
	// see Transform Quantity.
	TransformQuantity *PriceTransformQuantity `pulumi:"transformQuantity"`
	// Int. A positive integer in cents (or `-1` for a free
	// price) representing how much to charge.
	UnitAmount *int `pulumi:"unitAmount"`
	// Float. Same as `unitAmount`, but accepts a decimal value in cents with at most 12
	// decimal places. Only one of `unitAmount` and `unitAmountDecimal` can be set.
	UnitAmountDecimal *float64 `pulumi:"unitAmountDecimal"`
}

// The set of arguments for constructing a Price resource.
type PriceArgs struct {
	// Bool. Whether the price can be used for new purchases. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// String. Describes how to compute the price per period. Either `perUnit` or `tiered`
	// . `perUnit` indicates that the fixed amount (specified in `unitAmount` or `unitAmountDecimal`) will be charged per
	// unit in quantity (for prices with `usage_type=licensed`), or per unit of total usage (for prices
	// with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as
	// defined using the `tiers` and `tiersMode` attributes.
	BillingScheme pulumi.StringPtrInput
	// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
	Currency pulumi.StringInput
	// List(Resource). Prices defined in each available currency option. For details
	// of individual arguments see Currency Options.
	CurrencyOptions PriceCurrencyOptionArrayInput
	// String. A lookup key used to retrieve prices dynamically from a static string.
	LookupKey pulumi.StringPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// String. A brief description of the price, hidden from customers.
	Nickname pulumi.StringPtrInput
	// String. The ID of the product that this price will belong to.
	Product pulumi.StringInput
	// List(Resource). The recurring components of a price such as `interval` and `usageType`. For
	// details of individual arguments see Recurring.
	Recurring PriceRecurringPtrInput
	// String. Specifies whether the price is considered inclusive of taxes or exclusive of
	// taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it
	// cannot be changed, default is `unspecified`.
	TaxBehaviour pulumi.StringPtrInput
	// List(Resource). Each element represents a pricing tier. This parameter requires `billingScheme`
	// to be set to `tiered`. See also the documentation for `billingScheme`. For details of individual arguments
	// see Tiers.
	Tiers PriceTierArrayInput
	// String. Defines if the tiering price should be `graduated`
	// or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per-unit price,
	// in `graduated` tiering pricing can successively change as the quantity grows.
	TiersMode pulumi.StringPtrInput
	// Bool. If set to `true`, will atomically remove the lookup key from the existing
	// price, and assign it to this price.
	TransferLookupKey pulumi.BoolPtrInput
	// List(Resource). Apply a transformation to the reported usage or set quantity before
	// computing the billed price. Cannot be combined with `tiers`. For details of individual arguments
	// see Transform Quantity.
	TransformQuantity PriceTransformQuantityPtrInput
	// Int. A positive integer in cents (or `-1` for a free
	// price) representing how much to charge.
	UnitAmount pulumi.IntPtrInput
	// Float. Same as `unitAmount`, but accepts a decimal value in cents with at most 12
	// decimal places. Only one of `unitAmount` and `unitAmountDecimal` can be set.
	UnitAmountDecimal pulumi.Float64PtrInput
}

func (PriceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*priceArgs)(nil)).Elem()
}

type PriceInput interface {
	pulumi.Input

	ToPriceOutput() PriceOutput
	ToPriceOutputWithContext(ctx context.Context) PriceOutput
}

func (*Price) ElementType() reflect.Type {
	return reflect.TypeOf((**Price)(nil)).Elem()
}

func (i *Price) ToPriceOutput() PriceOutput {
	return i.ToPriceOutputWithContext(context.Background())
}

func (i *Price) ToPriceOutputWithContext(ctx context.Context) PriceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriceOutput)
}

// PriceArrayInput is an input type that accepts PriceArray and PriceArrayOutput values.
// You can construct a concrete instance of `PriceArrayInput` via:
//
//	PriceArray{ PriceArgs{...} }
type PriceArrayInput interface {
	pulumi.Input

	ToPriceArrayOutput() PriceArrayOutput
	ToPriceArrayOutputWithContext(context.Context) PriceArrayOutput
}

type PriceArray []PriceInput

func (PriceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Price)(nil)).Elem()
}

func (i PriceArray) ToPriceArrayOutput() PriceArrayOutput {
	return i.ToPriceArrayOutputWithContext(context.Background())
}

func (i PriceArray) ToPriceArrayOutputWithContext(ctx context.Context) PriceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriceArrayOutput)
}

// PriceMapInput is an input type that accepts PriceMap and PriceMapOutput values.
// You can construct a concrete instance of `PriceMapInput` via:
//
//	PriceMap{ "key": PriceArgs{...} }
type PriceMapInput interface {
	pulumi.Input

	ToPriceMapOutput() PriceMapOutput
	ToPriceMapOutputWithContext(context.Context) PriceMapOutput
}

type PriceMap map[string]PriceInput

func (PriceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Price)(nil)).Elem()
}

func (i PriceMap) ToPriceMapOutput() PriceMapOutput {
	return i.ToPriceMapOutputWithContext(context.Background())
}

func (i PriceMap) ToPriceMapOutputWithContext(ctx context.Context) PriceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PriceMapOutput)
}

type PriceOutput struct{ *pulumi.OutputState }

func (PriceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Price)(nil)).Elem()
}

func (o PriceOutput) ToPriceOutput() PriceOutput {
	return o
}

func (o PriceOutput) ToPriceOutputWithContext(ctx context.Context) PriceOutput {
	return o
}

// Bool. Whether the price can be used for new purchases. Defaults to `true`.
func (o PriceOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Price) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// String. Describes how to compute the price per period. Either `perUnit` or `tiered`
// . `perUnit` indicates that the fixed amount (specified in `unitAmount` or `unitAmountDecimal`) will be charged per
// unit in quantity (for prices with `usage_type=licensed`), or per unit of total usage (for prices
// with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as
// defined using the `tiers` and `tiersMode` attributes.
func (o PriceOutput) BillingScheme() pulumi.StringOutput {
	return o.ApplyT(func(v *Price) pulumi.StringOutput { return v.BillingScheme }).(pulumi.StringOutput)
}

// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
func (o PriceOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v *Price) pulumi.StringOutput { return v.Currency }).(pulumi.StringOutput)
}

// List(Resource). Prices defined in each available currency option. For details
// of individual arguments see Currency Options.
func (o PriceOutput) CurrencyOptions() PriceCurrencyOptionArrayOutput {
	return o.ApplyT(func(v *Price) PriceCurrencyOptionArrayOutput { return v.CurrencyOptions }).(PriceCurrencyOptionArrayOutput)
}

// String. A lookup key used to retrieve prices dynamically from a static string.
func (o PriceOutput) LookupKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Price) pulumi.StringPtrOutput { return v.LookupKey }).(pulumi.StringPtrOutput)
}

// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
// storing additional information about the object in a structured format.
func (o PriceOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Price) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// String. A brief description of the price, hidden from customers.
func (o PriceOutput) Nickname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Price) pulumi.StringPtrOutput { return v.Nickname }).(pulumi.StringPtrOutput)
}

// String. The ID of the product that this price will belong to.
func (o PriceOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v *Price) pulumi.StringOutput { return v.Product }).(pulumi.StringOutput)
}

// List(Resource). The recurring components of a price such as `interval` and `usageType`. For
// details of individual arguments see Recurring.
func (o PriceOutput) Recurring() PriceRecurringPtrOutput {
	return o.ApplyT(func(v *Price) PriceRecurringPtrOutput { return v.Recurring }).(PriceRecurringPtrOutput)
}

// String. Specifies whether the price is considered inclusive of taxes or exclusive of
// taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it
// cannot be changed, default is `unspecified`.
func (o PriceOutput) TaxBehaviour() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Price) pulumi.StringPtrOutput { return v.TaxBehaviour }).(pulumi.StringPtrOutput)
}

// List(Resource). Each element represents a pricing tier. This parameter requires `billingScheme`
// to be set to `tiered`. See also the documentation for `billingScheme`. For details of individual arguments
// see Tiers.
func (o PriceOutput) Tiers() PriceTierArrayOutput {
	return o.ApplyT(func(v *Price) PriceTierArrayOutput { return v.Tiers }).(PriceTierArrayOutput)
}

// String. Defines if the tiering price should be `graduated`
// or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per-unit price,
// in `graduated` tiering pricing can successively change as the quantity grows.
func (o PriceOutput) TiersMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Price) pulumi.StringPtrOutput { return v.TiersMode }).(pulumi.StringPtrOutput)
}

// Bool. If set to `true`, will atomically remove the lookup key from the existing
// price, and assign it to this price.
func (o PriceOutput) TransferLookupKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Price) pulumi.BoolPtrOutput { return v.TransferLookupKey }).(pulumi.BoolPtrOutput)
}

// List(Resource). Apply a transformation to the reported usage or set quantity before
// computing the billed price. Cannot be combined with `tiers`. For details of individual arguments
// see Transform Quantity.
func (o PriceOutput) TransformQuantity() PriceTransformQuantityPtrOutput {
	return o.ApplyT(func(v *Price) PriceTransformQuantityPtrOutput { return v.TransformQuantity }).(PriceTransformQuantityPtrOutput)
}

// String. One of `oneTime` or `recurring` depending on whether the price is for a one-time purchase or a
// recurring (subscription) purchase.
func (o PriceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Price) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Int. A positive integer in cents (or `-1` for a free
// price) representing how much to charge.
func (o PriceOutput) UnitAmount() pulumi.IntOutput {
	return o.ApplyT(func(v *Price) pulumi.IntOutput { return v.UnitAmount }).(pulumi.IntOutput)
}

// Float. Same as `unitAmount`, but accepts a decimal value in cents with at most 12
// decimal places. Only one of `unitAmount` and `unitAmountDecimal` can be set.
func (o PriceOutput) UnitAmountDecimal() pulumi.Float64Output {
	return o.ApplyT(func(v *Price) pulumi.Float64Output { return v.UnitAmountDecimal }).(pulumi.Float64Output)
}

type PriceArrayOutput struct{ *pulumi.OutputState }

func (PriceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Price)(nil)).Elem()
}

func (o PriceArrayOutput) ToPriceArrayOutput() PriceArrayOutput {
	return o
}

func (o PriceArrayOutput) ToPriceArrayOutputWithContext(ctx context.Context) PriceArrayOutput {
	return o
}

func (o PriceArrayOutput) Index(i pulumi.IntInput) PriceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Price {
		return vs[0].([]*Price)[vs[1].(int)]
	}).(PriceOutput)
}

type PriceMapOutput struct{ *pulumi.OutputState }

func (PriceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Price)(nil)).Elem()
}

func (o PriceMapOutput) ToPriceMapOutput() PriceMapOutput {
	return o
}

func (o PriceMapOutput) ToPriceMapOutputWithContext(ctx context.Context) PriceMapOutput {
	return o
}

func (o PriceMapOutput) MapIndex(k pulumi.StringInput) PriceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Price {
		return vs[0].(map[string]*Price)[vs[1].(string)]
	}).(PriceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PriceInput)(nil)).Elem(), &Price{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriceArrayInput)(nil)).Elem(), PriceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PriceMapInput)(nil)).Elem(), PriceMap{})
	pulumi.RegisterOutputType(PriceOutput{})
	pulumi.RegisterOutputType(PriceArrayOutput{})
	pulumi.RegisterOutputType(PriceMapOutput{})
}
