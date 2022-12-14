<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_albservicesjob"
sidebar_current: "docs-avi-datasource-albservicesjob"
description: |-
  Get information of Avi ALBServicesJob.
---

# avi_albservicesjob

This data source is used to to get avi_albservicesjob objects.

## Example Usage

```hcl
data "avi_albservicesjob" "foo_albservicesjob" {
    uuid = "albservicesjob-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ALBServicesJob by name.
* `uuid` - (Optional) Search ALBServicesJob by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `command` - The command to be triggered by the albservicesjob. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `end_time` - The time at which the albservicesjob is ended. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - The name of the albservicesjob. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `params` - Job params. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `pulse_job_id` - A unique identifier for this job entry on the pulse portal. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `pulse_sync_status` - Status of sync to pulse(result uploads/state updates). Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `result` - Job result. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `start_time` - The time at which the albservicesjob is started. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `status` - The status of the albservicesjob. Enum options - UNDETERMINED, PENDING, IN_PROGRESS, COMPLETED, FAILED, NOT_ENABLED. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - The unique identifier of the tenant to which this albservicesjob belongs. It is a reference to an object of type tenant. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `token` - Job token. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - A unique identifier for this albservicesjob entry. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.

