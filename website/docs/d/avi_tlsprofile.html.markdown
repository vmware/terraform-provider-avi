<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_tlsprofile"
sidebar_current: "docs-avi-datasource-tlsprofile"
description: |-
  Get information of Avi TLSProfile.
---

# avi_tlsprofile

This data source is used to to get avi_tlsprofile objects.

## Example Usage

```hcl
data "avi_tlsprofile" "foo_tlsprofile" {
    uuid = "tlsprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TLSProfile by name.
* `uuid` - (Optional) Search TLSProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `certificate_ref` - Client certificate (and private key) presented to the remote server during the tls handshake. Needed when a consumer requests mtls tls_mode against this tls profile. It is a reference to an object of type sslkeyandcertificate. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Human-readable description for this tls profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the tls profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `pki_profile_ref` - Pki profile containing the ca certificates used to validate the tls certificate presented by the remote server. Needed when a consumer (e.g. Authprofile) requests tls, mtls, or verify_only tls_mode against this tls profile. It is a reference to an object of type pkiprofile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the tls profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

