<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
data "avi_healthmonitor" "foo_healthmonitor" {
    uuid = "healthmonitor-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search HealthMonitor by name.
* `uuid` - (Optional) Search HealthMonitor by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allow_duplicate_monitors` - By default, multiple instances of the same healthmonitor to the same server are suppressed intelligently. In rare cases, the monitor may have specific constructs that go beyond the server keys (ip, port, etc.) during which such suppression is not desired. Use this knob to allow duplicates. Field introduced in 18.2.8. Allowed in enterprise edition with any value, essentials edition(allowed values- true), basic edition(allowed values- true), enterprise with cloud services edition.
* `authentication` - Authentication information for username/password. Field introduced in 20.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `disable_quickstart` - During addition of a server or healthmonitors or during bootup, avi performs sequential health checks rather than waiting for send-interval to kick in, to mark the server up as soon as possible. This knob may be used to turn this feature off. Field introduced in 18.2.7. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `dns_monitor` - Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `external_monitor` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `failed_checks` - Number of continuous failed health checks before the server is marked down. Allowed values are 1-50. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ftp_monitor` - Health monitor for ftp. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ftps_monitor` - Health monitor for ftps. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `http_monitor` - Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `https_monitor` - Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `imap_monitor` - Health monitor for imap. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `imaps_monitor` - Health monitor for imaps. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `is_federated` - This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `ldap_monitor` - Health monitor for ldap. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ldaps_monitor` - Health monitor for ldaps. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `monitor_port` - Use this port instead of the port defined for the server in the pool. If the monitor succeeds to this port, the load balanced traffic will still be sent to the port of the server defined within the pool. Allowed values are 1-65535. Special values are 0 - use server port. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - A user friendly name for this health monitor. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `pop3_monitor` - Health monitor for pop3. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `pop3s_monitor` - Health monitor for pop3s. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `radius_monitor` - Health monitor for radius. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `receive_timeout` - A valid response from the server is expected within the receive timeout window. This timeout must be less than the send interval. If server status is regularly flapping up and down, consider increasing this value. Allowed values are 1-2400. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `sctp_monitor` - Health monitor for sctp. Field introduced in 22.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `send_interval` - Frequency, in seconds, that monitors are sent to a server. Allowed values are 1-3600. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `sip_monitor` - Health monitor for sip. Field introduced in 17.2.8, 18.1.3, 18.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `smtp_monitor` - Health monitor for smtp. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `smtps_monitor` - Health monitor for smtps. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `successful_checks` - Number of continuous successful health checks before server is marked up. Allowed values are 1-50. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tcp_monitor` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `type` - Type of the health monitor. Enum options - HEALTH_MONITOR_PING, HEALTH_MONITOR_TCP, HEALTH_MONITOR_HTTP, HEALTH_MONITOR_HTTPS, HEALTH_MONITOR_EXTERNAL, HEALTH_MONITOR_UDP, HEALTH_MONITOR_DNS, HEALTH_MONITOR_GSLB, HEALTH_MONITOR_SIP, HEALTH_MONITOR_RADIUS, HEALTH_MONITOR_SMTP, HEALTH_MONITOR_SMTPS, HEALTH_MONITOR_POP3, HEALTH_MONITOR_POP3S, HEALTH_MONITOR_IMAP, HEALTH_MONITOR_IMAPS, HEALTH_MONITOR_FTP, HEALTH_MONITOR_FTPS, HEALTH_MONITOR_LDAP, HEALTH_MONITOR_LDAPS... Allowed in enterprise edition with any value, essentials edition(allowed values- health_monitor_ping,health_monitor_tcp,health_monitor_udp), basic edition(allowed values- health_monitor_ping,health_monitor_tcp,health_monitor_udp,health_monitor_http,health_monitor_https), enterprise with cloud services edition.
* `udp_monitor` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the health monitor. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

