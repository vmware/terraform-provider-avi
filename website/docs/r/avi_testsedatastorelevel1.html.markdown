############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_testsedatastorelevel1"
sidebar_current: "docs-avi-resource-testsedatastorelevel1"
description: |-
  Creates and manages Avi TestSeDatastoreLevel1.
---

# avi_testsedatastorelevel1

The TestSeDatastoreLevel1 resource allows the creation and management of Avi TestSeDatastoreLevel1

## Example Usage

```hcl
resource "avi_testsedatastorelevel1" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `test_se_datastore_level_2_ref` - (Optional) It is a reference to an object of type testsedatastorelevel2.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

