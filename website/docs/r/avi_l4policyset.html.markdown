############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_l4policyset"
sidebar_current: "docs-avi-resource-l4policyset"
description: |-
  Creates and manages Avi L4PolicySet.
---

# avi_l4policyset

The L4PolicySet resource allows the creation and management of Avi L4PolicySet

## Example Usage

```hcl
resource "avi_l4policyset" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `created_by` - (Optional) Creator name.
* `description` - (Optional) Field introduced in 17.2.7.
* `is_internal_policy` - (Optional) Field introduced in 17.2.7.
* `l4_connection_policy` - (Optional) Policy to apply when a new transport connection is setup.
* `labels` - (Optional) Key value pairs for granular object access control.
* `name` - (Optional) Name of the l4 policy set.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Id of the l4 policy set.

