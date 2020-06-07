---
layout: "avi"
page_title: "Avi: avi_backupconfiguration"
sidebar_current: "docs-avi-resource-backupconfiguration"
description: |-
  Creates and manages Avi BackupConfiguration.
---

# avi_backupconfiguration

The BackupConfiguration resource allows the creation and management of Avi BackupConfiguration

## Example Usage

```hcl
resource "avi_backupconfiguration" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of backup configuration.
* `aws_access_key` - (Optional) Aws access key id.
* `aws_bucket_id` - (Optional) Aws bucket.
* `aws_secret_access` - (Optional) Aws secret access key.
* `backup_file_prefix` - (Optional) Prefix of the exported configuration file.
* `backup_passphrase` - (Optional) Default passphrase for configuration export and periodic backup.
* `maximum_backups_stored` - (Optional) Rotate the backup files based on this count.
* `remote_directory` - (Optional) Directory at remote destination with write permission for ssh user.
* `remote_hostname` - (Optional) Remote destination.
* `save_local` - (Optional) Local backup.
* `ssh_user_ref` - (Optional) Access credentials for remote destination.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `upload_to_remote_host` - (Optional) Remote backup.
* `upload_to_s3` - (Optional) Cloud backup.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

