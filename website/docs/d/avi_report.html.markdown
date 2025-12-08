<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_report"
sidebar_current: "docs-avi-datasource-report"
description: |-
  Get information of Avi Report.
---

# avi_report

This data source is used to to get avi_report objects.

## Example Usage

```hcl
data "avi_report" "foo_report" {
    uuid = "report-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Report by name.
* `uuid` - (Optional) Search Report by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `duration` - Time taken to complete report generation in seconds. Field introduced in 31.2.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `end_time` - End time of the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `filename` - Name of the report artifact on reports repository. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - Name of the report. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `node` - Cluster member node on which the report is processed. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `pre_check` - Pre-check details for the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `progress` - Percentage of tasks completed. Allowed values are 0-100. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, enterprise with cloud services edition.
* `request` - Request for the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `start_time` - Start time of the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `state` - State of the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tasks` - List of tasks associated with the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tasks_completed` - No. Of tasks completed. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid of the report generation. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `total_tasks` - Total no. Of tasks. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Uuid identifier for the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

