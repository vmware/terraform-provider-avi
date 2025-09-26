<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_tenantbinding"
sidebar_current: "docs-avi-resource-tenantbinding"
description: |-
  Creates and manages Avi TenantBinding.
---

# avi_tenantbinding

The TenantBinding resource allows the creation and management of Avi TenantBinding

## Example Usage

```hcl
resource "avi_tenantbinding" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the tenant binding. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `shared_tenant_ref` - (Required) Uuid of the tenant to which the object is being shared. It is a reference to an object of type tenant. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 31.3.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `se_group_ref` - (Optional) Uuid of the service engine group being shared. It is a reference to an object of type serviceenginegroup. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant uuid. This field is used for admin tenant context. It is a reference to an object of type tenant. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the tenant binding. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.

