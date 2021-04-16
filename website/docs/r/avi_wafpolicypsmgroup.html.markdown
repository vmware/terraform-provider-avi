<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_wafpolicypsmgroup"
sidebar_current: "docs-avi-resource-wafpolicypsmgroup"
description: |-
  Creates and manages Avi WafPolicyPSMGroup.
---

# avi_wafpolicypsmgroup

The WafPolicyPSMGroup resource allows the creation and management of Avi WafPolicyPSMGroup

## Example Usage

```hcl
resource "avi_wafpolicypsmgroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) User defined name of the group. Field introduced in 18.2.3.
* `description` - (Optional) Free-text comment about this group. Field introduced in 18.2.3.
* `enable` - (Optional) Enable or disable this waf rule group. Field introduced in 18.2.3.
* `hit_action` - (Optional) If a rule in this group matches the match_value pattern, this action will be executed. Allowed actions are waf_action_no_op and waf_action_allow_parameter. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 18.2.3.
* `is_learning_group` - (Optional) This field indicates that this group is used for learning. Field introduced in 18.2.3.
* `locations` - (Optional) Positive security model locations. These are used to partition the application name space. Field introduced in 18.2.3. Maximum of 16384 items allowed.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `miss_action` - (Optional) If a rule in this group does not match the match_value pattern, this action will be executed. Allowed actions are waf_action_no_op and waf_action_block. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 18.2.3.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of this object. Field introduced in 18.2.3.

