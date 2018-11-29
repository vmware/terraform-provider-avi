---
layout: "avi"
page_title: "Avi: avi_vsdatascriptset"
sidebar_current: "docs-avi-resource-vsdatascriptset"
description: |-
  Creates and manages Avi VSDataScriptSet.
---

# avi_vsdatascriptset

The VSDataScriptSet resource allows the creation and management of Avi VSDataScriptSet

## Example Usage

```hcl
resource "VSDataScriptSet" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `created_by` - (Optional ) argument_description.
        * `datascript` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `ipgroup_refs` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `pool_group_refs` - (Optional ) argument_description.
        * `pool_refs` - (Optional ) argument_description.
        * `string_group_refs` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                        * `uuid` - argument_description.
