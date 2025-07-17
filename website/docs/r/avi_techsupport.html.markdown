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

* `case_number` - (Optional) 'customer case number for which this tech-upport is generated. ''useful for connected portal and other use-cases.'. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `description` - (Optional) User provided description to capture additional details and context regarding the tech-support invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `duration` - (Optional) Total time taken for tech-support collection. Field introduced in 31.2.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `end_time` - (Optional) End timestamp of tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `errors` - (Optional) Error logged during tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `level` - (Optional) Name of the tech-support level. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - (Optional) Name of tech-support invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `node` - (Optional) Cluster member node on which the tech-support tarball bundle is saved. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_name` - (Optional) Object name if one exists. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_uuid` - (Optional) Tech-support collection object uuid specified for different objects such as se/vs/pool etc. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `output` - (Optional) Tech-support collection output file path. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `params` - (Optional) Tech-support params associated with latest tech-support collection.user passed params will have more preference. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `progress` - (Optional) Tech-support collection progress which holds value between 0-100. Allowed values are 0-100. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, enterprise with cloud services edition.
* `size` - (Optional) Size of collected tech-support tarball. Field introduced in 31.2.1. Unit is mb. Allowed with any value in enterprise, enterprise with cloud services edition.
* `start_time` - (Optional) Start timestamp of tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `state` - (Optional) State of current/last tech-support invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tasks` - (Optional) Events performed for tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tasks_completed` - (Optional) Completed set of tasks in the tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `techsupport_readiness` - (Optional) Techsupport readiness checks execution details. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant uuid associated with the tech-support. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `total_tasks` - (Optional) Total number of tasks in the tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `warnings` - (Optional) Warning logged during tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the tech-support invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

