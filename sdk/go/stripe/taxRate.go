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

// With this resource, you can create a tax rate - [Stripe API tax rate  documentation](https://stripe.com/docs/api/tax_rates).
//
// Tax rates can be applied to invoices, subscriptions and Checkout Sessions to collect tax.
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
//			_, err := stripe.NewTaxRate(ctx, "taxRate", &stripe.TaxRateArgs{
//				Active:       pulumi.Bool(true),
//				Country:      pulumi.String("AU"),
//				Description:  pulumi.String("GST Australia"),
//				DisplayName:  pulumi.String("GST"),
//				Inclusive:    pulumi.Bool(true),
//				Jurisdiction: pulumi.String("AU"),
//				Metadata:     nil,
//				Percentage:   pulumi.Float64(10),
//				State:        pulumi.String(""),
//				TaxType:      pulumi.String(""),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TaxRate struct {
	pulumi.CustomResourceState

	// Bool. Flag determining whether the tax rate is active or inactive (archived). Inactive tax rates cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// String. Two-letter country code (ISO 3166-1 alpha-2).
	Country pulumi.StringPtrOutput `pulumi:"country"`
	// Int. Time at which the object was created. Measured in seconds since the Unix epoch.
	Created pulumi.IntOutput `pulumi:"created"`
	// String. An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// String. The display name of the tax rate, which will be shown to users.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Bool. This specifies if the tax rate is inclusive or exclusive.
	// * ` percentage  ` - (Required) Float. This represents the tax rate percent out of 100.
	Inclusive pulumi.BoolOutput `pulumi:"inclusive"`
	// String. The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.
	Jurisdiction pulumi.StringPtrOutput `pulumi:"jurisdiction"`
	// Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode pulumi.BoolOutput `pulumi:"livemode"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to metadata.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// String. String representing the object’s type. Objects of the same type share the same value.
	Object pulumi.StringOutput `pulumi:"object"`
	// This represents the tax rate percent out of 100.
	Percentage pulumi.Float64Output `pulumi:"percentage"`
	// String. ISO 3166-2 subdivision code, without country prefix. For example, “NY” for New York, United States.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// String. The high-level tax type, such as vat or sales_tax.
	TaxType pulumi.StringPtrOutput `pulumi:"taxType"`
}

