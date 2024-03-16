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
    /// With this resource, you can create a customer - [Stripe API customer documentation](https://stripe.com/docs/api/customers).
    /// 
    /// Customer objects allow you to perform recurring charges, and to track multiple charges,
    /// that are associated with the same customer.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Stripe = Pulumi.Stripe;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // A customer with other details
    ///     var customer = new Stripe.Customer("customer", new()
    ///     {
    ///         Balance = 10000,
    ///         InvoicePrefix = "LA000",
    ///         InvoiceSettings = 
    ///         {
    ///             { "customFieldName", "customFieldValue" },
    ///             { "footer", "--- Lukas Aron ---" },
    ///         },
    ///         NextInvoiceSequence = 1001,
    ///         PreferredLocales = new[]
    ///         {
    ///             "eng",
    ///             "esp",
    ///         },
    ///         Shipping = 
    ///         {
    ///             { "city", "Sydney" },
    ///             { "country", "AU" },
    ///             { "line1", "1 The Best Street" },
    ///             { "line2", "Apartment 401" },
    ///             { "name", "Lukas Aron" },
    ///             { "phone", "+610123456789" },
    ///             { "postal_code", "2000" },
    ///             { "state", "New South Wales" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [StripeResourceType("stripe:index/customer:Customer")]
    public partial class Customer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Map(String). The customer’s address, for all individual fields see: Address Fields.
        /// </summary>
        [Output("address")]
        public Output<ImmutableDictionary<string, string>?> Address { get; private set; } = null!;

        /// <summary>
        /// Int. Current balance, if any, being stored on the customer. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that will be added to their next invoice. The balance does not refer to any unpaid invoices; it solely takes into account amounts that have yet to be successfully applied to any invoice. This balance is only taken into account as invoices are finalized.
        /// </summary>
        [Output("balance")]
        public Output<int?> Balance { get; private set; } = null!;

        /// <summary>
        /// String. The default invoice prefix generated by Stripe when not individual invoice prefix provided.
        /// </summary>
        [Output("defaultInvoicePrefix")]
        public Output<string> DefaultInvoicePrefix { get; private set; } = null!;

        /// <summary>
        /// String. An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// String. Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching and tracking. This may be up to 512 characters.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// String. The prefix for the customer used to generate unique invoice numbers. Must be `3–12 uppercase letters or numbers`.
        /// </summary>
        [Output("invoicePrefix")]
        public Output<string?> InvoicePrefix { get; private set; } = null!;

        /// <summary>
        /// Map(String). Default invoice settings for this customer. For supported fields see: Invoice Settings Fields.
        /// </summary>
        [Output("invoiceSettings")]
        public Output<ImmutableDictionary<string, string>?> InvoiceSettings { get; private set; } = null!;

        /// <summary>
        /// Map(String). Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// String. The customer’s full name or business name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Int. The sequence to be used on the customer’s next invoice. Defaults to 1.
        /// </summary>
        [Output("nextInvoiceSequence")]
        public Output<int?> NextInvoiceSequence { get; private set; } = null!;

        /// <summary>
        /// String. The customer’s phone number.
        /// </summary>
        [Output("phone")]
        public Output<string?> Phone { get; private set; } = null!;

        /// <summary>
        /// List(String). Customer’s preferred languages, ordered by preference.
        /// </summary>
        [Output("preferredLocales")]
        public Output<ImmutableArray<string>> PreferredLocales { get; private set; } = null!;

        /// <summary>
        /// Map(String). Mailing and shipping address for the customer. Appears on invoices emailed to this customer. For all individual fields see: Shipping Fields.
        /// </summary>
        [Output("shipping")]
        public Output<ImmutableDictionary<string, string>?> Shipping { get; private set; } = null!;


        /// <summary>
        /// Create a Customer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Customer(string name, CustomerArgs? args = null, CustomResourceOptions? options = null)
            : base("stripe:index/customer:Customer", name, args ?? new CustomerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Customer(string name, Input<string> id, CustomerState? state = null, CustomResourceOptions? options = null)
            : base("stripe:index/customer:Customer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Customer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Customer Get(string name, Input<string> id, CustomerState? state = null, CustomResourceOptions? options = null)
        {
            return new Customer(name, id, state, options);
        }
    }

    public sealed class CustomerArgs : global::Pulumi.ResourceArgs
    {
        [Input("address")]
        private InputMap<string>? _address;

        /// <summary>
        /// Map(String). The customer’s address, for all individual fields see: Address Fields.
        /// </summary>
        public InputMap<string> Address
        {
            get => _address ?? (_address = new InputMap<string>());
            set => _address = value;
        }

        /// <summary>
        /// Int. Current balance, if any, being stored on the customer. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that will be added to their next invoice. The balance does not refer to any unpaid invoices; it solely takes into account amounts that have yet to be successfully applied to any invoice. This balance is only taken into account as invoices are finalized.
        /// </summary>
        [Input("balance")]
        public Input<int>? Balance { get; set; }

        /// <summary>
        /// String. An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// String. Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching and tracking. This may be up to 512 characters.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// String. The prefix for the customer used to generate unique invoice numbers. Must be `3–12 uppercase letters or numbers`.
        /// </summary>
        [Input("invoicePrefix")]
        public Input<string>? InvoicePrefix { get; set; }

        [Input("invoiceSettings")]
        private InputMap<string>? _invoiceSettings;

        /// <summary>
        /// Map(String). Default invoice settings for this customer. For supported fields see: Invoice Settings Fields.
        /// </summary>
        public InputMap<string> InvoiceSettings
        {
            get => _invoiceSettings ?? (_invoiceSettings = new InputMap<string>());
            set => _invoiceSettings = value;
        }

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
        /// String. The customer’s full name or business name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Int. The sequence to be used on the customer’s next invoice. Defaults to 1.
        /// </summary>
        [Input("nextInvoiceSequence")]
        public Input<int>? NextInvoiceSequence { get; set; }

        /// <summary>
        /// String. The customer’s phone number.
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        [Input("preferredLocales")]
        private InputList<string>? _preferredLocales;

        /// <summary>
        /// List(String). Customer’s preferred languages, ordered by preference.
        /// </summary>
        public InputList<string> PreferredLocales
        {
            get => _preferredLocales ?? (_preferredLocales = new InputList<string>());
            set => _preferredLocales = value;
        }

        [Input("shipping")]
        private InputMap<string>? _shipping;

        /// <summary>
        /// Map(String). Mailing and shipping address for the customer. Appears on invoices emailed to this customer. For all individual fields see: Shipping Fields.
        /// </summary>
        public InputMap<string> Shipping
        {
            get => _shipping ?? (_shipping = new InputMap<string>());
            set => _shipping = value;
        }

        public CustomerArgs()
        {
        }
        public static new CustomerArgs Empty => new CustomerArgs();
    }

    public sealed class CustomerState : global::Pulumi.ResourceArgs
    {
        [Input("address")]
        private InputMap<string>? _address;

        /// <summary>
        /// Map(String). The customer’s address, for all individual fields see: Address Fields.
        /// </summary>
        public InputMap<string> Address
        {
            get => _address ?? (_address = new InputMap<string>());
            set => _address = value;
        }

        /// <summary>
        /// Int. Current balance, if any, being stored on the customer. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that will be added to their next invoice. The balance does not refer to any unpaid invoices; it solely takes into account amounts that have yet to be successfully applied to any invoice. This balance is only taken into account as invoices are finalized.
        /// </summary>
        [Input("balance")]
        public Input<int>? Balance { get; set; }

        /// <summary>
        /// String. The default invoice prefix generated by Stripe when not individual invoice prefix provided.
        /// </summary>
        [Input("defaultInvoicePrefix")]
        public Input<string>? DefaultInvoicePrefix { get; set; }

        /// <summary>
        /// String. An arbitrary string that you can attach to a customer object. It is displayed alongside the customer in the dashboard.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// String. Customer’s email address. It’s displayed alongside the customer in your dashboard and can be useful for searching and tracking. This may be up to 512 characters.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// String. The prefix for the customer used to generate unique invoice numbers. Must be `3–12 uppercase letters or numbers`.
        /// </summary>
        [Input("invoicePrefix")]
        public Input<string>? InvoicePrefix { get; set; }

        [Input("invoiceSettings")]
        private InputMap<string>? _invoiceSettings;

        /// <summary>
        /// Map(String). Default invoice settings for this customer. For supported fields see: Invoice Settings Fields.
        /// </summary>
        public InputMap<string> InvoiceSettings
        {
            get => _invoiceSettings ?? (_invoiceSettings = new InputMap<string>());
            set => _invoiceSettings = value;
        }

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
        /// String. The customer’s full name or business name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Int. The sequence to be used on the customer’s next invoice. Defaults to 1.
        /// </summary>
        [Input("nextInvoiceSequence")]
        public Input<int>? NextInvoiceSequence { get; set; }

        /// <summary>
        /// String. The customer’s phone number.
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        [Input("preferredLocales")]
        private InputList<string>? _preferredLocales;

        /// <summary>
        /// List(String). Customer’s preferred languages, ordered by preference.
        /// </summary>
        public InputList<string> PreferredLocales
        {
            get => _preferredLocales ?? (_preferredLocales = new InputList<string>());
            set => _preferredLocales = value;
        }

        [Input("shipping")]
        private InputMap<string>? _shipping;

        /// <summary>
        /// Map(String). Mailing and shipping address for the customer. Appears on invoices emailed to this customer. For all individual fields see: Shipping Fields.
        /// </summary>
        public InputMap<string> Shipping
        {
            get => _shipping ?? (_shipping = new InputMap<string>());
            set => _shipping = value;
        }

        public CustomerState()
        {
        }
        public static new CustomerState Empty => new CustomerState();
    }
}
