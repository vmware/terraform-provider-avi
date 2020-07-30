############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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
resource "avi_actiongroupconfig" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `external_only` - (Required) Generate alert only to external destinations.
* `level` - (Required) When an alert is generated, mark its priority via the alert level.
* `name` - (Required) Name of the object.
* `action_script_config_ref` - (Optional) Reference of the action script configuration to be used.
* `autoscale_trigger_notification` - (Optional) Trigger notification to autoscale manager.
* `description` - (Optional) User defined description for the object.
* `email_config_ref` - (Optional) Select the email notification configuration to use when sending alerts via email.
* `snmp_trap_profile_ref` - (Optional) Select the snmp trap notification to use when sending alerts via snmp trap.
* `syslog_config_ref` - (Optional) Select the syslog notification configuration to use when sending alerts via syslog.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

