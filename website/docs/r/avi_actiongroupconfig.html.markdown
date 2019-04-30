---
layout: "avi"
page_title: "Avi: avi_actiongroupconfig"
sidebar_current: "docs-avi-resource-actiongroupconfig"
description: |-
  Creates and manages Avi ActionGroupConfig.
---

# avi_actiongroupconfig

The ActionGroupConfig resource allows the creation and management of Avi ActionGroupConfig

## Example Usage

```hcl
resource "ActionGroupConfig" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `action_script_config_ref` - (Optional ) argument_description.
        * `autoscale_trigger_notification` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `email_config_ref` - (Optional ) argument_description.
        * `external_only` - (Required) argument_description.
        * `level` - (Required) argument_description.
        * `name` - (Required) argument_description.
        * `snmp_trap_profile_ref` - (Optional ) argument_description.
        * `syslog_config_ref` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                            * `uuid` - argument_description.
    
