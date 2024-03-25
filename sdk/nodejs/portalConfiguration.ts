// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * With this resource, you can create a Customer Portal Configuration - [Stripe API portal configuration documentation](https://stripe.com/docs/api/customer_portal/configuration).
 *
 * The Billing customer portal is a Stripe-hosted UI for subscription and billing management.
 *
 * A portal configuration describes the functionality and features that you want to provide to your customers through the portal.
 *
 * > Removal of the Customer Portal isn't supported through the Stripe SDK. The best practice, which this provider follows,
 * is to deactivate the Customer Portal by marking it as inactive on destroy, which indicates that resource is no longer
 * available.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as stripe from "pulumi-stripe";
 *
 * // A billing portal using all the available options
 * const portalConfiguration = new stripe.PortalConfiguration("portalConfiguration", {
 *     businessProfile: {
 *         headline: "My special headline",
 *         privacyPolicyUrl: "https://example.com/privacy",
 *         termsOfServiceUrl: "https://example.com/terms",
 *     },
 *     defaultReturnUrl: "https://example.com/special_headline",
 *     features: {
 *         customerUpdate: {
 *             allowedUpdates: [
 *                 "email",
 *                 "address",
 *                 "shipping",
 *                 "phone",
 *                 "tax_id",
 *             ],
 *             enabled: true,
 *         },
 *         invoiceHistory: {
 *             enabled: true,
 *         },
 *         paymentMethodUpdate: {
 *             enabled: true,
 *         },
 *         subscriptionCancel: {
 *             cancellationReason: {
 *                 enabled: true,
 *                 options: [
 *                     "too_expensive",
 *                     "missing_features",
 *                     "switched_service",
 *                     "unused",
 *                     "customer_service",
 *                     "too_complex",
 *                     "low_quality",
 *                     "other",
 *                 ],
 *             },
 *             enabled: true,
 *             mode: "at_period_end",
 *             prorationBehavior: "none",
 *         },
 *         subscriptionPauses: [{
 *             enabled: true,
 *         }],
 *         subscriptionUpdates: [{
 *             defaultAllowedUpdates: [
 *                 "price",
 *                 "quantity",
 *                 "promotion_code",
 *             ],
 *             enabled: true,
 *             products: [{
 *                 prices: [
 *                     "my_price_id1",
 *                     "my_price_id2",
 *                 ],
 *                 product: "my_product_id",
 *             }],
 *             prorationBehavior: "none",
 *         }],
 *     },
 *     metadata: {
 *         foo: "bar",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class PortalConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing PortalConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortalConfigurationState, opts?: pulumi.CustomResourceOptions): PortalConfiguration {
        return new PortalConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'stripe:index/portalConfiguration:PortalConfiguration';

    /**
     * Returns true if the given object is an instance of PortalConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PortalConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PortalConfiguration.__pulumiType;
    }

    /**
     * Bool. Whether the configuration is active and can be used to create portal sessions. (On create it is always set as `true`)
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * List(Resource). The business information shown to customers in the portal. More details in Business Profile section
     */
    public readonly businessProfile!: pulumi.Output<outputs.PortalConfigurationBusinessProfile>;
    /**
     * String. The default URL to redirect customers to when they click on the portal’s link to return to your website. This can be overriden when creating the session.
     */
    public readonly defaultReturnUrl!: pulumi.Output<string | undefined>;
    /**
     * List(Resource). Information about the features available in the portal. Feature section described in Feature section
     */
    public readonly features!: pulumi.Output<outputs.PortalConfigurationFeatures>;
    /**
     * Bool. Whether the configuration is the default.
     */
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    /**
     * List(Resource). The hosted login page for this configuration. See details in Login Page Section.
     */
    public readonly loginPage!: pulumi.Output<outputs.PortalConfigurationLoginPage>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a PortalConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortalConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortalConfigurationArgs | PortalConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PortalConfigurationState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["businessProfile"] = state ? state.businessProfile : undefined;
            resourceInputs["defaultReturnUrl"] = state ? state.defaultReturnUrl : undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["loginPage"] = state ? state.loginPage : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
        } else {
            const args = argsOrState as PortalConfigurationArgs | undefined;
            if ((!args || args.businessProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'businessProfile'");
            }
            if ((!args || args.features === undefined) && !opts.urn) {
                throw new Error("Missing required property 'features'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["businessProfile"] = args ? args.businessProfile : undefined;
            resourceInputs["defaultReturnUrl"] = args ? args.defaultReturnUrl : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["loginPage"] = args ? args.loginPage : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["isDefault"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PortalConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PortalConfiguration resources.
 */
export interface PortalConfigurationState {
    /**
     * Bool. Whether the configuration is active and can be used to create portal sessions. (On create it is always set as `true`)
     */
    active?: pulumi.Input<boolean>;
    /**
     * List(Resource). The business information shown to customers in the portal. More details in Business Profile section
     */
    businessProfile?: pulumi.Input<inputs.PortalConfigurationBusinessProfile>;
    /**
     * String. The default URL to redirect customers to when they click on the portal’s link to return to your website. This can be overriden when creating the session.
     */
    defaultReturnUrl?: pulumi.Input<string>;
    /**
     * List(Resource). Information about the features available in the portal. Feature section described in Feature section
     */
    features?: pulumi.Input<inputs.PortalConfigurationFeatures>;
    /**
     * Bool. Whether the configuration is the default.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * List(Resource). The hosted login page for this configuration. See details in Login Page Section.
     */
    loginPage?: pulumi.Input<inputs.PortalConfigurationLoginPage>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a PortalConfiguration resource.
 */
export interface PortalConfigurationArgs {
    /**
     * Bool. Whether the configuration is active and can be used to create portal sessions. (On create it is always set as `true`)
     */
    active?: pulumi.Input<boolean>;
    /**
     * List(Resource). The business information shown to customers in the portal. More details in Business Profile section
     */
    businessProfile: pulumi.Input<inputs.PortalConfigurationBusinessProfile>;
    /**
     * String. The default URL to redirect customers to when they click on the portal’s link to return to your website. This can be overriden when creating the session.
     */
    defaultReturnUrl?: pulumi.Input<string>;
    /**
     * List(Resource). Information about the features available in the portal. Feature section described in Feature section
     */
    features: pulumi.Input<inputs.PortalConfigurationFeatures>;
    /**
     * List(Resource). The hosted login page for this configuration. See details in Login Page Section.
     */
    loginPage?: pulumi.Input<inputs.PortalConfigurationLoginPage>;
    /**
     * Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
