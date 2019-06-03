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
* `ldap` - Ldap server and directory settings.
* `name` - Name of the auth profile.
* `pa_agent_ref` - Pingaccessagent uuid.
* `saml` - Saml settings.
* `tacacs_plus` - Tacacs+ settings.
* `tenant_ref` - It is a reference to an object of type tenant.
* `type` - Type of the auth profile.
* `uuid` - Uuid of the auth profile.

