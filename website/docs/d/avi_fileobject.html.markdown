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

* `checksum` - Sha1 checksum of the file. Field introduced in 20.1.1.
* `compressed` - This field indicates whether the file is gzip-compressed. Field introduced in 20.1.1.
* `created` - Timestamp of creation for the file. Field introduced in 20.1.1.
* `description` - Description of the file. Field introduced in 20.1.1.
* `expires_at` - Timestamp when the file will be no longer needed and can be removed by the system. If this is set, a garbage collector process will try to remove the file after this time. Field introduced in 20.1.1.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.1.
* `name` - Name of the file object. Field introduced in 20.1.1.
* `path` - Path to the file. Field introduced in 20.1.1.
* `read_only` - Enforce read-only on the file. Field introduced in 20.1.1.
* `restrict_download` - Flag to allow/restrict download of the file. Field introduced in 20.1.1.
* `size` - Size of the file. Field introduced in 20.1.1.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `type` - Type of the file. Enum options - OTHER_FILE_TYPES, IP_REPUTATION, GEO_DB, TECH_SUPPORT, HSMPACKAGES, IPAMDNSSCRIPTS, CONTROLLER_IMAGE. Field introduced in 20.1.1. Allowed in basic(allowed values- other_file_types) edition, essentials(allowed values- other_file_types) edition, enterprise edition.
* `uuid` - Uuid of the file. Field introduced in 20.1.1.
* `version` - Version of the file. Field introduced in 20.1.1.

