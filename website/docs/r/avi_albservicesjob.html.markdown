<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_albservicesjob"
sidebar_current: "docs-avi-resource-albservicesjob"
description: |-
  Creates and manages Avi ALBServicesJob.
---

# avi_albservicesjob

The ALBServicesJob resource allows the creation and management of Avi ALBServicesJob

## Example Usage

```hcl
resource "avi_albservicesjob" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `command` - (Required) The command to be triggered by the albservicesjob. Field introduced in 21.1.3.
* `name` - (Required) The name of the albservicesjob. Field introduced in 21.1.3.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.3.
* `end_time` - (Optional) The time at which the albservicesjob is ended. Field introduced in 21.1.3.
* `pulse_job_id` - (Optional) A unique identifier for this job entry on the pulse portal. Field introduced in 21.1.3.
* `start_time` - (Optional) The time at which the albservicesjob is started. Field introduced in 21.1.3.
* `status` - (Optional) The status of the albservicesjob. Enum options - UNDETERMINED, PENDING, IN_PROGRESS, COMPLETED, FAILED. Field introduced in 21.1.3.
* `tenant_ref` - (Optional) The unique identifier of the tenant to which this albservicesjob belongs. It is a reference to an object of type tenant. Field introduced in 21.1.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  A unique identifier for this albservicesjob entry. Field introduced in 21.1.3.
