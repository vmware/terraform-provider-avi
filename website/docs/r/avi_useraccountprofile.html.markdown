<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_useraccountprofile"
sidebar_current: "docs-avi-resource-useraccountprofile"
description: |-
  Creates and manages Avi UserAccountProfile.
---

# avi_useraccountprofile

The UserAccountProfile resource allows the creation and management of Avi UserAccountProfile

## Example Usage

```hcl
resource "avi_useraccountprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `complexity_constraint` - (Required) Password complexity constraints for the user account profile. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `expiration_constraint` - (Required) Password expiration settings for the user account profile. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `lockout_constraint` - (Required) Account lockout settings for the user account profile. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `name` - (Required) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `max_concurrent_sessions` - (Optional) Maximum number of concurrent sessions allowed. There are unlimited sessions by default. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

