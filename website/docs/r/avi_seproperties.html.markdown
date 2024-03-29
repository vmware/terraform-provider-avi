<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `se_agent_properties` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_bootup_properties` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_runtime_properties` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

