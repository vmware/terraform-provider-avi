<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_cloudconnectoruser"
sidebar_current: "docs-avi-resource-cloudconnectoruser"
description: |-
  Creates and manages Avi CloudConnectorUser.
---

# avi_cloudconnectoruser

The CloudConnectorUser resource allows the creation and management of Avi CloudConnectorUser

## Example Usage

```hcl
resource "avi_cloudconnectoruser" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `azure_serviceprincipal` - (Optional) Field introduced in 17.2.1.
* `azure_userpass` - (Optional) Field introduced in 17.2.1.
* `gcp_credentials` - (Optional) Credentials for google cloud platform. Field introduced in 18.2.1.
* `oci_credentials` - (Optional) Credentials for oracle cloud infrastructure. Field introduced in 18.2.1,18.1.3.
* `password` - (Optional) Placeholder for description of property password of obj type cloudconnectoruser field type string  type str.
* `private_key` - (Optional) Placeholder for description of property private_key of obj type cloudconnectoruser field type string  type str.
* `public_key` - (Optional) Placeholder for description of property public_key of obj type cloudconnectoruser field type string  type str.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `tencent_credentials` - (Optional) Credentials for tencent cloud. Field introduced in 18.2.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

