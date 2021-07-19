<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_alertsyslogconfig"
sidebar_current: "docs-avi-datasource-alertsyslogconfig"
description: |-
  Get information of Avi AlertSyslogConfig.
---

# avi_alertsyslogconfig

This data source is used to to get avi_alertsyslogconfig objects.

## Example Usage

```hcl
data "avi_alertsyslogconfig" "foo_alertsyslogconfig" {
    uuid = "alertsyslogconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AlertSyslogConfig by name.
* `uuid` - (Optional) Search AlertSyslogConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - User defined description for alert syslog config.
* `name` - A user-friendly name of the syslog notification.
* `syslog_servers` - The list of syslog servers.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.

