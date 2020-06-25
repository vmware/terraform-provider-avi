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

* `cloud_ref` - (Optional) Vcenter belongs to cloud.
* `content_lib` - (Optional) Vcenter template to create service engine.
* `name` - (Optional) Availabilty zone where vcenter list belongs to.
* `tenant_ref` - (Optional) Vcenter belongs to tenant.
* `vcenter_credentials_ref` - (Optional) Credentials to access vcenter.
* `vcenter_url` - (Optional) Vcenter hostname or ip address.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Vcenter config uuid.

