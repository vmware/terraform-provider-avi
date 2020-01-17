---
layout: "avi"
page_title: "AVI: avi_testsedatastorelevel1"
sidebar_current: "docs-avi-datasource-testsedatastorelevel1"
description: |-
  Get information of Avi TestSeDatastoreLevel1.
---

# avi_testsedatastorelevel1

This data source is used to to get avi_testsedatastorelevel1 objects.

## Example Usage

```hcl
data "avi_testsedatastorelevel1" "foo_testsedatastorelevel1" {
    uuid = "testsedatastorelevel1-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TestSeDatastoreLevel1 by name.
* `uuid` - (Optional) Search TestSeDatastoreLevel1 by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the object.
* `tenant_ref` - It is a reference to an object of type tenant.
* `test_se_datastore_level_2_ref` - It is a reference to an object of type testsedatastorelevel2.
* `uuid` - Unique object identifier of the object.

