############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_networksecuritypolicy"
sidebar_current: "docs-avi-resource-networksecuritypolicy"
description: |-
  Creates and manages Avi NetworkSecurityPolicy.
---

# avi_networksecuritypolicy

The NetworkSecurityPolicy resource allows the creation and management of Avi NetworkSecurityPolicy

## Example Usage

```hcl
resource "avi_networksecuritypolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `cloud_config_cksum` - (Optional) Checksum of cloud configuration for network sec policy.
* `created_by` - (Optional) Creator name.
* `description` - (Optional) User defined description for the object.
* `name` - (Optional) Name of the object.
* `rules` - (Optional) List of list.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

