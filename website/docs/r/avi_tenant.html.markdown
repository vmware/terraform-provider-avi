<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_tenant"
sidebar_current: "docs-avi-resource-tenant"
description: |-
  Creates and manages Avi Tenant.
---

# avi_tenant

The Tenant resource allows the creation and management of Avi Tenant

## Example Usage

```hcl
resource "avi_tenant" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `config_settings` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `created_by` - (Optional) Creator of this tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enforce_label_group` - (Optional) The referred label groups are enforced on the tenant if this is set to true.if this is set to false, the label groups are suggested for the tenant. Field introduced in 20.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `label_group_refs` - (Optional) The label_groups to be enforced on the tenant. This is strictly enforced only if enforce_label_group is set to true. It is a reference to an object of type labelgroup. Field introduced in 20.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `local` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

