<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_serviceauthprofile"
sidebar_current: "docs-avi-datasource-serviceauthprofile"
description: |-
  Get information of Avi ServiceAuthProfile.
---

# avi_serviceauthprofile

This data source is used to to get avi_serviceauthprofile objects.

## Example Usage

```hcl
data "avi_serviceauthprofile" "foo_serviceauthprofile" {
    uuid = "serviceauthprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ServiceAuthProfile by name.
* `uuid` - (Optional) Search ServiceAuthProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 30.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Description for the service auth profile. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name of the service auth profile. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `service_oauth_profile` - Oauth profile - common endpoint information for service authentication. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant ref for the service auth profile. It is a reference to an object of type tenant. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `type` - Type of the service auth profile. Enum options - SERVICE_AUTH_OAUTH. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid of the service auth profile. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

