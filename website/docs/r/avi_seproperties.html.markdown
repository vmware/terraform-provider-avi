############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_seproperties"
sidebar_current: "docs-avi-resource-seproperties"
description: |-
  Creates and manages Avi SeProperties.
---

# avi_seproperties

The SeProperties resource allows the creation and management of Avi SeProperties

## Example Usage

```hcl
resource "avi_seproperties" "foo" {
    uuid = "default-uuid"
}
```

## Argument Reference

The following arguments are supported:

* `se_agent_properties` - (Optional) Dict settings for seproperties.
* `se_bootup_properties` - (Optional) Dict settings for seproperties.
* `se_runtime_properties` - (Optional) Dict settings for seproperties.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

