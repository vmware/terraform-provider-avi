<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_apispecgenerate"
sidebar_current: "docs-avi-resource-apispecgenerate"
description: |-
  Creates and manages Avi ApiSpecGenerate.
---

# avi_apispecgenerate

The ApiSpecGenerate resource allows the creation and management of Avi ApiSpecGenerate

## Example Usage

```hcl
resource "avi_apispecgenerate" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `completed_events` - (Optional) Number of tasks completed. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `duration` - (Optional) Spec generation duration in seconds. Field introduced in 32.1.4. Unit is sec. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `end_time` - (Optional) Time the spec generation completed or failed. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `name` - (Optional) Name of the spec generation object. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `params` - (Optional) Parameters for the spec generation. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `path` - (Optional) Path to the generated spec file. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `progress` - (Optional) Overall spec generation progress percentage. Allowed values are 0-100. Field introduced in 32.1.4. Unit is percent. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `start_time` - (Optional) Time the spec generation started. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `state` - (Optional) Current lifecycle state of the spec generation. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `task_events` - (Optional) Per-task status and event details. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `total_events` - (Optional) Total number of tasks. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the spec generation object. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

