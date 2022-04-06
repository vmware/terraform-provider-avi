<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_gslbgeodbprofile"
sidebar_current: "docs-avi-datasource-gslbgeodbprofile"
description: |-
  Get information of Avi GslbGeoDbProfile.
---

# avi_gslbgeodbprofile

This data source is used to to get avi_gslbgeodbprofile objects.

## Example Usage

```hcl
data "avi_gslbgeodbprofile" "foo_gslbgeodbprofile" {
    uuid = "gslbgeodbprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GslbGeoDbProfile by name.
* `uuid` - (Optional) Search GslbGeoDbProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `entries` - List of geodb entries. An entry can either be a geodb file or an ip address group with geo properties. Field introduced in 17.1.1. Minimum of 1 items required. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `is_federated` - This field indicates that this object is replicated across gslb federation. Field introduced in 17.1.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - A user-friendly name for the geodb profile. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Uuid of the geodb profile. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

