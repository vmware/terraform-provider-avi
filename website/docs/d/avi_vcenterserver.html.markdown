---
layout: "avi"
page_title: "AVI: avi_vcenterserver"
sidebar_current: "docs-avi-datasource-vcenterserver"
description: |-
  Get information of Avi VCenterServer.
---

# avi_vcenterserver

This data source is used to to get avi_vcenterserver objects.

## Example Usage

```hcl
data "avi_vcenterserver" "foo_vcenterserver" {
    uuid = "vcenterserver-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search VCenterServer by name.
* `uuid` - (Optional) Search VCenterServer by uuid.
* `cloud_ref` - (Optional) Search VCenterServer by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_ref` - Vcenter belongs to cloud.
* `content_lib` - Vcenter template to create service engine.
* `name` - Availabilty zone where vcenter list belongs to.
* `tenant_ref` - Vcenter belongs to tenant.
* `uuid` - Vcenter config uuid.
* `vcenter_credentials_ref` - Credentials to access vcenter.
* `vcenter_url` - Vcenter hostname or ip address.

