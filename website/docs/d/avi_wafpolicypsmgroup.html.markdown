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
data "WafPolicyPSMGroup" "foo_WafPolicyPSMGroup" {
    uuid = "WafPolicyPSMGroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafPolicyPSMGroup by name.
* `uuid` - (Optional) Search WafPolicyPSMGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Freetext comment about this group.
* `enable` - Enable or disable this waf rule group.
* `hit_action` - If a rule in this group matches the match_value pattern, this action will be executed.
* `is_learning_group` - This field indicates that this group is used for learning.
* `locations` - Positive security model locations.
* `miss_action` - If a rule in this group does not match the match_value pattern, this action will be executed.
* `name` - User defined name of the group.
* `tenant_ref` - Tenant that this object belongs to.
* `uuid` - Uuid of this object.

