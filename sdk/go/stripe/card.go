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
//			// card for the customer
//			_, err := stripe.NewCard(ctx, "cardCard", &stripe.CardArgs{
//				Customer: pulumi.Any(stripe_customer.Customer.Id),
//				Number:   pulumi.String("4242424242424242"),
//				Cvc:      pulumi.Int(123),
//				ExpMonth: pulumi.Int(8),
//				ExpYear:  pulumi.Int(2030),
//			})
//			if err != nil {
//				return err
//			}
//			// card for the customer with address
//			_, err = stripe.NewCard(ctx, "cardIndex/cardCard", &stripe.CardArgs{
//				Customer: pulumi.Any(stripe_customer.Customer.Id),
//				Number:   pulumi.String("4242424242424242"),
//				Cvc:      pulumi.Int(123),
//				ExpMonth: pulumi.Int(8),
//				ExpYear:  pulumi.Int(2030),
//				Address: pulumi.StringMap{
//					"line1":   pulumi.String("1 The Best Street"),
//					"line2":   pulumi.String("Apartment 401"),
//					"city":    pulumi.String("Sydney"),
//					"state":   pulumi.String("NSW"),
//					"zip":     pulumi.String("2000"),
//					"country": pulumi.String("Australia"),
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
type Card struct {
	pulumi.CustomResourceState

	// Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
	// , `zip` and `country`.
	Address pulumi.StringMapOutput `pulumi:"address"`
	// String. If address `line1` was provided, results of the check: `pass`, `fail`, `unavailable`,
	// or `unchecked`.
	AddressLine1Check pulumi.StringOutput `pulumi:"addressLine1Check"`
	// String. If address `zip` was provided, results of the check: `pass`, `fail`, `unavailable`,
	// or `unchecked`.
	AddressZipCheck pulumi.StringOutput `pulumi:"addressZipCheck"`
	// List(String). A set of available payout methods for this card. Only values from this set
	// should be passed as the method when creating a payout.
	AvailablePayoutMethods pulumi.StringArrayOutput `pulumi:"availablePayoutMethods"`
	// String. Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`
	// , `Visa`, or `Unknown`.
	Brand pulumi.StringOutput `pulumi:"brand"`
	// String. Two-letter ISO code representing the country of the card. You could use this attribute to get a
	// sense of the international breakdown of cards you’ve collected.
	Country pulumi.StringOutput `pulumi:"country"`
	// String. The customer that this card belongs to.
	Customer pulumi.StringOutput `pulumi:"customer"`
	// Int. Card security code. Highly recommended to always include this value, but it's required only
	// for accounts based in European countries.
	Cvc pulumi.IntPtrOutput `pulumi:"cvc"`
	// String. If a `cvc` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A
	// result of `unchecked` indicates that CVC was provided but hasn’t been checked yet
	CvcCheck pulumi.StringOutput `pulumi:"cvcCheck"`
	// Int. Number representing the card's expiration month.
	ExpMonth pulumi.IntOutput `pulumi:"expMonth"`
	// Int. Four-digit number representing the card's expiration year.
	ExpYear pulumi.IntOutput `pulumi:"expYear"`
	// String. Uniquely identifies this particular card number. You can use this attribute to check whether
	// two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize
	// card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card
	// number.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// String. Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding pulumi.StringOutput `pulumi:"funding"`
	// String. The last four digits of the card.
	Last4 pulumi.StringOutput `pulumi:"last4"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// String. Cardholder name.
	Name pulumi.StringOutput `pulumi:"name"`
	// String. The card number, as a string without any separators.
	Number pulumi.StringOutput `pulumi:"number"`
	// String. If the card number is tokenized, this is the method that was used. Can
	// be `androidPay` (includes Google Pay), `applePay`, `masterpass`, `visaCheckout`, or `null`.
	TokenizationMethod pulumi.StringOutput `pulumi:"tokenizationMethod"`
}

// NewCard registers a new resource with the given unique name, arguments, and options.
func NewCard(ctx *pulumi.Context,
	name string, args *CardArgs, opts ...pulumi.ResourceOption) (*Card, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Customer == nil {
		return nil, errors.New("invalid value for required argument 'Customer'")
	}
	if args.ExpMonth == nil {
		return nil, errors.New("invalid value for required argument 'ExpMonth'")
	}
	if args.ExpYear == nil {
		return nil, errors.New("invalid value for required argument 'ExpYear'")
	}
	if args.Number == nil {
		return nil, errors.New("invalid value for required argument 'Number'")
	}
	if args.Cvc != nil {
		args.Cvc = pulumi.ToSecret(args.Cvc).(pulumi.IntPtrInput)
	}
	if args.Number != nil {
		args.Number = pulumi.ToSecret(args.Number).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"cvc",
		"number",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Card
	err := ctx.RegisterResource("stripe:index/card:Card", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCard gets an existing Card resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CardState, opts ...pulumi.ResourceOption) (*Card, error) {
	var resource Card
	err := ctx.ReadResource("stripe:index/card:Card", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Card resources.
type cardState struct {
	// Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
	// , `zip` and `country`.
	Address map[string]string `pulumi:"address"`
	// String. If address `line1` was provided, results of the check: `pass`, `fail`, `unavailable`,
	// or `unchecked`.
	AddressLine1Check *string `pulumi:"addressLine1Check"`
	// String. If address `zip` was provided, results of the check: `pass`, `fail`, `unavailable`,
	// or `unchecked`.
	AddressZipCheck *string `pulumi:"addressZipCheck"`
	// List(String). A set of available payout methods for this card. Only values from this set
	// should be passed as the method when creating a payout.
	AvailablePayoutMethods []string `pulumi:"availablePayoutMethods"`
	// String. Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`
	// , `Visa`, or `Unknown`.
	Brand *string `pulumi:"brand"`
	// String. Two-letter ISO code representing the country of the card. You could use this attribute to get a
	// sense of the international breakdown of cards you’ve collected.
	Country *string `pulumi:"country"`
	// String. The customer that this card belongs to.
	Customer *string `pulumi:"customer"`
	// Int. Card security code. Highly recommended to always include this value, but it's required only
	// for accounts based in European countries.
	Cvc *int `pulumi:"cvc"`
	// String. If a `cvc` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A
	// result of `unchecked` indicates that CVC was provided but hasn’t been checked yet
	CvcCheck *string `pulumi:"cvcCheck"`
	// Int. Number representing the card's expiration month.
	ExpMonth *int `pulumi:"expMonth"`
	// Int. Four-digit number representing the card's expiration year.
	ExpYear *int `pulumi:"expYear"`
	// String. Uniquely identifies this particular card number. You can use this attribute to check whether
	// two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize
	// card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card
	// number.
	Fingerprint *string `pulumi:"fingerprint"`
	// String. Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding *string `pulumi:"funding"`
	// String. The last four digits of the card.
	Last4 *string `pulumi:"last4"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// String. Cardholder name.
	Name *string `pulumi:"name"`
	// String. The card number, as a string without any separators.
	Number *string `pulumi:"number"`
	// String. If the card number is tokenized, this is the method that was used. Can
	// be `androidPay` (includes Google Pay), `applePay`, `masterpass`, `visaCheckout`, or `null`.
	TokenizationMethod *string `pulumi:"tokenizationMethod"`
}

type CardState struct {
	// Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
	// , `zip` and `country`.
	Address pulumi.StringMapInput
	// String. If address `line1` was provided, results of the check: `pass`, `fail`, `unavailable`,
	// or `unchecked`.
	AddressLine1Check pulumi.StringPtrInput
	// String. If address `zip` was provided, results of the check: `pass`, `fail`, `unavailable`,
	// or `unchecked`.
	AddressZipCheck pulumi.StringPtrInput
	// List(String). A set of available payout methods for this card. Only values from this set
	// should be passed as the method when creating a payout.
	AvailablePayoutMethods pulumi.StringArrayInput
	// String. Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`
	// , `Visa`, or `Unknown`.
	Brand pulumi.StringPtrInput
	// String. Two-letter ISO code representing the country of the card. You could use this attribute to get a
	// sense of the international breakdown of cards you’ve collected.
	Country pulumi.StringPtrInput
	// String. The customer that this card belongs to.
	Customer pulumi.StringPtrInput
	// Int. Card security code. Highly recommended to always include this value, but it's required only
	// for accounts based in European countries.
	Cvc pulumi.IntPtrInput
	// String. If a `cvc` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A
	// result of `unchecked` indicates that CVC was provided but hasn’t been checked yet
	CvcCheck pulumi.StringPtrInput
	// Int. Number representing the card's expiration month.
	ExpMonth pulumi.IntPtrInput
	// Int. Four-digit number representing the card's expiration year.
	ExpYear pulumi.IntPtrInput
	// String. Uniquely identifies this particular card number. You can use this attribute to check whether
	// two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize
	// card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card
	// number.
	Fingerprint pulumi.StringPtrInput
	// String. Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding pulumi.StringPtrInput
	// String. The last four digits of the card.
	Last4 pulumi.StringPtrInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// String. Cardholder name.
	Name pulumi.StringPtrInput
	// String. The card number, as a string without any separators.
	Number pulumi.StringPtrInput
	// String. If the card number is tokenized, this is the method that was used. Can
	// be `androidPay` (includes Google Pay), `applePay`, `masterpass`, `visaCheckout`, or `null`.
	TokenizationMethod pulumi.StringPtrInput
}

