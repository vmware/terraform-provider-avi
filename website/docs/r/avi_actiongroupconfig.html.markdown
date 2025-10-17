<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `external_only` - (Required) Generate alert only to external destinations. Allowed with any value in enterprise, enterprise with cloud services edition. Allowed in essentials (allowed values- false), basic (allowed values- false) edition.
* `level` - (Required) When an alert is generated, mark its priority via the alert level. Enum options - ALERT_LOW, ALERT_MEDIUM, ALERT_HIGH. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - (Required) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `action_script_config_ref` - (Optional) Reference of the action script configuration to be used. It is a reference to an object of type alertscriptconfig. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `autoscale_trigger_notification` - (Optional) Trigger notification to autoscale manager. Allowed with any value in enterprise, enterprise with cloud services edition. Allowed in essentials (allowed values- false), basic (allowed values- false) edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `email_config_ref` - (Optional) Select the email notification configuration to use when sending alerts via email. It is a reference to an object of type alertemailconfig. Allowed with any value in enterprise, enterprise with cloud services edition.
* `snmp_trap_profile_ref` - (Optional) Select the snmp trap notification to use when sending alerts via snmp trap. It is a reference to an object of type snmptrapprofile. Allowed with any value in enterprise, enterprise with cloud services edition.
* `syslog_config_ref` - (Optional) Select the syslog notification configuration to use when sending alerts via syslog. It is a reference to an object of type alertsyslogconfig. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

