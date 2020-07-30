############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `name` - (Required) Field introduced in 17.2.1.
* `config` - (Optional) Config params for waf.
* `description` - (Optional) Field introduced in 17.2.1.
* `files` - (Optional) List of data files used for waf rules.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.2.1.

