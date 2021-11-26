############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_ipamdnsproviderprofile"
sidebar_current: "docs-avi-datasource-ipamdnsproviderprofile"
description: |-
  Get information of Avi IpamDnsProviderProfile.
---

# avi_ipamdnsproviderprofile

This data source is used to to get avi_ipamdnsproviderprofile objects.

## Example Usage

```hcl
data "avi_ipamdnsproviderprofile" "foo_ipamdnsproviderprofile" {
    uuid = "ipamdnsproviderprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search IpamDnsProviderProfile by name.
* `uuid` - (Optional) Search IpamDnsProviderProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allocate_ip_in_vrf` - If this flag is set, only allocate ip from networks in the virtual service vrf. Applicable for avi vantage ipam only. Field introduced in 17.2.4.
* `aws_profile` - Provider details if type is aws.
* `azure_profile` - Provider details if type is microsoft azure. Field introduced in 17.2.1.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `custom_profile` - Provider details if type is custom. Field introduced in 17.1.1.
* `gcp_profile` - Provider details if type is google cloud.
* `infoblox_profile` - Provider details if type is infoblox.
* `internal_profile` - Provider details if type is avi.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name for the ipam/dns provider profile.
* `oci_profile` - Provider details for oracle cloud. Field introduced in 18.2.1,18.1.3.
* `openstack_profile` - Provider details if type is openstack.
* `proxy_configuration` - Field introduced in 17.1.1.
* `tenant_ref` - It is a reference to an object of type tenant.
* `tencent_profile` - Provider details for tencent cloud. Field introduced in 18.2.3.
* `type` - Provider type for the ipam/dns provider profile. Enum options - IPAMDNS_TYPE_INFOBLOX, IPAMDNS_TYPE_AWS, IPAMDNS_TYPE_OPENSTACK, IPAMDNS_TYPE_GCP, IPAMDNS_TYPE_INFOBLOX_DNS, IPAMDNS_TYPE_CUSTOM, IPAMDNS_TYPE_CUSTOM_DNS, IPAMDNS_TYPE_AZURE, IPAMDNS_TYPE_OCI, IPAMDNS_TYPE_TENCENT, IPAMDNS_TYPE_INTERNAL, IPAMDNS_TYPE_INTERNAL_DNS, IPAMDNS_TYPE_AWS_DNS, IPAMDNS_TYPE_AZURE_DNS. Allowed in basic(allowed values- ipamdns_type_internal) edition, essentials(allowed values- ipamdns_type_internal) edition, enterprise edition.
* `uuid` - Uuid of the ipam/dns provider profile.

