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

* `name` - (Required) Ip reputation db name. Field introduced in 20.1.1.
* `vendor` - (Required) Organization providing ip reputation data. Enum options - IP_REPUTATION_VENDOR_WEBROOT. Field introduced in 20.1.1.
* `base_file_refs` - (Optional) Ip reputation db base file. It is a reference to an object of type fileobject. Field introduced in 20.1.1. Maximum of 1 items allowed.
* `description` - (Optional) Description. Field introduced in 20.1.1.
* `incremental_file_refs` - (Optional) Ip reputation db incremental update files. It is a reference to an object of type fileobject. Field introduced in 20.1.1.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `service_status` - (Optional) If this object is managed by the ip reputation service, this field contain the status of this syncronization. Field introduced in 20.1.1.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `version` - (Optional) A version number for this database object. This is informal for the consumer of this api only, a tool which manages this object can store version information here. Field introduced in 20.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of this object. Field introduced in 20.1.1.

