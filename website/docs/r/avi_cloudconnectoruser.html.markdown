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

* `name` - (Required) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `azure_serviceprincipal` - (Optional) Field introduced in 17.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `azure_userpass` - (Optional) Field introduced in 17.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `gcp_credentials` - (Optional) Credentials for google cloud platform. Field introduced in 18.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `nsxt_credentials` - (Optional) Credentials to talk to nsx-t manager. Field introduced in 20.1.1. Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `oci_credentials` - (Optional) Credentials for oracle cloud infrastructure. Field introduced in 18.2.1,18.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `password` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `private_key` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `public_key` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tencent_credentials` - (Optional) Credentials for tencent cloud. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vcenter_credentials` - (Optional) Credentials to talk to vcenter. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

