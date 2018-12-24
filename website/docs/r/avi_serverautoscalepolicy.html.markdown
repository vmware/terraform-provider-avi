---
layout: "avi"
page_title: "Avi: avi_serverautoscalepolicy"
sidebar_current: "docs-avi-resource-serverautoscalepolicy"
description: |-
  Creates and manages Avi ServerAutoScalePolicy.
---

# avi_serverautoscalepolicy

The ServerAutoScalePolicy resource allows the creation and management of Avi ServerAutoScalePolicy

## Example Usage

```hcl
resource "ServerAutoScalePolicy" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `description` - (Optional ) argument_description.
        * `intelligent_autoscale` - (Optional ) argument_description.
        * `intelligent_scalein_margin` - (Optional ) argument_description.
        * `intelligent_scaleout_margin` - (Optional ) argument_description.
        * `max_scalein_adjustment_step` - (Optional ) argument_description.
        * `max_scaleout_adjustment_step` - (Optional ) argument_description.
        * `max_size` - (Optional ) argument_description.
        * `min_size` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `scalein_alertconfig_refs` - (Optional ) argument_description.
        * `scalein_cooldown` - (Optional ) argument_description.
        * `scaleout_alertconfig_refs` - (Optional ) argument_description.
        * `scaleout_cooldown` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `use_predicted_load` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                * `uuid` - argument_description.
    
