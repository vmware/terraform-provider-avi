---
layout: "avi"
page_title: "Avi: avi_gslbgeodbprofile"
sidebar_current: "docs-avi-resource-gslbgeodbprofile"
description: |-
  Creates and manages Avi GslbGeoDbProfile.
---

# avi_gslbgeodbprofile

The GslbGeoDbProfile resource allows the creation and management of Avi GslbGeoDbProfile

## Example Usage

```hcl
resource "avi_gslbgeodbprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user-friendly name for the geodb profile.
* `description` - (Optional) Field introduced in 17.1.1.
* `entries` - (Optional) List of geodb entries.
* `is_federated` - (Optional) This field indicates that this object is replicated across gslb federation.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the geodb profile.

