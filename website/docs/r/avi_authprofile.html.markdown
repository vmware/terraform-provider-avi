<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_authprofile"
sidebar_current: "docs-avi-resource-authprofile"
description: |-
  Creates and manages Avi AuthProfile.
---

# avi_authprofile

The AuthProfile resource allows the creation and management of Avi AuthProfile

## Example Usage

```hcl
resource "avi_authprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the auth profile. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `type` - (Required) Type of the auth profile. Enum options - AUTH_PROFILE_LDAP, AUTH_PROFILE_TACACS_PLUS, AUTH_PROFILE_SAML, AUTH_PROFILE_PINGACCESS, AUTH_PROFILE_JWT, AUTH_PROFILE_OAUTH. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `http` - (Optional) Http user authentication params. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `jwt_profile_ref` - (Optional) Jwtserverprofile to be used for authentication. It is a reference to an object of type jwtserverprofile. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `ldap` - (Optional) Ldap server and directory settings. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.6. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `oauth_profile` - (Optional) Oauth profile - common endpoint information. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `pa_agent_ref` - (Optional) Pingaccessagent uuid. It is a reference to an object of type pingaccessagent. Field introduced in 18.2.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `saml` - (Optional) Saml settings. Field introduced in 17.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tacacs_plus` - (Optional) Tacacs+ settings. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the auth profile. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

