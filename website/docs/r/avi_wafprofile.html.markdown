<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_wafprofile"
sidebar_current: "docs-avi-resource-wafprofile"
description: |-
  Creates and manages Avi WafProfile.
---

# avi_wafprofile

The WafProfile resource allows the creation and management of Avi WafProfile

## Example Usage

```hcl
resource "avi_wafprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `config` - (Required) Config params for waf. Field introduced in 17.2.1.
* `name` - (Required) Field introduced in 17.2.1.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - (Optional) Field introduced in 17.2.1.
* `files` - (Optional) List of data files used for waf rules. Field introduced in 17.2.1. Maximum of 64 items allowed.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.2.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.2.1.

