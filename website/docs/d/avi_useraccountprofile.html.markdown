<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_useraccountprofile"
sidebar_current: "docs-avi-datasource-useraccountprofile"
description: |-
  Get information of Avi UserAccountProfile.
---

# avi_useraccountprofile

This data source is used to to get avi_useraccountprofile objects.

## Example Usage

```hcl
data "avi_useraccountprofile" "foo_useraccountprofile" {
    uuid = "useraccountprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search UserAccountProfile by name.
* `uuid` - (Optional) Search UserAccountProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `complexity_constraint` - Password complexity constraints for the user account profile. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `expiration_constraint` - Password expiration settings for the user account profile. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `lockout_constraint` - Account lockout settings for the user account profile. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `max_concurrent_sessions` - Maximum number of concurrent sessions allowed. There are unlimited sessions by default. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

