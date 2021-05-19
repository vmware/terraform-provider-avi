<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_jwtprofile"
sidebar_current: "docs-avi-datasource-jwtprofile"
description: |-
  Get information of Avi JWTProfile.
---

# avi_jwtprofile

This data source is used to to get avi_jwtprofile objects.

## Example Usage

```hcl
data "avi_jwtprofile" "foo_jwtprofile" {
    uuid = "jwtprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search JWTProfile by name.
* `uuid` - (Optional) Search JWTProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.5. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `jwks_keys` - Jwk keys used for signing/validating the jwt. Field introduced in 20.1.5. Minimum of 1 items required. Maximum of 1 items allowed.
* `jwt_auth_type` - Jwt auth type for jwt validation. Enum options - JWT_TYPE_JWS. Field introduced in 20.1.5.
* `name` - A user friendly name for this jwt profile. Field introduced in 20.1.5.
* `tenant_ref` - Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 20.1.5.
* `uuid` - Uuid of the jwt profile. Field introduced in 20.1.5.

