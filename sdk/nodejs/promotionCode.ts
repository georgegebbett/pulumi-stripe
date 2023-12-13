// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * With this resource, you can create a promotion code - [Stripe API promotion code documentation](https://stripe.com/docs/api/promotion_codes).
 *
 * A Promotion Code represents a customer-redeemable code for a coupon. It can be used to create multiple codes for a single coupon.
 *
 * > Removal of the promotion code isn't supported through the Stripe SDK.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as stripe from "pulumi-stripe";
 *
 * // promotion code for the coupon
 * const codePromotionCode = new stripe.PromotionCode("codePromotionCode", {
 *     coupon: stripe_coupon.coupon.id,
 *     code: "FREE",
 * });
 * // promotion code for the coupon with limitations
 * const codeIndex_promotionCodePromotionCode = new stripe.PromotionCode("codeIndex/promotionCodePromotionCode", {
 *     coupon: stripe_coupon.coupon.id,
 *     code: "FREE",
 *     maxRedemptions: 5,
 *     expiresAt: "2025-08-03T08:37:18+00:00",
 * });
 * // promotion code for the coupon to customer
 * const codeStripeIndex_promotionCodePromotionCode = new stripe.PromotionCode("codeStripeIndex/promotionCodePromotionCode", {
 *     coupon: stripe_coupon.coupon.id,
 *     code: "FREE",
 *     customer: "cus...",
 * });
 * // promotion code for the coupon with restrictions
 * const codeStripeIndex_promotionCodePromotionCode1 = new stripe.PromotionCode("codeStripeIndex/promotionCodePromotionCode1", {
 *     coupon: stripe_coupon.coupon.id,
 *     code: "FREE",
 *     restrictions: {
 *         firstTimeTransaction: true,
 *         minimumAmount: 100,
 *         minimumAmountCurrency: "aud",
 *     },
 * });
 * ```
 */
export class PromotionCode extends pulumi.CustomResource {
    /**
     * Get an existing PromotionCode resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PromotionCodeState, opts?: pulumi.CustomResourceOptions): PromotionCode {
        return new PromotionCode(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'stripe:index/promotionCode:PromotionCode';

    /**
     * Returns true if the given object is an instance of PromotionCode.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PromotionCode {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PromotionCode.__pulumiType;
    }

    /**
     * Bool. Whether the promotion code is currently active. Defaults to `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
     */
    public readonly code!: pulumi.Output<string | undefined>;
    /**
     * String. The coupon for this promotion code.
     */
    public readonly coupon!: pulumi.Output<string>;
    /**
     * String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
     */
    public readonly customer!: pulumi.Output<string | undefined>;
    /**
     * String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
     */
    public readonly expiresAt!: pulumi.Output<string | undefined>;
    /**
     * Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
     */
    public readonly maxRedemptions!: pulumi.Output<number | undefined>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
     */
    public readonly restrictions!: pulumi.Output<outputs.PromotionCodeRestrictions | undefined>;

    /**
     * Create a PromotionCode resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PromotionCodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PromotionCodeArgs | PromotionCodeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PromotionCodeState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["code"] = state ? state.code : undefined;
            resourceInputs["coupon"] = state ? state.coupon : undefined;
            resourceInputs["customer"] = state ? state.customer : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["maxRedemptions"] = state ? state.maxRedemptions : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["restrictions"] = state ? state.restrictions : undefined;
        } else {
            const args = argsOrState as PromotionCodeArgs | undefined;
            if ((!args || args.coupon === undefined) && !opts.urn) {
                throw new Error("Missing required property 'coupon'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["code"] = args ? args.code : undefined;
            resourceInputs["coupon"] = args ? args.coupon : undefined;
            resourceInputs["customer"] = args ? args.customer : undefined;
            resourceInputs["expiresAt"] = args ? args.expiresAt : undefined;
            resourceInputs["maxRedemptions"] = args ? args.maxRedemptions : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["restrictions"] = args ? args.restrictions : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PromotionCode.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PromotionCode resources.
 */
export interface PromotionCodeState {
    /**
     * Bool. Whether the promotion code is currently active. Defaults to `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
     */
    code?: pulumi.Input<string>;
    /**
     * String. The coupon for this promotion code.
     */
    coupon?: pulumi.Input<string>;
    /**
     * String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
     */
    customer?: pulumi.Input<string>;
    /**
     * String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
     */
    maxRedemptions?: pulumi.Input<number>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
     */
    restrictions?: pulumi.Input<inputs.PromotionCodeRestrictions>;
}

/**
 * The set of arguments for constructing a PromotionCode resource.
 */
export interface PromotionCodeArgs {
    /**
     * Bool. Whether the promotion code is currently active. Defaults to `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * String. The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. If left blank, we will generate one automatically.
     */
    code?: pulumi.Input<string>;
    /**
     * String. The coupon for this promotion code.
     */
    coupon: pulumi.Input<string>;
    /**
     * String. The customer that this promotion code can be used by. If not set, the promotion code can be used by all customers.
     */
    customer?: pulumi.Input<string>;
    /**
     * String. The timestamp at which this promotion code will expire. If the coupon has specified a `redeemsBy`, then this value cannot be after the coupon’s `redeemsBy`. Expected format is `RFC3339`.
     */
    expiresAt?: pulumi.Input<string>;
    /**
     * Int. A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `maxRedemptions`, then this value cannot be greater than the coupon’s `maxRedemptions`.
     */
    maxRedemptions?: pulumi.Input<number>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List(Resource). Settings that restrict the redemption of the promotion code. For details of individual arguments see Restrictions.
     */
    restrictions?: pulumi.Input<inputs.PromotionCodeRestrictions>;
}