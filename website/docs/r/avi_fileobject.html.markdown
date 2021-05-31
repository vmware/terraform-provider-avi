<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_fileobject"
sidebar_current: "docs-avi-resource-fileobject"
description: |-
  Creates and manages Avi FileObject.
---

# avi_fileobject

The FileObject resource allows the creation and management of Avi FileObject

## Example Usage

```hcl
resource "avi_fileobject" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the file object. Field introduced in 20.1.1.
* `type` - (Required) Type of the file. Enum options - OTHER_FILE_TYPES, IP_REPUTATION, GEO_DB, TECH_SUPPORT, HSMPACKAGES, IPAMDNSSCRIPTS, CONTROLLER_IMAGE. Field introduced in 20.1.1. Allowed in basic(allowed values- other_file_types) edition, essentials(allowed values- other_file_types) edition, enterprise edition.
* `checksum` - (Optional) Sha1 checksum of the file. Field introduced in 20.1.1.
* `compressed` - (Optional) This field indicates whether the file is gzip-compressed. Field introduced in 20.1.1.
* `created` - (Optional) Timestamp of creation for the file. Field introduced in 20.1.1.
* `description` - (Optional) Description of the file. Field introduced in 20.1.1.
* `expires_at` - (Optional) Timestamp when the file will be no longer needed and can be removed by the system. If this is set, a garbage collector process will try to remove the file after this time. Field introduced in 20.1.1.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.1.
* `path` - (Optional) Path to the file. Field introduced in 20.1.1.
* `read_only` - (Optional) Enforce read-only on the file. Field introduced in 20.1.1.
* `restrict_download` - (Optional) Flag to allow/restrict download of the file. Field introduced in 20.1.1.
* `size` - (Optional) Size of the file. Field introduced in 20.1.1.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `version` - (Optional) Version of the file. Field introduced in 20.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the file. Field introduced in 20.1.1.

