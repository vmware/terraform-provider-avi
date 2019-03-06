
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
page_title: "Avi: avi_alertconfig"
sidebar_current: "docs-avi-resource-alertconfig"
description: |-
  Creates and manages Avi AlertConfig.
---

# avi_alertconfig

The AlertConfig resource allows the creation and management of Avi AlertConfig

## Example Usage

```hcl
resource "AlertConfig" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `action_group_ref` - (Optional ) argument_description.
        * `alert_rule` - (Required) argument_description.
        * `autoscale_alert` - (Optional ) argument_description.
        * `category` - (Required) argument_description.
        * `description` - (Optional ) argument_description.
        * `enabled` - (Optional ) argument_description.
        * `expiry_time` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `obj_uuid` - (Optional ) argument_description.
        * `object_type` - (Optional ) argument_description.
        * `recommendation` - (Optional ) argument_description.
        * `rolling_window` - (Optional ) argument_description.
        * `source` - (Required) argument_description.
        * `summary` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `threshold` - (Optional ) argument_description.
        * `throttle` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                        * `uuid` - argument_description.
    
