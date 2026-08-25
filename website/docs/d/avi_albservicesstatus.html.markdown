<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_albservicesstatus"
sidebar_current: "docs-avi-datasource-albservicesstatus"
description: |-
  Get information of Avi ALBServicesStatus.
---

# avi_albservicesstatus

This data source is used to to get avi_albservicesstatus objects.

## Example Usage

```hcl
data "avi_albservicesstatus" "foo_albservicesstatus" {
    uuid = "albservicesstatus-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ALBServicesStatus by name.
* `uuid` - (Optional) Search ALBServicesStatus by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `asset_details` - Asset details corresponding to this controller cluster, on registering with pulse. Field introduced in 22.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `connected_at` - Timestamp of last successful connection. Field introduced in 22.1.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `connectivity_status` - Connectivity status of controller with albservices. Enum options - ALBSERVICES_CONNECTIVITY_UNKNOWN, ALBSERVICES_DISCONNECTED, ALBSERVICES_CONNECTED. Field introduced in 18.2.6. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `error` - Descriptive error message. Field introduced in 18.2.6. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the albservicesstatus object. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `registration_status` - Registration status of the controller with albservices. Enum options - ALBSERVICES_REGISTRATION_UNKNOWN, ALBSERVICES_REGISTERED, ALBSERVICES_DEREGISTERED. Field introduced in 18.2.6. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `services_health` - Health of hosted services. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 30.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_status` - Tenant based status information. Field introduced in 30.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Unique identifier of customer portal status object in the database and datastore. Field introduced in 18.2.6. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

