
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
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
data "Backup" "foo_Backup" {
    uuid = "Backup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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
* `uuid` - General description.

