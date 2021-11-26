############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_backup"
sidebar_current: "docs-avi-resource-backup"
description: |-
  Creates and manages Avi Backup.
---

# avi_backup

The Backup resource allows the creation and management of Avi Backup

## Example Usage

```hcl
resource "avi_backup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `file_name` - (Required) The file name of backup.
* `backup_config_ref` - (Optional) Backupconfiguration information. It is a reference to an object of type backupconfiguration.
* `local_file_url` - (Optional) Url to download the backup file.
* `remote_file_url` - (Optional) Url to download the backup file.
* `scheduler_ref` - (Optional) Scheduler information. It is a reference to an object of type scheduler.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `timestamp` - (Optional) Unix timestamp of when the backup file is created.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

