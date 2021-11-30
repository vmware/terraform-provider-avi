<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_scheduler"
sidebar_current: "docs-avi-datasource-scheduler"
description: |-
  Get information of Avi Scheduler.
---

# avi_scheduler

This data source is used to to get avi_scheduler objects.

## Example Usage

```hcl
data "avi_scheduler" "foo_scheduler" {
    uuid = "scheduler-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Scheduler by name.
* `uuid` - (Optional) Search Scheduler by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `backup_config_ref` - Backup configuration to be executed by this scheduler. It is a reference to an object of type backupconfiguration.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `enabled` - Boolean flag to set enabled.
* `end_date_time` - Scheduler end date and time.
* `frequency` - Frequency at which custom scheduler will run. Allowed values are 0-60.
* `frequency_unit` - Unit at which custom scheduler will run. Enum options - SCHEDULER_FREQUENCY_UNIT_MIN, SCHEDULER_FREQUENCY_UNIT_HOUR, SCHEDULER_FREQUENCY_UNIT_DAY, SCHEDULER_FREQUENCY_UNIT_WEEK, SCHEDULER_FREQUENCY_UNIT_MONTH.
* `name` - Name of scheduler.
* `run_mode` - Scheduler run mode. Enum options - RUN_MODE_PERIODIC, RUN_MODE_AT, RUN_MODE_NOW.
* `run_script_ref` - Control script to be executed by this scheduler. It is a reference to an object of type alertscriptconfig.
* `scheduler_action` - Define scheduler action. Enum options - SCHEDULER_ACTION_RUN_A_SCRIPT, SCHEDULER_ACTION_BACKUP.
* `start_date_time` - Scheduler start date and time.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.

