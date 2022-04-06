<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_certificatemanagementprofile"
sidebar_current: "docs-avi-datasource-certificatemanagementprofile"
description: |-
  Get information of Avi CertificateManagementProfile.
---

# avi_certificatemanagementprofile

This data source is used to to get avi_certificatemanagementprofile objects.

## Example Usage

```hcl
data "avi_certificatemanagementprofile" "foo_certificatemanagementprofile" {
    uuid = "certificatemanagementprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CertificateManagementProfile by name.
* `uuid` - (Optional) Search CertificateManagementProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.6. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Name of the pki profile. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `run_script_ref` - Alert script config object for certificate management profile. It is a reference to an object of type alertscriptconfig. Field introduced in 20.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `script_params` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

