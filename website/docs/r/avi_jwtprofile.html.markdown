<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_jwtprofile"
sidebar_current: "docs-avi-resource-jwtprofile"
description: |-
  Creates and manages Avi JWTProfile.
---

# avi_jwtprofile

The JWTProfile resource allows the creation and management of Avi JWTProfile

## Example Usage

```hcl
resource "avi_jwtprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `jwks_keys` - (Required) Jwk keys used for signing/validating the jwt. Field introduced in 20.1.5. Minimum of 1 items required. Maximum of 1 items allowed.
* `jwt_auth_type` - (Required) Jwt auth type for jwt validation. Enum options - JWT_TYPE_JWS. Field introduced in 20.1.5.
* `name` - (Required) A user friendly name for this jwt profile. Field introduced in 20.1.5.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.5. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `tenant_ref` - (Optional) Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 20.1.5.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the jwt profile. Field introduced in 20.1.5.

