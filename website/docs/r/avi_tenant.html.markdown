############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_tenant"
sidebar_current: "docs-avi-resource-tenant"
description: |-
  Creates and manages Avi Tenant.
---

# avi_tenant

The Tenant resource allows the creation and management of Avi Tenant

## Example Usage

```hcl
resource "avi_tenant" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `config_settings` - (Optional) Dict settings for tenant.
* `created_by` - (Optional) Creator of this tenant.
* `description` - (Optional) User defined description for the object.
* `local` - (Optional) Boolean flag to set local.
* `suggested_object_labels` - (Optional) Suggestive pool of key value pairs for recommending assignment of labels to objects in the user interface.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

