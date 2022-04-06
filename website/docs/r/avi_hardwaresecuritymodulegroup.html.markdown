<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_hardwaresecuritymodulegroup"
sidebar_current: "docs-avi-resource-hardwaresecuritymodulegroup"
description: |-
  Creates and manages Avi HardwareSecurityModuleGroup.
---

# avi_hardwaresecuritymodulegroup

The HardwareSecurityModuleGroup resource allows the creation and management of Avi HardwareSecurityModuleGroup

## Example Usage

```hcl
resource "avi_hardwaresecuritymodulegroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `hsm` - (Required) Hardware security module configuration. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `name` - (Required) Name of the hsm group configuration object. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the hsm group configuration object. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

