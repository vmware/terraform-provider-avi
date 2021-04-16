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

* `azure_serviceprincipal` - Field introduced in 17.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `azure_userpass` - Field introduced in 17.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `gcp_credentials` - Credentials for google cloud platform. Field introduced in 18.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the object.
* `nsxt_credentials` - Credentials to talk to nsx-t manager. Field introduced in 20.1.1. Allowed in essentials edition, enterprise edition.
* `oci_credentials` - Credentials for oracle cloud infrastructure. Field introduced in 18.2.1,18.1.3. Allowed in basic edition, essentials edition, enterprise edition.
* `password` - Placeholder for description of property password of obj type cloudconnectoruser field type string  type str.
* `private_key` - Placeholder for description of property private_key of obj type cloudconnectoruser field type string  type str.
* `public_key` - Placeholder for description of property public_key of obj type cloudconnectoruser field type string  type str.
* `tenant_ref` - It is a reference to an object of type tenant.
* `tencent_credentials` - Credentials for tencent cloud. Field introduced in 18.2.3. Allowed in basic edition, essentials edition, enterprise edition.
* `uuid` - Unique object identifier of the object.
* `vcenter_credentials` - Credentials to talk to vcenter. Field introduced in 20.1.1.

