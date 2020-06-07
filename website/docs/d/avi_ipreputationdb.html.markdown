---
layout: "avi"
page_title: "AVI: avi_ipreputationdb"
sidebar_current: "docs-avi-datasource-ipreputationdb"
description: |-
  Get information of Avi IPReputationDB.
---

# avi_ipreputationdb

This data source is used to to get avi_ipreputationdb objects.

## Example Usage

```hcl
data "avi_ipreputationdb" "foo_ipreputationdb" {
    uuid = "ipreputationdb-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search IPReputationDB by name.
* `uuid` - (Optional) Search IPReputationDB by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `base_file_refs` - Ip reputation db base file.
* `description` - Description.
* `incremental_file_refs` - Ip reputation db incremental update files.
* `name` - Ip reputation db name.
* `service_status` - If this object is managed by the ip reputation service, this field contain the status of this syncronization.
* `tenant_ref` - Tenant that this object belongs to.
* `uuid` - Uuid of this object.
* `vendor` - Organization providing ip reputation data.
* `version` - A version number for this database object.

