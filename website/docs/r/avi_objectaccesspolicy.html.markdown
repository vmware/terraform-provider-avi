---
layout: "avi"
page_title: "Avi: avi_objectaccesspolicy"
sidebar_current: "docs-avi-resource-objectaccesspolicy"
description: |-
  Creates and manages Avi ObjectAccessPolicy.
---

# avi_objectaccesspolicy

The ObjectAccessPolicy resource allows the creation and management of Avi ObjectAccessPolicy

## Example Usage

```hcl
resource "avi_objectaccesspolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of the object access policy.
* `rules` - (Optional) Rules which grant access to specific objects.
* `tenant_ref` - (Optional) Tenant that this object belongs to.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the object access policy.

