<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_upgradeprofile"
sidebar_current: "docs-avi-datasource-upgradeprofile"
description: |-
  Get information of Avi UpgradeProfile.
---

# avi_upgradeprofile

This data source is used to to get avi_upgradeprofile objects.

## Example Usage

```hcl
data "avi_upgradeprofile" "foo_upgradeprofile" {
    uuid = "upgradeprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search UpgradeProfile by name.
* `uuid` - (Optional) Search UpgradeProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `controller_params` - List of controller upgrade related configurable parameters. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `dry_run` - List of dryrun related configurable parameters. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `image` - List of image related configurable parameters. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `pre_checks` - List of upgrade pre-checks related configurable parameters. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `service_engine` - List of service engine upgrade related configurable parameters. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid identifier for the upgradeprofile object. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

