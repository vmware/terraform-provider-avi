<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_apiratelimitprofile"
sidebar_current: "docs-avi-datasource-apiratelimitprofile"
description: |-
  Get information of Avi ApiRateLimitProfile.
---

# avi_apiratelimitprofile

This data source is used to to get avi_apiratelimitprofile objects.

## Example Usage

```hcl
data "avi_apiratelimitprofile" "foo_apiratelimitprofile" {
    uuid = "apiratelimitprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApiRateLimitProfile by name.
* `uuid` - (Optional) Search ApiRateLimitProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Description for the api rate limit profile. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `enabled` - Activate/deactivate the api rate limit profile. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the api rate limit profile. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `rate_limit_configuration_refs` - List of the rate limiter configuration uuids. It is a reference to an object of type ratelimitconfiguration. Field introduced in 31.2.1. Minimum of 1 items required. Maximum of 100 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant ref for the api rate limit profile. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the api rate limit profile. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

