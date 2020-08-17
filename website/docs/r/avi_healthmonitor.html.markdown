############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
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
resource "avi_healthmonitor" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user friendly name for this health monitor.
* `type` - (Required) Type of the health monitor.
* `allow_duplicate_monitors` - (Optional) By default, multiple instances of the same healthmonitor to the same server are suppressed intelligently.
* `description` - (Optional) User defined description for the object.
* `disable_quickstart` - (Optional) During addition of a server or healthmonitors or during bootup, avi performs sequential health checks rather than waiting for send-interval to kick in, to mark the server up as soon as possible.
* `dns_monitor` - (Optional) Dict settings for healthmonitor.
* `external_monitor` - (Optional) Dict settings for healthmonitor.
* `failed_checks` - (Optional) Number of continuous failed health checks before the server is marked down.
* `http_monitor` - (Optional) Dict settings for healthmonitor.
* `https_monitor` - (Optional) Dict settings for healthmonitor.
* `is_federated` - (Optional) This field describes the object's replication scope.
* `monitor_port` - (Optional) Use this port instead of the port defined for the server in the pool.
* `radius_monitor` - (Optional) Health monitor for radius.
* `receive_timeout` - (Optional) A valid response from the server is expected within the receive timeout window.
* `send_interval` - (Optional) Frequency, in seconds, that monitors are sent to a server.
* `sip_monitor` - (Optional) Health monitor for sip.
* `successful_checks` - (Optional) Number of continuous successful health checks before server is marked up.
* `tcp_monitor` - (Optional) Dict settings for healthmonitor.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `udp_monitor` - (Optional) Dict settings for healthmonitor.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the health monitor.

