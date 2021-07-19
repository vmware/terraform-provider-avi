<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `aws_access_key` - Aws access key id. Field introduced in 18.2.3.
* `aws_bucket_id` - Aws bucket. Field introduced in 18.2.3.
* `aws_secret_access` - Aws secret access key. Field introduced in 18.2.3.
* `backup_file_prefix` - Prefix of the exported configuration file. Field introduced in 17.1.1.
* `backup_passphrase` - Passphrase of backup configuration.
* `maximum_backups_stored` - Rotate the backup files based on this count. Allowed values are 1-20.
* `name` - Name of backup configuration.
* `remote_directory` - Directory at remote destination with write permission for ssh user.
* `remote_hostname` - Remote destination.
* `save_local` - Local backup.
* `ssh_user_ref` - Access credentials for remote destination. It is a reference to an object of type cloudconnectoruser.
* `tenant_ref` - It is a reference to an object of type tenant.
* `upload_to_remote_host` - Remote backup.
* `upload_to_s3` - Cloud backup. Field introduced in 18.2.3.
* `uuid` - Unique object identifier of the object.

