<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_testsedatastorelevel2"
sidebar_current: "docs-avi-resource-testsedatastorelevel2"
description: |-
  Creates and manages Avi TestSeDatastoreLevel2.
---

# avi_testsedatastorelevel2

The TestSeDatastoreLevel2 resource allows the creation and management of Avi TestSeDatastoreLevel2

## Example Usage

```hcl
resource "avi_testsedatastorelevel2" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `test_se_datastore_level_3_refs` - (Optional) It is a reference to an object of type testsedatastorelevel3. Field introduced in 18.2.6.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

