# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CardArgs', 'Card']

@pulumi.input_type
class CardArgs:
    def __init__(__self__, *,
                 customer: pulumi.Input[str],
                 exp_month: pulumi.Input[int],
                 exp_year: pulumi.Input[int],
                 number: pulumi.Input[str],
                 address: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 cvc: Optional[pulumi.Input[int]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Card resource.
        :param pulumi.Input[str] customer: String. The customer that this card belongs to.
        :param pulumi.Input[int] exp_month: Int. Number representing the card's expiration month.
        :param pulumi.Input[int] exp_year: Int. Four-digit number representing the card's expiration year.
        :param pulumi.Input[str] number: String. The card number, as a string without any separators.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] address: Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
               , `zip` and `country`.
        :param pulumi.Input[int] cvc: Int. Card security code. Highly recommended to always include this value, but it's required only
               for accounts based in European countries.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
               storing additional information about the object in a structured format.
        :param pulumi.Input[str] name: String. Cardholder name.
        """
        pulumi.set(__self__, "customer", customer)
        pulumi.set(__self__, "exp_month", exp_month)
        pulumi.set(__self__, "exp_year", exp_year)
        pulumi.set(__self__, "number", number)
        if address is not None:
            pulumi.set(__self__, "address", address)
        if cvc is not None:
            pulumi.set(__self__, "cvc", cvc)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def customer(self) -> pulumi.Input[str]:
        """
        String. The customer that this card belongs to.
        """
        return pulumi.get(self, "customer")

    @customer.setter
    def customer(self, value: pulumi.Input[str]):
        pulumi.set(self, "customer", value)

    @property
    @pulumi.getter(name="expMonth")
    def exp_month(self) -> pulumi.Input[int]:
        """
        Int. Number representing the card's expiration month.
        """
        return pulumi.get(self, "exp_month")

    @exp_month.setter
    def exp_month(self, value: pulumi.Input[int]):
        pulumi.set(self, "exp_month", value)

    @property
    @pulumi.getter(name="expYear")
    def exp_year(self) -> pulumi.Input[int]:
        """
        Int. Four-digit number representing the card's expiration year.
        """
        return pulumi.get(self, "exp_year")

    @exp_year.setter
    def exp_year(self, value: pulumi.Input[int]):
        pulumi.set(self, "exp_year", value)

    @property
    @pulumi.getter
    def number(self) -> pulumi.Input[str]:
        """
        String. The card number, as a string without any separators.
        """
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: pulumi.Input[str]):
        pulumi.set(self, "number", value)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
        , `zip` and `country`.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def cvc(self) -> Optional[pulumi.Input[int]]:
        """
        Int. Card security code. Highly recommended to always include this value, but it's required only
        for accounts based in European countries.
        """
        return pulumi.get(self, "cvc")

    @cvc.setter
    def cvc(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cvc", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
        storing additional information about the object in a structured format.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        String. Cardholder name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _CardState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 address_line1_check: Optional[pulumi.Input[str]] = None,
                 address_zip_check: Optional[pulumi.Input[str]] = None,
                 available_payout_methods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 brand: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 cvc: Optional[pulumi.Input[int]] = None,
                 cvc_check: Optional[pulumi.Input[str]] = None,
                 exp_month: Optional[pulumi.Input[int]] = None,
                 exp_year: Optional[pulumi.Input[int]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 funding: Optional[pulumi.Input[str]] = None,
                 last4: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 tokenization_method: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Card resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] address: Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
               , `zip` and `country`.
        :param pulumi.Input[str] address_line1_check: String. If address `line1` was provided, results of the check: `pass`, `fail`, `unavailable`,
               or `unchecked`.
        :param pulumi.Input[str] address_zip_check: String. If address `zip` was provided, results of the check: `pass`, `fail`, `unavailable`,
               or `unchecked`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] available_payout_methods: List(String). A set of available payout methods for this card. Only values from this set
               should be passed as the method when creating a payout.
        :param pulumi.Input[str] brand: String. Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`
               , `Visa`, or `Unknown`.
        :param pulumi.Input[str] country: String. Two-letter ISO code representing the country of the card. You could use this attribute to get a
               sense of the international breakdown of cards you’ve collected.
        :param pulumi.Input[str] customer: String. The customer that this card belongs to.
        :param pulumi.Input[int] cvc: Int. Card security code. Highly recommended to always include this value, but it's required only
               for accounts based in European countries.
        :param pulumi.Input[str] cvc_check: String. If a `cvc` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A
               result of `unchecked` indicates that CVC was provided but hasn’t been checked yet
        :param pulumi.Input[int] exp_month: Int. Number representing the card's expiration month.
        :param pulumi.Input[int] exp_year: Int. Four-digit number representing the card's expiration year.
        :param pulumi.Input[str] fingerprint: String. Uniquely identifies this particular card number. You can use this attribute to check whether
               two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize
               card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card
               number.
        :param pulumi.Input[str] funding: String. Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
        :param pulumi.Input[str] last4: String. The last four digits of the card.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
               storing additional information about the object in a structured format.
        :param pulumi.Input[str] name: String. Cardholder name.
        :param pulumi.Input[str] number: String. The card number, as a string without any separators.
        :param pulumi.Input[str] tokenization_method: String. If the card number is tokenized, this is the method that was used. Can
               be `android_pay` (includes Google Pay), `apple_pay`, `masterpass`, `visa_checkout`, or `null`.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if address_line1_check is not None:
            pulumi.set(__self__, "address_line1_check", address_line1_check)
        if address_zip_check is not None:
            pulumi.set(__self__, "address_zip_check", address_zip_check)
        if available_payout_methods is not None:
            pulumi.set(__self__, "available_payout_methods", available_payout_methods)
        if brand is not None:
            pulumi.set(__self__, "brand", brand)
        if country is not None:
            pulumi.set(__self__, "country", country)
        if customer is not None:
            pulumi.set(__self__, "customer", customer)
        if cvc is not None:
            pulumi.set(__self__, "cvc", cvc)
        if cvc_check is not None:
            pulumi.set(__self__, "cvc_check", cvc_check)
        if exp_month is not None:
            pulumi.set(__self__, "exp_month", exp_month)
        if exp_year is not None:
            pulumi.set(__self__, "exp_year", exp_year)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if funding is not None:
            pulumi.set(__self__, "funding", funding)
        if last4 is not None:
            pulumi.set(__self__, "last4", last4)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if number is not None:
            pulumi.set(__self__, "number", number)
        if tokenization_method is not None:
            pulumi.set(__self__, "tokenization_method", tokenization_method)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
        , `zip` and `country`.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="addressLine1Check")
    def address_line1_check(self) -> Optional[pulumi.Input[str]]:
        """
        String. If address `line1` was provided, results of the check: `pass`, `fail`, `unavailable`,
        or `unchecked`.
        """
        return pulumi.get(self, "address_line1_check")

    @address_line1_check.setter
    def address_line1_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_line1_check", value)

    @property
    @pulumi.getter(name="addressZipCheck")
    def address_zip_check(self) -> Optional[pulumi.Input[str]]:
        """
        String. If address `zip` was provided, results of the check: `pass`, `fail`, `unavailable`,
        or `unchecked`.
        """
        return pulumi.get(self, "address_zip_check")

    @address_zip_check.setter
    def address_zip_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_zip_check", value)

    @property
    @pulumi.getter(name="availablePayoutMethods")
    def available_payout_methods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List(String). A set of available payout methods for this card. Only values from this set
        should be passed as the method when creating a payout.
        """
        return pulumi.get(self, "available_payout_methods")

    @available_payout_methods.setter
    def available_payout_methods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "available_payout_methods", value)

    @property
    @pulumi.getter
    def brand(self) -> Optional[pulumi.Input[str]]:
        """
        String. Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`
        , `Visa`, or `Unknown`.
        """
        return pulumi.get(self, "brand")

    @brand.setter
    def brand(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "brand", value)

    @property
    @pulumi.getter
    def country(self) -> Optional[pulumi.Input[str]]:
        """
        String. Two-letter ISO code representing the country of the card. You could use this attribute to get a
        sense of the international breakdown of cards you’ve collected.
        """
        return pulumi.get(self, "country")

    @country.setter
    def country(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "country", value)

    @property
    @pulumi.getter
    def customer(self) -> Optional[pulumi.Input[str]]:
        """
        String. The customer that this card belongs to.
        """
        return pulumi.get(self, "customer")

    @customer.setter
    def customer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer", value)

    @property
    @pulumi.getter
    def cvc(self) -> Optional[pulumi.Input[int]]:
        """
        Int. Card security code. Highly recommended to always include this value, but it's required only
        for accounts based in European countries.
        """
        return pulumi.get(self, "cvc")

    @cvc.setter
    def cvc(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cvc", value)

    @property
    @pulumi.getter(name="cvcCheck")
    def cvc_check(self) -> Optional[pulumi.Input[str]]:
        """
        String. If a `cvc` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A
        result of `unchecked` indicates that CVC was provided but hasn’t been checked yet
        """
        return pulumi.get(self, "cvc_check")

    @cvc_check.setter
    def cvc_check(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cvc_check", value)

    @property
    @pulumi.getter(name="expMonth")
    def exp_month(self) -> Optional[pulumi.Input[int]]:
        """
        Int. Number representing the card's expiration month.
        """
        return pulumi.get(self, "exp_month")

    @exp_month.setter
    def exp_month(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "exp_month", value)

    @property
    @pulumi.getter(name="expYear")
    def exp_year(self) -> Optional[pulumi.Input[int]]:
        """
        Int. Four-digit number representing the card's expiration year.
        """
        return pulumi.get(self, "exp_year")

    @exp_year.setter
    def exp_year(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "exp_year", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        String. Uniquely identifies this particular card number. You can use this attribute to check whether
        two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize
        card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card
        number.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def funding(self) -> Optional[pulumi.Input[str]]:
        """
        String. Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
        """
        return pulumi.get(self, "funding")

    @funding.setter
    def funding(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "funding", value)

    @property
    @pulumi.getter
    def last4(self) -> Optional[pulumi.Input[str]]:
        """
        String. The last four digits of the card.
        """
        return pulumi.get(self, "last4")

    @last4.setter
    def last4(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last4", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
        storing additional information about the object in a structured format.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        String. Cardholder name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def number(self) -> Optional[pulumi.Input[str]]:
        """
        String. The card number, as a string without any separators.
        """
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "number", value)

    @property
    @pulumi.getter(name="tokenizationMethod")
    def tokenization_method(self) -> Optional[pulumi.Input[str]]:
        """
        String. If the card number is tokenized, this is the method that was used. Can
        be `android_pay` (includes Google Pay), `apple_pay`, `masterpass`, `visa_checkout`, or `null`.
        """
        return pulumi.get(self, "tokenization_method")

    @tokenization_method.setter
    def tokenization_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tokenization_method", value)


class Card(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 cvc: Optional[pulumi.Input[int]] = None,
                 exp_month: Optional[pulumi.Input[int]] = None,
                 exp_year: Optional[pulumi.Input[int]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_stripe as stripe

        # card for the customer
        card_card = stripe.Card("cardCard",
            customer=stripe_customer["customer"]["id"],
            number="4242424242424242",
            cvc=123,
            exp_month=8,
            exp_year=2030)
        # card for the customer with address
        card_index_card_card = stripe.Card("cardIndex/cardCard",
            customer=stripe_customer["customer"]["id"],
            number="4242424242424242",
            cvc=123,
            exp_month=8,
            exp_year=2030,
            address={
                "line1": "1 The Best Street",
                "line2": "Apartment 401",
                "city": "Sydney",
                "state": "NSW",
                "zip": "2000",
                "country": "Australia",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] address: Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
               , `zip` and `country`.
        :param pulumi.Input[str] customer: String. The customer that this card belongs to.
        :param pulumi.Input[int] cvc: Int. Card security code. Highly recommended to always include this value, but it's required only
               for accounts based in European countries.
        :param pulumi.Input[int] exp_month: Int. Number representing the card's expiration month.
        :param pulumi.Input[int] exp_year: Int. Four-digit number representing the card's expiration year.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
               storing additional information about the object in a structured format.
        :param pulumi.Input[str] name: String. Cardholder name.
        :param pulumi.Input[str] number: String. The card number, as a string without any separators.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_stripe as stripe

        # card for the customer
        card_card = stripe.Card("cardCard",
            customer=stripe_customer["customer"]["id"],
            number="4242424242424242",
            cvc=123,
            exp_month=8,
            exp_year=2030)
        # card for the customer with address
        card_index_card_card = stripe.Card("cardIndex/cardCard",
            customer=stripe_customer["customer"]["id"],
            number="4242424242424242",
            cvc=123,
            exp_month=8,
            exp_year=2030,
            address={
                "line1": "1 The Best Street",
                "line2": "Apartment 401",
                "city": "Sydney",
                "state": "NSW",
                "zip": "2000",
                "country": "Australia",
            })
        ```

        :param str resource_name: The name of the resource.
        :param CardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 cvc: Optional[pulumi.Input[int]] = None,
                 exp_month: Optional[pulumi.Input[int]] = None,
                 exp_year: Optional[pulumi.Input[int]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CardArgs.__new__(CardArgs)

            __props__.__dict__["address"] = address
            if customer is None and not opts.urn:
                raise TypeError("Missing required property 'customer'")
            __props__.__dict__["customer"] = customer
            __props__.__dict__["cvc"] = None if cvc is None else pulumi.Output.secret(cvc)
            if exp_month is None and not opts.urn:
                raise TypeError("Missing required property 'exp_month'")
            __props__.__dict__["exp_month"] = exp_month
            if exp_year is None and not opts.urn:
                raise TypeError("Missing required property 'exp_year'")
            __props__.__dict__["exp_year"] = exp_year
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["name"] = name
            if number is None and not opts.urn:
                raise TypeError("Missing required property 'number'")
            __props__.__dict__["number"] = None if number is None else pulumi.Output.secret(number)
            __props__.__dict__["address_line1_check"] = None
            __props__.__dict__["address_zip_check"] = None
            __props__.__dict__["available_payout_methods"] = None
            __props__.__dict__["brand"] = None
            __props__.__dict__["country"] = None
            __props__.__dict__["cvc_check"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["funding"] = None
            __props__.__dict__["last4"] = None
            __props__.__dict__["tokenization_method"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["cvc", "number"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Card, __self__).__init__(
            'stripe:index/card:Card',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            address_line1_check: Optional[pulumi.Input[str]] = None,
            address_zip_check: Optional[pulumi.Input[str]] = None,
            available_payout_methods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            brand: Optional[pulumi.Input[str]] = None,
            country: Optional[pulumi.Input[str]] = None,
            customer: Optional[pulumi.Input[str]] = None,
            cvc: Optional[pulumi.Input[int]] = None,
            cvc_check: Optional[pulumi.Input[str]] = None,
            exp_month: Optional[pulumi.Input[int]] = None,
            exp_year: Optional[pulumi.Input[int]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            funding: Optional[pulumi.Input[str]] = None,
            last4: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            number: Optional[pulumi.Input[str]] = None,
            tokenization_method: Optional[pulumi.Input[str]] = None) -> 'Card':
        """
        Get an existing Card resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] address: Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
               , `zip` and `country`.
        :param pulumi.Input[str] address_line1_check: String. If address `line1` was provided, results of the check: `pass`, `fail`, `unavailable`,
               or `unchecked`.
        :param pulumi.Input[str] address_zip_check: String. If address `zip` was provided, results of the check: `pass`, `fail`, `unavailable`,
               or `unchecked`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] available_payout_methods: List(String). A set of available payout methods for this card. Only values from this set
               should be passed as the method when creating a payout.
        :param pulumi.Input[str] brand: String. Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`
               , `Visa`, or `Unknown`.
        :param pulumi.Input[str] country: String. Two-letter ISO code representing the country of the card. You could use this attribute to get a
               sense of the international breakdown of cards you’ve collected.
        :param pulumi.Input[str] customer: String. The customer that this card belongs to.
        :param pulumi.Input[int] cvc: Int. Card security code. Highly recommended to always include this value, but it's required only
               for accounts based in European countries.
        :param pulumi.Input[str] cvc_check: String. If a `cvc` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A
               result of `unchecked` indicates that CVC was provided but hasn’t been checked yet
        :param pulumi.Input[int] exp_month: Int. Number representing the card's expiration month.
        :param pulumi.Input[int] exp_year: Int. Four-digit number representing the card's expiration year.
        :param pulumi.Input[str] fingerprint: String. Uniquely identifies this particular card number. You can use this attribute to check whether
               two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize
               card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card
               number.
        :param pulumi.Input[str] funding: String. Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
        :param pulumi.Input[str] last4: String. The last four digits of the card.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
               storing additional information about the object in a structured format.
        :param pulumi.Input[str] name: String. Cardholder name.
        :param pulumi.Input[str] number: String. The card number, as a string without any separators.
        :param pulumi.Input[str] tokenization_method: String. If the card number is tokenized, this is the method that was used. Can
               be `android_pay` (includes Google Pay), `apple_pay`, `masterpass`, `visa_checkout`, or `null`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CardState.__new__(_CardState)

        __props__.__dict__["address"] = address
        __props__.__dict__["address_line1_check"] = address_line1_check
        __props__.__dict__["address_zip_check"] = address_zip_check
        __props__.__dict__["available_payout_methods"] = available_payout_methods
        __props__.__dict__["brand"] = brand
        __props__.__dict__["country"] = country
        __props__.__dict__["customer"] = customer
        __props__.__dict__["cvc"] = cvc
        __props__.__dict__["cvc_check"] = cvc_check
        __props__.__dict__["exp_month"] = exp_month
        __props__.__dict__["exp_year"] = exp_year
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["funding"] = funding
        __props__.__dict__["last4"] = last4
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["name"] = name
        __props__.__dict__["number"] = number
        __props__.__dict__["tokenization_method"] = tokenization_method
        return Card(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map(String). Address map with fields related to the address: `line1`, `line2`, `city`, `state`
        , `zip` and `country`.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="addressLine1Check")
    def address_line1_check(self) -> pulumi.Output[str]:
        """
        String. If address `line1` was provided, results of the check: `pass`, `fail`, `unavailable`,
        or `unchecked`.
        """
        return pulumi.get(self, "address_line1_check")

    @property
    @pulumi.getter(name="addressZipCheck")
    def address_zip_check(self) -> pulumi.Output[str]:
        """
        String. If address `zip` was provided, results of the check: `pass`, `fail`, `unavailable`,
        or `unchecked`.
        """
        return pulumi.get(self, "address_zip_check")

    @property
    @pulumi.getter(name="availablePayoutMethods")
    def available_payout_methods(self) -> pulumi.Output[Sequence[str]]:
        """
        List(String). A set of available payout methods for this card. Only values from this set
        should be passed as the method when creating a payout.
        """
        return pulumi.get(self, "available_payout_methods")

    @property
    @pulumi.getter
    def brand(self) -> pulumi.Output[str]:
        """
        String. Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`
        , `Visa`, or `Unknown`.
        """
        return pulumi.get(self, "brand")

    @property
    @pulumi.getter
    def country(self) -> pulumi.Output[str]:
        """
        String. Two-letter ISO code representing the country of the card. You could use this attribute to get a
        sense of the international breakdown of cards you’ve collected.
        """
        return pulumi.get(self, "country")

    @property
    @pulumi.getter
    def customer(self) -> pulumi.Output[str]:
        """
        String. The customer that this card belongs to.
        """
        return pulumi.get(self, "customer")

    @property
    @pulumi.getter
    def cvc(self) -> pulumi.Output[Optional[int]]:
        """
        Int. Card security code. Highly recommended to always include this value, but it's required only
        for accounts based in European countries.
        """
        return pulumi.get(self, "cvc")

    @property
    @pulumi.getter(name="cvcCheck")
    def cvc_check(self) -> pulumi.Output[str]:
        """
        String. If a `cvc` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A
        result of `unchecked` indicates that CVC was provided but hasn’t been checked yet
        """
        return pulumi.get(self, "cvc_check")

    @property
    @pulumi.getter(name="expMonth")
    def exp_month(self) -> pulumi.Output[int]:
        """
        Int. Number representing the card's expiration month.
        """
        return pulumi.get(self, "exp_month")

    @property
    @pulumi.getter(name="expYear")
    def exp_year(self) -> pulumi.Output[int]:
        """
        Int. Four-digit number representing the card's expiration year.
        """
        return pulumi.get(self, "exp_year")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        String. Uniquely identifies this particular card number. You can use this attribute to check whether
        two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize
        card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card
        number.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def funding(self) -> pulumi.Output[str]:
        """
        String. Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
        """
        return pulumi.get(self, "funding")

    @property
    @pulumi.getter
    def last4(self) -> pulumi.Output[str]:
        """
        String. The last four digits of the card.
        """
        return pulumi.get(self, "last4")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
        storing additional information about the object in a structured format.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        String. Cardholder name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def number(self) -> pulumi.Output[str]:
        """
        String. The card number, as a string without any separators.
        """
        return pulumi.get(self, "number")

    @property
    @pulumi.getter(name="tokenizationMethod")
    def tokenization_method(self) -> pulumi.Output[str]:
        """
        String. If the card number is tokenized, this is the method that was used. Can
        be `android_pay` (includes Google Pay), `apple_pay`, `masterpass`, `visa_checkout`, or `null`.
        """
        return pulumi.get(self, "tokenization_method")

