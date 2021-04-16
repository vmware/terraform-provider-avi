<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_customipamdnsprofile"
sidebar_current: "docs-avi-resource-customipamdnsprofile"
description: |-
  Creates and manages Avi CustomIpamDnsProfile.
---

# avi_customipamdnsprofile

The CustomIpamDnsProfile resource allows the creation and management of Avi CustomIpamDnsProfile

## Example Usage

```hcl
resource "avi_customipamdnsprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the custom ipam dns profile. Field introduced in 17.1.1.
* `script_uri` - (Required) Script uri of form controller //ipamdnsscripts/<file-name>. Field introduced in 17.1.1.
* `script_params` - (Optional) Parameters that are always passed to the ipam/dns script. Field introduced in 17.1.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.1.1.

