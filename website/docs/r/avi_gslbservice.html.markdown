---
layout: "avi"
page_title: "Avi: avi_gslbservice"
sidebar_current: "docs-avi-resource-gslbservice"
description: |-
  Creates and manages Avi GslbService.
---

# avi_gslbservice

The GslbService resource allows the creation and management of Avi GslbService

## Example Usage

```hcl
resource "GslbService" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `application_persistence_profile_ref` - (Optional ) argument_description.
        * `controller_health_status_enabled` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `domain_names` - (Optional ) argument_description.
        * `down_response` - (Optional ) argument_description.
        * `enabled` - (Optional ) argument_description.
        * `groups` - (Optional ) argument_description.
        * `health_monitor_refs` - (Optional ) argument_description.
        * `health_monitor_scope` - (Optional ) argument_description.
        * `is_federated` - (Optional ) argument_description.
        * `min_members` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `num_dns_ip` - (Optional ) argument_description.
        * `pool_algorithm` - (Optional ) argument_description.
        * `site_persistence_enabled` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `ttl` - (Optional ) argument_description.
        * `use_edns_client_subnet` - (Optional ) argument_description.
            * `wildcard_match` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                                * `uuid` - argument_description.
        
