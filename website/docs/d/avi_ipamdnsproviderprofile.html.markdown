############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `allocate_ip_in_vrf` - If this flag is set, only allocate ip from networks in the virtual service vrf.
* `aws_profile` - Provider details if type is aws.
* `azure_profile` - Provider details if type is microsoft azure.
* `custom_profile` - Provider details if type is custom.
* `gcp_profile` - Provider details if type is google cloud.
* `infoblox_profile` - Provider details if type is infoblox.
* `internal_profile` - Provider details if type is avi.
* `name` - Name for the ipam/dns provider profile.
* `oci_profile` - Provider details for oracle cloud.
* `openstack_profile` - Provider details if type is openstack.
* `proxy_configuration` - Field introduced in 17.1.1.
* `tenant_ref` - It is a reference to an object of type tenant.
* `tencent_profile` - Provider details for tencent cloud.
* `type` - Provider type for the ipam/dns provider profile.
* `uuid` - Uuid of the ipam/dns provider profile.

