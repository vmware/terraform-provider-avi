
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
page_title: "Avi: avi_healthmonitor"
sidebar_current: "docs-avi-resource-healthmonitor"
description: |-
  Creates and manages Avi HealthMonitor.
---

# avi_healthmonitor

The HealthMonitor resource allows the creation and management of Avi HealthMonitor

## Example Usage

```hcl
resource "HealthMonitor" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `description` - (Optional ) argument_description.
        * `dns_monitor` - (Optional ) argument_description.
        * `external_monitor` - (Optional ) argument_description.
        * `failed_checks` - (Optional ) argument_description.
        * `http_monitor` - (Optional ) argument_description.
        * `https_monitor` - (Optional ) argument_description.
        * `is_federated` - (Optional ) argument_description.
        * `monitor_port` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `receive_timeout` - (Optional ) argument_description.
        * `send_interval` - (Optional ) argument_description.
        * `sip_monitor` - (Optional ) argument_description.
        * `successful_checks` - (Optional ) argument_description.
        * `tcp_monitor` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `type` - (Required) argument_description.
        * `udp_monitor` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                        * `uuid` - argument_description.
    
