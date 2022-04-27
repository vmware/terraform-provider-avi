<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_ipamdnsproviderprofile"
sidebar_current: "docs-avi-resource-ipamdnsproviderprofile"
description: |-
  Creates and manages Avi IpamDnsProviderProfile.
---

# avi_ipamdnsproviderprofile

The IpamDnsProviderProfile resource allows the creation and management of Avi IpamDnsProviderProfile

## Example Usage

```hcl
resource "avi_ipamdnsproviderprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the ipam/dns provider profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `type` - (Required) Provider type for the ipam/dns provider profile. Enum options - IPAMDNS_TYPE_INFOBLOX, IPAMDNS_TYPE_AWS, IPAMDNS_TYPE_OPENSTACK, IPAMDNS_TYPE_GCP, IPAMDNS_TYPE_INFOBLOX_DNS, IPAMDNS_TYPE_CUSTOM, IPAMDNS_TYPE_CUSTOM_DNS, IPAMDNS_TYPE_AZURE, IPAMDNS_TYPE_OCI, IPAMDNS_TYPE_TENCENT, IPAMDNS_TYPE_INTERNAL, IPAMDNS_TYPE_INTERNAL_DNS, IPAMDNS_TYPE_AWS_DNS, IPAMDNS_TYPE_AZURE_DNS. Allowed in enterprise edition with any value, essentials edition(allowed values- ipamdns_type_internal), basic edition(allowed values- ipamdns_type_internal), enterprise with cloud services edition.
* `allocate_ip_in_vrf` - (Optional) If this flag is set, only allocate ip from networks in the virtual service vrf. Applicable for avi vantage ipam only. Field introduced in 17.2.4. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `aws_profile` - (Optional) Provider details if type is aws. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `azure_profile` - (Optional) Provider details if type is microsoft azure. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `custom_profile` - (Optional) Provider details if type is custom. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `gcp_profile` - (Optional) Provider details if type is google cloud. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `infoblox_profile` - (Optional) Provider details if type is infoblox. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `internal_profile` - (Optional) Provider details if type is avi. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `oci_profile` - (Optional) Provider details for oracle cloud. Field introduced in 18.2.1,18.1.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `openstack_profile` - (Optional) Provider details if type is openstack. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `proxy_configuration` - (Optional) Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tencent_profile` - (Optional) Provider details for tencent cloud. Field introduced in 18.2.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the ipam/dns provider profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

