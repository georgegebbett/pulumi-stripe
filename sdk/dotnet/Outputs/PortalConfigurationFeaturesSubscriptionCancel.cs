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
    public sealed class PortalConfigurationFeaturesSubscriptionCancel
    {
        /// <summary>
        /// Whether the cancellation reasons will be collected in the portal and which options are exposed to the customer
        /// </summary>
        public readonly Outputs.PortalConfigurationFeaturesSubscriptionCancelCancellationReason? CancellationReason;
        /// <summary>
        /// Whether the feature is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Whether to cancel subscriptions immediately or at the end of the billing period.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// Whether to create prorations when canceling subscriptions.
        /// </summary>
        public readonly string? ProrationBehavior;

        [OutputConstructor]
        private PortalConfigurationFeaturesSubscriptionCancel(
            Outputs.PortalConfigurationFeaturesSubscriptionCancelCancellationReason? cancellationReason,

            bool enabled,

            string? mode,

            string? prorationBehavior)
        {
            CancellationReason = cancellationReason;
            Enabled = enabled;
            Mode = mode;
            ProrationBehavior = prorationBehavior;
        }
    }
}
