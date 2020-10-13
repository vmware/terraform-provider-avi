############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_ssopolicy"
sidebar_current: "docs-avi-resource-ssopolicy"
description: |-
  Creates and manages Avi SSOPolicy.
---

# avi_ssopolicy

The SSOPolicy resource allows the creation and management of Avi SSOPolicy

## Example Usage

```hcl
resource "avi_ssopolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `authentication_policy` - (Optional) Authentication policy settings.
* `authorization_policy` - (Optional) Authorization policy settings.
* `labels` - (Optional) Key value pairs for granular object access control.
* `name` - (Optional) Name of the sso policy.
* `tenant_ref` - (Optional) Uuid of the tenant.
* `type` - (Optional) Sso policy type.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the sso policy.

