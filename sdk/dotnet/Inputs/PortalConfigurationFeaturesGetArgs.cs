// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe.Inputs
{

    public sealed class PortalConfigurationFeaturesGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// List(Resource). Information about updating the customer details in the portal. See Customer Update.
        /// </summary>
        [Input("customerUpdate")]
        public Input<Inputs.PortalConfigurationFeaturesCustomerUpdateGetArgs>? CustomerUpdate { get; set; }

        /// <summary>
        /// List(Resource). Information about showing the billing history in the portal. See Invoice History.
        /// </summary>
        [Input("invoiceHistory")]
        public Input<Inputs.PortalConfigurationFeaturesInvoiceHistoryGetArgs>? InvoiceHistory { get; set; }

        /// <summary>
        /// List(Resource). Information about updating payment methods in the portal. See Payment Method Update.
        /// </summary>
        [Input("paymentMethodUpdate")]
        public Input<Inputs.PortalConfigurationFeaturesPaymentMethodUpdateGetArgs>? PaymentMethodUpdate { get; set; }

        /// <summary>
        /// List(Resource). Information about canceling subscriptions in the portal. See Subscription Cancel.
        /// </summary>
        [Input("subscriptionCancel")]
        public Input<Inputs.PortalConfigurationFeaturesSubscriptionCancelGetArgs>? SubscriptionCancel { get; set; }

        [Input("subscriptionPauses")]
        private InputList<Inputs.PortalConfigurationFeaturesSubscriptionPauseGetArgs>? _subscriptionPauses;

        /// <summary>
        /// List(Resource). Information about pausing subscriptions in the portal. See Subscription Pause.
        /// </summary>
        public InputList<Inputs.PortalConfigurationFeaturesSubscriptionPauseGetArgs> SubscriptionPauses
        {
            get => _subscriptionPauses ?? (_subscriptionPauses = new InputList<Inputs.PortalConfigurationFeaturesSubscriptionPauseGetArgs>());
            set => _subscriptionPauses = value;
        }

        [Input("subscriptionUpdates")]
        private InputList<Inputs.PortalConfigurationFeaturesSubscriptionUpdateGetArgs>? _subscriptionUpdates;

        /// <summary>
        /// List(Resource). Information about updating subscriptions in the portal. See Subscription Update.
        /// </summary>
        public InputList<Inputs.PortalConfigurationFeaturesSubscriptionUpdateGetArgs> SubscriptionUpdates
        {
            get => _subscriptionUpdates ?? (_subscriptionUpdates = new InputList<Inputs.PortalConfigurationFeaturesSubscriptionUpdateGetArgs>());
            set => _subscriptionUpdates = value;
        }

        public PortalConfigurationFeaturesGetArgs()
        {
        }
        public static new PortalConfigurationFeaturesGetArgs Empty => new PortalConfigurationFeaturesGetArgs();
    }
}
