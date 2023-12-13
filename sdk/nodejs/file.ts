// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This object represents files hosted on Stripe's servers - [Stripe API file documentation](https://stripe.com/docs/api/files).
 * You can upload files with the create file request (for example, when uploading dispute evidence).
 *
 * Stripe File upload [guide](https://stripe.com/docs/file-upload#uploading-a-file)
 *
 * > Update or removal of the file isn't supported through the Stripe SDK.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as stripe from "pulumi-stripe";
 *
 * // minimal file field
 * const logoFile = new stripe.File("logoFile", {
 *     filename: "logo.jpg",
 *     purpose: "business_logo",
 *     base64content: Buffer.from(fs.readFileSync(`${HOME}/logo.jpg`), 'binary').toString('base64'),
 * });
 * // file with file link
 * const logoIndex_fileFile = new stripe.File("logoIndex/fileFile", {
 *     filename: "logo.jpg",
 *     purpose: "business_logo",
 *     base64content: Buffer.from(fs.readFileSync(`${HOME}/logo.jpg`), 'binary').toString('base64'),
 *     linkData: {
 *         create: true,
 *         expiresAt: 1826659124,
 *     },
 * });
 * ```
 */
export class File extends pulumi.CustomResource {
    /**
     * Get an existing File resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileState, opts?: pulumi.CustomResourceOptions): File {
        return new File(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'stripe:index/file:File';

    /**
     * Returns true if the given object is an instance of File.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is File {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === File.__pulumiType;
    }

    /**
     * A content file to upload encoded by base64.
     */
    public readonly base64content!: pulumi.Output<string>;
    /**
     * String. Time at which the object was created. Measured in seconds since the Unix epoch.
     */
    public /*out*/ readonly created!: pulumi.Output<number>;
    /**
     * Int. The link isn’t available after this future timestamp.
     */
    public /*out*/ readonly expiresAt!: pulumi.Output<number>;
    /**
     * String. The suitable name for saving the file to a filesystem.
     */
    public readonly filename!: pulumi.Output<string>;
    /**
     * Optional parameters that automatically create a file link for the newly created file.
     */
    public readonly linkData!: pulumi.Output<outputs.FileLinkData | undefined>;
    /**
     * List(Resource). A list of [file links](https://stripe.com/docs/api/files/object#file_links) that point at this file.
     * Please see details of links.
     */
    public /*out*/ readonly links!: pulumi.Output<outputs.FileLink[]>;
    /**
     * String. String representing the object’s type. Objects of the same type share the same value.
     */
    public /*out*/ readonly object!: pulumi.Output<string>;
    /**
     * String. The purpose of the uploaded file. One of these values are accepted: `accountRequirement`,
     * `additionalVerification`, `businessIcon`, `businessLogo`, `customerSignature`, `disputeEvidence`,
     * `documentProviderIdentityDocument`, `financeReportRun`, `identityDocument`, `identityDocumentDownloadable`,
     * `pciDocument`, `selfie`, `sigmaScheduledQuery`, `taxDocumentUserUpload`, `terminalReaderSplashscreen`
     */
    public readonly purpose!: pulumi.Output<string>;
    /**
     * Int. The size of the file object in bytes.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * String. The returned file type (for example, `csv`, `pdf`, `jpg`, or `png`).
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * String. The publicly accessible URL to download the file.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a File resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileArgs | FileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FileState | undefined;
            resourceInputs["base64content"] = state ? state.base64content : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["expiresAt"] = state ? state.expiresAt : undefined;
            resourceInputs["filename"] = state ? state.filename : undefined;
            resourceInputs["linkData"] = state ? state.linkData : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["object"] = state ? state.object : undefined;
            resourceInputs["purpose"] = state ? state.purpose : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as FileArgs | undefined;
            if ((!args || args.base64content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'base64content'");
            }
            if ((!args || args.filename === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filename'");
            }
            if ((!args || args.purpose === undefined) && !opts.urn) {
                throw new Error("Missing required property 'purpose'");
            }
            resourceInputs["base64content"] = args ? args.base64content : undefined;
            resourceInputs["filename"] = args ? args.filename : undefined;
            resourceInputs["linkData"] = args ? args.linkData : undefined;
            resourceInputs["purpose"] = args ? args.purpose : undefined;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["links"] = undefined /*out*/;
            resourceInputs["object"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(File.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering File resources.
 */
export interface FileState {
    /**
     * A content file to upload encoded by base64.
     */
    base64content?: pulumi.Input<string>;
    /**
     * String. Time at which the object was created. Measured in seconds since the Unix epoch.
     */
    created?: pulumi.Input<number>;
    /**
     * Int. The link isn’t available after this future timestamp.
     */
    expiresAt?: pulumi.Input<number>;
    /**
     * String. The suitable name for saving the file to a filesystem.
     */
    filename?: pulumi.Input<string>;
    /**
     * Optional parameters that automatically create a file link for the newly created file.
     */
    linkData?: pulumi.Input<inputs.FileLinkData>;
    /**
     * List(Resource). A list of [file links](https://stripe.com/docs/api/files/object#file_links) that point at this file.
     * Please see details of links.
     */
    links?: pulumi.Input<pulumi.Input<inputs.FileLink>[]>;
    /**
     * String. String representing the object’s type. Objects of the same type share the same value.
     */
    object?: pulumi.Input<string>;
    /**
     * String. The purpose of the uploaded file. One of these values are accepted: `accountRequirement`,
     * `additionalVerification`, `businessIcon`, `businessLogo`, `customerSignature`, `disputeEvidence`,
     * `documentProviderIdentityDocument`, `financeReportRun`, `identityDocument`, `identityDocumentDownloadable`,
     * `pciDocument`, `selfie`, `sigmaScheduledQuery`, `taxDocumentUserUpload`, `terminalReaderSplashscreen`
     */
    purpose?: pulumi.Input<string>;
    /**
     * Int. The size of the file object in bytes.
     */
    size?: pulumi.Input<number>;
    /**
     * String. The returned file type (for example, `csv`, `pdf`, `jpg`, or `png`).
     */
    type?: pulumi.Input<string>;
    /**
     * String. The publicly accessible URL to download the file.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a File resource.
 */
export interface FileArgs {
    /**
     * A content file to upload encoded by base64.
     */
    base64content: pulumi.Input<string>;
    /**
     * String. The suitable name for saving the file to a filesystem.
     */
    filename: pulumi.Input<string>;
    /**
     * Optional parameters that automatically create a file link for the newly created file.
     */
    linkData?: pulumi.Input<inputs.FileLinkData>;
    /**
     * String. The purpose of the uploaded file. One of these values are accepted: `accountRequirement`,
     * `additionalVerification`, `businessIcon`, `businessLogo`, `customerSignature`, `disputeEvidence`,
     * `documentProviderIdentityDocument`, `financeReportRun`, `identityDocument`, `identityDocumentDownloadable`,
     * `pciDocument`, `selfie`, `sigmaScheduledQuery`, `taxDocumentUserUpload`, `terminalReaderSplashscreen`
     */
    purpose: pulumi.Input<string>;
}