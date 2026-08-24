<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_albservicesstatus"
sidebar_current: "docs-avi-resource-albservicesstatus"
description: |-
  Creates and manages Avi ALBServicesStatus.
---

# avi_albservicesstatus

The ALBServicesStatus resource allows the creation and management of Avi ALBServicesStatus

## Example Usage

```hcl
resource "avi_albservicesstatus" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the albservicesstatus object. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `asset_details` - (Optional) Asset details corresponding to this controller cluster, on registering with pulse. Field introduced in 22.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `connected_at` - (Optional) Timestamp of last successful connection. Field introduced in 22.1.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `connectivity_status` - (Optional) Connectivity status of controller with albservices. Enum options - ALBSERVICES_CONNECTIVITY_UNKNOWN, ALBSERVICES_DISCONNECTED, ALBSERVICES_CONNECTED. Field introduced in 18.2.6. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `error` - (Optional) Descriptive error message. Field introduced in 18.2.6. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `registration_status` - (Optional) Registration status of the controller with albservices. Enum options - ALBSERVICES_REGISTRATION_UNKNOWN, ALBSERVICES_REGISTERED, ALBSERVICES_DEREGISTERED. Field introduced in 18.2.6. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `services_health` - (Optional) Health of hosted services. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 30.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_status` - (Optional) Tenant based status information. Field introduced in 30.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique identifier of customer portal status object in the database and datastore. Field introduced in 18.2.6. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

