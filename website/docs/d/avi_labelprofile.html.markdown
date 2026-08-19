<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_labelprofile"
sidebar_current: "docs-avi-datasource-labelprofile"
description: |-
  Get information of Avi LabelProfile.
---

# avi_labelprofile

This data source is used to to get avi_labelprofile objects.

## Example Usage

```hcl
data "avi_labelprofile" "foo_labelprofile" {
    uuid = "labelprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search LabelProfile by name.
* `uuid` - (Optional) Search LabelProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Description of this label profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `label_definitions` - Labels available in this profile. Label names must be unique (case-insensitive). Field introduced in 32.2.1. Maximum of 256 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `log_labels` - Enables logging of waap labels effective for a request into apilog.effective_labels in the application log. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of this object, unique per tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - The object uuid. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

