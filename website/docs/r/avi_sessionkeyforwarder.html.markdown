<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_sessionkeyforwarder"
sidebar_current: "docs-avi-resource-sessionkeyforwarder"
description: |-
  Creates and manages Avi SessionKeyForwarder.
---

# avi_sessionkeyforwarder

The SessionKeyForwarder resource allows the creation and management of Avi SessionKeyForwarder

## Example Usage

```hcl
resource "avi_sessionkeyforwarder" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `ip_ports` - (Required) Ip addresses and ports to be used for connection with session key forwarder. At least one entry required; maximum 16 (matches the per-core stats slot limit). Field introduced in 32.2.1. Minimum of 1 items required. Maximum of 16 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `name` - (Required) Name of the session key forwarder profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `enable` - (Optional) Enable or disable session key forwarder. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `pki_profile_ref` - (Optional) Pki profile used to validate the ssl certificate presented by a server. It is a reference to an object of type pkiprofile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `ssl_key_and_certificate_ref` - (Optional) Service engines will present this ssl certificate to the server. It is a reference to an object of type sslkeyandcertificate. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `ssl_profile_ref` - (Optional) Ssl profile defines ciphers and ssl versions to be used for session key forwarder. It is a reference to an object of type sslprofile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Tenant reference for the session key forwarder object. It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `use_mgmt` - (Optional) If enabled, connection with session key forwarder will use the management network. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the session key forwarder profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

