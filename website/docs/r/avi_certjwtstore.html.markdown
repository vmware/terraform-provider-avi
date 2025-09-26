<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_certjwtstore"
sidebar_current: "docs-avi-resource-certjwtstore"
description: |-
  Creates and manages Avi CertJwtStore.
---

# avi_certjwtstore

The CertJwtStore resource allows the creation and management of Avi CertJwtStore

## Example Usage

```hcl
resource "avi_certjwtstore" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `jwt` - (Required) Jwt containing current portal certificate along with the full certificate bundle chain, signed by the private key of previous portal certificate. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `kid` - (Required) Sha256 thumbprint of the previous old portal certificate. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `last_rotated_at` - (Required) Timestamp of certificate rotation. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `public_key_algorithm` - (Required) Public key algorithm. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 31.3.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of jwt. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.

