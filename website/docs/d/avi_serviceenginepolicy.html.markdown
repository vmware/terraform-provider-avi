---
layout: "avi"
page_title: "AVI: avi_serviceenginepolicy"
sidebar_current: "docs-avi-datasource-serviceenginepolicy"
description: |-
  Get information of Avi ServiceEnginePolicy.
---

# avi_serviceenginepolicy

This data source is used to to get avi_serviceenginepolicy objects.

## Example Usage

```hcl
data "ServiceEnginePolicy" "foo_ServiceEnginePolicy" {
    uuid = "ServiceEnginePolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ServiceEnginePolicy by name.
* `uuid` - (Optional) Search ServiceEnginePolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the service engine policy.
* `nat_policy_ref` - Nat policy.
* `se_group_ref` - Service engine group to which the policy is applied.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the service engine policy.
* `vrf_ref` - Vrf context to which the policy is scoped.

