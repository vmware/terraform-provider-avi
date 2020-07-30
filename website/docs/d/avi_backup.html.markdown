############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_backup"
sidebar_current: "docs-avi-datasource-backup"
description: |-
  Get information of Avi Backup.
---

# avi_backup

This data source is used to to get avi_backup objects.

## Example Usage

```hcl
data "avi_backup" "foo_backup" {
    uuid = "backup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Backup by name.
* `uuid` - (Optional) Search Backup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `backup_config_ref` - Backupconfiguration information.
* `file_name` - The file name of backup.
* `local_file_url` - Url to download the backup file.
* `remote_file_url` - Url to download the backup file.
* `scheduler_ref` - Scheduler information.
* `tenant_ref` - It is a reference to an object of type tenant.
* `timestamp` - Unix timestamp of when the backup file is created.
* `uuid` - Unique object identifier of the object.

