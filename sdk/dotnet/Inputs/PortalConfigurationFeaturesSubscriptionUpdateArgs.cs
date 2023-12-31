// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe.Inputs
{

    public sealed class PortalConfigurationFeaturesSubscriptionUpdateArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultAllowedUpdates", required: true)]
        private InputList<string>? _defaultAllowedUpdates;

        /// <summary>
        /// List(String). The types of subscription updates that are supported. When empty, subscriptions are not updatable. Supported values are `price`, `quantity`, and `promotion_code`.
        /// </summary>
        public InputList<string> DefaultAllowedUpdates
        {
            get => _defaultAllowedUpdates ?? (_defaultAllowedUpdates = new InputList<string>());
            set => _defaultAllowedUpdates = value;
        }

        /// <summary>
        /// Bool. Whether the feature is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("products", required: true)]
        private InputList<Inputs.PortalConfigurationFeaturesSubscriptionUpdateProductArgs>? _products;

        /// <summary>
        /// List(Resource). The list of products that support subscription updates. See details Products.
        /// </summary>
        public InputList<Inputs.PortalConfigurationFeaturesSubscriptionUpdateProductArgs> Products
        {
            get => _products ?? (_products = new InputList<Inputs.PortalConfigurationFeaturesSubscriptionUpdateProductArgs>());
            set => _products = value;
        }

        /// <summary>
        /// String. Whether to create prorations when canceling subscriptions. Possible values are `none` and `create_prorations`, which is only compatible with `mode=immediately`. No prorations are generated when canceling a subscription at the end of its natural billing period.
        /// </summary>
        [Input("prorationBehavior")]
        public Input<string>? ProrationBehavior { get; set; }

        public PortalConfigurationFeaturesSubscriptionUpdateArgs()
        {
        }
        public static new PortalConfigurationFeaturesSubscriptionUpdateArgs Empty => new PortalConfigurationFeaturesSubscriptionUpdateArgs();
    }
}
