<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_apiratelimitprofile"
sidebar_current: "docs-avi-resource-apiratelimitprofile"
description: |-
  Creates and manages Avi ApiRateLimitProfile.
---

# avi_apiratelimitprofile

The ApiRateLimitProfile resource allows the creation and management of Avi ApiRateLimitProfile

## Example Usage

```hcl
resource "avi_apiratelimitprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the api rate limit profile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `rate_limit_configuration_refs` - (Required) List of the rate limiter configuration uuids. It is a reference to an object of type ratelimitconfiguration. Field introduced in 31.2.1. Minimum of 1 items required. Maximum of 100 items allowed. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - (Optional) Description for the api rate limit profile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enabled` - (Optional) Activate/deactivate the api rate limit profile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant ref for the api rate limit profile. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the api rate limit profile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

