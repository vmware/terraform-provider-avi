<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `base_file_refs` - Ip reputation db base file. It is a reference to an object of type fileobject. Field introduced in 20.1.1. Maximum of 1 items allowed.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - Description. Field introduced in 20.1.1.
* `incremental_file_refs` - Ip reputation db incremental update files. It is a reference to an object of type fileobject. Field introduced in 20.1.1.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `name` - Ip reputation db name. Field introduced in 20.1.1.
* `service_status` - If this object is managed by the ip reputation service, this field contain the status of this syncronization. Field introduced in 20.1.1.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `uuid` - Uuid of this object. Field introduced in 20.1.1.
* `vendor` - Organization providing ip reputation data. Enum options - IP_REPUTATION_VENDOR_WEBROOT. Field introduced in 20.1.1.
* `version` - A version number for this database object. This is informal for the consumer of this api only, a tool which manages this object can store version information here. Field introduced in 20.1.1.

