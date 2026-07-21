<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_sspinstance"
sidebar_current: "docs-avi-resource-sspinstance"
description: |-
  Creates and manages Avi SspInstance.
---

# avi_sspinstance

The SspInstance resource allows the creation and management of Avi SspInstance

## Example Usage

```hcl
resource "avi_sspinstance" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `feature` - (Required) Type of the ssp feature instance. Enum options - SSP_AVI_OPERATIONS, SSP_CENTRAL_LICENSING_SERVICE. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Changing this value forces the resource to be recreated.
* `hostname` - (Required) Hostname of the ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `name` - (Required) Name of the onboarded ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Changing this value forces the resource to be recreated.
* `status` - (Required) Status of the ssp feature instance. Enum options - SSP_STATUS_IN_PROGRESS, SSP_STATUS_ACTIVE, SSP_STATUS_CERT_UPDATE_FAILED. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `avi_client_cert` - (Optional) Client certificate that avi uses to authenticate with the ssp instance. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `client_cert` - (Optional) Client certificate that the ssp instance uses to authenticate with avi. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Description of the onboarded ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `ingress_cert` - (Optional) Ingress (server) certificate chain that the ssp endpoint uses. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `resources` - (Optional) Resources associated with the ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Tenant reference for the ssp object. It is a reference to an object of type tenant. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid for the onboarded ssp feature instance. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

