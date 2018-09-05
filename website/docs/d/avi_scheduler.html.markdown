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
data "Scheduler" "foo_Scheduler" {
    uuid = "Scheduler-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Scheduler by name.
* `uuid` - (Optional) Search Scheduler by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `backup_config_ref` - Backup configuration to be executed by this scheduler.
* `enabled` - General description.
* `end_date_time` - Scheduler end date and time.
* `frequency` - Frequency at which custom scheduler will run.
* `frequency_unit` - Unit at which custom scheduler will run.
* `name` - Name of scheduler.
* `run_mode` - Scheduler run mode.
* `run_script_ref` - Control script to be executed by this scheduler.
* `scheduler_action` - Define scheduler action.
* `start_date_time` - Scheduler start date and time.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.

