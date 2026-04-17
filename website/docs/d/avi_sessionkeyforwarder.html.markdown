<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_sessionkeyforwarder"
sidebar_current: "docs-avi-datasource-sessionkeyforwarder"
description: |-
  Get information of Avi SessionKeyForwarder.
---

# avi_sessionkeyforwarder

This data source is used to to get avi_sessionkeyforwarder objects.

## Example Usage

```hcl
data "avi_sessionkeyforwarder" "foo_sessionkeyforwarder" {
    uuid = "sessionkeyforwarder-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SessionKeyForwarder by name.
* `uuid` - (Optional) Search SessionKeyForwarder by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `enable` - Enable or disable session key forwarder. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `ip_ports` - Ip addresses and ports to be used for connection with session key forwarder. At least one entry required; maximum 16 (matches the per-core stats slot limit). Field introduced in 32.2.1. Minimum of 1 items required. Maximum of 16 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the session key forwarder profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `pki_profile_ref` - Pki profile used to validate the ssl certificate presented by a server. It is a reference to an object of type pkiprofile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `ssl_key_and_certificate_ref` - Service engines will present this ssl certificate to the server. It is a reference to an object of type sslkeyandcertificate. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `ssl_profile_ref` - Ssl profile defines ciphers and ssl versions to be used for session key forwarder. It is a reference to an object of type sslprofile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant reference for the session key forwarder object. It is a reference to an object of type tenant. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `use_mgmt` - If enabled, connection with session key forwarder will use the management network. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the session key forwarder profile. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

