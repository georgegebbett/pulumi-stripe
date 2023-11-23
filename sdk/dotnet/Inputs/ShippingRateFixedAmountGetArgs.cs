// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe.Inputs
{

    public sealed class ShippingRateFixedAmountGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Int. A non-negative integer in cents representing how much to charge.
        /// </summary>
        [Input("amount", required: true)]
        public Input<int> Amount { get; set; } = null!;

        /// <summary>
        /// String. Three-letter ISO currency code, in lowercase - [supported currencies](https://stripe.com/docs/currencies).
        /// </summary>
        [Input("currency", required: true)]
        public Input<string> Currency { get; set; } = null!;

        [Input("currencyOptions")]
        private InputList<Inputs.ShippingRateFixedAmountCurrencyOptionGetArgs>? _currencyOptions;

        /// <summary>
        /// List(Resource). Please see argument details Currency Option
        /// </summary>
        public InputList<Inputs.ShippingRateFixedAmountCurrencyOptionGetArgs> CurrencyOptions
        {
            get => _currencyOptions ?? (_currencyOptions = new InputList<Inputs.ShippingRateFixedAmountCurrencyOptionGetArgs>());
            set => _currencyOptions = value;
        }

        public ShippingRateFixedAmountGetArgs()
        {
        }
        public static new ShippingRateFixedAmountGetArgs Empty => new ShippingRateFixedAmountGetArgs();
    }
}
