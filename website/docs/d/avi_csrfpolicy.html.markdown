<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_csrfpolicy"
sidebar_current: "docs-avi-datasource-csrfpolicy"
description: |-
  Get information of Avi CSRFPolicy.
---

# avi_csrfpolicy

This data source is used to to get avi_csrfpolicy objects.

## Example Usage

```hcl
data "avi_csrfpolicy" "foo_csrfpolicy" {
    uuid = "csrfpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CSRFPolicy by name.
* `uuid` - (Optional) Search CSRFPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `cookie_name` - Name of the cookie to be used for csrf token. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `description` - Human-readable description of this csrf protection policy. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - The name of this csrf protection policy. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `rules` - Rules to control which requests undergo csrf protection.if the client's request doesn't match with any rules matchtarget, bypass_csrf action is applied. Field introduced in 30.2.1. Minimum of 1 items required. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - The unique identifier of the tenant to which this policy belongs. It is a reference to an object of type tenant. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `token_validity_time_min` - Csrf token is rotated when this time expires. Tokens will be acceptable for twice the token_validity_time time. Allowed values are 10-1440. Special values are 0- unlimited. Field introduced in 30.2.1. Unit is min. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - A unique identifier to this csrf protection policy. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

