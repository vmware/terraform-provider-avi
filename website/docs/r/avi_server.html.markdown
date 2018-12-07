---
layout: "avi"
page_title: "Avi: avi_server"
sidebar_current: "docs-avi-resource-server"
description: |-
  Creates and manages Avi Server.
---

# avi_server

The Server resource allows the creation and management of Avi Server

## Example Usage

```hcl
resource "avi_server" "foo" {
    pool_ref = "/api/pool/Pool-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    ip = "10.0.0.3"
}
```

## Argument Reference

The following arguments are supported:

    * `pool_ref` - (Required) argument_description.
    * `ip` - (Required) argument_description.
    * `port` - (Optional ) argument_description.
    * `type` - (Optional ) argument_description.
    * `autoscaling_group_name` - (Optional ) argument_description.
    * `description` - (Optional ) argument_description.
    * `enabled` - (Optional ) argument_description.
    * `external_orchestration_id` - (Optional ) argument_description.
    * `external_uuid` - (Optional ) argument_description.
    * `hostname` - (Optional ) argument_description.
    * `location` - (Optional ) argument_description.
    * `nw_ref` - (Optional ) argument_description.
    * `prst_hdr_val` - (Optional ) argument_description.
    * `rewrite_host_header` - (Optional ) argument_description.
    * `vm_ref` - (Optional ) argument_description.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                                                                                                                                                        * `uuid` - argument_description.
