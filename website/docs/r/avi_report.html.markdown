<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_report"
sidebar_current: "docs-avi-resource-report"
description: |-
  Creates and manages Avi Report.
---

# avi_report

The Report resource allows the creation and management of Avi Report

## Example Usage

```hcl
resource "avi_report" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `duration` - (Optional) Time taken to complete report generation in seconds. Field introduced in 31.2.1. Unit is sec. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `end_time` - (Optional) End time of the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `filename` - (Optional) Name of the report artifact on reports repository. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `name` - (Optional) Name of the report. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `node` - (Optional) Cluster member node on which the report is processed. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `pre_check` - (Optional) Pre-check details for the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `progress` - (Optional) Percentage of tasks completed. Allowed values are 0-100. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `request` - (Optional) Request for the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `start_time` - (Optional) Start time of the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `state` - (Optional) State of the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tasks` - (Optional) List of tasks associated with the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tasks_completed` - (Optional) No. Of tasks completed. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Tenant uuid of the report generation. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `total_tasks` - (Optional) Total no. Of tasks. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the report generation. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

