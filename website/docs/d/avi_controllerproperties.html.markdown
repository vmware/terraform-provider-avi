---
layout: "avi"
page_title: "AVI: avi_controllerproperties"
sidebar_current: "docs-avi-datasource-controllerproperties"
description: |-
  Get information of Avi ControllerProperties.
---

# avi_controllerproperties

This data source is used to to get avi_controllerproperties objects.

## Example Usage

```hcl
data "ControllerProperties" "foo_ControllerProperties" {
    uuid = "ControllerProperties-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ControllerProperties by name.
* `uuid` - (Optional) Search ControllerProperties by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allow_ip_forwarding` - Field introduced in 17.1.1.
* `allow_unauthenticated_apis` - Allow unauthenticated access for special apis.
* `allow_unauthenticated_nodes` - General description.
* `api_idle_timeout` - Allowed values are 0-1440.
* `api_perf_logging_threshold` - Threshold to log request timing in portal_performance.log and server-timing response header.
* `appviewx_compat_mode` - Export configuration in appviewx compatibility mode.
* `attach_ip_retry_interval` - General description.
* `attach_ip_retry_limit` - General description.
* `bm_use_ansible` - Use ansible for se creation in baremetal.
* `cleanup_expired_authtoken_timeout_period` - Period for auth token cleanup job.
* `cleanup_sessions_timeout_period` - Period for sessions cleanup job.
* `cluster_ip_gratuitous_arp_period` - Period for cluster ip gratuitous arp job.
* `consistency_check_timeout_period` - Period for consistency check job.
* `crashed_se_reboot` - General description.
* `dead_se_detection_timer` - General description.
* `dns_refresh_period` - Period for refresh pool and gslb dns job.
* `dummy` - General description.
* `enable_memory_balancer` - Enable/disable memory balancer.
* `fatal_error_lease_time` - General description.
* `max_dead_se_in_grp` - General description.
* `max_pcap_per_tenant` - Maximum number of pcap files stored per tenant.
* `max_seq_attach_ip_failures` - Maximum number of consecutive attach ip failures that halts vs placement.
* `max_seq_vnic_failures` - General description.
* `persistence_key_rotate_period` - Period for rotate app persistence keys job.
* `portal_token` - Token used for uploading tech-support to portal.
* `process_locked_useraccounts_timeout_period` - Period for process locked user accounts job.
* `process_pki_profile_timeout_period` - Period for process pki profile job.
* `query_host_fail` - General description.
* `safenet_hsm_version` - Version of the safenet package installed on the controller.
* `se_create_timeout` - General description.
* `se_failover_attempt_interval` - Interval between attempting failovers to an se.
* `se_from_marketplace` - This setting decides whether se is to be deployed from the cloud marketplace or to be created by the controller.
* `se_offline_del` - General description.
* `se_vnic_cooldown` - General description.
* `secure_channel_cleanup_timeout` - Period for secure channel cleanup job.
* `secure_channel_controller_token_timeout` - General description.
* `secure_channel_se_token_timeout` - General description.
* `seupgrade_fabric_pool_size` - Pool size used for all fabric commands during se upgrade.
* `seupgrade_segroup_min_dead_timeout` - Time to wait before marking segroup upgrade as stuck.
* `ssl_certificate_expiry_warning_days` - Number of days for ssl certificate expiry warning.
* `unresponsive_se_reboot` - General description.
* `upgrade_dns_ttl` - Time to account for dns ttl during upgrade.
* `upgrade_lease_time` - General description.
* `uuid` - General description.
* `vnic_op_fail_time` - General description.
* `vs_apic_scaleout_timeout` - Time to wait for the scaled out se to become ready before marking the scaleout done, applies to apic configuration only.
* `vs_awaiting_se_timeout` - General description.
* `vs_key_rotate_period` - Period for rotate vs keys job.
* `vs_se_attach_ip_fail` - Time to wait before marking attach ip operation on an se as failed.
* `vs_se_bootup_fail` - General description.
* `vs_se_create_fail` - General description.
* `vs_se_ping_fail` - General description.
* `vs_se_vnic_fail` - General description.
* `vs_se_vnic_ip_fail` - General description.
* `warmstart_se_reconnect_wait_time` - General description.
* `warmstart_vs_resync_wait_time` - Timeout for warmstart vs resync.
