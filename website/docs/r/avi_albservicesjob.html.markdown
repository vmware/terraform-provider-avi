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

* `command` - (Required) The command to be triggered by the albservicesjob. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - (Required) The name of the albservicesjob. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `end_time` - (Optional) The time at which the albservicesjob is ended. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `params` - (Optional) Job params. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `pulse_job_id` - (Optional) A unique identifier for this job entry on the pulse portal. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `pulse_sync_status` - (Optional) Status of sync to pulse(result uploads/state updates). Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `result` - (Optional) Job result. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `start_time` - (Optional) The time at which the albservicesjob is started. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `status` - (Optional) The status of the albservicesjob. Enum options - UNDETERMINED, PENDING, IN_PROGRESS, COMPLETED, FAILED, NOT_ENABLED. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `status_update_time` - (Optional) Time at which the status of albservicesjob updated. Field introduced in 22.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) The unique identifier of the tenant to which this albservicesjob belongs. It is a reference to an object of type tenant. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `token` - (Optional) Job token. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  A unique identifier for this albservicesjob entry. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.

