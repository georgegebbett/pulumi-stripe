# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FileArgs', 'File']

@pulumi.input_type
class FileArgs:
    def __init__(__self__, *,
                 base64content: pulumi.Input[str],
                 filename: pulumi.Input[str],
                 purpose: pulumi.Input[str],
                 link_data: Optional[pulumi.Input['FileLinkDataArgs']] = None):
        """
        The set of arguments for constructing a File resource.
        :param pulumi.Input[str] base64content: A content file to upload encoded by base64.
        :param pulumi.Input[str] filename: String. The suitable name for saving the file to a filesystem.
        :param pulumi.Input[str] purpose: String. The purpose of the uploaded file. One of these values are accepted: `account_requirement`,
               `additional_verification`, `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`,
               `document_provider_identity_document`, `finance_report_run`, `identity_document`, `identity_document_downloadable`,
               `pci_document`, `selfie`, `sigma_scheduled_query`, `tax_document_user_upload`, `terminal_reader_splashscreen`
        :param pulumi.Input['FileLinkDataArgs'] link_data: Optional parameters that automatically create a file link for the newly created file.
        """
        pulumi.set(__self__, "base64content", base64content)
        pulumi.set(__self__, "filename", filename)
        pulumi.set(__self__, "purpose", purpose)
        if link_data is not None:
            pulumi.set(__self__, "link_data", link_data)

    @property
    @pulumi.getter
    def base64content(self) -> pulumi.Input[str]:
        """
        A content file to upload encoded by base64.
        """
        return pulumi.get(self, "base64content")

    @base64content.setter
    def base64content(self, value: pulumi.Input[str]):
        pulumi.set(self, "base64content", value)

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Input[str]:
        """
        String. The suitable name for saving the file to a filesystem.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: pulumi.Input[str]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter
    def purpose(self) -> pulumi.Input[str]:
        """
        String. The purpose of the uploaded file. One of these values are accepted: `account_requirement`,
        `additional_verification`, `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`,
        `document_provider_identity_document`, `finance_report_run`, `identity_document`, `identity_document_downloadable`,
        `pci_document`, `selfie`, `sigma_scheduled_query`, `tax_document_user_upload`, `terminal_reader_splashscreen`
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: pulumi.Input[str]):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter(name="linkData")
    def link_data(self) -> Optional[pulumi.Input['FileLinkDataArgs']]:
        """
        Optional parameters that automatically create a file link for the newly created file.
        """
        return pulumi.get(self, "link_data")

    @link_data.setter
    def link_data(self, value: Optional[pulumi.Input['FileLinkDataArgs']]):
        pulumi.set(self, "link_data", value)


