<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_authmappingprofile"
sidebar_current: "docs-avi-resource-authmappingprofile"
description: |-
  Creates and manages Avi AuthMappingProfile.
---

# avi_authmappingprofile

The AuthMappingProfile resource allows the creation and management of Avi AuthMappingProfile

## Example Usage

```hcl
resource "avi_authmappingprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `mapping_rules` - (Required) Rules list for tenant or role mapping. Field introduced in 22.1.1. Minimum of 1 items required. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - (Required) Name of the authmappingprofile. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `type` - (Required) Type of the auth profile for which these rules can be linked. Enum options - AUTH_PROFILE_LDAP, AUTH_PROFILE_TACACS_PLUS, AUTH_PROFILE_SAML, AUTH_PROFILE_PINGACCESS, AUTH_PROFILE_JWT, AUTH_PROFILE_OAUTH. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 22.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - (Optional) Description for the authmappingprofile. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant ref for the auth mapping profile. It is a reference to an object of type tenant. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the authmappingprofile. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

