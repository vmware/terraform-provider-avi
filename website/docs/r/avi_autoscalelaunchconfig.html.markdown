---
layout: "avi"
page_title: "Avi: avi_autoscalelaunchconfig"
sidebar_current: "docs-avi-resource-autoscalelaunchconfig"
description: |-
  Creates and manages Avi AutoScaleLaunchConfig.
---

# avi_autoscalelaunchconfig

The AutoScaleLaunchConfig resource allows the creation and management of Avi AutoScaleLaunchConfig

## Example Usage

```hcl
resource "AutoScaleLaunchConfig" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `description` - (Optional ) argument_description.
        * `image_id` - (Optional ) argument_description.
        * `mesos` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `openstack` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `use_external_asg` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                * `uuid` - argument_description.
    
