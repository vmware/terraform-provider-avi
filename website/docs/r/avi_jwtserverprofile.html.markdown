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

* `issuer` - (Required) Uniquely identifiable name of the token issuer. Field introduced in 20.1.3.
* `jwks_keys` - (Required) Jwks key set used for validating the jwt. Field introduced in 20.1.3.
* `name` - (Required) Name of the jwt profile. Field introduced in 20.1.3.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `tenant_ref` - (Optional) Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 20.1.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the jwtprofile. Field introduced in 20.1.3.

