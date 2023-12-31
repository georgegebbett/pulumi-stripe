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
    public sealed class PriceCurrencyOption
    {
        /// <summary>
        /// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
        /// </summary>
        public readonly string Currency;
        /// <summary>
        /// List(Resource). When set, 
        /// provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
        /// See details in custom unit amount.
        /// </summary>
        public readonly Outputs.PriceCurrencyOptionCustomUnitAmount? CustomUnitAmount;
        /// <summary>
        /// String. Only required if a default tax behavior was not provided in the Stripe Tax settings. 
        /// Specifies whether the price is considered inclusive of taxes or exclusive of taxes.
        /// One of `inclusive`, `exclusive`, or `unspecified`.
        /// Once specified as either inclusive or exclusive, it cannot be changed.
        /// </summary>
        public readonly string? TaxBehavior;
        /// <summary>
        /// List(Resource). Each element represents a pricing tier. 
        /// This parameter requires `billing_scheme` to be set to `tiered`. This resource can be used more than once and follows
        /// the same fields as the root tiers block
        /// </summary>
        public readonly ImmutableArray<Outputs.PriceCurrencyOptionTier> Tiers;
        /// <summary>
        /// Int. A positive integer in cents (or -1 for a free price) representing how much to charge.
        /// </summary>
        public readonly int? UnitAmount;
        /// <summary>
        /// Float. Same as unit_amount, but accepts a decimal value in cents with at most 12
        /// decimal places. Only one of unit_amount and unit_amount_decimal can be set.
        /// </summary>
        public readonly double? UnitAmountDecimal;

        [OutputConstructor]
        private PriceCurrencyOption(
            string currency,

            Outputs.PriceCurrencyOptionCustomUnitAmount? customUnitAmount,

            string? taxBehavior,

            ImmutableArray<Outputs.PriceCurrencyOptionTier> tiers,

            int? unitAmount,

            double? unitAmountDecimal)
        {
            Currency = currency;
            CustomUnitAmount = customUnitAmount;
            TaxBehavior = taxBehavior;
            Tiers = tiers;
            UnitAmount = unitAmount;
            UnitAmountDecimal = unitAmountDecimal;
        }
    }
}
