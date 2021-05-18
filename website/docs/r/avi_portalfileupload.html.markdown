############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_portalfileupload"
sidebar_current: "docs-avi-resource-portalfileupload"
description: |-
  Creates and manages Avi PortalFileUpload.
---

# avi_portalfileupload

The PortalFileUpload resource allows the creation and management of Avi PortalFileUpload

## Example Usage

```hcl
resource "avi_portalfileupload" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `case_id` - (Optional) Salesforce alphanumeric caseid to attach uploaded file to.
* `error` - (Optional) Error reported during file upload.
* `file_path` - (Optional) Stores output file path, for upload to aws s3.
* `name` - (Optional) Field introduced in 18.2.6.
* `s3_directory` - (Optional) Custom aws s3 directory path to upload file.
* `status` - (Optional) Captures status for file upload.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

