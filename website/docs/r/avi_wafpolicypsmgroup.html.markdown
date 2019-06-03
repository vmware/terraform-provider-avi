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

* `description` - (Optional) Freetext comment about this group.
* `enable` - (Optional) Enable or disable this waf rule group.
* `hit_action` - (Optional) If a rule in this group matches the match_value pattern, this action will be executed.
* `is_learning_group` - (Optional) This field indicates that this group is used for learning.
* `locations` - (Optional) Positive security model locations.
* `miss_action` - (Optional) If a rule in this group does not match the match_value pattern, this action will be executed.
* `name` - (Optional) User defined name of the group.
* `tenant_ref` - (Optional) Tenant that this object belongs to.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of this object.

