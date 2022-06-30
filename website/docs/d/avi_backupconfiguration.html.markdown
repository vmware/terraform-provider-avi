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

* `aws_access_key` - Aws access key id. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `aws_bucket_id` - Aws bucket. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `aws_bucket_region` - The name of the aws region associated with the bucket. Field introduced in 21.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `aws_secret_access` - Aws secret access key. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `backup_file_prefix` - Prefix of the exported configuration file. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `backup_passphrase` - Default passphrase for configuration export and periodic backup. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `maximum_backups_stored` - Rotate the backup files based on this count. Allowed values are 1-20. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Name of backup configuration. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `remote_directory` - Directory at remote destination with write permission for ssh user. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `remote_hostname` - Remote destination. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `save_local` - Local backup. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ssh_user_ref` - Access credentials for remote destination. It is a reference to an object of type cloudconnectoruser. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `upload_to_remote_host` - Remote backup. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `upload_to_s3` - Cloud backup. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