func (CardState) ElementType() reflect.Type {
	return reflect.TypeOf((*cardState)(nil)).Elem()
}

type cardArgs struct {
	// Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
	// , `zip` and `country`.
	Address map[string]string `pulumi:"address"`
	// String. The customer that this card belongs to.
	Customer string `pulumi:"customer"`
	// Int. Card security code. Highly recommended to always include this value, but it's required only
	// for accounts based in European countries.
	Cvc *int `pulumi:"cvc"`
	// Int. Number representing the card's expiration month.
	ExpMonth int `pulumi:"expMonth"`
	// Int. Four-digit number representing the card's expiration year.
	ExpYear int `pulumi:"expYear"`
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata map[string]string `pulumi:"metadata"`
	// String. Cardholder name.
	Name *string `pulumi:"name"`
	// String. The card number, as a string without any separators.
	Number string `pulumi:"number"`
}

// The set of arguments for constructing a Card resource.
type CardArgs struct {
	// Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
	// , `zip` and `country`.
	Address pulumi.StringMapInput
	// String. The customer that this card belongs to.
	Customer pulumi.StringInput
	// Int. Card security code. Highly recommended to always include this value, but it's required only
	// for accounts based in European countries.
	Cvc pulumi.IntPtrInput
	// Int. Number representing the card's expiration month.
	ExpMonth pulumi.IntInput
	// Int. Four-digit number representing the card's expiration year.
	ExpYear pulumi.IntInput
	// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
	// storing additional information about the object in a structured format.
	Metadata pulumi.StringMapInput
	// String. Cardholder name.
	Name pulumi.StringPtrInput
	// String. The card number, as a string without any separators.
	Number pulumi.StringInput
}

