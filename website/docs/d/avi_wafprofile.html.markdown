<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_wafprofile"
sidebar_current: "docs-avi-datasource-wafprofile"
description: |-
  Get information of Avi WafProfile.
---

# avi_wafprofile

This data source is used to to get avi_wafprofile objects.

## Example Usage

```hcl
data "avi_wafprofile" "foo_wafprofile" {
    uuid = "wafprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafProfile by name.
* `uuid` - (Optional) Search WafProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `config` - Config params for waf. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `files` - List of data files used for waf rules. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

