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

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - User defined description for alert syslog config. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `name` - A user-friendly name of the syslog notification. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `syslog_servers` - The list of syslog servers. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

