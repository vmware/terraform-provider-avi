---
layout: "avi"
page_title: "Avi: avi_trafficcloneprofile"
sidebar_current: "docs-avi-resource-trafficcloneprofile"
description: |-
  Creates and manages Avi TrafficCloneProfile.
---

# avi_trafficcloneprofile

The TrafficCloneProfile resource allows the creation and management of Avi TrafficCloneProfile

## Example Usage

```hcl
resource "avi_trafficcloneprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the traffic clone profile.
* `clone_servers` - (Optional) Field introduced in 17.1.1.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `preserve_client_ip` - (Optional) Specifies if client ip needs to be preserved to clone destination.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the traffic clone profile.

