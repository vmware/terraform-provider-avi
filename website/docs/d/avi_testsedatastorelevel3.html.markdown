<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_testsedatastorelevel3"
sidebar_current: "docs-avi-datasource-testsedatastorelevel3"
description: |-
  Get information of Avi TestSeDatastoreLevel3.
---

# avi_testsedatastorelevel3

This data source is used to to get avi_testsedatastorelevel3 objects.

## Example Usage

```hcl
data "avi_testsedatastorelevel3" "foo_testsedatastorelevel3" {
    uuid = "testsedatastorelevel3-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TestSeDatastoreLevel3 by name.
* `uuid` - (Optional) Search TestSeDatastoreLevel3 by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

