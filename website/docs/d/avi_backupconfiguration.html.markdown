############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_backupconfiguration"
sidebar_current: "docs-avi-datasource-backupconfiguration"
description: |-
  Get information of Avi BackupConfiguration.
---

# avi_backupconfiguration

This data source is used to to get avi_backupconfiguration objects.

## Example Usage

```hcl
data "avi_backupconfiguration" "foo_backupconfiguration" {
    uuid = "backupconfiguration-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search BackupConfiguration by name.
* `uuid` - (Optional) Search BackupConfiguration by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `aws_access_key` - Aws access key id.
* `aws_bucket_id` - Aws bucket.
* `aws_secret_access` - Aws secret access key.
* `backup_file_prefix` - Prefix of the exported configuration file.
* `backup_passphrase` - Passphrase of backup configuration.
* `maximum_backups_stored` - Rotate the backup files based on this count.
* `name` - Name of backup configuration.
* `remote_directory` - Directory at remote destination with write permission for ssh user.
* `remote_hostname` - Remote destination.
* `save_local` - Local backup.
* `ssh_user_ref` - Access credentials for remote destination.
* `tenant_ref` - It is a reference to an object of type tenant.
* `upload_to_remote_host` - Remote backup.
* `upload_to_s3` - Cloud backup.
* `uuid` - Unique object identifier of the object.

