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
resource "BackupConfiguration" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `backup_file_prefix` - (Optional ) argument_description.
        * `backup_passphrase` - (Optional ) argument_description.
        * `maximum_backups_stored` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `remote_directory` - (Optional ) argument_description.
        * `remote_hostname` - (Optional ) argument_description.
        * `save_local` - (Optional ) argument_description.
        * `ssh_user_ref` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `upload_to_remote_host` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                            * `uuid` - argument_description.
    
