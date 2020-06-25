---
layout: "avi"
page_title: "Avi: avi_licenseledgerdetails"
sidebar_current: "docs-avi-resource-licenseledgerdetails"
description: |-
  Creates and manages Avi LicenseLedgerDetails.
---

# avi_licenseledgerdetails

The LicenseLedgerDetails resource allows the creation and management of Avi LicenseLedgerDetails

## Example Usage

```hcl
resource "avi_licenseledgerdetails" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `escrow_infos` - (Optional) Maintain information about reservation against cookie.
* `se_infos` - (Optional) Maintain information about consumed licenses against se_uuid.
* `tier_usages` - (Optional) License usage per tier.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid for reference.

