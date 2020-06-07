---
layout: "avi"
page_title: "Avi: avi_wafapplicationsignatureprovider"
sidebar_current: "docs-avi-resource-wafapplicationsignatureprovider"
description: |-
  Creates and manages Avi WafApplicationSignatureProvider.
---

# avi_wafapplicationsignatureprovider

The WafApplicationSignatureProvider resource allows the creation and management of Avi WafApplicationSignatureProvider

## Example Usage

```hcl
resource "avi_wafapplicationsignatureprovider" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Name of application specific ruleset provider.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 20.1.1.

