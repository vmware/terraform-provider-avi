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
data "CertificateManagementProfile" "foo_CertificateManagementProfile" {
    uuid = "CertificateManagementProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CertificateManagementProfile by name.
* `uuid` - (Optional) Search CertificateManagementProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the pki profile.
* `script_params` - General description.
* `script_path` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.

