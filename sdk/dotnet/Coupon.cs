// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe
{
    /// <summary>
    /// With this resource, you can create a coupon - [Stripe API coupon documentation](https://stripe.com/docs/api/coupons).
    /// 
    /// A coupon contains information about a percent-off or amount-off discount you might want to apply to a customer.
    /// 
    /// A coupon has either a `percent_off` or an `amount_off` and `currency`. If you set an `amount_off`, that amount will be subtracted from any invoice’s subtotal.
    /// 
    /// For example, an invoice with a subtotal of $100 will have a final total of $0 if a coupon with an amount_off of 20000 is applied to it and an invoice with a subtotal of $300 will have a final total of $100 if a coupon with an amount_off of 20000 is applied to it.
    /// </summary>
    [StripeResourceType("stripe:index/coupon:Coupon")]
    public partial class Coupon : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
        /// </summary>
        [Output("amountOff")]
        public Output<int?> AmountOff { get; private set; } = null!;

        /// <summary>
        /// List(String). A list of product IDs this coupon applies to.
        /// </summary>
        [Output("appliesTos")]
        public Output<ImmutableArray<string>> AppliesTos { get; private set; } = null!;

        /// <summary>
        /// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
        /// </summary>
        [Output("couponId")]
        public Output<string> CouponId { get; private set; } = null!;

        /// <summary>
        /// String. Required if `amount_off` has been set, the three-letter ISO code for the currency of the amount to take off.
        /// </summary>
        [Output("currency")]
        public Output<string?> Currency { get; private set; } = null!;

        /// <summary>
        /// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
        /// </summary>
        [Output("duration")]
        public Output<string?> Duration { get; private set; } = null!;

        /// <summary>
        /// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
        /// </summary>
        [Output("durationInMonths")]
        public Output<int?> DurationInMonths { get; private set; } = null!;

        /// <summary>
        /// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
        /// </summary>
        [Output("maxRedemptions")]
        public Output<int?> MaxRedemptions { get; private set; } = null!;

        /// <summary>
        /// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// String. Name of the coupon displayed to customers on for instance invoices or receipts.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a $100 invoice $50 instead.
        /// </summary>
        [Output("percentOff")]
        public Output<double?> PercentOff { get; private set; } = null!;

        /// <summary>
        /// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
        /// </summary>
        [Output("redeemBy")]
        public Output<string?> RedeemBy { get; private set; } = null!;

        /// <summary>
        /// Int. Number of times this coupon has been applied to a customer.
        /// </summary>
        [Output("timesRedeemed")]
        public Output<int> TimesRedeemed { get; private set; } = null!;

        /// <summary>
        /// Bool. Taking account of the above properties, whether this coupon can still be applied to a customer.
        /// </summary>
        [Output("valid")]
        public Output<bool> Valid { get; private set; } = null!;


        /// <summary>
        /// Create a Coupon resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Coupon(string name, CouponArgs? args = null, CustomResourceOptions? options = null)
            : base("stripe:index/coupon:Coupon", name, args ?? new CouponArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Coupon(string name, Input<string> id, CouponState? state = null, CustomResourceOptions? options = null)
            : base("stripe:index/coupon:Coupon", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/georgegebbett",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Coupon resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Coupon Get(string name, Input<string> id, CouponState? state = null, CustomResourceOptions? options = null)
        {
            return new Coupon(name, id, state, options);
        }
    }

    public sealed class CouponArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
        /// </summary>
        [Input("amountOff")]
        public Input<int>? AmountOff { get; set; }

        [Input("appliesTos")]
        private InputList<string>? _appliesTos;

        /// <summary>
        /// List(String). A list of product IDs this coupon applies to.
        /// </summary>
        public InputList<string> AppliesTos
        {
            get => _appliesTos ?? (_appliesTos = new InputList<string>());
            set => _appliesTos = value;
        }

        /// <summary>
        /// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
        /// </summary>
        [Input("couponId")]
        public Input<string>? CouponId { get; set; }

        /// <summary>
        /// String. Required if `amount_off` has been set, the three-letter ISO code for the currency of the amount to take off.
        /// </summary>
        [Input("currency")]
        public Input<string>? Currency { get; set; }

        /// <summary>
        /// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
        /// </summary>
        [Input("durationInMonths")]
        public Input<int>? DurationInMonths { get; set; }

        /// <summary>
        /// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
        /// </summary>
        [Input("maxRedemptions")]
        public Input<int>? MaxRedemptions { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// String. Name of the coupon displayed to customers on for instance invoices or receipts.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a $100 invoice $50 instead.
        /// </summary>
        [Input("percentOff")]
        public Input<double>? PercentOff { get; set; }

        /// <summary>
        /// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
        /// </summary>
        [Input("redeemBy")]
        public Input<string>? RedeemBy { get; set; }

        public CouponArgs()
        {
        }
        public static new CouponArgs Empty => new CouponArgs();
    }

    public sealed class CouponState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Int. Amount (in the currency specified) that will be taken off the subtotal of any invoices for this customer.
        /// </summary>
        [Input("amountOff")]
        public Input<int>? AmountOff { get; set; }

        [Input("appliesTos")]
        private InputList<string>? _appliesTos;

        /// <summary>
        /// List(String). A list of product IDs this coupon applies to.
        /// </summary>
        public InputList<string> AppliesTos
        {
            get => _appliesTos ?? (_appliesTos = new InputList<string>());
            set => _appliesTos = value;
        }

        /// <summary>
        /// String. Unique string of your choice that will be used to identify this coupon when applying it to a customer.
        /// </summary>
        [Input("couponId")]
        public Input<string>? CouponId { get; set; }

        /// <summary>
        /// String. Required if `amount_off` has been set, the three-letter ISO code for the currency of the amount to take off.
        /// </summary>
        [Input("currency")]
        public Input<string>? Currency { get; set; }

        /// <summary>
        /// String. Describes how long a customer who applies this coupon will get the discount. One of `forever`, `once`, and `repeating`.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// If duration is repeating, the number of months the coupon applies. Null if coupon duration is forever or once.
        /// </summary>
        [Input("durationInMonths")]
        public Input<int>? DurationInMonths { get; set; }

        /// <summary>
        /// Int. Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.
        /// </summary>
        [Input("maxRedemptions")]
        public Input<int>? MaxRedemptions { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// String. Name of the coupon displayed to customers on for instance invoices or receipts.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Float. Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a $100 invoice $50 instead.
        /// </summary>
        [Input("percentOff")]
        public Input<double>? PercentOff { get; set; }

        /// <summary>
        /// String. Date after which the coupon can no longer be redeemed. Expected format is in the `RFC3339`.
        /// </summary>
        [Input("redeemBy")]
        public Input<string>? RedeemBy { get; set; }

        /// <summary>
        /// Int. Number of times this coupon has been applied to a customer.
        /// </summary>
        [Input("timesRedeemed")]
        public Input<int>? TimesRedeemed { get; set; }

        /// <summary>
        /// Bool. Taking account of the above properties, whether this coupon can still be applied to a customer.
        /// </summary>
        [Input("valid")]
        public Input<bool>? Valid { get; set; }

        public CouponState()
        {
        }
        public static new CouponState Empty => new CouponState();
    }
}