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

* `api_spec_detail` - Api specification details extracted from the file, populated for open_api_spec type only. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `checksum` - Sha1 checksum of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `child_refs` - Avi internal formatted/converted files. It is a reference to an object of type fileobject. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `completed_events` - Number of processing events that have completed. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `compressed` - This field indicates whether the file is gzip-compressed. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 30.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `created` - Timestamp of creation for the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `crl_info` - This field contains crl metadata. Field introduced in 30.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Description of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `duration` - Time taken to complete the operation in seconds. Field introduced in 32.2.1. Unit is sec. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `end_time` - End time of the file object processing operation. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `events` - List of all fileobject events. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `expires_at` - Timestamp when the crl contents are no longer valid and hence crl-file will be no longer needed and can be removed by the system. If this is set, a garbage collector process shall remove the crl-file after this time. This field is applicable in the crl context. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `gslb_geodb_format` - This field indicates the file format(avi/maxmind and v4/v6/v4-v6) of gslb geodb file type. Enum options - GSLB_GEODB_FILE_FORMAT_AVI, GSLB_GEODB_FILE_FORMAT_MAXMIND_CITY, GSLB_GEODB_FILE_FORMAT_MAXMIND_CITY_V6, GSLB_GEODB_FILE_FORMAT_MAXMIND_CITY_V4_AND_V6, GSLB_GEODB_FILE_FORMAT_AVI_V6, GSLB_GEODB_FILE_FORMAT_AVI_V4_AND_V6. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `has_parent` - This field indicates if the the given fileobjecthas a parent fileobject or not. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `history` - File object processing events history for the version specified. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the gslb federation. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the file object. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `path` - Path to the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `progress` - Percentage of completed events. Allowed values are 0-100. Field introduced in 32.2.1. Unit is percent. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `read_only` - Enforce read-only on the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `restrict_download` - Flag to allow/restrict download of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `size` - Size of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `start_time` - Start time of the file object processing operation. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `state` - State of the file object. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `task_events` - File object processing events for the version specified. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `total_events` - Total number of processing events for this file object. Field introduced in 32.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `type` - Type of the file. Enum options - OTHER_FILE_TYPES, IP_REPUTATION, GEO_DB, TECH_SUPPORT, HSMPACKAGES, IPAMDNSSCRIPTS, CONTROLLER_IMAGE, CRL_DATA, IP_REPUTATION_IPV6, GSLB_GEO_DB, CSRF_JS, KNOWN_HOSTS, OPEN_API_SPEC. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `version` - Version of the file. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

