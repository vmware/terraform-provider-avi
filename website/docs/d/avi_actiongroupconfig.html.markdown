<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `action_script_config_ref` - Reference of the action script configuration to be used. It is a reference to an object of type alertscriptconfig. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `autoscale_trigger_notification` - Trigger notification to autoscale manager. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `email_config_ref` - Select the email notification configuration to use when sending alerts via email. It is a reference to an object of type alertemailconfig. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `external_only` - Generate alert only to external destinations. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `level` - When an alert is generated, mark its priority via the alert level. Enum options - ALERT_LOW, ALERT_MEDIUM, ALERT_HIGH. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `snmp_trap_profile_ref` - Select the snmp trap notification to use when sending alerts via snmp trap. It is a reference to an object of type snmptrapprofile. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `syslog_config_ref` - Select the syslog notification configuration to use when sending alerts via syslog. It is a reference to an object of type alertsyslogconfig. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

