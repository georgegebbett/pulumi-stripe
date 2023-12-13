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
    public sealed class PortalConfigurationLoginPage
    {
        /// <summary>
        /// Bool. Set to true to generate a shareable URL login_page.url that will take your customers to a hosted login page for the customer portal.
        /// </summary>
        public readonly bool? Enabled;
        public readonly string? Url;

        [OutputConstructor]
        private PortalConfigurationLoginPage(
            bool? enabled,

            string? url)
        {
            Enabled = enabled;
            Url = url;
        }
    }
}