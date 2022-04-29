<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_certificatemanagementprofile"
sidebar_current: "docs-avi-resource-certificatemanagementprofile"
description: |-
  Creates and manages Avi CertificateManagementProfile.
---

# avi_certificatemanagementprofile

The CertificateManagementProfile resource allows the creation and management of Avi CertificateManagementProfile

## Example Usage

```hcl
resource "avi_certificatemanagementprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the pki profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `run_script_ref` - (Required) Alert script config object for certificate management profile. It is a reference to an object of type alertscriptconfig. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.6. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `script_params` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

