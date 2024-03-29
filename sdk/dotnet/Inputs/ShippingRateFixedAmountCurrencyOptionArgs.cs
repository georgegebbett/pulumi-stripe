// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe.Inputs
{

    public sealed class ShippingRateFixedAmountCurrencyOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Int. (Required) Int. A non-negative integer in cents representing how much to charge.
        /// </summary>
        [Input("amount", required: true)]
        public Input<int> Amount { get; set; } = null!;

        /// <summary>
        /// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
        /// </summary>
        [Input("currency", required: true)]
        public Input<string> Currency { get; set; } = null!;

        /// <summary>
        /// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of inclusive, exclusive, or unspecified.
        /// </summary>
        [Input("taxBehavior")]
        public Input<string>? TaxBehavior { get; set; }

        public ShippingRateFixedAmountCurrencyOptionArgs()
        {
        }
        public static new ShippingRateFixedAmountCurrencyOptionArgs Empty => new ShippingRateFixedAmountCurrencyOptionArgs();
    }
}