func (CardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cardArgs)(nil)).Elem()
}

type CardInput interface {
	pulumi.Input

	ToCardOutput() CardOutput
	ToCardOutputWithContext(ctx context.Context) CardOutput
}

func (*Card) ElementType() reflect.Type {
	return reflect.TypeOf((**Card)(nil)).Elem()
}

func (i *Card) ToCardOutput() CardOutput {
	return i.ToCardOutputWithContext(context.Background())
}

func (i *Card) ToCardOutputWithContext(ctx context.Context) CardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CardOutput)
}

// CardArrayInput is an input type that accepts CardArray and CardArrayOutput values.
// You can construct a concrete instance of `CardArrayInput` via:
//
//	CardArray{ CardArgs{...} }
type CardArrayInput interface {
	pulumi.Input

	ToCardArrayOutput() CardArrayOutput
	ToCardArrayOutputWithContext(context.Context) CardArrayOutput
}

type CardArray []CardInput

func (CardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Card)(nil)).Elem()
}

func (i CardArray) ToCardArrayOutput() CardArrayOutput {
	return i.ToCardArrayOutputWithContext(context.Background())
}

func (i CardArray) ToCardArrayOutputWithContext(ctx context.Context) CardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CardArrayOutput)
}

// CardMapInput is an input type that accepts CardMap and CardMapOutput values.
// You can construct a concrete instance of `CardMapInput` via:
//
//	CardMap{ "key": CardArgs{...} }
type CardMapInput interface {
	pulumi.Input

	ToCardMapOutput() CardMapOutput
	ToCardMapOutputWithContext(context.Context) CardMapOutput
}

type CardMap map[string]CardInput

func (CardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Card)(nil)).Elem()
}

func (i CardMap) ToCardMapOutput() CardMapOutput {
	return i.ToCardMapOutputWithContext(context.Background())
}

func (i CardMap) ToCardMapOutputWithContext(ctx context.Context) CardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CardMapOutput)
}

type CardOutput struct{ *pulumi.OutputState }

func (CardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Card)(nil)).Elem()
}

func (o CardOutput) ToCardOutput() CardOutput {
	return o
}

func (o CardOutput) ToCardOutputWithContext(ctx context.Context) CardOutput {
	return o
}

// Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
// , `zip` and `country`.
func (o CardOutput) Address() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Card) pulumi.StringMapOutput { return v.Address }).(pulumi.StringMapOutput)
}

