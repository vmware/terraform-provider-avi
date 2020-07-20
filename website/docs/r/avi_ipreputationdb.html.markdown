############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_ipreputationdb"
sidebar_current: "docs-avi-resource-ipreputationdb"
description: |-
  Creates and manages Avi IPReputationDB.
---

# avi_ipreputationdb

The IPReputationDB resource allows the creation and management of Avi IPReputationDB

## Example Usage

```hcl
resource "avi_ipreputationdb" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `base_file_refs` - (Optional) Ip reputation db base file.
* `description` - (Optional) Description.
* `incremental_file_refs` - (Optional) Ip reputation db incremental update files.
* `name` - (Optional) Ip reputation db name.
* `service_status` - (Optional) If this object is managed by the ip reputation service, this field contain the status of this syncronization.
* `tenant_ref` - (Optional) Tenant that this object belongs to.
* `vendor` - (Optional) Organization providing ip reputation data.
* `version` - (Optional) A version number for this database object.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of this object.

