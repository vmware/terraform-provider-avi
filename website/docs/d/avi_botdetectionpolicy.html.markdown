############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_botdetectionpolicy"
sidebar_current: "docs-avi-datasource-botdetectionpolicy"
description: |-
  Get information of Avi BotDetectionPolicy.
---

# avi_botdetectionpolicy

This data source is used to to get avi_botdetectionpolicy objects.

## Example Usage

```hcl
data "avi_botdetectionpolicy" "foo_botdetectionpolicy" {
    uuid = "botdetectionpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search BotDetectionPolicy by name.
* `uuid` - (Optional) Search BotDetectionPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allow_list` - Allow the user to skip botmanagement for selected requests. Field introduced in 21.1.1.
* `bot_mapping_uuids` - System- and user-defined rules for classification. Field introduced in 21.1.1.
* `consolidator_ref` - The installation provides an updated ruleset for consolidating the results of different decider phases. It is a reference to an object of type botconfigconsolidator. Field introduced in 21.1.1.
* `description` - Human-readable description of this bot detection policy. Field introduced in 21.1.1.
* `ip_location_detector` - The ip location configuration used in this policy. Field introduced in 21.1.1.
* `ip_reputation_detector` - The ip reputation configuration used in this policy. Field introduced in 21.1.1.
* `name` - The name of this bot detection policy. Field introduced in 21.1.1.
* `tenant_ref` - The unique identifier of the tenant to which this policy belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1.
* `user_agent_detector` - The user-agent configuration used in this policy. Field introduced in 21.1.1.
* `uuid` - A unique identifier to this bot detection policy. Field introduced in 21.1.1.

