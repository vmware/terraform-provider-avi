############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_actiongroupconfig"
sidebar_current: "docs-avi-datasource-actiongroupconfig"
description: |-
  Get information of Avi ActionGroupConfig.
---

# avi_actiongroupconfig

This data source is used to to get avi_actiongroupconfig objects.

## Example Usage

```hcl
data "avi_actiongroupconfig" "foo_actiongroupconfig" {
    uuid = "actiongroupconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ActionGroupConfig by name.
* `uuid` - (Optional) Search ActionGroupConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `action_script_config_ref` - Reference of the action script configuration to be used.
* `autoscale_trigger_notification` - Trigger notification to autoscale manager.
* `description` - User defined description for the object.
* `email_config_ref` - Select the email notification configuration to use when sending alerts via email.
* `external_only` - Generate alert only to external destinations.
* `level` - When an alert is generated, mark its priority via the alert level.
* `name` - Name of the object.
* `snmp_trap_profile_ref` - Select the snmp trap notification to use when sending alerts via snmp trap.
* `syslog_config_ref` - Select the syslog notification configuration to use when sending alerts via syslog.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.

