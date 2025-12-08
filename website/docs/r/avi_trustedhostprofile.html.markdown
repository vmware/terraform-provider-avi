<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_trustedhostprofile"
sidebar_current: "docs-avi-resource-trustedhostprofile"
description: |-
  Creates and manages Avi TrustedHostProfile.
---

# avi_trustedhostprofile

The TrustedHostProfile resource allows the creation and management of Avi TrustedHostProfile

## Example Usage

```hcl
resource "avi_trustedhostprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `hosts` - (Required) List of host ip(v4/v6) addresses or fqdns. Field introduced in 31.1.1. Minimum of 1 items required. Maximum of 20 items allowed. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - (Required) Trustedhostprofile name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant ref for trusted host profile. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Trustedhostprofile uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

