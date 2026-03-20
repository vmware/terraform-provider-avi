<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_tenantbinding"
sidebar_current: "docs-avi-datasource-tenantbinding"
description: |-
  Get information of Avi TenantBinding.
---

# avi_tenantbinding

This data source is used to to get avi_tenantbinding objects.

## Example Usage

```hcl
data "avi_tenantbinding" "foo_tenantbinding" {
    uuid = "tenantbinding-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TenantBinding by name.
* `uuid` - (Optional) Search TenantBinding by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the tenant binding. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `se_group_ref` - Uuid of the service engine group being shared. It is a reference to an object of type serviceenginegroup. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `shared_tenant_ref` - Uuid of the tenant to which the object is being shared. It is a reference to an object of type tenant. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid. This field is used for admin tenant context. It is a reference to an object of type tenant. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the tenant binding. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

