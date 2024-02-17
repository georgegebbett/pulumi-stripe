// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package stripe

import (
	"context"
	"reflect"

	"github.com/georgegebbett/pulumi-stripe/sdk/go/stripe/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			_, err := stripe.NewProduct(ctx, "product", &stripe.ProductArgs{
//				Description: pulumi.String("fantastic product"),
//				UnitLabel:   pulumi.String("piece"),
//				Url:         pulumi.String("https://www.terraform.io"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Product struct {
	pulumi.CustomResourceState

	// Bool. Whether the product is currently available for purchase. Defaults to `true`.
	Active pulumi.BoolPtrOutput `pulumi:"active"`
	// String. The product’s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List(String). A list of up to 15 features for this product. These are displayed in pricing tables.
	Features pulumi.StringArrayOutput `pulumi:"features"`
	// List(String). A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images pulumi.StringArrayOutput `pulumi:"images"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// String. The product’s name, meant to be displayable to the customer. Whenever this product is sold via a subscription, name will show up on associated invoice line item descriptions.
	Name pulumi.StringOutput `pulumi:"name"`
	// Map(Float). The dimensions of this product for shipping purposes. When used these fields are required: `height`,`length`,`width` and `weight`; the precision is 2 decimal places.
	PackageDimensions pulumi.Float64MapOutput `pulumi:"packageDimensions"`
	// String. The bespoke unique identifier for the object.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// Bool. Whether this product is shipped (i.e., physical goods).
	Shippable pulumi.BoolPtrOutput `pulumi:"shippable"`
	// String. An arbitrary string to be displayed on your customer’s credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all. This may be up to 22 characters. The statement description may not include `<`,`  > `, `\`, `"`, `’` characters, and will appear on your customer’s statement in capital letters. Non-ASCII characters are automatically stripped. It must contain at least one letter.
	StatementDescriptor pulumi.StringPtrOutput `pulumi:"statementDescriptor"`
	// String. A tax code ID. Supported values are listed in the TaxCode resource and at https://stripe.com/docs/tax/tax-categories.
	TaxCode pulumi.StringPtrOutput `pulumi:"taxCode"`
	// String. A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.
	UnitLabel pulumi.StringPtrOutput `pulumi:"unitLabel"`
	// String. A URL of a publicly-accessible webpage for this product.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewProduct registers a new resource with the given unique name, arguments, and options.
func NewProduct(ctx *pulumi.Context,
	name string, args *ProductArgs, opts ...pulumi.ResourceOption) (*Product, error) {
	if args == nil {
		args = &ProductArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Product
	err := ctx.RegisterResource("stripe:index/product:Product", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProduct gets an existing Product resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductState, opts ...pulumi.ResourceOption) (*Product, error) {
	var resource Product
	err := ctx.ReadResource("stripe:index/product:Product", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Product resources.
type productState struct {
	// Bool. Whether the product is currently available for purchase. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// String. The product’s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `pulumi:"description"`
	// List(String). A list of up to 15 features for this product. These are displayed in pricing tables.
	Features []string `pulumi:"features"`
	// List(String). A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []string `pulumi:"images"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// String. The product’s name, meant to be displayable to the customer. Whenever this product is sold via a subscription, name will show up on associated invoice line item descriptions.
	Name *string `pulumi:"name"`
	// Map(Float). The dimensions of this product for shipping purposes. When used these fields are required: `height`,`length`,`width` and `weight`; the precision is 2 decimal places.
	PackageDimensions map[string]float64 `pulumi:"packageDimensions"`
	// String. The bespoke unique identifier for the object.
	ProductId *string `pulumi:"productId"`
	// Bool. Whether this product is shipped (i.e., physical goods).
	Shippable *bool `pulumi:"shippable"`
	// String. An arbitrary string to be displayed on your customer’s credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all. This may be up to 22 characters. The statement description may not include `<`,`  > `, `\`, `"`, `’` characters, and will appear on your customer’s statement in capital letters. Non-ASCII characters are automatically stripped. It must contain at least one letter.
	StatementDescriptor *string `pulumi:"statementDescriptor"`
	// String. A tax code ID. Supported values are listed in the TaxCode resource and at https://stripe.com/docs/tax/tax-categories.
	TaxCode *string `pulumi:"taxCode"`
	// String. A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.
	UnitLabel *string `pulumi:"unitLabel"`
	// String. A URL of a publicly-accessible webpage for this product.
	Url *string `pulumi:"url"`
}

type ProductState struct {
	// Bool. Whether the product is currently available for purchase. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// String. The product’s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description pulumi.StringPtrInput
	// List(String). A list of up to 15 features for this product. These are displayed in pricing tables.
	Features pulumi.StringArrayInput
	// List(String). A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images pulumi.StringArrayInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// String. The product’s name, meant to be displayable to the customer. Whenever this product is sold via a subscription, name will show up on associated invoice line item descriptions.
	Name pulumi.StringPtrInput
	// Map(Float). The dimensions of this product for shipping purposes. When used these fields are required: `height`,`length`,`width` and `weight`; the precision is 2 decimal places.
	PackageDimensions pulumi.Float64MapInput
	// String. The bespoke unique identifier for the object.
	ProductId pulumi.StringPtrInput
	// Bool. Whether this product is shipped (i.e., physical goods).
	Shippable pulumi.BoolPtrInput
	// String. An arbitrary string to be displayed on your customer’s credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all. This may be up to 22 characters. The statement description may not include `<`,`  > `, `\`, `"`, `’` characters, and will appear on your customer’s statement in capital letters. Non-ASCII characters are automatically stripped. It must contain at least one letter.
	StatementDescriptor pulumi.StringPtrInput
	// String. A tax code ID. Supported values are listed in the TaxCode resource and at https://stripe.com/docs/tax/tax-categories.
	TaxCode pulumi.StringPtrInput
	// String. A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.
	UnitLabel pulumi.StringPtrInput
	// String. A URL of a publicly-accessible webpage for this product.
	Url pulumi.StringPtrInput
}

func (ProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*productState)(nil)).Elem()
}

type productArgs struct {
	// Bool. Whether the product is currently available for purchase. Defaults to `true`.
	Active *bool `pulumi:"active"`
	// String. The product’s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `pulumi:"description"`
	// List(String). A list of up to 15 features for this product. These are displayed in pricing tables.
	Features []string `pulumi:"features"`
	// List(String). A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []string `pulumi:"images"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// String. The product’s name, meant to be displayable to the customer. Whenever this product is sold via a subscription, name will show up on associated invoice line item descriptions.
	Name *string `pulumi:"name"`
	// Map(Float). The dimensions of this product for shipping purposes. When used these fields are required: `height`,`length`,`width` and `weight`; the precision is 2 decimal places.
	PackageDimensions map[string]float64 `pulumi:"packageDimensions"`
	// String. The bespoke unique identifier for the object.
	ProductId *string `pulumi:"productId"`
	// Bool. Whether this product is shipped (i.e., physical goods).
	Shippable *bool `pulumi:"shippable"`
	// String. An arbitrary string to be displayed on your customer’s credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all. This may be up to 22 characters. The statement description may not include `<`,`  > `, `\`, `"`, `’` characters, and will appear on your customer’s statement in capital letters. Non-ASCII characters are automatically stripped. It must contain at least one letter.
	StatementDescriptor *string `pulumi:"statementDescriptor"`
	// String. A tax code ID. Supported values are listed in the TaxCode resource and at https://stripe.com/docs/tax/tax-categories.
	TaxCode *string `pulumi:"taxCode"`
	// String. A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.
	UnitLabel *string `pulumi:"unitLabel"`
	// String. A URL of a publicly-accessible webpage for this product.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Product resource.
type ProductArgs struct {
	// Bool. Whether the product is currently available for purchase. Defaults to `true`.
	Active pulumi.BoolPtrInput
	// String. The product’s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description pulumi.StringPtrInput
	// List(String). A list of up to 15 features for this product. These are displayed in pricing tables.
	Features pulumi.StringArrayInput
	// List(String). A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images pulumi.StringArrayInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// String. The product’s name, meant to be displayable to the customer. Whenever this product is sold via a subscription, name will show up on associated invoice line item descriptions.
	Name pulumi.StringPtrInput
	// Map(Float). The dimensions of this product for shipping purposes. When used these fields are required: `height`,`length`,`width` and `weight`; the precision is 2 decimal places.
	PackageDimensions pulumi.Float64MapInput
	// String. The bespoke unique identifier for the object.
	ProductId pulumi.StringPtrInput
	// Bool. Whether this product is shipped (i.e., physical goods).
	Shippable pulumi.BoolPtrInput
	// String. An arbitrary string to be displayed on your customer’s credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all. This may be up to 22 characters. The statement description may not include `<`,`  > `, `\`, `"`, `’` characters, and will appear on your customer’s statement in capital letters. Non-ASCII characters are automatically stripped. It must contain at least one letter.
	StatementDescriptor pulumi.StringPtrInput
	// String. A tax code ID. Supported values are listed in the TaxCode resource and at https://stripe.com/docs/tax/tax-categories.
	TaxCode pulumi.StringPtrInput
	// String. A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.
	UnitLabel pulumi.StringPtrInput
	// String. A URL of a publicly-accessible webpage for this product.
	Url pulumi.StringPtrInput
}

func (ProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productArgs)(nil)).Elem()
}

type ProductInput interface {
	pulumi.Input

	ToProductOutput() ProductOutput
	ToProductOutputWithContext(ctx context.Context) ProductOutput
}

func (*Product) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil)).Elem()
}

