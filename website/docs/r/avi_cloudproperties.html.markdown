---
layout: "avi"
page_title: "Avi: avi_cloudproperties"
sidebar_current: "docs-avi-resource-cloudproperties"
description: |-
  Creates and manages Avi CloudProperties.
---

# avi_cloudproperties

The CloudProperties resource allows the creation and management of Avi CloudProperties

## Example Usage

```hcl
resource "avi_cloudproperties" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `cc_props` - (Optional) Cloudconnector properties.
* `cc_vtypes` - (Optional) Cloud types supported by cloudconnector.
* `hyp_props` - (Optional) Hypervisor properties.
* `info` - (Optional) Properties specific to a cloud type.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

