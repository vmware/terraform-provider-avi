<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_jwtserverprofile"
sidebar_current: "docs-avi-datasource-jwtserverprofile"
description: |-
  Get information of Avi JWTServerProfile.
---

# avi_jwtserverprofile

This data source is used to to get avi_jwtserverprofile objects.

## Example Usage

```hcl
data "avi_jwtserverprofile" "foo_jwtserverprofile" {
    uuid = "jwtserverprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search JWTServerProfile by name.
* `uuid` - (Optional) Search JWTServerProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `controller_internal_auth` - Jwt auth configuration for profile_type controller_internal_auth. Field introduced in 20.1.6.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.6.
* `issuer` - Uniquely identifiable name of the token issuer, only allowed with profile_type client_auth. Field introduced in 20.1.3.
* `jwks_keys` - Jwks key set used for validating the jwt, only allowed with profile_type client_auth. Field introduced in 20.1.3.
* `jwt_profile_type` - Type of jwt server profile which defines the usage type. Enum options - CLIENT_AUTH, CONTROLLER_INTERNAL_AUTH. Field introduced in 20.1.6.
* `name` - Name of the jwt profile. Field introduced in 20.1.3.
* `tenant_ref` - Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 20.1.3.
* `uuid` - Uuid of the jwtprofile. Field introduced in 20.1.3.

