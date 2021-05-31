<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_wafpolicypsmgroup"
sidebar_current: "docs-avi-datasource-wafpolicypsmgroup"
description: |-
  Get information of Avi WafPolicyPSMGroup.
---

# avi_wafpolicypsmgroup

This data source is used to to get avi_wafpolicypsmgroup objects.

## Example Usage

```hcl
data "avi_wafpolicypsmgroup" "foo_wafpolicypsmgroup" {
    uuid = "wafpolicypsmgroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafPolicyPSMGroup by name.
* `uuid` - (Optional) Search WafPolicyPSMGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - Free-text comment about this group. Field introduced in 18.2.3.
* `enable` - Enable or disable this waf rule group. Field introduced in 18.2.3.
* `hit_action` - If a rule in this group matches the match_value pattern, this action will be executed. Allowed actions are waf_action_no_op and waf_action_allow_parameter. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 18.2.3.
* `is_learning_group` - This field indicates that this group is used for learning. Field introduced in 18.2.3.
* `locations` - Positive security model locations. These are used to partition the application name space. Field introduced in 18.2.3. Maximum of 16384 items allowed.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `miss_action` - If a rule in this group does not match the match_value pattern, this action will be executed. Allowed actions are waf_action_no_op and waf_action_block. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 18.2.3.
* `name` - User defined name of the group. Field introduced in 18.2.3.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.3.
* `uuid` - Uuid of this object. Field introduced in 18.2.3.

