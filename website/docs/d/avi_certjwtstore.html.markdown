<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_certjwtstore"
sidebar_current: "docs-avi-datasource-certjwtstore"
description: |-
  Get information of Avi CertJwtStore.
---

# avi_certjwtstore

This data source is used to to get avi_certjwtstore objects.

## Example Usage

```hcl
data "avi_certjwtstore" "foo_certjwtstore" {
    uuid = "certjwtstore-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CertJwtStore by name.
* `uuid` - (Optional) Search CertJwtStore by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `jwt` - Jwt containing current portal certificate along with the full certificate bundle chain, signed by the private key of previous portal certificate. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `key` - Private key. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `key_passphrase` - Private key passphrase. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `kid` - Sha256 thumbprint of the previous old portal certificate. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `last_rotated_at` - Timestamp of certificate rotation. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `public_key_algorithm` - Public key algorithm. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `type` - Type of ssl certificate. Enum options - SSL_CERTIFICATE_TYPE_VIRTUALSERVICE, SSL_CERTIFICATE_TYPE_SYSTEM, SSL_CERTIFICATE_TYPE_CA, SSL_CERTIFICATE_TYPE_CLIENT, SSL_CERTIFICATE_TYPE_SECURE_CHANNEL. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of jwt. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

