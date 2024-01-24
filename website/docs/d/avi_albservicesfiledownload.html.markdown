<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_albservicesfiledownload"
sidebar_current: "docs-avi-datasource-albservicesfiledownload"
description: |-
  Get information of Avi ALBServicesFileDownload.
---

# avi_albservicesfiledownload

This data source is used to to get avi_albservicesfiledownload objects.

## Example Usage

```hcl
data "avi_albservicesfiledownload" "foo_albservicesfiledownload" {
    uuid = "albservicesfiledownload-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ALBServicesFileDownload by name.
* `uuid` - (Optional) Search ALBServicesFileDownload by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `destination_dir` - Destination of the file to be saved. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `file_type` - Software / crs/ inventory. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `file_uri` - File uri on the cloud bucket. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `message` - Download's success / failure message. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `metadata` - Metadata of the file from pulse. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - The name of the file with which it is saved to the disk. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `status` - Status of file download. Enum options - SYSERR_SUCCESS, SYSERR_FAILURE, SYSERR_OUT_OF_MEMORY, SYSERR_NO_ENT, SYSERR_INVAL, SYSERR_ACCESS, SYSERR_FAULT, SYSERR_IO, SYSERR_TIMEOUT, SYSERR_NOT_SUPPORTED, SYSERR_NOT_READY, SYSERR_UPGRADE_IN_PROGRESS, SYSERR_WARM_START_IN_PROGRESS, SYSERR_TRY_AGAIN, SYSERR_NOT_UPGRADING, SYSERR_PENDING, SYSERR_EVENT_GEN_FAILURE, SYSERR_CONFIG_PARAM_MISSING, SYSERR_RANGE, SYSERR_BAD_REQUEST... Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Unique id of the object. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

