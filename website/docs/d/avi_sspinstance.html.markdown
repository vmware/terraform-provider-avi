<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_sspinstance"
sidebar_current: "docs-avi-datasource-sspinstance"
description: |-
  Get information of Avi SspInstance.
---

# avi_sspinstance

This data source is used to to get avi_sspinstance objects.

## Example Usage

```hcl
data "avi_sspinstance" "foo_sspinstance" {
    uuid = "sspinstance-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SspInstance by name.
* `uuid` - (Optional) Search SspInstance by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `avi_client_cert` - Client certificate that avi uses to authenticate with the ssp instance. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `client_cert` - Client certificate that the ssp instance uses to authenticate with avi. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Description of the onboarded ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `feature` - Type of the ssp feature instance. Enum options - SSP_INTELLIGENT_ASSIST, SSP_CENTRAL_LICENSING_SERVICE. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `hostname` - Hostname of the ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `ingress_cert` - Ingress (server) certificate chain that the ssp endpoint uses. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - Name of the onboarded ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `resources` - Resources associated with the ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `status` - Status of the ssp feature instance. Enum options - SSP_STATUS_IN_PROGRESS, SSP_STATUS_ACTIVE, SSP_STATUS_CERT_UPDATE_FAILED. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Tenant reference for the ssp object. It is a reference to an object of type tenant. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Uuid for the onboarded ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

