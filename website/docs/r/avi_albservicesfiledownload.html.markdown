<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_albservicesfiledownload"
sidebar_current: "docs-avi-resource-albservicesfiledownload"
description: |-
  Creates and manages Avi ALBServicesFileDownload.
---

# avi_albservicesfiledownload

The ALBServicesFileDownload resource allows the creation and management of Avi ALBServicesFileDownload

## Example Usage

```hcl
resource "avi_albservicesfiledownload" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `destination_dir` - (Required) Destination of the file to be saved. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `file_type` - (Required) Software / crs/ inventory. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `file_uri` - (Required) File uri on the cloud bucket. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - (Required) The name of the file with which it is saved to the disk. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `message` - (Optional) Download's success / failure message. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `metadata` - (Optional) Metadata of the file from pulse. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `status` - (Optional) Status of file download. Enum options - SYSERR_SUCCESS, SYSERR_FAILURE, SYSERR_OUT_OF_MEMORY, SYSERR_NO_ENT, SYSERR_INVAL, SYSERR_ACCESS, SYSERR_FAULT, SYSERR_IO, SYSERR_TIMEOUT, SYSERR_NOT_SUPPORTED, SYSERR_NOT_READY, SYSERR_UPGRADE_IN_PROGRESS, SYSERR_WARM_START_IN_PROGRESS, SYSERR_TRY_AGAIN, SYSERR_NOT_UPGRADING, SYSERR_PENDING, SYSERR_EVENT_GEN_FAILURE, SYSERR_CONFIG_PARAM_MISSING, SYSERR_RANGE, SYSERR_BAD_REQUEST... Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique id of the object. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

