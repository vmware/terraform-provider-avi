---
layout: "avi"
page_title: "Avi: avi_federationcheckpoint"
sidebar_current: "docs-avi-resource-federationcheckpoint"
description: |-
  Creates and manages Avi FederationCheckpoint.
---

# avi_federationcheckpoint

The FederationCheckpoint resource allows the creation and management of Avi FederationCheckpoint

## Example Usage

```hcl
resource "avi_federationcheckpoint" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `date` - (Optional) Date when the checkpoint was created.
* `description` - (Optional) Description for this checkpoint.
* `is_federated` - (Optional) This field describes the object's replication scope.
* `name` - (Optional) Name of the checkpoint.
* `tenant_ref` - (Optional) Tenant that this object belongs to.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the checkpoint.

