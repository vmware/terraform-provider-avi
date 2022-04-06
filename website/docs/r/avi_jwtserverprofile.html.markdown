<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_jwtserverprofile"
sidebar_current: "docs-avi-resource-jwtserverprofile"
description: |-
  Creates and manages Avi JWTServerProfile.
---

# avi_jwtserverprofile

The JWTServerProfile resource allows the creation and management of Avi JWTServerProfile

## Example Usage

```hcl
resource "avi_jwtserverprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the jwt profile. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `controller_internal_auth` - (Optional) Jwt auth configuration for profile_type controller_internal_auth. Field introduced in 20.1.6. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.6. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `issuer` - (Optional) Uniquely identifiable name of the token issuer, only allowed with profile_type client_auth. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `jwks_keys` - (Optional) Jwks key set used for validating the jwt, only allowed with profile_type client_auth. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `jwt_profile_type` - (Optional) Type of jwt server profile which defines the usage type. Enum options - CLIENT_AUTH, CONTROLLER_INTERNAL_AUTH. Field introduced in 20.1.6. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the jwtprofile. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.

