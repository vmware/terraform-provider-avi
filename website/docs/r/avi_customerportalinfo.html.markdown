---
layout: "avi"
page_title: "Avi: avi_customerportalinfo"
sidebar_current: "docs-avi-resource-customerportalinfo"
description: |-
  Creates and manages Avi CustomerPortalInfo.
---

# avi_customerportalinfo

The CustomerPortalInfo resource allows the creation and management of Avi CustomerPortalInfo

## Example Usage

```hcl
resource "avi_customerportalinfo" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `polling_interval` - (Optional) Time interval in minutes.
* `portal_url` - (Optional) The fqdn or ip address of the customer portal.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 18.2.6.

