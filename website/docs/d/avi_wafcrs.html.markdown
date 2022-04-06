<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_wafcrs"
sidebar_current: "docs-avi-datasource-wafcrs"
description: |-
  Get information of Avi WafCRS.
---

# avi_wafcrs

This data source is used to to get avi_wafcrs objects.

## Example Usage

```hcl
data "avi_wafcrs" "foo_wafcrs" {
    uuid = "wafcrs-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafCRS by name.
* `uuid` - (Optional) Search WafCRS by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - A short description of this ruleset. Field introduced in 18.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `groups` - Waf rules are sorted in groups based on their characterization. Field introduced in 18.1.1. Maximum of 64 items allowed. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `integrity` - Integrity protection value. Field introduced in 18.2.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.6. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - The name of this ruleset object. Field introduced in 18.2.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `release_date` - The release date of this version in rfc 3339 / iso 8601 format. Field introduced in 18.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Field introduced in 18.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `version` - The version of this ruleset object. Field introduced in 18.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

