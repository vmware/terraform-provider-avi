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

* `name` - (Required) Name of the file object. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `type` - (Required) Type of the file. Enum options - OTHER_FILE_TYPES, IP_REPUTATION, GEO_DB, TECH_SUPPORT, HSMPACKAGES, IPAMDNSSCRIPTS, CONTROLLER_IMAGE, CRL_DATA, IP_REPUTATION_IPV6, GSLB_GEO_DB, CSRF_JS. Field introduced in 20.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. Allowed in essentials (allowed values- other_file_types), basic (allowed values- other_file_types) edition.
* `checksum` - (Optional) Sha1 checksum of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `child_refs` - (Optional) Avi internal formatted/converted files. It is a reference to an object of type fileobject. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `compressed` - (Optional) This field indicates whether the file is gzip-compressed. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 30.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `created` - (Optional) Timestamp of creation for the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `crl_info` - (Optional) This field contains certificate revocation list metadata. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `description` - (Optional) Description of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `events` - (Optional) List of all fileobject events. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `expires_at` - (Optional) Timestamp when the file will be no longer needed and can be removed by the system. If this is set, a garbage collector process will try to remove the file after this time. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `gslb_geodb_format` - (Optional) This field indicates the file format of gslb geodb file type. Enum options - GSLB_GEODB_FILE_FORMAT_AVI, GSLB_GEODB_FILE_FORMAT_MAXMIND_CITY, GSLB_GEODB_FILE_FORMAT_MAXMIND_CITY_V6, GSLB_GEODB_FILE_FORMAT_MAXMIND_CITY_V4_AND_V6, GSLB_GEODB_FILE_FORMAT_AVI_V6, GSLB_GEODB_FILE_FORMAT_AVI_V4_AND_V6. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `has_parent` - (Optional) This field indicates if the the given fileobjecthas a parent fileobject or not. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `path` - (Optional) Path to the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `read_only` - (Optional) Enforce read-only on the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `restrict_download` - (Optional) Flag to allow/restrict download of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `size` - (Optional) Size of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `version` - (Optional) Version of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

