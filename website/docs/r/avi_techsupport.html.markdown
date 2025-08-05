<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_techsupport"
sidebar_current: "docs-avi-resource-techsupport"
description: |-
  Creates and manages Avi TechSupport.
---

# avi_techsupport

The TechSupport resource allows the creation and management of Avi TechSupport

## Example Usage

```hcl
resource "avi_techsupport" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `case_number` - (Optional) 'customer case number for which this techsupport is generated. ''useful for connected portal and other use-cases.'. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `description` - (Optional) User provided description to capture additional details and context regarding the techsupport invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `duration` - (Optional) Total time taken for techsupport collection. Field introduced in 31.2.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `end_time` - (Optional) End timestamp of techsupport collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `errors` - (Optional) Error logged during techsupport collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `level` - (Optional) Name of the techsupport level. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - (Optional) Name of techsupport invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `node` - (Optional) Cluster member node on which the techsupport tarball bundle is saved. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_name` - (Optional) Object name if one exists. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_uuid` - (Optional) Techsupport collection object uuid specified for different objects such as se/vs/pool etc. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `output` - (Optional) Techsupport collection output file path. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `params` - (Optional) Techsupport params associated with latest techsupport collection. User passed params will have more preference. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `progress` - (Optional) Techsupport collection progress which holds value between 0-100. Allowed values are 0-100. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, enterprise with cloud services edition.
* `size` - (Optional) Size of collected techsupport tarball. Field introduced in 31.2.1. Unit is mb. Allowed with any value in enterprise, enterprise with cloud services edition.
* `start_time` - (Optional) Start timestamp of techsupport collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `state` - (Optional) State of current/last techsupport invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tasks` - (Optional) Events performed for techsupport collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tasks_completed` - (Optional) Completed set of tasks in the techsupport collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `techsupport_readiness` - (Optional) Techsupport readiness checks execution details. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant uuid associated with the techsupport. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `total_tasks` - (Optional) Total number of tasks in the techsupport collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `warnings` - (Optional) Warning logged during techsupport collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the techsupport invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

