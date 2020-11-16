############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `name` - (Required) Name for the ipam/dns provider profile.
* `type` - (Required) Provider type for the ipam/dns provider profile.
* `allocate_ip_in_vrf` - (Optional) If this flag is set, only allocate ip from networks in the virtual service vrf.
* `aws_profile` - (Optional) Provider details if type is aws.
* `azure_profile` - (Optional) Provider details if type is microsoft azure.
* `custom_profile` - (Optional) Provider details if type is custom.
* `gcp_profile` - (Optional) Provider details if type is google cloud.
* `infoblox_profile` - (Optional) Provider details if type is infoblox.
* `internal_profile` - (Optional) Provider details if type is avi.
* `labels` - (Optional) Key value pairs for granular object access control.
* `oci_profile` - (Optional) Provider details for oracle cloud.
* `openstack_profile` - (Optional) Provider details if type is openstack.
* `proxy_configuration` - (Optional) Field introduced in 17.1.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `tencent_profile` - (Optional) Provider details for tencent cloud.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the ipam/dns provider profile.

