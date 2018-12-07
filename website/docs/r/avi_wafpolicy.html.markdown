---
layout: "avi"
page_title: "Avi: avi_wafpolicy"
sidebar_current: "docs-avi-resource-wafpolicy"
description: |-
  Creates and manages Avi WafPolicy.
---

# avi_wafpolicy

The WafPolicy resource allows the creation and management of Avi WafPolicy

## Example Usage

```hcl
resource "WafPolicy" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `created_by` - (Optional ) argument_description.
        * `crs_groups` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `failure_mode` - (Optional ) argument_description.
        * `learning` - (Optional ) argument_description.
        * `mode` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `paranoia_level` - (Optional ) argument_description.
        * `post_crs_groups` - (Optional ) argument_description.
        * `pre_crs_groups` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
            * `waf_crs_ref` - (Optional ) argument_description.
        * `waf_profile_ref` - (Optional ) argument_description.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                * `uuid` - argument_description.
