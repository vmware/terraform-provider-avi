<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_upgradeprofile"
sidebar_current: "docs-avi-resource-upgradeprofile"
description: |-
  Creates and manages Avi UpgradeProfile.
---

# avi_upgradeprofile

The UpgradeProfile resource allows the creation and management of Avi UpgradeProfile

## Example Usage

```hcl
resource "avi_upgradeprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `controller_params` - (Optional) List of controller upgrade related configurable parameters. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `dry_run` - (Optional) List of dryrun related configurable parameters. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `image` - (Optional) List of image related configurable parameters. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `pre_checks` - (Optional) List of upgrade pre-checks related configurable parameters. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `service_engine` - (Optional) List of service engine upgrade related configurable parameters. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the upgradeprofile object. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

