<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_trustedhostprofile"
sidebar_current: "docs-avi-datasource-trustedhostprofile"
description: |-
  Get information of Avi TrustedHostProfile.
---

# avi_trustedhostprofile

This data source is used to to get avi_trustedhostprofile objects.

## Example Usage

```hcl
data "avi_trustedhostprofile" "foo_trustedhostprofile" {
    uuid = "trustedhostprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TrustedHostProfile by name.
* `uuid` - (Optional) Search TrustedHostProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `hosts` - List of host ip(v4/v6) addresses or fqdns. Field introduced in 31.1.1. Minimum of 1 items required. Maximum of 20 items allowed. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - Trustedhostprofile name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Tenant ref for trusted host profile. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Trustedhostprofile uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

