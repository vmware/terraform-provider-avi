---
layout: "avi"
page_title: "Avi: avi_testsedatastorelevel3"
sidebar_current: "docs-avi-resource-testsedatastorelevel3"
description: |-
  Creates and manages Avi TestSeDatastoreLevel3.
---

# avi_testsedatastorelevel3

The TestSeDatastoreLevel3 resource allows the creation and management of Avi TestSeDatastoreLevel3

## Example Usage

```hcl
resource "avi_testsedatastorelevel3" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

