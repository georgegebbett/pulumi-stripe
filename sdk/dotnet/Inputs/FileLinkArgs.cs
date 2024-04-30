// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Stripe.Inputs
{

    public sealed class FileLinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// String. Time at which the object was created. Measured in seconds since the Unix epoch.
        /// </summary>
        [Input("created")]
        public Input<int>? Created { get; set; }

        /// <summary>
        /// Bool. Returns if the link is already expired.
        /// </summary>
        [Input("expired")]
        public Input<bool>? Expired { get; set; }

        /// <summary>
        /// Int. Time that the link expires.
        /// </summary>
        [Input("expiresAt")]
        public Input<int>? ExpiresAt { get; set; }

        /// <summary>
        /// String. Unique identifier for the object.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Bool. Has the value `true` if the object exists in live mode or the value `false` 
        /// if the object exists in test mode.
        /// </summary>
        [Input("livemode")]
        public Input<bool>? Livemode { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Map(String). Set of key-value pairs that you can attach to an object.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// String. String representing the object’s type. Objects of the same type share the same value.
        /// </summary>
        [Input("object")]
        public Input<string>? Object { get; set; }

        /// <summary>
        /// String. The publicly accessible URL to download the file.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public FileLinkArgs()
        {
        }
        public static new FileLinkArgs Empty => new FileLinkArgs();
    }
}
