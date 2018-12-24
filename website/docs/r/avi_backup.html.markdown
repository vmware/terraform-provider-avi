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
resource "Backup" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `backup_config_ref` - (Optional ) argument_description.
        * `file_name` - (Required) argument_description.
        * `local_file_url` - (Optional ) argument_description.
        * `remote_file_url` - (Optional ) argument_description.
        * `scheduler_ref` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `timestamp` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                * `uuid` - argument_description.
    