@pulumi.input_type
class _FileState:
    def __init__(__self__, *,
                 base64content: Optional[pulumi.Input[str]] = None,
                 created: Optional[pulumi.Input[int]] = None,
                 expires_at: Optional[pulumi.Input[int]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 link_data: Optional[pulumi.Input['FileLinkDataArgs']] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input['FileLinkArgs']]]] = None,
                 object: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering File resources.
        :param pulumi.Input[str] base64content: A content file to upload encoded by base64.
        :param pulumi.Input[int] created: String. Time at which the object was created. Measured in seconds since the Unix epoch.
        :param pulumi.Input[int] expires_at: Int. Time that the link expires.
        :param pulumi.Input[str] filename: String. The suitable name for saving the file to a filesystem.
        :param pulumi.Input['FileLinkDataArgs'] link_data: Optional parameters that automatically create a file link for the newly created file.
        :param pulumi.Input[Sequence[pulumi.Input['FileLinkArgs']]] links: List(Resource). A list of [file links](https://stripe.com/docs/api/files/object#file_links) that point at this file.
               Please see details of links.
        :param pulumi.Input[str] object: String. String representing the object’s type. Objects of the same type share the same value.
        :param pulumi.Input[str] purpose: String. The purpose of the uploaded file. One of these values are accepted: `account_requirement`,
               `additional_verification`, `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`,
               `document_provider_identity_document`, `finance_report_run`, `identity_document`, `identity_document_downloadable`,
               `pci_document`, `selfie`, `sigma_scheduled_query`, `tax_document_user_upload`, `terminal_reader_splashscreen`
        :param pulumi.Input[int] size: Int. The size of the file object in bytes.
        :param pulumi.Input[str] type: String. The returned file type (for example, `csv`, `pdf`, `jpg`, or `png`).
        :param pulumi.Input[str] url: String. The publicly accessible URL to download the file.
        """
        if base64content is not None:
            pulumi.set(__self__, "base64content", base64content)
        if created is not None:
            pulumi.set(__self__, "created", created)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if filename is not None:
            pulumi.set(__self__, "filename", filename)
        if link_data is not None:
            pulumi.set(__self__, "link_data", link_data)
        if links is not None:
            pulumi.set(__self__, "links", links)
        if object is not None:
            pulumi.set(__self__, "object", object)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def base64content(self) -> Optional[pulumi.Input[str]]:
        """
        A content file to upload encoded by base64.
        """
        return pulumi.get(self, "base64content")

    @base64content.setter
    def base64content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base64content", value)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[int]]:
        """
        String. Time at which the object was created. Measured in seconds since the Unix epoch.
        """
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[int]]:
        """
        Int. Time that the link expires.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def filename(self) -> Optional[pulumi.Input[str]]:
        """
        String. The suitable name for saving the file to a filesystem.
        """
        return pulumi.get(self, "filename")

    @filename.setter
    def filename(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filename", value)

    @property
    @pulumi.getter(name="linkData")
    def link_data(self) -> Optional[pulumi.Input['FileLinkDataArgs']]:
        """
        Optional parameters that automatically create a file link for the newly created file.
        """
        return pulumi.get(self, "link_data")

    @link_data.setter
    def link_data(self, value: Optional[pulumi.Input['FileLinkDataArgs']]):
        pulumi.set(self, "link_data", value)

    @property
    @pulumi.getter
    def links(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FileLinkArgs']]]]:
        """
        List(Resource). A list of [file links](https://stripe.com/docs/api/files/object#file_links) that point at this file.
        Please see details of links.
        """
        return pulumi.get(self, "links")

    @links.setter
    def links(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FileLinkArgs']]]]):
        pulumi.set(self, "links", value)

    @property
    @pulumi.getter
    def object(self) -> Optional[pulumi.Input[str]]:
        """
        String. String representing the object’s type. Objects of the same type share the same value.
        """
        return pulumi.get(self, "object")

    @object.setter
    def object(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object", value)

    @property
    @pulumi.getter
    def purpose(self) -> Optional[pulumi.Input[str]]:
        """
        String. The purpose of the uploaded file. One of these values are accepted: `account_requirement`,
        `additional_verification`, `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`,
        `document_provider_identity_document`, `finance_report_run`, `identity_document`, `identity_document_downloadable`,
        `pci_document`, `selfie`, `sigma_scheduled_query`, `tax_document_user_upload`, `terminal_reader_splashscreen`
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Int. The size of the file object in bytes.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        String. The returned file type (for example, `csv`, `pdf`, `jpg`, or `png`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        String. The publicly accessible URL to download the file.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class File(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base64content: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 link_data: Optional[pulumi.Input[pulumi.InputType['FileLinkDataArgs']]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This object represents files hosted on Stripe's servers - [Stripe API file documentation](https://stripe.com/docs/api/files).
        You can upload files with the create file request (for example, when uploading dispute evidence).

        Stripe File upload [guide](https://stripe.com/docs/file-upload#uploading-a-file)

        > Update or removal of the file isn't supported through the Stripe SDK.

        ## Example Usage

        ```python
        import pulumi
        import base64
        import pulumi_stripe as stripe

        # minimal file field
        logo_file = stripe.File("logoFile",
            filename="logo.jpg",
            purpose="business_logo",
            base64content=(lambda path: base64.b64encode(open(path).read().encode()).decode())(f"{home}/logo.jpg"))
        # file with file link
        logo_index_file_file = stripe.File("logoIndex/fileFile",
            filename="logo.jpg",
            purpose="business_logo",
            base64content=(lambda path: base64.b64encode(open(path).read().encode()).decode())(f"{home}/logo.jpg"),
            link_data=stripe.FileLinkDataArgs(
                create=True,
                expires_at=1826659124,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] base64content: A content file to upload encoded by base64.
        :param pulumi.Input[str] filename: String. The suitable name for saving the file to a filesystem.
        :param pulumi.Input[pulumi.InputType['FileLinkDataArgs']] link_data: Optional parameters that automatically create a file link for the newly created file.
        :param pulumi.Input[str] purpose: String. The purpose of the uploaded file. One of these values are accepted: `account_requirement`,
               `additional_verification`, `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`,
               `document_provider_identity_document`, `finance_report_run`, `identity_document`, `identity_document_downloadable`,
               `pci_document`, `selfie`, `sigma_scheduled_query`, `tax_document_user_upload`, `terminal_reader_splashscreen`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This object represents files hosted on Stripe's servers - [Stripe API file documentation](https://stripe.com/docs/api/files).
        You can upload files with the create file request (for example, when uploading dispute evidence).

        Stripe File upload [guide](https://stripe.com/docs/file-upload#uploading-a-file)

        > Update or removal of the file isn't supported through the Stripe SDK.

        ## Example Usage

        ```python
        import pulumi
        import base64
        import pulumi_stripe as stripe

        # minimal file field
        logo_file = stripe.File("logoFile",
            filename="logo.jpg",
            purpose="business_logo",
            base64content=(lambda path: base64.b64encode(open(path).read().encode()).decode())(f"{home}/logo.jpg"))
        # file with file link
        logo_index_file_file = stripe.File("logoIndex/fileFile",
            filename="logo.jpg",
            purpose="business_logo",
            base64content=(lambda path: base64.b64encode(open(path).read().encode()).decode())(f"{home}/logo.jpg"),
            link_data=stripe.FileLinkDataArgs(
                create=True,
                expires_at=1826659124,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param FileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 base64content: Optional[pulumi.Input[str]] = None,
                 filename: Optional[pulumi.Input[str]] = None,
                 link_data: Optional[pulumi.Input[pulumi.InputType['FileLinkDataArgs']]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FileArgs.__new__(FileArgs)

            if base64content is None and not opts.urn:
                raise TypeError("Missing required property 'base64content'")
            __props__.__dict__["base64content"] = base64content
            if filename is None and not opts.urn:
                raise TypeError("Missing required property 'filename'")
            __props__.__dict__["filename"] = filename
            __props__.__dict__["link_data"] = link_data
            if purpose is None and not opts.urn:
                raise TypeError("Missing required property 'purpose'")
            __props__.__dict__["purpose"] = purpose
            __props__.__dict__["created"] = None
            __props__.__dict__["expires_at"] = None
            __props__.__dict__["links"] = None
            __props__.__dict__["object"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["url"] = None
        super(File, __self__).__init__(
            'stripe:index/file:File',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            base64content: Optional[pulumi.Input[str]] = None,
            created: Optional[pulumi.Input[int]] = None,
            expires_at: Optional[pulumi.Input[int]] = None,
            filename: Optional[pulumi.Input[str]] = None,
            link_data: Optional[pulumi.Input[pulumi.InputType['FileLinkDataArgs']]] = None,
            links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileLinkArgs']]]]] = None,
            object: Optional[pulumi.Input[str]] = None,
            purpose: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'File':
        """
        Get an existing File resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] base64content: A content file to upload encoded by base64.
        :param pulumi.Input[int] created: String. Time at which the object was created. Measured in seconds since the Unix epoch.
        :param pulumi.Input[int] expires_at: Int. Time that the link expires.
        :param pulumi.Input[str] filename: String. The suitable name for saving the file to a filesystem.
        :param pulumi.Input[pulumi.InputType['FileLinkDataArgs']] link_data: Optional parameters that automatically create a file link for the newly created file.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileLinkArgs']]]] links: List(Resource). A list of [file links](https://stripe.com/docs/api/files/object#file_links) that point at this file.
               Please see details of links.
        :param pulumi.Input[str] object: String. String representing the object’s type. Objects of the same type share the same value.
        :param pulumi.Input[str] purpose: String. The purpose of the uploaded file. One of these values are accepted: `account_requirement`,
               `additional_verification`, `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`,
               `document_provider_identity_document`, `finance_report_run`, `identity_document`, `identity_document_downloadable`,
               `pci_document`, `selfie`, `sigma_scheduled_query`, `tax_document_user_upload`, `terminal_reader_splashscreen`
        :param pulumi.Input[int] size: Int. The size of the file object in bytes.
        :param pulumi.Input[str] type: String. The returned file type (for example, `csv`, `pdf`, `jpg`, or `png`).
        :param pulumi.Input[str] url: String. The publicly accessible URL to download the file.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FileState.__new__(_FileState)

        __props__.__dict__["base64content"] = base64content
        __props__.__dict__["created"] = created
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["filename"] = filename
        __props__.__dict__["link_data"] = link_data
        __props__.__dict__["links"] = links
        __props__.__dict__["object"] = object
        __props__.__dict__["purpose"] = purpose
        __props__.__dict__["size"] = size
        __props__.__dict__["type"] = type
        __props__.__dict__["url"] = url
        return File(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def base64content(self) -> pulumi.Output[str]:
        """
        A content file to upload encoded by base64.
        """
        return pulumi.get(self, "base64content")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[int]:
        """
        String. Time at which the object was created. Measured in seconds since the Unix epoch.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[int]:
        """
        Int. Time that the link expires.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def filename(self) -> pulumi.Output[str]:
        """
        String. The suitable name for saving the file to a filesystem.
        """
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter(name="linkData")
    def link_data(self) -> pulumi.Output[Optional['outputs.FileLinkData']]:
        """
        Optional parameters that automatically create a file link for the newly created file.
        """
        return pulumi.get(self, "link_data")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Sequence['outputs.FileLink']]:
        """
        List(Resource). A list of [file links](https://stripe.com/docs/api/files/object#file_links) that point at this file.
        Please see details of links.
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def object(self) -> pulumi.Output[str]:
        """
        String. String representing the object’s type. Objects of the same type share the same value.
        """
        return pulumi.get(self, "object")

    @property
    @pulumi.getter
    def purpose(self) -> pulumi.Output[str]:
        """
        String. The purpose of the uploaded file. One of these values are accepted: `account_requirement`,
        `additional_verification`, `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`,
        `document_provider_identity_document`, `finance_report_run`, `identity_document`, `identity_document_downloadable`,
        `pci_document`, `selfie`, `sigma_scheduled_query`, `tax_document_user_upload`, `terminal_reader_splashscreen`
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Int. The size of the file object in bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        String. The returned file type (for example, `csv`, `pdf`, `jpg`, or `png`).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        String. The publicly accessible URL to download the file.
        """
        return pulumi.get(self, "url")

