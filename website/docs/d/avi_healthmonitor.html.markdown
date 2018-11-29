---
layout: "avi"
page_title: "AVI: avi_healthmonitor"
sidebar_current: "docs-avi-datasource-healthmonitor"
description: |-
  Get information of Avi HealthMonitor.
---

# avi_healthmonitor

This data source is used to to get avi_healthmonitor objects.

## Example Usage

```hcl
data "HealthMonitor" "foo_HealthMonitor" {
    uuid = "HealthMonitor-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search HealthMonitor by name.
* `uuid` - (Optional) Search HealthMonitor by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - General description.
* `dns_monitor` - General description.
* `external_monitor` - General description.
* `failed_checks` - Number of continuous failed health checks before the server is marked down.
* `http_monitor` - General description.
* `https_monitor` - General description.
* `is_federated` - This field describes the object's replication scope.
* `monitor_port` - Use this port instead of the port defined for the server in the pool.
* `name` - A user friendly name for this health monitor.
* `receive_timeout` - A valid response from the server is expected within the receive timeout window.
* `send_interval` - Frequency, in seconds, that monitors are sent to a server.
* `sip_monitor` - Health monitor for sip.
* `successful_checks` - Number of continuous successful health checks before server is marked up.
* `tcp_monitor` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `type` - Type of the health monitor.
* `udp_monitor` - General description.
* `uuid` - Uuid of the health monitor.
