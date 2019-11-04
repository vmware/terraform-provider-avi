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

* `case_id` - Salesforce alphanumeric caseid to attach uploaded file to.
* `error` - Error reported during file upload.
* `file_path` - Stores output file path, for upload to aws s3.
* `name` - Field introduced in 18.2.6.
* `s3_directory` - Custom aws s3 directory path to upload file.
* `status` - Captures status for file upload.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.

