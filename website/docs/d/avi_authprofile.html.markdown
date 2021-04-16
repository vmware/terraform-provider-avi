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

* `description` - User defined description for the object.
* `http` - Http user authentication params.
* `jwt_profile_ref` - Jwtserverprofile to be used for authentication. It is a reference to an object of type jwtserverprofile. Field introduced in 20.1.3.
* `ldap` - Ldap server and directory settings.
* `name` - Name of the auth profile.
* `pa_agent_ref` - Pingaccessagent uuid. It is a reference to an object of type pingaccessagent. Field introduced in 18.2.3. Allowed in basic edition, essentials edition, enterprise edition.
* `saml` - Saml settings. Field introduced in 17.2.3.
* `tacacs_plus` - Tacacs+ settings.
* `tenant_ref` - It is a reference to an object of type tenant.
* `type` - Type of the auth profile. Enum options - AUTH_PROFILE_LDAP, AUTH_PROFILE_TACACS_PLUS, AUTH_PROFILE_SAML, AUTH_PROFILE_PINGACCESS, AUTH_PROFILE_JWT.
* `uuid` - Uuid of the auth profile.

