<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_serviceauthprofile"
sidebar_current: "docs-avi-resource-serviceauthprofile"
description: |-
  Creates and manages Avi ServiceAuthProfile.
---

# avi_serviceauthprofile

The ServiceAuthProfile resource allows the creation and management of Avi ServiceAuthProfile

## Example Usage

```hcl
resource "avi_serviceauthprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the service auth profile. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `type` - (Required) Type of the service auth profile. Enum options - SERVICE_AUTH_OAUTH. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 30.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - (Optional) Description for the service auth profile. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `service_oauth_profile` - (Optional) Oauth profile - common endpoint information for service authentication. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant ref for the service auth profile. It is a reference to an object of type tenant. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the service auth profile. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

