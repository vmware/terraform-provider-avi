
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
###

---
layout: "avi"
page_title: "Avi: avi_vsvip"
sidebar_current: "docs-avi-resource-vsvip"
description: |-
  Creates and manages Avi VsVip.
---

# avi_vsvip

The VsVip resource allows the creation and management of Avi VsVip

## Example Usage

```hcl
resource "VsVip" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `cloud_ref` - (Optional ) argument_description.
        * `dns_info` - (Optional ) argument_description.
        * `east_west_placement` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
            * `vip` - (Optional ) argument_description.
        * `vrf_context_ref` - (Optional ) argument_description.
        * `vsvip_cloud_config_cksum` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                        * `uuid` - argument_description.
                
