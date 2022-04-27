<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `backup_config_ref` - Backupconfiguration information. It is a reference to an object of type backupconfiguration. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `file_name` - The file name of backup. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `local_file_url` - Url to download the backup file. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `remote_file_url` - Url to download the backup file. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `scheduler_ref` - Scheduler information. It is a reference to an object of type scheduler. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `timestamp` - Unix timestamp of when the backup file is created. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