func (i *Product) ToProductOutput() ProductOutput {
	return i.ToProductOutputWithContext(context.Background())
}

func (i *Product) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductOutput)
}

// ProductArrayInput is an input type that accepts ProductArray and ProductArrayOutput values.
// You can construct a concrete instance of `ProductArrayInput` via:
//
//	ProductArray{ ProductArgs{...} }
type ProductArrayInput interface {
	pulumi.Input

	ToProductArrayOutput() ProductArrayOutput
	ToProductArrayOutputWithContext(context.Context) ProductArrayOutput
}

type ProductArray []ProductInput

func (ProductArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Product)(nil)).Elem()
}

func (i ProductArray) ToProductArrayOutput() ProductArrayOutput {
	return i.ToProductArrayOutputWithContext(context.Background())
}

func (i ProductArray) ToProductArrayOutputWithContext(ctx context.Context) ProductArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductArrayOutput)
}

// ProductMapInput is an input type that accepts ProductMap and ProductMapOutput values.
// You can construct a concrete instance of `ProductMapInput` via:
//
//	ProductMap{ "key": ProductArgs{...} }
type ProductMapInput interface {
	pulumi.Input

	ToProductMapOutput() ProductMapOutput
	ToProductMapOutputWithContext(context.Context) ProductMapOutput
}

