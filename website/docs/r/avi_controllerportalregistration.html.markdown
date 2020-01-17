---
layout: "avi"
page_title: "Avi: avi_controllerportalregistration"
sidebar_current: "docs-avi-resource-controllerportalregistration"
description: |-
  Creates and manages Avi ControllerPortalRegistration.
---

# avi_controllerportalregistration

The ControllerPortalRegistration resource allows the creation and management of Avi ControllerPortalRegistration

## Example Usage

```hcl
resource "avi_controllerportalregistration" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `asset` - (Optional) Field introduced in 18.2.6.
* `name` - (Optional) Field introduced in 18.2.6.
* `portal_auth` - (Optional) Field introduced in 18.2.6.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 18.2.6.

