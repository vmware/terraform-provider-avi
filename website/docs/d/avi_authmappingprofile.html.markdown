<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_authmappingprofile"
sidebar_current: "docs-avi-datasource-authmappingprofile"
description: |-
  Get information of Avi AuthMappingProfile.
---

# avi_authmappingprofile

This data source is used to to get avi_authmappingprofile objects.

## Example Usage

```hcl
data "avi_authmappingprofile" "foo_authmappingprofile" {
    uuid = "authmappingprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AuthMappingProfile by name.
* `uuid` - (Optional) Search AuthMappingProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 22.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Description for the authmappingprofile. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `mapping_rules` - Rules list for tenant or role mapping. Field introduced in 22.1.1. Minimum of 1 items required. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name of the authmappingprofile. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant ref for the auth mapping profile. It is a reference to an object of type tenant. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `type` - Type of the auth profile for which these rules can be linked. Enum options - AUTH_PROFILE_LDAP, AUTH_PROFILE_TACACS_PLUS, AUTH_PROFILE_SAML, AUTH_PROFILE_PINGACCESS, AUTH_PROFILE_JWT, AUTH_PROFILE_OAUTH. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid of the authmappingprofile. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

