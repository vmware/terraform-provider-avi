---
layout: "avi"
page_title: "Avi: avi_clusterclouddetails"
sidebar_current: "docs-avi-resource-clusterclouddetails"
description: |-
Creates and manages Avi ClusterCloudDetails.
---

# avi_clusterclouddetails

The ClusterCloudDetails resource allows the creation and management of Avi ClusterCloudDetails

## Example Usage

```hcl
resource "ClusterCloudDetails" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `azure_info` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                * `uuid` - argument_description.
    
