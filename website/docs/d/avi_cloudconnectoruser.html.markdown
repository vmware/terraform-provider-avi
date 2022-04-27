<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_cloudconnectoruser"
sidebar_current: "docs-avi-datasource-cloudconnectoruser"
description: |-
  Get information of Avi CloudConnectorUser.
---

# avi_cloudconnectoruser

This data source is used to to get avi_cloudconnectoruser objects.

## Example Usage

```hcl
data "avi_cloudconnectoruser" "foo_cloudconnectoruser" {
    uuid = "cloudconnectoruser-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CloudConnectorUser by name.
* `uuid` - (Optional) Search CloudConnectorUser by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `azure_serviceprincipal` - Field introduced in 17.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `azure_userpass` - Field introduced in 17.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `gcp_credentials` - Credentials for google cloud platform. Field introduced in 18.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `nsxt_credentials` - Credentials to talk to nsx-t manager. Field introduced in 20.1.1. Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `oci_credentials` - Credentials for oracle cloud infrastructure. Field introduced in 18.2.1,18.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `password` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `private_key` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `public_key` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tencent_credentials` - Credentials for tencent cloud. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vcenter_credentials` - Credentials to talk to vcenter. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

