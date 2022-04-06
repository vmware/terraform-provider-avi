<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_natpolicy"
sidebar_current: "docs-avi-datasource-natpolicy"
description: |-
  Get information of Avi NatPolicy.
---

# avi_natpolicy

This data source is used to to get avi_natpolicy objects.

## Example Usage

```hcl
data "avi_natpolicy" "foo_natpolicy" {
    uuid = "natpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search NatPolicy by name.
* `uuid` - (Optional) Search NatPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `created_by` - Creator name. Field introduced in 18.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `description` - Field introduced in 18.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Name of the nat policy. Field introduced in 18.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `rules` - Nat policy rules. Field introduced in 18.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 18.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Uuid of the nat policy. Field introduced in 18.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

