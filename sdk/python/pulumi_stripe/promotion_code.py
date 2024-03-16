# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PromotionCodeArgs', 'PromotionCode']

@pulumi.input_type
class PromotionCodeArgs:
    def __init__(__self__, *,
                 coupon: pulumi.Input[str],
                 active: Optional[pulumi.Input[bool]] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 max_redemptions: Optional[pulumi.Input[int]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 restrictions: Optional[pulumi.Input['PromotionCodeRestrictionsArgs']] = None):
        """
        The set of arguments for constructing a PromotionCode resource.
        :param pulumi.Input[str] coupon: String. The coupon for this promotion code.
        :param pulumi.Input[bool] active: Bool. Whether the promotion code is currently active. Defaults to `true`.
        :param pulumi.Input[str] code: String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
        :param pulumi.Input[str] customer: String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
        :param pulumi.Input[str] expires_at: String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon’s `redeems_by`. Expected format is `RFC3339`.
        :param pulumi.Input[int] max_redemptions: Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon’s `max_redemptions`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        :param pulumi.Input['PromotionCodeRestrictionsArgs'] restrictions: List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
        """
        pulumi.set(__self__, "coupon", coupon)
        if active is not None:
            pulumi.set(__self__, "active", active)
        if code is not None:
            pulumi.set(__self__, "code", code)
        if customer is not None:
            pulumi.set(__self__, "customer", customer)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if max_redemptions is not None:
            pulumi.set(__self__, "max_redemptions", max_redemptions)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if restrictions is not None:
            pulumi.set(__self__, "restrictions", restrictions)

    @property
    @pulumi.getter
    def coupon(self) -> pulumi.Input[str]:
        """
        String. The coupon for this promotion code.
        """
        return pulumi.get(self, "coupon")

    @coupon.setter
    def coupon(self, value: pulumi.Input[str]):
        pulumi.set(self, "coupon", value)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool. Whether the promotion code is currently active. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input[str]]:
        """
        String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter
    def customer(self) -> Optional[pulumi.Input[str]]:
        """
        String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
        """
        return pulumi.get(self, "customer")

    @customer.setter
    def customer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon’s `redeems_by`. Expected format is `RFC3339`.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="maxRedemptions")
    def max_redemptions(self) -> Optional[pulumi.Input[int]]:
        """
        Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon’s `max_redemptions`.
        """
        return pulumi.get(self, "max_redemptions")

    @max_redemptions.setter
    def max_redemptions(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_redemptions", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def restrictions(self) -> Optional[pulumi.Input['PromotionCodeRestrictionsArgs']]:
        """
        List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
        """
        return pulumi.get(self, "restrictions")

    @restrictions.setter
    def restrictions(self, value: Optional[pulumi.Input['PromotionCodeRestrictionsArgs']]):
        pulumi.set(self, "restrictions", value)


@pulumi.input_type
class _PromotionCodeState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 coupon: Optional[pulumi.Input[str]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 max_redemptions: Optional[pulumi.Input[int]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 restrictions: Optional[pulumi.Input['PromotionCodeRestrictionsArgs']] = None):
        """
        Input properties used for looking up and filtering PromotionCode resources.
        :param pulumi.Input[bool] active: Bool. Whether the promotion code is currently active. Defaults to `true`.
        :param pulumi.Input[str] code: String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
        :param pulumi.Input[str] coupon: String. The coupon for this promotion code.
        :param pulumi.Input[str] customer: String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
        :param pulumi.Input[str] expires_at: String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon’s `redeems_by`. Expected format is `RFC3339`.
        :param pulumi.Input[int] max_redemptions: Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon’s `max_redemptions`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        :param pulumi.Input['PromotionCodeRestrictionsArgs'] restrictions: List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if code is not None:
            pulumi.set(__self__, "code", code)
        if coupon is not None:
            pulumi.set(__self__, "coupon", coupon)
        if customer is not None:
            pulumi.set(__self__, "customer", customer)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if max_redemptions is not None:
            pulumi.set(__self__, "max_redemptions", max_redemptions)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if restrictions is not None:
            pulumi.set(__self__, "restrictions", restrictions)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        Bool. Whether the promotion code is currently active. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def code(self) -> Optional[pulumi.Input[str]]:
        """
        String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
        """
        return pulumi.get(self, "code")

    @code.setter
    def code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "code", value)

    @property
    @pulumi.getter
    def coupon(self) -> Optional[pulumi.Input[str]]:
        """
        String. The coupon for this promotion code.
        """
        return pulumi.get(self, "coupon")

    @coupon.setter
    def coupon(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "coupon", value)

    @property
    @pulumi.getter
    def customer(self) -> Optional[pulumi.Input[str]]:
        """
        String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
        """
        return pulumi.get(self, "customer")

    @customer.setter
    def customer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon’s `redeems_by`. Expected format is `RFC3339`.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="maxRedemptions")
    def max_redemptions(self) -> Optional[pulumi.Input[int]]:
        """
        Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon’s `max_redemptions`.
        """
        return pulumi.get(self, "max_redemptions")

    @max_redemptions.setter
    def max_redemptions(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_redemptions", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def restrictions(self) -> Optional[pulumi.Input['PromotionCodeRestrictionsArgs']]:
        """
        List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
        """
        return pulumi.get(self, "restrictions")

    @restrictions.setter
    def restrictions(self, value: Optional[pulumi.Input['PromotionCodeRestrictionsArgs']]):
        pulumi.set(self, "restrictions", value)


class PromotionCode(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 coupon: Optional[pulumi.Input[str]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 max_redemptions: Optional[pulumi.Input[int]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 restrictions: Optional[pulumi.Input[pulumi.InputType['PromotionCodeRestrictionsArgs']]] = None,
                 __props__=None):
        """
        With this resource, you can create a promotion code - [Stripe API promotion code documentation](https://stripe.com/docs/api/promotion_codes).

        A Promotion Code represents a customer-redeemable code for a coupon. It can be used to create multiple codes for a single coupon.

        > Removal of the promotion code isn't supported through the Stripe SDK.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_stripe as stripe

        # promotion code for the coupon
        code_promotion_code = stripe.PromotionCode("codePromotionCode",
            coupon=stripe_coupon["coupon"]["id"],
            code="FREE")
        # promotion code for the coupon with limitations
        code_index_promotion_code_promotion_code = stripe.PromotionCode("codeIndex/promotionCodePromotionCode",
            coupon=stripe_coupon["coupon"]["id"],
            code="FREE",
            max_redemptions=5,
            expires_at="2025-08-03T08:37:18+00:00")
        # promotion code for the coupon to customer
        code_stripe_index_promotion_code_promotion_code = stripe.PromotionCode("codeStripeIndex/promotionCodePromotionCode",
            coupon=stripe_coupon["coupon"]["id"],
            code="FREE",
            customer="cus...")
        # promotion code for the coupon with restrictions
        code_stripe_index_promotion_code_promotion_code1 = stripe.PromotionCode("codeStripeIndex/promotionCodePromotionCode1",
            coupon=stripe_coupon["coupon"]["id"],
            code="FREE",
            restrictions=stripe.PromotionCodeRestrictionsArgs(
                first_time_transaction=True,
                minimum_amount=100,
                minimum_amount_currency="aud",
            ))
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Bool. Whether the promotion code is currently active. Defaults to `true`.
        :param pulumi.Input[str] code: String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
        :param pulumi.Input[str] coupon: String. The coupon for this promotion code.
        :param pulumi.Input[str] customer: String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
        :param pulumi.Input[str] expires_at: String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon’s `redeems_by`. Expected format is `RFC3339`.
        :param pulumi.Input[int] max_redemptions: Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon’s `max_redemptions`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        :param pulumi.Input[pulumi.InputType['PromotionCodeRestrictionsArgs']] restrictions: List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PromotionCodeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        With this resource, you can create a promotion code - [Stripe API promotion code documentation](https://stripe.com/docs/api/promotion_codes).

        A Promotion Code represents a customer-redeemable code for a coupon. It can be used to create multiple codes for a single coupon.

        > Removal of the promotion code isn't supported through the Stripe SDK.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_stripe as stripe

        # promotion code for the coupon
        code_promotion_code = stripe.PromotionCode("codePromotionCode",
            coupon=stripe_coupon["coupon"]["id"],
            code="FREE")
        # promotion code for the coupon with limitations
        code_index_promotion_code_promotion_code = stripe.PromotionCode("codeIndex/promotionCodePromotionCode",
            coupon=stripe_coupon["coupon"]["id"],
            code="FREE",
            max_redemptions=5,
            expires_at="2025-08-03T08:37:18+00:00")
        # promotion code for the coupon to customer
        code_stripe_index_promotion_code_promotion_code = stripe.PromotionCode("codeStripeIndex/promotionCodePromotionCode",
            coupon=stripe_coupon["coupon"]["id"],
            code="FREE",
            customer="cus...")
        # promotion code for the coupon with restrictions
        code_stripe_index_promotion_code_promotion_code1 = stripe.PromotionCode("codeStripeIndex/promotionCodePromotionCode1",
            coupon=stripe_coupon["coupon"]["id"],
            code="FREE",
            restrictions=stripe.PromotionCodeRestrictionsArgs(
                first_time_transaction=True,
                minimum_amount=100,
                minimum_amount_currency="aud",
            ))
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param PromotionCodeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PromotionCodeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 code: Optional[pulumi.Input[str]] = None,
                 coupon: Optional[pulumi.Input[str]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 max_redemptions: Optional[pulumi.Input[int]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 restrictions: Optional[pulumi.Input[pulumi.InputType['PromotionCodeRestrictionsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PromotionCodeArgs.__new__(PromotionCodeArgs)

            __props__.__dict__["active"] = active
            __props__.__dict__["code"] = code
            if coupon is None and not opts.urn:
                raise TypeError("Missing required property 'coupon'")
            __props__.__dict__["coupon"] = coupon
            __props__.__dict__["customer"] = customer
            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["max_redemptions"] = max_redemptions
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["restrictions"] = restrictions
        super(PromotionCode, __self__).__init__(
            'stripe:index/promotionCode:PromotionCode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            code: Optional[pulumi.Input[str]] = None,
            coupon: Optional[pulumi.Input[str]] = None,
            customer: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            max_redemptions: Optional[pulumi.Input[int]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            restrictions: Optional[pulumi.Input[pulumi.InputType['PromotionCodeRestrictionsArgs']]] = None) -> 'PromotionCode':
        """
        Get an existing PromotionCode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Bool. Whether the promotion code is currently active. Defaults to `true`.
        :param pulumi.Input[str] code: String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
        :param pulumi.Input[str] coupon: String. The coupon for this promotion code.
        :param pulumi.Input[str] customer: String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
        :param pulumi.Input[str] expires_at: String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon’s `redeems_by`. Expected format is `RFC3339`.
        :param pulumi.Input[int] max_redemptions: Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon’s `max_redemptions`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        :param pulumi.Input[pulumi.InputType['PromotionCodeRestrictionsArgs']] restrictions: List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PromotionCodeState.__new__(_PromotionCodeState)

        __props__.__dict__["active"] = active
        __props__.__dict__["code"] = code
        __props__.__dict__["coupon"] = coupon
        __props__.__dict__["customer"] = customer
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["max_redemptions"] = max_redemptions
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["restrictions"] = restrictions
        return PromotionCode(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[Optional[bool]]:
        """
        Bool. Whether the promotion code is currently active. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter
    def code(self) -> pulumi.Output[Optional[str]]:
        """
        String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def coupon(self) -> pulumi.Output[str]:
        """
        String. The coupon for this promotion code.
        """
        return pulumi.get(self, "coupon")

    @property
    @pulumi.getter
    def customer(self) -> pulumi.Output[Optional[str]]:
        """
        String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
        """
        return pulumi.get(self, "customer")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon’s `redeems_by`. Expected format is `RFC3339`.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="maxRedemptions")
    def max_redemptions(self) -> pulumi.Output[Optional[int]]:
        """
        Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon’s `max_redemptions`.
        """
        return pulumi.get(self, "max_redemptions")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def restrictions(self) -> pulumi.Output[Optional['outputs.PromotionCodeRestrictions']]:
        """
        List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
        """
        return pulumi.get(self, "restrictions")

