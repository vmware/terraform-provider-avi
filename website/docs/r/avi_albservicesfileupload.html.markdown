############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_albservicesfileupload"
sidebar_current: "docs-avi-resource-albservicesfileupload"
description: |-
  Creates and manages Avi ALBServicesFileUpload.
---

# avi_albservicesfileupload

The ALBServicesFileUpload resource allows the creation and management of Avi ALBServicesFileUpload

## Example Usage

```hcl
resource "avi_albservicesfileupload" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `case_id` - (Optional) Salesforce alphanumeric caseid to attach uploaded file to.
* `file_path` - (Optional) Stores output file path, for upload to aws s3.
* `name` - (Optional) Field introduced in 18.2.6.
* `s3_directory` - (Optional) Custom aws s3 directory path to upload file.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

