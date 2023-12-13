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
    public sealed class ShippingRateFixedAmountCurrencyOption
    {
        /// <summary>
        /// Int. (Required) Int. A non-negative integer in cents representing how much to charge.
        /// </summary>
        public readonly int Amount;
        /// <summary>
        /// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
        /// </summary>
        public readonly string Currency;
        public readonly string? TaxBehavior;

        [OutputConstructor]
        private ShippingRateFixedAmountCurrencyOption(
            int amount,

            string currency,

            string? taxBehavior)
        {
            Amount = amount;
            Currency = currency;
            TaxBehavior = taxBehavior;
        }
    }
}