---
layout: "avi"
page_title: "Avi: avi_serviceenginepolicy"
sidebar_current: "docs-avi-resource-serviceenginepolicy"
description: |-
  Creates and manages Avi ServiceEnginePolicy.
---

# avi_serviceenginepolicy

The ServiceEnginePolicy resource allows the creation and management of Avi ServiceEnginePolicy

## Example Usage

```hcl
resource "avi_serviceenginepolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the service engine policy.
* `nat_policy_ref` - (Optional) Nat policy.
* `se_group_ref` - (Optional) Service engine group to which the policy is applied.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `vrf_ref` - (Optional) Vrf context to which the policy is scoped.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the service engine policy.

