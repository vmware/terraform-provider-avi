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

* `checksum` - Sha1 checksum of the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `compressed` - This field indicates whether the file is gzip-compressed. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 30.2.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `created` - Timestamp of creation for the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `crl_info` - This field contains certificate revocation list metadata. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `description` - Description of the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `expires_at` - Timestamp when the file will be no longer needed and can be removed by the system. If this is set, a garbage collector process will try to remove the file after this time. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the file object. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `path` - Path to the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `read_only` - Enforce read-only on the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `restrict_download` - Flag to allow/restrict download of the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `size` - Size of the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `type` - Type of the file. Enum options - OTHER_FILE_TYPES, IP_REPUTATION, GEO_DB, TECH_SUPPORT, HSMPACKAGES, IPAMDNSSCRIPTS, CONTROLLER_IMAGE, CRL_DATA. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- other_file_types), basic edition(allowed values- other_file_types), enterprise with cloud services edition.
* `uuid` - Uuid of the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `version` - Version of the file. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