// String. If address `line1` was provided, results of the check: `pass`, `fail`, `unavailable`,
// or `unchecked`.
func (o CardOutput) AddressLine1Check() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.AddressLine1Check }).(pulumi.StringOutput)
}

// String. If address `zip` was provided, results of the check: `pass`, `fail`, `unavailable`,
// or `unchecked`.
func (o CardOutput) AddressZipCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.AddressZipCheck }).(pulumi.StringOutput)
}

// List(String). A set of available payout methods for this card. Only values from this set
// should be passed as the method when creating a payout.
func (o CardOutput) AvailablePayoutMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Card) pulumi.StringArrayOutput { return v.AvailablePayoutMethods }).(pulumi.StringArrayOutput)
}

// String. Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`
// , `Visa`, or `Unknown`.
func (o CardOutput) Brand() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.Brand }).(pulumi.StringOutput)
}

// String. Two-letter ISO code representing the country of the card. You could use this attribute to get a
// sense of the international breakdown of cards you’ve collected.
func (o CardOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// String. The customer that this card belongs to.
func (o CardOutput) Customer() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.Customer }).(pulumi.StringOutput)
}

// Int. Card security code. Highly recommended to always include this value, but it's required only
// for accounts based in European countries.
func (o CardOutput) Cvc() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Card) pulumi.IntPtrOutput { return v.Cvc }).(pulumi.IntPtrOutput)
}

// String. If a `cvc` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A
// result of `unchecked` indicates that CVC was provided but hasn’t been checked yet
func (o CardOutput) CvcCheck() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.CvcCheck }).(pulumi.StringOutput)
}

// Int. Number representing the card's expiration month.
func (o CardOutput) ExpMonth() pulumi.IntOutput {
	return o.ApplyT(func(v *Card) pulumi.IntOutput { return v.ExpMonth }).(pulumi.IntOutput)
}

// Int. Four-digit number representing the card's expiration year.
func (o CardOutput) ExpYear() pulumi.IntOutput {
	return o.ApplyT(func(v *Card) pulumi.IntOutput { return v.ExpYear }).(pulumi.IntOutput)
}

// String. Uniquely identifies this particular card number. You can use this attribute to check whether
// two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize
// card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card
// number.
func (o CardOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// String. Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
func (o CardOutput) Funding() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.Funding }).(pulumi.StringOutput)
}

// String. The last four digits of the card.
func (o CardOutput) Last4() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.Last4 }).(pulumi.StringOutput)
}

// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
// storing additional information about the object in a structured format.
func (o CardOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Card) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// String. Cardholder name.
func (o CardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// String. The card number, as a string without any separators.
func (o CardOutput) Number() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.Number }).(pulumi.StringOutput)
}

// String. If the card number is tokenized, this is the method that was used. Can
// be `androidPay` (includes Google Pay), `applePay`, `masterpass`, `visaCheckout`, or `null`.
func (o CardOutput) TokenizationMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Card) pulumi.StringOutput { return v.TokenizationMethod }).(pulumi.StringOutput)
}

type CardArrayOutput struct{ *pulumi.OutputState }

func (CardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Card)(nil)).Elem()
}

func (o CardArrayOutput) ToCardArrayOutput() CardArrayOutput {
	return o
}

func (o CardArrayOutput) ToCardArrayOutputWithContext(ctx context.Context) CardArrayOutput {
	return o
}

func (o CardArrayOutput) Index(i pulumi.IntInput) CardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Card {
		return vs[0].([]*Card)[vs[1].(int)]
	}).(CardOutput)
}

type CardMapOutput struct{ *pulumi.OutputState }

func (CardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Card)(nil)).Elem()
}

func (o CardMapOutput) ToCardMapOutput() CardMapOutput {
	return o
}

func (o CardMapOutput) ToCardMapOutputWithContext(ctx context.Context) CardMapOutput {
	return o
}

func (o CardMapOutput) MapIndex(k pulumi.StringInput) CardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Card {
		return vs[0].(map[string]*Card)[vs[1].(string)]
	}).(CardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CardInput)(nil)).Elem(), &Card{})
	pulumi.RegisterInputType(reflect.TypeOf((*CardArrayInput)(nil)).Elem(), CardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CardMapInput)(nil)).Elem(), CardMap{})
	pulumi.RegisterOutputType(CardOutput{})
	pulumi.RegisterOutputType(CardArrayOutput{})
	pulumi.RegisterOutputType(CardMapOutput{})
}
