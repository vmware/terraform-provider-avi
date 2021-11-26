############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_albservicesfileupload"
sidebar_current: "docs-avi-datasource-albservicesfileupload"
description: |-
  Get information of Avi ALBServicesFileUpload.
---

# avi_albservicesfileupload

This data source is used to to get avi_albservicesfileupload objects.

## Example Usage

```hcl
data "avi_albservicesfileupload" "foo_albservicesfileupload" {
    uuid = "albservicesfileupload-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ALBServicesFileUpload by name.
* `uuid` - (Optional) Search ALBServicesFileUpload by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `case_id` - Salesforce alphanumeric caseid to attach uploaded file to. Field introduced in 18.2.6.
* `file_path` - Stores output file path, for upload to aws s3. Field introduced in 18.2.6.
* `name` - Field introduced in 18.2.6.
* `s3_directory` - Custom aws s3 directory path to upload file. Field introduced in 18.2.6.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `uuid` - Unique object identifier of the object.

