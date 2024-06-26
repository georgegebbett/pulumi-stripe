// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe.Outputs
{

    [OutputType]
    public sealed class PortalConfigurationFeaturesSubscriptionUpdate
    {
        /// <summary>
        /// The types of subscription updates that are supported. When empty, subscriptions are not updateable.
        /// </summary>
        public readonly ImmutableArray<string> DefaultAllowedUpdates;
        /// <summary>
        /// Whether the feature is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The list of products that support subscription updates.
        /// </summary>
        public readonly ImmutableArray<Outputs.PortalConfigurationFeaturesSubscriptionUpdateProduct> Products;
        /// <summary>
        /// Determines how to handle prorations resulting from subscription updates
        /// </summary>
        public readonly string? ProrationBehavior;

        [OutputConstructor]
        private PortalConfigurationFeaturesSubscriptionUpdate(
            ImmutableArray<string> defaultAllowedUpdates,

            bool enabled,

            ImmutableArray<Outputs.PortalConfigurationFeaturesSubscriptionUpdateProduct> products,

            string? prorationBehavior)
        {
            DefaultAllowedUpdates = defaultAllowedUpdates;
            Enabled = enabled;
            Products = products;
            ProrationBehavior = prorationBehavior;
        }
    }
}
