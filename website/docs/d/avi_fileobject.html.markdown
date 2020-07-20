############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_fileobject"
sidebar_current: "docs-avi-datasource-fileobject"
description: |-
  Get information of Avi FileObject.
---

# avi_fileobject

This data source is used to to get avi_fileobject objects.

## Example Usage

```hcl
data "avi_fileobject" "foo_fileobject" {
    uuid = "fileobject-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search FileObject by name.
* `uuid` - (Optional) Search FileObject by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `checksum` - Sha1 checksum of the file.
* `compressed` - This field indicates whether the file is gzip-compressed.
* `created` - Timestamp of creation for the file.
* `description` - Description of the file.
* `expires_at` - Timestamp when the file will be no longer needed and can be removed by the system.
* `is_federated` - This field describes the object's replication scope.
* `name` - Name of the file object.
* `path` - Path to the file.
* `read_only` - Enforce read-only on the file.
* `restrict_download` - Flag to allow/restrict download of the file.
* `size` - Size of the file.
* `tenant_ref` - Tenant that this object belongs to.
* `type` - Type of the file.
* `uuid` - Uuid of the file.
* `version` - Version of the file.

