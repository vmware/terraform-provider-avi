<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_tenantsystemconfiguration"
sidebar_current: "docs-avi-datasource-tenantsystemconfiguration"
description: |-
  Get information of Avi TenantSystemConfiguration.
---

# avi_tenantsystemconfiguration

This data source is used to to get avi_tenantsystemconfiguration objects.

## Example Usage

```hcl
data "avi_tenantsystemconfiguration" "foo_tenantsystemconfiguration" {
    uuid = "tenantsystemconfiguration-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TenantSystemConfiguration by name.
* `uuid` - (Optional) Search TenantSystemConfiguration by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 30.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `dns_virtualservice_refs` - Dns virtual services hosting fqdn records for applications configured within this tenant. It is a reference to an object of type virtualservice. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name of the tenant system configuration object. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Unique identifier of the tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Tenant system configuration uuid. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

