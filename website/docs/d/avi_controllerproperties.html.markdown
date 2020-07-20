############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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
data "avi_controllerproperties" "foo_controllerproperties" {
    uuid = "controllerproperties-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ControllerProperties by name.
* `uuid` - (Optional) Search ControllerProperties by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allow_admin_network_updates` - Allow non-admin tenants to update admin vrfcontext and network objects.
* `allow_ip_forwarding` - Field introduced in 17.1.1.
* `allow_unauthenticated_apis` - Allow unauthenticated access for special apis.
* `allow_unauthenticated_nodes` - Boolean flag to set allow_unauthenticated_nodes.
* `api_idle_timeout` - Allowed values are 0-1440.
* `api_perf_logging_threshold` - Threshold to log request timing in portal_performance.log and server-timing response header.
* `appviewx_compat_mode` - Export configuration in appviewx compatibility mode.
* `attach_ip_retry_interval` - Placeholder for description of property attach_ip_retry_interval of obj type controllerproperties field type integer  type int.
* `attach_ip_retry_limit` - Placeholder for description of property attach_ip_retry_limit of obj type controllerproperties field type integer  type int.
* `bm_use_ansible` - Use ansible for se creation in baremetal.
* `cleanup_expired_authtoken_timeout_period` - Period for auth token cleanup job.
* `cleanup_sessions_timeout_period` - Period for sessions cleanup job.
* `cloud_reconcile` - Enable/disable periodic reconcile for all the clouds.
* `cluster_ip_gratuitous_arp_period` - Period for cluster ip gratuitous arp job.
* `consistency_check_timeout_period` - Period for consistency check job.
* `crashed_se_reboot` - Placeholder for description of property crashed_se_reboot of obj type controllerproperties field type integer  type int.
* `dead_se_detection_timer` - Placeholder for description of property dead_se_detection_timer of obj type controllerproperties field type integer  type int.
* `default_minimum_api_timeout` - Minimum api timeout value.if this value is not 60, it will be the default timeout for all apis that do not have a specific timeout.if an api has a specific timeout but is less than this value, this value will become the new timeout.
* `dns_refresh_period` - Period for refresh pool and gslb dns job.
* `dummy` - Placeholder for description of property dummy of obj type controllerproperties field type integer  type int.
* `edit_system_limits` - Allow editing of system limits.
* `enable_api_sharding` - This setting enables the controller leader to shard api requests to the followers (if any).
* `enable_memory_balancer` - Enable/disable memory balancer.
* `fatal_error_lease_time` - Placeholder for description of property fatal_error_lease_time of obj type controllerproperties field type integer  type int.
* `federated_datastore_cleanup_duration` - Federated datastore will not cleanup diffs unless they are at least this duration in the past.
* `file_object_cleanup_period` - Period for file object cleanup job.
* `max_dead_se_in_grp` - Placeholder for description of property max_dead_se_in_grp of obj type controllerproperties field type integer  type int.
* `max_pcap_per_tenant` - Maximum number of pcap files stored per tenant.
* `max_se_spawn_interval_delay` - Maximum delay possible to add to se_spawn_retry_interval after successive se spawn failure.
* `max_seq_attach_ip_failures` - Maximum number of consecutive attach ip failures that halts vs placement.
* `max_seq_vnic_failures` - Placeholder for description of property max_seq_vnic_failures of obj type controllerproperties field type integer  type int.
* `permission_scoped_shared_admin_networks` - Network and vrfcontext objects from the admin tenant will not be shared to non-admin tenants unless admin permissions are granted.
* `persistence_key_rotate_period` - Period for rotate app persistence keys job.
* `portal_request_burst_limit` - Burst limit on number of incoming requests0 to disable.
* `portal_request_rate_limit` - Maximum average number of requests allowed per second0 to disable.
* `portal_token` - Token used for uploading tech-support to portal.
* `process_locked_useraccounts_timeout_period` - Period for process locked user accounts job.
* `process_pki_profile_timeout_period` - Period for process pki profile job.
* `query_host_fail` - Placeholder for description of property query_host_fail of obj type controllerproperties field type integer  type int.
* `safenet_hsm_version` - Version of the safenet package installed on the controller.
* `se_create_timeout` - Placeholder for description of property se_create_timeout of obj type controllerproperties field type integer  type int.
* `se_failover_attempt_interval` - Interval between attempting failovers to an se.
* `se_from_marketplace` - This setting decides whether se is to be deployed from the cloud marketplace or to be created by the controller.
* `se_offline_del` - Placeholder for description of property se_offline_del of obj type controllerproperties field type integer  type int.
* `se_spawn_retry_interval` - Default retry period before attempting another service engine spawn in se group.
* `se_vnic_cooldown` - Placeholder for description of property se_vnic_cooldown of obj type controllerproperties field type integer  type int.
* `secure_channel_cleanup_timeout` - Period for secure channel cleanup job.
* `secure_channel_controller_token_timeout` - Placeholder for description of property secure_channel_controller_token_timeout of obj type controllerproperties field type integer  type int.
* `secure_channel_se_token_timeout` - Placeholder for description of property secure_channel_se_token_timeout of obj type controllerproperties field type integer  type int.
* `seupgrade_copy_pool_size` - This parameter defines the number of simultaneous se image downloads in a segroup.
* `seupgrade_fabric_pool_size` - Pool size used for all fabric commands during se upgrade.
* `seupgrade_segroup_min_dead_timeout` - Time to wait before marking segroup upgrade as stuck.
* `shared_ssl_certificates` - Ssl certificates in the admin tenant can be used in non-admin tenants.
* `ssl_certificate_expiry_warning_days` - Number of days for ssl certificate expiry warning.
* `unresponsive_se_reboot` - Placeholder for description of property unresponsive_se_reboot of obj type controllerproperties field type integer  type int.
* `upgrade_dns_ttl` - Time to account for dns ttl during upgrade.
* `upgrade_fat_se_lease_time` - Amount of time controller waits for a large-sized se (>=128gb memory) to reconnect after it is rebooted during upgrade.
* `upgrade_lease_time` - Amount of time controller waits for a regular-sized se (<128gb memory) to reconnect after it is rebooted during upgrade.
* `upgrade_se_per_vs_scale_ops_txn_time` - This parameter defines the upper-bound value of the vs scale-in or vs scale-out operation executed in the sescalein and sescale context.
* `uuid` - Unique object identifier of the object.
* `vnic_op_fail_time` - Placeholder for description of property vnic_op_fail_time of obj type controllerproperties field type integer  type int.
* `vs_apic_scaleout_timeout` - Time to wait for the scaled out se to become ready before marking the scaleout done, applies to apic configuration only.
* `vs_awaiting_se_timeout` - Placeholder for description of property vs_awaiting_se_timeout of obj type controllerproperties field type integer  type int.
* `vs_key_rotate_period` - Period for rotate vs keys job.
* `vs_scaleout_ready_check_interval` - Interval for checking scaleout_ready status while controller is waiting for scaleoutready rpc from the service engine.
* `vs_se_attach_ip_fail` - Time to wait before marking attach ip operation on an se as failed.
* `vs_se_bootup_fail` - Placeholder for description of property vs_se_bootup_fail of obj type controllerproperties field type integer  type int.
* `vs_se_create_fail` - Placeholder for description of property vs_se_create_fail of obj type controllerproperties field type integer  type int.
* `vs_se_ping_fail` - Placeholder for description of property vs_se_ping_fail of obj type controllerproperties field type integer  type int.
* `vs_se_vnic_fail` - Placeholder for description of property vs_se_vnic_fail of obj type controllerproperties field type integer  type int.
* `vs_se_vnic_ip_fail` - Placeholder for description of property vs_se_vnic_ip_fail of obj type controllerproperties field type integer  type int.
* `warmstart_se_reconnect_wait_time` - Placeholder for description of property warmstart_se_reconnect_wait_time of obj type controllerproperties field type integer  type int.
* `warmstart_vs_resync_wait_time` - Timeout for warmstart vs resync.

