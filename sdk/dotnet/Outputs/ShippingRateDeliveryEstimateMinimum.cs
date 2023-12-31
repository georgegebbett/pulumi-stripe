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
    public sealed class ShippingRateDeliveryEstimateMinimum
    {
        /// <summary>
        /// String. A unit of time. Possible values `hour`, `day`, `business_day`, `week` and `month`.
        /// </summary>
        public readonly string Unit;
        /// <summary>
        /// Int. Must be greater than 0.
        /// </summary>
        public readonly int Value;

        [OutputConstructor]
        private ShippingRateDeliveryEstimateMinimum(
            string unit,

            int value)
        {
            Unit = unit;
            Value = value;
        }
    }
}
