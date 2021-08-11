<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_portalfileupload"
sidebar_current: "docs-avi-datasource-portalfileupload"
description: |-
  Get information of Avi PortalFileUpload.
---

# avi_portalfileupload

This data source is used to to get avi_portalfileupload objects.

## Example Usage

```hcl
data "avi_portalfileupload" "foo_portalfileupload" {
    uuid = "portalfileupload-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search PortalFileUpload by name.
* `uuid` - (Optional) Search PortalFileUpload by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `case_id` - Salesforce alphanumeric caseid to attach uploaded file to. Field introduced in 18.2.6.
* `error` - Error reported during file upload. Field introduced in 18.2.6.
* `file_path` - Stores output file path, for upload to aws s3. Field introduced in 18.2.6.
* `name` - Field introduced in 18.2.6.
* `s3_directory` - Custom aws s3 directory path to upload file. Field introduced in 18.2.6.
* `status` - Captures status for file upload. Enum options - SYSERR_SUCCESS, SYSERR_FAILURE, SYSERR_OUT_OF_MEMORY, SYSERR_NO_ENT, SYSERR_INVAL, SYSERR_ACCESS, SYSERR_FAULT, SYSERR_IO, SYSERR_TIMEOUT, SYSERR_NOT_SUPPORTED, SYSERR_NOT_READY, SYSERR_UPGRADE_IN_PROGRESS, SYSERR_WARM_START_IN_PROGRESS, SYSERR_TRY_AGAIN, SYSERR_NOT_UPGRADING, SYSERR_PENDING, SYSERR_EVENT_GEN_FAILURE, SYSERR_CONFIG_PARAM_MISSING, SYSERR_BAD_REQUEST, SYSERR_TEST1... Field introduced in 18.2.6.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `uuid` - Unique object identifier of the object.

