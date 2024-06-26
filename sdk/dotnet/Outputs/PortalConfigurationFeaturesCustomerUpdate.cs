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
    public sealed class PortalConfigurationFeaturesCustomerUpdate
    {
        /// <summary>
        /// The types of customer updates that are supported. When empty, customers are not updatable.
        /// </summary>
        public readonly ImmutableArray<string> AllowedUpdates;
        /// <summary>
        /// Whether the feature is enabled.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private PortalConfigurationFeaturesCustomerUpdate(
            ImmutableArray<string> allowedUpdates,

            bool enabled)
        {
            AllowedUpdates = allowedUpdates;
            Enabled = enabled;
        }
    }
}
