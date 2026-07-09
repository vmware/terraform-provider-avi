<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_tlsprofile"
sidebar_current: "docs-avi-resource-tlsprofile"
description: |-
  Creates and manages Avi TLSProfile.
---

# avi_tlsprofile

The TLSProfile resource allows the creation and management of Avi TLSProfile

## Example Usage

```hcl
resource "avi_tlsprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the tls profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `certificate_ref` - (Optional) Client certificate (and private key) presented to the remote server during the tls handshake. Needed when a consumer requests mtls tls_mode against this tls profile. It is a reference to an object of type sslkeyandcertificate. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Human-readable description for this tls profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `pki_profile_ref` - (Optional) Pki profile containing the ca certificates used to validate the tls certificate presented by the remote server. Needed when a consumer (e.g. Authprofile) requests tls, mtls, or verify_only tls_mode against this tls profile. It is a reference to an object of type pkiprofile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Changing this value forces the resource to be recreated.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the tls profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