// NewTaxRate registers a new resource with the given unique name, arguments, and options.
func NewTaxRate(ctx *pulumi.Context,
	name string, args *TaxRateArgs, opts ...pulumi.ResourceOption) (*TaxRate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Inclusive == nil {
		return nil, errors.New("invalid value for required argument 'Inclusive'")
	}
	if args.Percentage == nil {
		return nil, errors.New("invalid value for required argument 'Percentage'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TaxRate
	err := ctx.RegisterResource("stripe:index/taxRate:TaxRate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaxRate gets an existing TaxRate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaxRate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaxRateState, opts ...pulumi.ResourceOption) (*TaxRate, error) {
	var resource TaxRate
	err := ctx.ReadResource("stripe:index/taxRate:TaxRate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaxRate resources.
type taxRateState struct {
	// Bool. Flag determining whether the tax rate is active or inactive (archived). Inactive tax rates cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.
	Active *bool `pulumi:"active"`
	// String. Two-letter country code (ISO 3166-1 alpha-2).
	Country *string `pulumi:"country"`
	// Int. Time at which the object was created. Measured in seconds since the Unix epoch.
	Created *int `pulumi:"created"`
	// String. An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.
	Description *string `pulumi:"description"`
	// String. The display name of the tax rate, which will be shown to users.
	DisplayName *string `pulumi:"displayName"`
	// Bool. This specifies if the tax rate is inclusive or exclusive.
	// * ` percentage  ` - (Required) Float. This represents the tax rate percent out of 100.
	Inclusive *bool `pulumi:"inclusive"`
	// String. The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.
	Jurisdiction *string `pulumi:"jurisdiction"`
	// Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode *bool `pulumi:"livemode"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// String. String representing the object’s type. Objects of the same type share the same value.
	Object *string `pulumi:"object"`
	// This represents the tax rate percent out of 100.
	Percentage *float64 `pulumi:"percentage"`
	// String. ISO 3166-2 subdivision code, without country prefix. For example, “NY” for New York, United States.
	State *string `pulumi:"state"`
	// String. The high-level tax type, such as vat or sales_tax.
	TaxType *string `pulumi:"taxType"`
}

type TaxRateState struct {
	// Bool. Flag determining whether the tax rate is active or inactive (archived). Inactive tax rates cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.
	Active pulumi.BoolPtrInput
	// String. Two-letter country code (ISO 3166-1 alpha-2).
	Country pulumi.StringPtrInput
	// Int. Time at which the object was created. Measured in seconds since the Unix epoch.
	Created pulumi.IntPtrInput
	// String. An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.
	Description pulumi.StringPtrInput
	// String. The display name of the tax rate, which will be shown to users.
	DisplayName pulumi.StringPtrInput
	// Bool. This specifies if the tax rate is inclusive or exclusive.
	// * ` percentage  ` - (Required) Float. This represents the tax rate percent out of 100.
	Inclusive pulumi.BoolPtrInput
	// String. The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.
	Jurisdiction pulumi.StringPtrInput
	// Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
	Livemode pulumi.BoolPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to metadata.
	Metadata pulumi.StringMapInput
	// String. String representing the object’s type. Objects of the same type share the same value.
	Object pulumi.StringPtrInput
	// This represents the tax rate percent out of 100.
	Percentage pulumi.Float64PtrInput
	// String. ISO 3166-2 subdivision code, without country prefix. For example, “NY” for New York, United States.
	State pulumi.StringPtrInput
	// String. The high-level tax type, such as vat or sales_tax.
	TaxType pulumi.StringPtrInput
}

func (TaxRateState) ElementType() reflect.Type {
	return reflect.TypeOf((*taxRateState)(nil)).Elem()
}

type taxRateArgs struct {
	// Bool. Flag determining whether the tax rate is active or inactive (archived). Inactive tax rates cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.
	Active *bool `pulumi:"active"`
	// String. Two-letter country code (ISO 3166-1 alpha-2).
	Country *string `pulumi:"country"`
	// String. An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.
	Description *string `pulumi:"description"`
	// String. The display name of the tax rate, which will be shown to users.
	DisplayName string `pulumi:"displayName"`
	// Bool. This specifies if the tax rate is inclusive or exclusive.
	// * ` percentage  ` - (Required) Float. This represents the tax rate percent out of 100.
	Inclusive bool `pulumi:"inclusive"`
	// String. The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.
	Jurisdiction *string `pulumi:"jurisdiction"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// This represents the tax rate percent out of 100.
	Percentage float64 `pulumi:"percentage"`
	// String. ISO 3166-2 subdivision code, without country prefix. For example, “NY” for New York, United States.
	State *string `pulumi:"state"`
	// String. The high-level tax type, such as vat or sales_tax.
	TaxType *string `pulumi:"taxType"`
}

// The set of arguments for constructing a TaxRate resource.
type TaxRateArgs struct {
	// Bool. Flag determining whether the tax rate is active or inactive (archived). Inactive tax rates cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.
	Active pulumi.BoolPtrInput
	// String. Two-letter country code (ISO 3166-1 alpha-2).
	Country pulumi.StringPtrInput
	// String. An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.
	Description pulumi.StringPtrInput
	// String. The display name of the tax rate, which will be shown to users.
	DisplayName pulumi.StringInput
	// Bool. This specifies if the tax rate is inclusive or exclusive.
	// * ` percentage  ` - (Required) Float. This represents the tax rate percent out of 100.
	Inclusive pulumi.BoolInput
	// String. The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.
	Jurisdiction pulumi.StringPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to metadata.
	Metadata pulumi.StringMapInput
	// This represents the tax rate percent out of 100.
	Percentage pulumi.Float64Input
	// String. ISO 3166-2 subdivision code, without country prefix. For example, “NY” for New York, United States.
	State pulumi.StringPtrInput
	// String. The high-level tax type, such as vat or sales_tax.
	TaxType pulumi.StringPtrInput
}

func (TaxRateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taxRateArgs)(nil)).Elem()
}

type TaxRateInput interface {
	pulumi.Input

	ToTaxRateOutput() TaxRateOutput
	ToTaxRateOutputWithContext(ctx context.Context) TaxRateOutput
}

func (*TaxRate) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxRate)(nil)).Elem()
}

func (i *TaxRate) ToTaxRateOutput() TaxRateOutput {
	return i.ToTaxRateOutputWithContext(context.Background())
}

func (i *TaxRate) ToTaxRateOutputWithContext(ctx context.Context) TaxRateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxRateOutput)
}

// TaxRateArrayInput is an input type that accepts TaxRateArray and TaxRateArrayOutput values.
// You can construct a concrete instance of `TaxRateArrayInput` via:
//
//	TaxRateArray{ TaxRateArgs{...} }
type TaxRateArrayInput interface {
	pulumi.Input

	ToTaxRateArrayOutput() TaxRateArrayOutput
	ToTaxRateArrayOutputWithContext(context.Context) TaxRateArrayOutput
}

type TaxRateArray []TaxRateInput

func (TaxRateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaxRate)(nil)).Elem()
}

func (i TaxRateArray) ToTaxRateArrayOutput() TaxRateArrayOutput {
	return i.ToTaxRateArrayOutputWithContext(context.Background())
}

func (i TaxRateArray) ToTaxRateArrayOutputWithContext(ctx context.Context) TaxRateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxRateArrayOutput)
}

// TaxRateMapInput is an input type that accepts TaxRateMap and TaxRateMapOutput values.
// You can construct a concrete instance of `TaxRateMapInput` via:
//
//	TaxRateMap{ "key": TaxRateArgs{...} }
type TaxRateMapInput interface {
	pulumi.Input

	ToTaxRateMapOutput() TaxRateMapOutput
	ToTaxRateMapOutputWithContext(context.Context) TaxRateMapOutput
}

type TaxRateMap map[string]TaxRateInput

func (TaxRateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaxRate)(nil)).Elem()
}

func (i TaxRateMap) ToTaxRateMapOutput() TaxRateMapOutput {
	return i.ToTaxRateMapOutputWithContext(context.Background())
}

func (i TaxRateMap) ToTaxRateMapOutputWithContext(ctx context.Context) TaxRateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxRateMapOutput)
}

type TaxRateOutput struct{ *pulumi.OutputState }

func (TaxRateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxRate)(nil)).Elem()
}

func (o TaxRateOutput) ToTaxRateOutput() TaxRateOutput {
	return o
}

func (o TaxRateOutput) ToTaxRateOutputWithContext(ctx context.Context) TaxRateOutput {
	return o
}

// Bool. Flag determining whether the tax rate is active or inactive (archived). Inactive tax rates cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.
func (o TaxRateOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// String. Two-letter country code (ISO 3166-1 alpha-2).
func (o TaxRateOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.StringPtrOutput { return v.Country }).(pulumi.StringPtrOutput)
}

// Int. Time at which the object was created. Measured in seconds since the Unix epoch.
func (o TaxRateOutput) Created() pulumi.IntOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.IntOutput { return v.Created }).(pulumi.IntOutput)
}

// String. An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.
func (o TaxRateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// String. The display name of the tax rate, which will be shown to users.
func (o TaxRateOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Bool. This specifies if the tax rate is inclusive or exclusive.
// * ` percentage  ` - (Required) Float. This represents the tax rate percent out of 100.
func (o TaxRateOutput) Inclusive() pulumi.BoolOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.BoolOutput { return v.Inclusive }).(pulumi.BoolOutput)
}

// String. The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.
func (o TaxRateOutput) Jurisdiction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.StringPtrOutput { return v.Jurisdiction }).(pulumi.StringPtrOutput)
}

// Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
func (o TaxRateOutput) Livemode() pulumi.BoolOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.BoolOutput { return v.Livemode }).(pulumi.BoolOutput)
}

// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to metadata.
func (o TaxRateOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// String. String representing the object’s type. Objects of the same type share the same value.
func (o TaxRateOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.StringOutput { return v.Object }).(pulumi.StringOutput)
}

// This represents the tax rate percent out of 100.
func (o TaxRateOutput) Percentage() pulumi.Float64Output {
	return o.ApplyT(func(v *TaxRate) pulumi.Float64Output { return v.Percentage }).(pulumi.Float64Output)
}

// String. ISO 3166-2 subdivision code, without country prefix. For example, “NY” for New York, United States.
func (o TaxRateOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// String. The high-level tax type, such as vat or sales_tax.
func (o TaxRateOutput) TaxType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaxRate) pulumi.StringPtrOutput { return v.TaxType }).(pulumi.StringPtrOutput)
}

type TaxRateArrayOutput struct{ *pulumi.OutputState }

func (TaxRateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaxRate)(nil)).Elem()
}

func (o TaxRateArrayOutput) ToTaxRateArrayOutput() TaxRateArrayOutput {
	return o
}

func (o TaxRateArrayOutput) ToTaxRateArrayOutputWithContext(ctx context.Context) TaxRateArrayOutput {
	return o
}

func (o TaxRateArrayOutput) Index(i pulumi.IntInput) TaxRateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TaxRate {
		return vs[0].([]*TaxRate)[vs[1].(int)]
	}).(TaxRateOutput)
}

type TaxRateMapOutput struct{ *pulumi.OutputState }

func (TaxRateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaxRate)(nil)).Elem()
}

func (o TaxRateMapOutput) ToTaxRateMapOutput() TaxRateMapOutput {
	return o
}

func (o TaxRateMapOutput) ToTaxRateMapOutputWithContext(ctx context.Context) TaxRateMapOutput {
	return o
}

func (o TaxRateMapOutput) MapIndex(k pulumi.StringInput) TaxRateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TaxRate {
		return vs[0].(map[string]*TaxRate)[vs[1].(string)]
	}).(TaxRateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaxRateInput)(nil)).Elem(), &TaxRate{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaxRateArrayInput)(nil)).Elem(), TaxRateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaxRateMapInput)(nil)).Elem(), TaxRateMap{})
	pulumi.RegisterOutputType(TaxRateOutput{})
	pulumi.RegisterOutputType(TaxRateArrayOutput{})
	pulumi.RegisterOutputType(TaxRateMapOutput{})
}
