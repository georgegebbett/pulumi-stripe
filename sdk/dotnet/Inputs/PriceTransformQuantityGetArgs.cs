// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe.Inputs
{

    public sealed class PriceTransformQuantityGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Int. Divide usage by this number.
        /// </summary>
        [Input("divideBy", required: true)]
        public Input<int> DivideBy { get; set; } = null!;

        /// <summary>
        /// String. After division, either round the result `up` or `down`.
        /// </summary>
        [Input("round", required: true)]
        public Input<string> Round { get; set; } = null!;

        public PriceTransformQuantityGetArgs()
        {
        }
        public static new PriceTransformQuantityGetArgs Empty => new PriceTransformQuantityGetArgs();
    }
}