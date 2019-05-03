---
layout: "avi"
page_title: "Avi: avi_gslb"
sidebar_current: "docs-avi-resource-gslb"
description: |-
  Creates and manages Avi Gslb.
---

# avi_gslb

The Gslb resource allows the creation and management of Avi Gslb

## Example Usage

```hcl
resource "Gslb" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `async_interval` - (Optional ) argument_description.
        * `clear_on_max_retries` - (Optional ) argument_description.
        * `client_ip_addr_group` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `dns_configs` - (Optional ) argument_description.
        * `error_resync_interval` - (Optional ) argument_description.
        * `is_federated` - (Optional ) argument_description.
        * `leader_cluster_uuid` - (Optional ) argument_description.
        * `maintenance_mode` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `send_interval` - (Optional ) argument_description.
        * `send_interval_prior_to_maintenance_mode` - (Optional ) argument_description.
        * `sites` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `third_party_sites` - (Optional ) argument_description.
            * `view_id` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                * `uuid` - argument_description.
        
