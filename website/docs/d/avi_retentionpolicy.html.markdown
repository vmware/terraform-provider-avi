<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_retentionpolicy"
sidebar_current: "docs-avi-datasource-retentionpolicy"
description: |-
  Get information of Avi RetentionPolicy.
---

# avi_retentionpolicy

This data source is used to to get avi_retentionpolicy objects.

## Example Usage

```hcl
data "avi_retentionpolicy" "foo_retentionpolicy" {
    uuid = "retentionpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search RetentionPolicy by name.
* `uuid` - (Optional) Search RetentionPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `enabled` - Enables the policy. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `history` - History of previous runs. Field introduced in 31.1.1. Maximum of 10 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the policy. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `policy` - Policy specification. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `summary` - Details of most recent run. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Uuid identifier for the policy. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

