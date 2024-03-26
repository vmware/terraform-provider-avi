<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_authprofile"
sidebar_current: "docs-avi-datasource-authprofile"
description: |-
  Get information of Avi AuthProfile.
---

# avi_authprofile

This data source is used to to get avi_authprofile objects.

## Example Usage

```hcl
data "avi_authprofile" "foo_authprofile" {
    uuid = "authprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AuthProfile by name.
* `uuid` - (Optional) Search AuthProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `http` - Http user authentication params. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `jwt_profile_ref` - Jwtserverprofile to be used for authentication. It is a reference to an object of type jwtserverprofile. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ldap` - Ldap server and directory settings. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.6. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Name of the auth profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `oauth_profile` - Oauth profile - common endpoint information. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `saml` - Saml settings. Field introduced in 17.2.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tacacs_plus` - Tacacs+ settings. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `type` - Type of the auth profile. Enum options - AUTH_PROFILE_LDAP, AUTH_PROFILE_TACACS_PLUS, AUTH_PROFILE_SAML, AUTH_PROFILE_PINGACCESS, AUTH_PROFILE_JWT, AUTH_PROFILE_OAUTH. Allowed in enterprise edition with any value, essentials edition(allowed values- auth_profile_ldap,auth_profile_tacacs_plus,auth_profile_saml,auth_profile_jwt,auth_profile_oauth), basic edition(allowed values- auth_profile_ldap,auth_profile_tacacs_plus,auth_profile_saml,auth_profile_jwt,auth_profile_oauth), enterprise with cloud services edition.
* `uuid` - Uuid of the auth profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

