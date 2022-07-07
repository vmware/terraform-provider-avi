<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_webapput"
sidebar_current: "docs-avi-datasource-webapput"
description: |-
  Get information of Avi WebappUT.
---

# avi_webapput

This data source is used to to get avi_webapput objects.

## Example Usage

```hcl
data "avi_webapput" "foo_webapput" {
    uuid = "webapput-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WebappUT by name.
* `uuid` - (Optional) Search WebappUT by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `mandatory_test` - Optional message for nested f_mandatory test cases defined at level0. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `mandatory_tests` - Repeated message for nested f_mandatory test cases-level1. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name of the webapput object-level0. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `string_length_test` - Optional message for nested  max string length test cases. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `string_length_tests` - Repeated message for nested  max string length test cases. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant of the webapput object-level0. It is a reference to an object of type tenant. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `test_string` - The maximum string length. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid of the webapput object-level0. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.

