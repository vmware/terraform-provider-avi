<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_tenantsystemconfiguration"
sidebar_current: "docs-avi-resource-tenantsystemconfiguration"
description: |-
  Creates and manages Avi TenantSystemConfiguration.
---

# avi_tenantsystemconfiguration

The TenantSystemConfiguration resource allows the creation and management of Avi TenantSystemConfiguration

## Example Usage

```hcl
resource "avi_tenantsystemconfiguration" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the tenant system configuration object. Field introduced in 23.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 23.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `dns_virtualservice_refs` - (Optional) Dns virtual services hosting fqdn records for applications configured within this tenant. It is a reference to an object of type virtualservice. Field introduced in 23.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Unique identifier of the tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 23.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Tenant system configuration uuid. Field introduced in 23.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

