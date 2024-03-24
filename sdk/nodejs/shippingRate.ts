// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * With this resource, you can create a shipping rate - [Stripe API shipping rate documentation](https://stripe.com/docs/api/shipping_rates).
 *
 * Shipping rates let you display various shipping options—like standard, express, and overnight—with more accurate delivery estimates.
 * Charge your customer for shipping using different Stripe products, some of which require coding.
 *
 * > Removal of the shipping rate isn't supported through the Stripe SDK. The best practice, which this provider follows,
 * is to archive the shipping rate by marking it as inactive on destroy, which indicates that the shipping rate is no longer
 * available.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as stripe from "pulumi-stripe";
 *
 * // shipping rate with delivery estimate
 * const shippingRate = new stripe.ShippingRate("shippingRate", {
 *     deliveryEstimates: [{
 *         maximum: {
 *             unit: "day",
 *             value: 4,
 *         },
 *         minimum: {
 *             unit: "hour",
 *             value: 24,
 *         },
 *     }],
 *     displayName: "shipping rate",
 *     fixedAmount: {
 *         amount: 1000,
 *         currency: "aud",
 *     },
 * });
 * // shipping rate with currency options
 * // !!! Currency options have to be sorted alphabetically 
 * // !!! by the currency field
 * const shipping = new stripe.ShippingRate("shipping", {
 *     displayName: "shipping rate",
 *     fixedAmount: {
 *         amount: 1000,
 *         currency: "aud",
 *         currencyOptions: [
 *             {
 *                 amount: 350,
 *                 currency: "eur",
 *             },
 *             {
 *                 amount: 500,
 *                 currency: "usd",
 *             },
 *         ],
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class ShippingRate extends pulumi.CustomResource {
    /**
     * Get an existing ShippingRate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ShippingRateState, opts?: pulumi.CustomResourceOptions): ShippingRate {
        return new ShippingRate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'stripe:index/shippingRate:ShippingRate';

    /**
     * Returns true if the given object is an instance of ShippingRate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ShippingRate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ShippingRate.__pulumiType;
    }

    /**
     * Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * List(Resource). The estimated range for how long shipping will take, 
     * meant to be displayable to the customer. This will appear on CheckoutSessions.
     * For details please see Delivery Estimate.
     */
    public readonly deliveryEstimates!: pulumi.Output<outputs.ShippingRateDeliveryEstimate[] | undefined>;
    /**
     * String. The name of the shipping rate, meant to be displayable to the customer. 
     * This will appear on CheckoutSessions.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * List(Resource). Describes a fixed amount to charge for shipping. 
     * Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
     */
    public readonly fixedAmount!: pulumi.Output<outputs.ShippingRateFixedAmount>;
    /**
     * Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
     */
    public /*out*/ readonly livemode!: pulumi.Output<boolean>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
     * storing additional information about the object in a structured format.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
     * unspecified.
     */
    public readonly taxBehavior!: pulumi.Output<string | undefined>;
    /**
     * String. A tax code ID. The Shipping tax code is `txcd92010001`.
     */
    public readonly taxCode!: pulumi.Output<string | undefined>;
    /**
     * String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a ShippingRate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ShippingRateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ShippingRateArgs | ShippingRateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ShippingRateState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["deliveryEstimates"] = state ? state.deliveryEstimates : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["fixedAmount"] = state ? state.fixedAmount : undefined;
            resourceInputs["livemode"] = state ? state.livemode : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["taxBehavior"] = state ? state.taxBehavior : undefined;
            resourceInputs["taxCode"] = state ? state.taxCode : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ShippingRateArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.fixedAmount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fixedAmount'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["deliveryEstimates"] = args ? args.deliveryEstimates : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["fixedAmount"] = args ? args.fixedAmount : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["taxBehavior"] = args ? args.taxBehavior : undefined;
            resourceInputs["taxCode"] = args ? args.taxCode : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["livemode"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ShippingRate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ShippingRate resources.
 */
export interface ShippingRateState {
    /**
     * Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * List(Resource). The estimated range for how long shipping will take, 
     * meant to be displayable to the customer. This will appear on CheckoutSessions.
     * For details please see Delivery Estimate.
     */
    deliveryEstimates?: pulumi.Input<pulumi.Input<inputs.ShippingRateDeliveryEstimate>[]>;
    /**
     * String. The name of the shipping rate, meant to be displayable to the customer. 
     * This will appear on CheckoutSessions.
     */
    displayName?: pulumi.Input<string>;
    /**
     * List(Resource). Describes a fixed amount to charge for shipping. 
     * Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
     */
    fixedAmount?: pulumi.Input<inputs.ShippingRateFixedAmount>;
    /**
     * Bool. Has the value true if the object exists in live mode or the value false if the object exists in test mode.
     */
    livemode?: pulumi.Input<boolean>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
     * storing additional information about the object in a structured format.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
     * unspecified.
     */
    taxBehavior?: pulumi.Input<string>;
    /**
     * String. A tax code ID. The Shipping tax code is `txcd92010001`.
     */
    taxCode?: pulumi.Input<string>;
    /**
     * String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ShippingRate resource.
 */
export interface ShippingRateArgs {
    /**
     * Bool. Whether the shipping rate is active (can't be used when creating). Defaults to `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * List(Resource). The estimated range for how long shipping will take, 
     * meant to be displayable to the customer. This will appear on CheckoutSessions.
     * For details please see Delivery Estimate.
     */
    deliveryEstimates?: pulumi.Input<pulumi.Input<inputs.ShippingRateDeliveryEstimate>[]>;
    /**
     * String. The name of the shipping rate, meant to be displayable to the customer. 
     * This will appear on CheckoutSessions.
     */
    displayName: pulumi.Input<string>;
    /**
     * List(Resource). Describes a fixed amount to charge for shipping. 
     * Must be present if type is `fixedAmount`. For details of individual arguments see Fixed Amount.
     */
    fixedAmount: pulumi.Input<inputs.ShippingRateFixedAmount>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for
     * storing additional information about the object in a structured format.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or
     * unspecified.
     */
    taxBehavior?: pulumi.Input<string>;
    /**
     * String. A tax code ID. The Shipping tax code is `txcd92010001`.
     */
    taxCode?: pulumi.Input<string>;
    /**
     * String. The type of calculation to use on the shipping rate. Can only be `fixedAmount` for now.
     */
    type?: pulumi.Input<string>;
}
