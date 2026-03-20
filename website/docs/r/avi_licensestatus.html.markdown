<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_licensestatus"
sidebar_current: "docs-avi-resource-licensestatus"
description: |-
  Creates and manages Avi LicenseStatus.
---

# avi_licensestatus

The LicenseStatus resource allows the creation and management of Avi LicenseStatus

## Example Usage

```hcl
resource "avi_licensestatus" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `cls_status` - (Optional) Cls licensing status. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `legacy_license_grace_period` - (Optional) Legacy license grace period, when controller upgrades with existing legacy licenses. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `saas_status` - (Optional) Saas licensing status. Field introduced in 21.1.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `service_update` - (Optional) Pulse license service update. Field introduced in 21.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `support_legacy_license` - (Optional) Indicates if legacy licenses are supported. When false, legacy licenses have been cleaned up after grace period expiry. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_uuid` - (Optional) Tenant uuid. Field introduced in 30.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid. Field introduced in 21.1.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

