---
layout: "avi"
page_title: "Avi: avi_poolgroupdeploymentpolicy"
sidebar_current: "docs-avi-resource-poolgroupdeploymentpolicy"
description: |-
  Creates and manages Avi PoolGroupDeploymentPolicy.
---

# avi_poolgroupdeploymentpolicy

The PoolGroupDeploymentPolicy resource allows the creation and management of Avi PoolGroupDeploymentPolicy

## Example Usage

```hcl
resource "PoolGroupDeploymentPolicy" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `auto_disable_old_prod_pools` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `evaluation_duration` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `rules` - (Optional ) argument_description.
        * `scheme` - (Optional ) argument_description.
        * `target_test_traffic_ratio` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `test_traffic_ratio_rampup` - (Optional ) argument_description.
            * `webhook_ref` - (Optional ) argument_description.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                        * `uuid` - argument_description.
