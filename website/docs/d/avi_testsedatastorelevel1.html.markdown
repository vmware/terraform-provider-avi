<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `test_se_datastore_level_2_ref` - It is a reference to an object of type testsedatastorelevel2. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

