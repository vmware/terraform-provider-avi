---
layout: "avi"
page_title: "Avi: avi_securitypolicy"
sidebar_current: "docs-avi-resource-securitypolicy"
description: |-
  Creates and manages Avi SecurityPolicy.
---

# avi_securitypolicy

The SecurityPolicy resource allows the creation and management of Avi SecurityPolicy

## Example Usage

```hcl
resource "SecurityPolicy" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `description` - (Optional ) argument_description.
        * `dns_attacks` - (Optional ) argument_description.
        * `dns_policy_index` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `network_security_policy_index` - (Optional ) argument_description.
        * `oper_mode` - (Optional ) argument_description.
        * `tcp_attacks` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `udp_attacks` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                        * `uuid` - argument_description.
    
