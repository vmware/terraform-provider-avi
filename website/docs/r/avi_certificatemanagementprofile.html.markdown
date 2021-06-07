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

* `name` - (Required) Name of the pki profile.
* `run_script_ref` - (Required) Alert script config object for certificate management profile. It is a reference to an object of type alertscriptconfig. Field introduced in 20.1.3.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.6.
* `script_params` - (Optional) List of list.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

