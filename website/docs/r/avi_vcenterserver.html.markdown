############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_vcenterserver"
sidebar_current: "docs-avi-resource-vcenterserver"
description: |-
  Creates and manages Avi VCenterServer.
---

# avi_vcenterserver

The VCenterServer resource allows the creation and management of Avi VCenterServer

## Example Usage

```hcl
resource "avi_vcenterserver" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `content_lib` - (Required) Vcenter template to create service engine. Field introduced in 20.1.1.
* `name` - (Required) Availabilty zone where vcenter list belongs to. Field introduced in 20.1.1.
* `vcenter_credentials_ref` - (Required) Credentials to access vcenter. It is a reference to an object of type cloudconnectoruser. Field introduced in 20.1.1.
* `vcenter_url` - (Required) Vcenter hostname or ip address. Field introduced in 20.1.1.
* `cloud_ref` - (Optional) Vcenter belongs to cloud. It is a reference to an object of type cloud. Field introduced in 20.1.1.
* `tenant_ref` - (Optional) Vcenter belongs to tenant. It is a reference to an object of type tenant. Field introduced in 20.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Vcenter config uuid. Field introduced in 20.1.1.

