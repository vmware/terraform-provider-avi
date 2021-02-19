############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_botdetectionpolicy"
sidebar_current: "docs-avi-resource-botdetectionpolicy"
description: |-
  Creates and manages Avi BotDetectionPolicy.
---

# avi_botdetectionpolicy

The BotDetectionPolicy resource allows the creation and management of Avi BotDetectionPolicy

## Example Usage

```hcl
resource "avi_botdetectionpolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `ip_location_detector` - (Required) The ip location configuration used in this policy. Field introduced in 21.1.1.
* `ip_reputation_detector` - (Required) The ip reputation configuration used in this policy. Field introduced in 21.1.1.
* `name` - (Required) The name of this bot detection policy. Field introduced in 21.1.1.
* `user_agent_detector` - (Required) The user-agent configuration used in this policy. Field introduced in 21.1.1.
* `allow_list` - (Optional) Allow the user to skip botmanagement for selected requests. Field introduced in 21.1.1.
* `bot_mapping_uuids` - (Optional) System- and user-defined rules for classification. Field introduced in 21.1.1.
* `consolidator_ref` - (Optional) The installation provides an updated ruleset for consolidating the results of different decider phases. It is a reference to an object of type botconfigconsolidator. Field introduced in 21.1.1.
* `description` - (Optional) Human-readable description of this bot detection policy. Field introduced in 21.1.1.
* `tenant_ref` - (Optional) The unique identifier of the tenant to which this policy belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  A unique identifier to this bot detection policy. Field introduced in 21.1.1.

