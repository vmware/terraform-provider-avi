---
layout: "avi"
page_title: "AVI: avi_testsedatastorelevel2"
sidebar_current: "docs-avi-datasource-testsedatastorelevel2"
description: |-
  Get information of Avi TestSeDatastoreLevel2.
---

# avi_testsedatastorelevel2

This data source is used to to get avi_testsedatastorelevel2 objects.

## Example Usage

```hcl
data "avi_testsedatastorelevel2" "foo_testsedatastorelevel2" {
    uuid = "testsedatastorelevel2-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TestSeDatastoreLevel2 by name.
* `uuid` - (Optional) Search TestSeDatastoreLevel2 by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the object.
* `tenant_ref` - It is a reference to an object of type tenant.
* `test_se_datastore_level_3_refs` - It is a reference to an object of type testsedatastorelevel3.
* `uuid` - Unique object identifier of the object.