type ProductMap map[string]ProductInput

func (ProductMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Product)(nil)).Elem()
}

func (i ProductMap) ToProductMapOutput() ProductMapOutput {
	return i.ToProductMapOutputWithContext(context.Background())
}

func (i ProductMap) ToProductMapOutputWithContext(ctx context.Context) ProductMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductMapOutput)
}

type ProductOutput struct{ *pulumi.OutputState }

func (ProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil)).Elem()
}

func (o ProductOutput) ToProductOutput() ProductOutput {
	return o
}

func (o ProductOutput) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return o
}

// Bool. Whether the product is currently available for purchase. Defaults to `true`.
func (o ProductOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.BoolPtrOutput { return v.Active }).(pulumi.BoolPtrOutput)
}

// String. The product’s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
func (o ProductOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List(String). A list of up to 15 features for this product. These are displayed in pricing tables.
func (o ProductOutput) Features() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Product) pulumi.StringArrayOutput { return v.Features }).(pulumi.StringArrayOutput)
}

// List(String). A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
func (o ProductOutput) Images() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Product) pulumi.StringArrayOutput { return v.Images }).(pulumi.StringArrayOutput)
}

// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
func (o ProductOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Product) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// String. The product’s name, meant to be displayable to the customer. Whenever this product is sold via a subscription, name will show up on associated invoice line item descriptions.
func (o ProductOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Map(Float). The dimensions of this product for shipping purposes. When used these fields are required: `height`,`length`,`width` and `weight`; the precision is 2 decimal places.
func (o ProductOutput) PackageDimensions() pulumi.Float64MapOutput {
	return o.ApplyT(func(v *Product) pulumi.Float64MapOutput { return v.PackageDimensions }).(pulumi.Float64MapOutput)
}

// String. The bespoke unique identifier for the object.
func (o ProductOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

// Bool. Whether this product is shipped (i.e., physical goods).
func (o ProductOutput) Shippable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.BoolPtrOutput { return v.Shippable }).(pulumi.BoolPtrOutput)
}

// String. An arbitrary string to be displayed on your customer’s credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all. This may be up to 22 characters. The statement description may not include `<`,`  > `, `\`, `"`, `’` characters, and will appear on your customer’s statement in capital letters. Non-ASCII characters are automatically stripped. It must contain at least one letter.
func (o ProductOutput) StatementDescriptor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.StringPtrOutput { return v.StatementDescriptor }).(pulumi.StringPtrOutput)
}

// String. A tax code ID. Supported values are listed in the TaxCode resource and at https://stripe.com/docs/tax/tax-categories.
func (o ProductOutput) TaxCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.StringPtrOutput { return v.TaxCode }).(pulumi.StringPtrOutput)
}

// String. A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.
func (o ProductOutput) UnitLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.StringPtrOutput { return v.UnitLabel }).(pulumi.StringPtrOutput)
}

// String. A URL of a publicly-accessible webpage for this product.
func (o ProductOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

type ProductArrayOutput struct{ *pulumi.OutputState }

func (ProductArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Product)(nil)).Elem()
}

func (o ProductArrayOutput) ToProductArrayOutput() ProductArrayOutput {
	return o
}

func (o ProductArrayOutput) ToProductArrayOutputWithContext(ctx context.Context) ProductArrayOutput {
	return o
}

func (o ProductArrayOutput) Index(i pulumi.IntInput) ProductOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Product {
		return vs[0].([]*Product)[vs[1].(int)]
	}).(ProductOutput)
}

type ProductMapOutput struct{ *pulumi.OutputState }

func (ProductMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Product)(nil)).Elem()
}

func (o ProductMapOutput) ToProductMapOutput() ProductMapOutput {
	return o
}

func (o ProductMapOutput) ToProductMapOutputWithContext(ctx context.Context) ProductMapOutput {
	return o
}

func (o ProductMapOutput) MapIndex(k pulumi.StringInput) ProductOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Product {
		return vs[0].(map[string]*Product)[vs[1].(string)]
	}).(ProductOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProductInput)(nil)).Elem(), &Product{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductArrayInput)(nil)).Elem(), ProductArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProductMapInput)(nil)).Elem(), ProductMap{})
	pulumi.RegisterOutputType(ProductOutput{})
	pulumi.RegisterOutputType(ProductArrayOutput{})
	pulumi.RegisterOutputType(ProductMapOutput{})
}
