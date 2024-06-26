// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface FileLink {
    /**
     * String. Time at which the object was created. Measured in seconds since the Unix epoch.
     */
    created?: pulumi.Input<number>;
    /**
     * Bool. Returns if the link is already expired.
     */
    expired?: pulumi.Input<boolean>;
    /**
     * Int. Time that the link expires.
     */
    expiresAt?: pulumi.Input<number>;
    /**
     * String. Unique identifier for the object.
     */
    id?: pulumi.Input<string>;
    /**
     * Bool. Has the value `true` if the object exists in live mode or the value `false` 
     * if the object exists in test mode.
     */
    livemode?: pulumi.Input<boolean>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * String. String representing the object’s type. Objects of the same type share the same value.
     */
    object?: pulumi.Input<string>;
    /**
     * String. The publicly accessible URL to download the file.
     */
    url?: pulumi.Input<string>;
}

export interface FileLinkData {
    /**
     * Set this to true to create a file link for the newly created file. Creating a link is only possible when the file’s purpose is one of the following: business_icon, business_logo, customer_signature, dispute_evidence, pci_document, tax_document_user_upload, or terminal_reader_splashscreen.
     */
    create: pulumi.Input<boolean>;
    /**
     * Int. Time that the link expires.
     */
    expiresAt?: pulumi.Input<number>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface PortalConfigurationBusinessProfile {
    /**
     * String. The messaging shown to customers in the portal.
     */
    headline?: pulumi.Input<string>;
    /**
     * String. A link to the business's publicly available privacy policy.
     */
    privacyPolicyUrl?: pulumi.Input<string>;
    /**
     * String. A link to the business's publicly available terms of service.
     */
    termsOfServiceUrl?: pulumi.Input<string>;
}

export interface PortalConfigurationFeatures {
    /**
     * List(Resource). Information about updating the customer details in the portal. See Customer Update.
     */
    customerUpdate?: pulumi.Input<inputs.PortalConfigurationFeaturesCustomerUpdate>;
    /**
     * List(Resource). Information about showing the billing history in the portal. See Invoice History.
     */
    invoiceHistory?: pulumi.Input<inputs.PortalConfigurationFeaturesInvoiceHistory>;
    /**
     * List(Resource). Information about updating payment methods in the portal. See Payment Method Update.
     */
    paymentMethodUpdate?: pulumi.Input<inputs.PortalConfigurationFeaturesPaymentMethodUpdate>;
    /**
     * List(Resource). Information about canceling subscriptions in the portal. See Subscription Cancel.
     */
    subscriptionCancel?: pulumi.Input<inputs.PortalConfigurationFeaturesSubscriptionCancel>;
    /**
     * List(Resource). Information about pausing subscriptions in the portal. See Subscription Pause.
     */
    subscriptionPauses?: pulumi.Input<pulumi.Input<inputs.PortalConfigurationFeaturesSubscriptionPause>[]>;
    /**
     * List(Resource). Information about updating subscriptions in the portal. See Subscription Update.
     */
    subscriptionUpdates?: pulumi.Input<pulumi.Input<inputs.PortalConfigurationFeaturesSubscriptionUpdate>[]>;
}

export interface PortalConfigurationFeaturesCustomerUpdate {
    /**
     * The types of customer updates that are supported. When empty, customers are not updatable.
     */
    allowedUpdates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the feature is enabled.
     */
    enabled: pulumi.Input<boolean>;
}

export interface PortalConfigurationFeaturesInvoiceHistory {
    /**
     * Whether the feature is enabled.
     */
    enabled: pulumi.Input<boolean>;
}

export interface PortalConfigurationFeaturesPaymentMethodUpdate {
    /**
     * Whether the feature is enabled.
     */
    enabled: pulumi.Input<boolean>;
}

export interface PortalConfigurationFeaturesSubscriptionCancel {
    /**
     * Whether the cancellation reasons will be collected in the portal and which options are exposed to the customer
     */
    cancellationReason?: pulumi.Input<inputs.PortalConfigurationFeaturesSubscriptionCancelCancellationReason>;
    /**
     * Whether the feature is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Whether to cancel subscriptions immediately or at the end of the billing period.
     */
    mode?: pulumi.Input<string>;
    /**
     * Whether to create prorations when canceling subscriptions.
     */
    prorationBehavior?: pulumi.Input<string>;
}

export interface PortalConfigurationFeaturesSubscriptionCancelCancellationReason {
    /**
     * Bool. Whether the feature is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * List(String). Which cancellation reasons will be given as options to the customer. Supported values are `tooExpensive`, `missingFeatures`, `switchedService`, `unused`, `customerService`, `tooComplex`, `lowQuality`, and `other`.
     */
    options: pulumi.Input<pulumi.Input<string>[]>;
}

export interface PortalConfigurationFeaturesSubscriptionPause {
    /**
     * Whether the feature is enabled.
     */
    enabled?: pulumi.Input<boolean>;
}

export interface PortalConfigurationFeaturesSubscriptionUpdate {
    /**
     * The types of subscription updates that are supported. When empty, subscriptions are not updateable.
     */
    defaultAllowedUpdates: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the feature is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The list of products that support subscription updates.
     */
    products: pulumi.Input<pulumi.Input<inputs.PortalConfigurationFeaturesSubscriptionUpdateProduct>[]>;
    /**
     * Determines how to handle prorations resulting from subscription updates
     */
    prorationBehavior?: pulumi.Input<string>;
}

export interface PortalConfigurationFeaturesSubscriptionUpdateProduct {
    /**
     * List(String). The list of price IDs for the product that a subscription can be updated to.
     */
    prices: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String. The product id.
     */
    product: pulumi.Input<string>;
}

export interface PortalConfigurationLoginPage {
    /**
     * Bool. Set to true to generate a shareable URL login_page.url that will take your customers to a hosted login page for the customer portal.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * A shareable URL to the hosted portal login page. Your customers will be able to log in with their email and receive a link to their customer portal.
     */
    url?: pulumi.Input<string>;
}

export interface PriceCurrencyOption {
    /**
     * String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
     */
    currency: pulumi.Input<string>;
    /**
     * List(Resource). When set, 
     * provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
     * See details in custom unit amount.
     */
    customUnitAmount?: pulumi.Input<inputs.PriceCurrencyOptionCustomUnitAmount>;
    /**
     * String. Only required if a default tax behavior was not provided in the Stripe Tax settings. 
     * Specifies whether the price is considered inclusive of taxes or exclusive of taxes.
     * One of `inclusive`, `exclusive`, or `unspecified`.
     * Once specified as either inclusive or exclusive, it cannot be changed.
     */
    taxBehavior?: pulumi.Input<string>;
    /**
     * List(Resource). Each element represents a pricing tier. 
     * This parameter requires `billingScheme` to be set to `tiered`. This resource can be used more than once and follows
     * the same fields as the root tiers block
     */
    tiers?: pulumi.Input<pulumi.Input<inputs.PriceCurrencyOptionTier>[]>;
    /**
     * Int. A positive integer in cents (or -1 for a free price) representing how much to charge.
     */
    unitAmount?: pulumi.Input<number>;
    /**
     * Float. Same as unit_amount, but accepts a decimal value in cents with at most 12
     * decimal places. Only one of unitAmount and unitAmountDecimal can be set.
     */
    unitAmountDecimal?: pulumi.Input<number>;
}

export interface PriceCurrencyOptionCustomUnitAmount {
    /**
     * Bool. Pass in `true` to enable `customUnitAmount`, otherwise omit `customUnitAmount`.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Int. The maximum unit amount the customer can specify for this item.
     */
    maximum?: pulumi.Input<number>;
    /**
     * Int. The minimum unit amount the customer can specify for this item. 
     * Must be at least the minimum charge amount.
     */
    minimum?: pulumi.Input<number>;
    /**
     * Int. The starting unit amount which can be updated by the customer.
     */
    preset?: pulumi.Input<number>;
}

export interface PriceCurrencyOptionTier {
    /**
     * Int. The flat billing amount for an entire tier, regardless of the number of units in the
     * tier.
     */
    flatAmount?: pulumi.Input<number>;
    /**
     * Float. Same as `flatAmount`, but accepts a decimal value representing an integer
     * in the minor units of the currency. Only one of `flatAmount` and `flatAmountDecimal` can be set.
     */
    flatAmountDecimal?: pulumi.Input<number>;
    /**
     * Int. The per-unit billing amount for each individual unit for which this tier applies.
     */
    unitAmount?: pulumi.Input<number>;
    /**
     * Float. Same as `unitAmount`, but accepts a decimal value in cents with at most 12
     * decimal places. Only one of `unitAmount` and `unitAmountDecimal` can be set.
     */
    unitAmountDecimal?: pulumi.Input<number>;
    /**
     * Int. Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the
     * previous tier adding one. Use `-1` to define a fallback tier.
     */
    upTo?: pulumi.Input<number>;
}

export interface PriceRecurring {
    /**
     * String. Specifies a usage of aggregation strategy for prices of `usage_type=metered`.
     * Allowed values are `sum` for summing up all usage during a period, `lastDuringPeriod` for using the last usage
     * record reported within a period, `lastEver` for using the last usage record ever (across period bounds) or `max`
     * which uses the usage record with the maximum reported usage during a period.
     */
    aggregateUsage?: pulumi.Input<string>;
    /**
     * String. Specifies billing frequency. Either `day`, `week`, `month` or `year`.
     */
    interval: pulumi.Input<string>;
    /**
     * Int. This parameter is (Required) when interval value is set. The number of intervals between subscription billings. For
     * example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year,
     * 12 months, or 52 weeks).
     */
    intervalCount?: pulumi.Input<number>;
    /**
     * String. Configures how the quantity per period should be determined. Can be either `metered`
     * or `licensed`. `licensed` automatically bills the quantity set when adding it to a subscription. `metered` aggregates
     * the total usage based on usage records. Defaults to `licensed`.
     */
    usageType?: pulumi.Input<string>;
}

export interface PriceTier {
    /**
     * Int. The flat billing amount for an entire tier, regardless of the number of units in the
     * tier.
     */
    flatAmount?: pulumi.Input<number>;
    /**
     * Float. Same as `flatAmount`, but accepts a decimal value representing an integer
     * in the minor units of the currency. Only one of `flatAmount` and `flatAmountDecimal` can be set.
     */
    flatAmountDecimal?: pulumi.Input<number>;
    /**
     * Int. The per-unit billing amount for each individual unit for which this tier applies.
     */
    unitAmount?: pulumi.Input<number>;
    /**
     * Float. Same as `unitAmount`, but accepts a decimal value in cents with at most 12
     * decimal places. Only one of `unitAmount` and `unitAmountDecimal` can be set.
     */
    unitAmountDecimal?: pulumi.Input<number>;
    /**
     * Int. Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the
     * previous tier adding one. Use `-1` to define a fallback tier.
     */
    upTo?: pulumi.Input<number>;
}

export interface PriceTransformQuantity {
    /**
     * Int. Divide usage by this number.
     */
    divideBy: pulumi.Input<number>;
    /**
     * String. After division, either round the result `up` or `down`.
     */
    round: pulumi.Input<string>;
}

export interface PromotionCodeRestrictions {
    /**
     * Bool. A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices.
     */
    firstTimeTransaction: pulumi.Input<boolean>;
    /**
     * Int. Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).
     */
    minimumAmount: pulumi.Input<number>;
    /**
     * String. Three-letter ISO code for `minimumAmount`.
     */
    minimumAmountCurrency: pulumi.Input<string>;
}

export interface ShippingRateDeliveryEstimate {
    /**
     * List(Resource. The upper bound of the estimated range.
     * Please see Delivery Estimate Definition.
     */
    maximum?: pulumi.Input<inputs.ShippingRateDeliveryEstimateMaximum>;
    /**
     * List(Resource). The lower bound of the estimated range. 
     * Please see Delivery Estimate Definition.
     */
    minimum?: pulumi.Input<inputs.ShippingRateDeliveryEstimateMinimum>;
}

export interface ShippingRateDeliveryEstimateMaximum {
    /**
     * The upper bound of the estimated range. If empty, represents no lower bound.
     */
    unit: pulumi.Input<string>;
    /**
     * Must be greater than 0.
     */
    value: pulumi.Input<number>;
}

export interface ShippingRateDeliveryEstimateMinimum {
    /**
     * The lower bound of the estimated range. If empty, represents no lower bound.
     */
    unit: pulumi.Input<string>;
    /**
     * Must be greater than 0.
     */
    value: pulumi.Input<number>;
}

export interface ShippingRateFixedAmount {
    /**
     * Int. A non-negative integer in cents representing how much to charge.
     */
    amount: pulumi.Input<number>;
    /**
     * String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
     */
    currency: pulumi.Input<string>;
    /**
     * List(Resource). Please see argument details Currency Option
     */
    currencyOptions?: pulumi.Input<pulumi.Input<inputs.ShippingRateFixedAmountCurrencyOption>[]>;
}

export interface ShippingRateFixedAmountCurrencyOption {
    /**
     * Int. (Required) Int. A non-negative integer in cents representing how much to charge.
     */
    amount: pulumi.Input<number>;
    /**
     * String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
     */
    currency: pulumi.Input<string>;
    /**
     * Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or unspecified.
     */
    taxBehavior?: pulumi.Input<string>;
}
