<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `allow_admin_network_updates` - Allow non-admin tenants to update admin vrfcontext and network objects. Field introduced in 18.2.7.
* `allow_ip_forwarding` - Field introduced in 17.1.1.
* `allow_unauthenticated_apis` - Allow unauthenticated access for special apis.
* `allow_unauthenticated_nodes` - Boolean flag to set allow_unauthenticated_nodes.
* `api_idle_timeout` - Allowed values are 0-1440.
* `api_perf_logging_threshold` - Threshold to log request timing in portal_performance.log and server-timing response header. Any stage taking longer than 1% of the threshold will be included in the server-timing header. Field introduced in 18.1.4, 18.2.1.
* `appviewx_compat_mode` - Export configuration in appviewx compatibility mode. Field introduced in 17.1.1.
* `async_patch_merge_period` - Period for which asynchronous patch requests are queued. Allowed values are 30-120. Special values are 0 - 'deactivated'. Field introduced in 18.2.11.
* `async_patch_request_cleanup_duration` - Duration for which asynchronous patch requests should be kept, after being marked as success or fail. Allowed values are 5-120. Field introduced in 18.2.11.
* `attach_ip_retry_interval` - Placeholder for description of property attach_ip_retry_interval of obj type controllerproperties field type integer  type int.
* `attach_ip_retry_limit` - Placeholder for description of property attach_ip_retry_limit of obj type controllerproperties field type integer  type int.
* `bm_use_ansible` - Use ansible for se creation in baremetal. Field introduced in 17.2.2.
* `cleanup_expired_authtoken_timeout_period` - Period for auth token cleanup job. Field introduced in 18.1.1.
* `cleanup_sessions_timeout_period` - Period for sessions cleanup job. Field introduced in 18.1.1.
* `cloud_reconcile` - Enable/disable periodic reconcile for all the clouds. Field introduced in 17.2.14,18.1.5,18.2.1.
* `cluster_ip_gratuitous_arp_period` - Period for cluster ip gratuitous arp job.
* `consistency_check_timeout_period` - Period for consistency check job. Field introduced in 18.1.1.
* `crashed_se_reboot` - Placeholder for description of property crashed_se_reboot of obj type controllerproperties field type integer  type int.
* `dead_se_detection_timer` - Placeholder for description of property dead_se_detection_timer of obj type controllerproperties field type integer  type int.
* `default_minimum_api_timeout` - Minimum api timeout value.if this value is not 60, it will be the default timeout for all apis that do not have a specific timeout.if an api has a specific timeout but is less than this value, this value will become the new timeout. Allowed values are 60-3600. Field introduced in 18.2.6.
* `dns_refresh_period` - Period for refresh pool and gslb dns job.
* `dummy` - Placeholder for description of property dummy of obj type controllerproperties field type integer  type int.
* `enable_api_sharding` - This setting enables the controller leader to shard api requests to the followers (if any). Field introduced in 18.1.5, 18.2.1.
* `enable_memory_balancer` - Enable/disable memory balancer. Field introduced in 17.2.8.
* `fatal_error_lease_time` - Placeholder for description of property fatal_error_lease_time of obj type controllerproperties field type integer  type int.
* `max_dead_se_in_grp` - Placeholder for description of property max_dead_se_in_grp of obj type controllerproperties field type integer  type int.
* `max_pcap_per_tenant` - Maximum number of pcap files stored per tenant.
* `max_seq_attach_ip_failures` - Maximum number of consecutive attach ip failures that halts vs placement. Field introduced in 17.2.2.
* `max_seq_vnic_failures` - Placeholder for description of property max_seq_vnic_failures of obj type controllerproperties field type integer  type int.
* `permission_scoped_shared_admin_networks` - Network and vrfcontext objects from the admin tenant will not be shared to non-admin tenants unless admin permissions are granted. Field introduced in 18.2.7.
* `persistence_key_rotate_period` - Period for rotate app persistence keys job. Allowed values are 1-1051200. Special values are 0 - 'disabled'.
* `portal_token` - Token used for uploading tech-support to portal. Field introduced in 16.4.6,17.1.2.
* `process_locked_useraccounts_timeout_period` - Period for process locked user accounts job. Field introduced in 18.1.1.
* `process_pki_profile_timeout_period` - Period for process pki profile job. Field introduced in 18.1.1.
* `query_host_fail` - Placeholder for description of property query_host_fail of obj type controllerproperties field type integer  type int.
* `safenet_hsm_version` - Version of the safenet package installed on the controller. Field introduced in 16.5.2,17.2.3.
* `se_create_timeout` - Placeholder for description of property se_create_timeout of obj type controllerproperties field type integer  type int.
* `se_failover_attempt_interval` - Interval between attempting failovers to an se.
* `se_from_marketplace` - This setting decides whether se is to be deployed from the cloud marketplace or to be created by the controller. The setting is applicable only when byol license is selected. Enum options - MARKETPLACE, IMAGE. Field introduced in 18.1.4, 18.2.1.
* `se_offline_del` - Placeholder for description of property se_offline_del of obj type controllerproperties field type integer  type int.
* `se_vnic_cooldown` - Placeholder for description of property se_vnic_cooldown of obj type controllerproperties field type integer  type int.
* `secure_channel_cleanup_timeout` - Period for secure channel cleanup job.
* `secure_channel_controller_token_timeout` - Placeholder for description of property secure_channel_controller_token_timeout of obj type controllerproperties field type integer  type int.
* `secure_channel_se_token_timeout` - Placeholder for description of property secure_channel_se_token_timeout of obj type controllerproperties field type integer  type int.
* `seupgrade_copy_pool_size` - This parameter defines the number of simultaneous se image downloads in a segroup. It is used to pace the se downloads so that controller network/cpu bandwidth is a bounded operation. A value of 0 will disable the pacing scheme and all the se(s) in the segroup will attempt to download the image. Field introduced in 18.2.6.
* `seupgrade_fabric_pool_size` - Pool size used for all fabric commands during se upgrade.
* `seupgrade_segroup_min_dead_timeout` - Time to wait before marking segroup upgrade as stuck.
* `shared_ssl_certificates` - Ssl certificates in the admin tenant can be used in non-admin tenants. Field introduced in 18.2.5.
* `ssl_certificate_expiry_warning_days` - Number of days for ssl certificate expiry warning.
* `unresponsive_se_reboot` - Placeholder for description of property unresponsive_se_reboot of obj type controllerproperties field type integer  type int.
* `upgrade_dns_ttl` - Time to account for dns ttl during upgrade. This is in addition to vs_scalein_timeout_for_upgrade in se_group. Field introduced in 17.1.1.
* `upgrade_fat_se_lease_time` - Amount of time controller waits for a large-sized se (>=128gb memory) to reconnect after it is rebooted during upgrade. Field introduced in 18.2.10.
* `upgrade_lease_time` - Amount of time controller waits for a regular-sized se (<128gb memory) to reconnect after it is rebooted during upgrade. Starting 18.2.10/20.1.1, the default time has increased from 360 seconds to 600 seconds.
* `upgrade_se_per_vs_scale_ops_txn_time` - This parameter defines the upper-bound value of the vs scale-in or vs scale-out operation executed in the sescalein and sescale context. User can tweak this parameter to a higher value if the segroup gets suspended due to sescalein or sescaleout timeout failure typically associated with high number of vs(es) scaled out. Field introduced in 18.2.10.
* `uuid` - Unique object identifier of the object.
* `vnic_op_fail_time` - Placeholder for description of property vnic_op_fail_time of obj type controllerproperties field type integer  type int.
* `vs_apic_scaleout_timeout` - Time to wait for the scaled out se to become ready before marking the scaleout done, applies to apic configuration only.
* `vs_awaiting_se_timeout` - Placeholder for description of property vs_awaiting_se_timeout of obj type controllerproperties field type integer  type int.
* `vs_key_rotate_period` - Period for rotate vs keys job. Allowed values are 1-1051200. Special values are 0 - 'disabled'.
* `vs_scaleout_ready_check_interval` - Interval for checking scaleout_ready status while controller is waiting for scaleoutready rpc from the service engine. Field introduced in 18.2.2.
* `vs_se_attach_ip_fail` - Time to wait before marking attach ip operation on an se as failed. Field introduced in 17.2.2.
* `vs_se_bootup_fail` - Placeholder for description of property vs_se_bootup_fail of obj type controllerproperties field type integer  type int.
* `vs_se_create_fail` - Placeholder for description of property vs_se_create_fail of obj type controllerproperties field type integer  type int.
* `vs_se_ping_fail` - Placeholder for description of property vs_se_ping_fail of obj type controllerproperties field type integer  type int.
* `vs_se_vnic_fail` - Placeholder for description of property vs_se_vnic_fail of obj type controllerproperties field type integer  type int.
* `vs_se_vnic_ip_fail` - Placeholder for description of property vs_se_vnic_ip_fail of obj type controllerproperties field type integer  type int.
* `warmstart_se_reconnect_wait_time` - Placeholder for description of property warmstart_se_reconnect_wait_time of obj type controllerproperties field type integer  type int.
* `warmstart_vs_resync_wait_time` - Timeout for warmstart vs resync. Field introduced in 18.1.4, 18.2.1.

