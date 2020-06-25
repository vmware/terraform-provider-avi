---
layout: "avi"
page_title: "Avi: avi_systemlimits"
sidebar_current: "docs-avi-resource-systemlimits"
description: |-
  Creates and manages Avi SystemLimits.
---

# avi_systemlimits

The SystemLimits resource allows the creation and management of Avi SystemLimits

## Example Usage

```hcl
resource "avi_systemlimits" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `controller_limits` - (Optional) System limits for the entire controller cluster.
* `controller_sizes` - (Optional) Possible controller sizes.
* `serviceengine_limits` - (Optional) System limits that apply to a serviceengine.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid for the system limits object.

