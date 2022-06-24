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

* `allow_admin_network_updates` - Allow non-admin tenants to update admin vrfcontext and network objects. Field introduced in 18.2.7, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `allow_ip_forwarding` - Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `allow_unauthenticated_apis` - Allow unauthenticated access for special apis. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `allow_unauthenticated_nodes` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `api_idle_timeout` - Allowed values are 0-1440. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `api_perf_logging_threshold` - Threshold to log request timing in portal_performance.log and server-timing response header. Any stage taking longer than 1% of the threshold will be included in the server-timing header. Field introduced in 18.1.4, 18.2.1. Unit is milliseconds. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `appviewx_compat_mode` - Export configuration in appviewx compatibility mode. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `async_patch_merge_period` - Period for which asynchronous patch requests are queued. Allowed values are 30-120. Special values are 0 - deactivated. Field introduced in 18.2.11, 20.1.3. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `async_patch_request_cleanup_duration` - Duration for which asynchronous patch requests should be kept, after being marked as success or fail. Allowed values are 5-120. Field introduced in 18.2.11, 20.1.3. Unit is min. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `attach_ip_retry_interval` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `attach_ip_retry_limit` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `bm_use_ansible` - Use ansible for se creation in baremetal. Field introduced in 17.2.2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `check_vsvip_fqdn_syntax` - Enforce vsvip fqdn syntax checks. Field introduced in 20.1.6. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `cleanup_expired_authtoken_timeout_period` - Period for auth token cleanup job. Field introduced in 18.1.1. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cleanup_sessions_timeout_period` - Period for sessions cleanup job. Field introduced in 18.1.1. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cloud_reconcile` - Enable/disable periodic reconcile for all the clouds. Field introduced in 17.2.14,18.1.5,18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cluster_ip_gratuitous_arp_period` - Period for cluster ip gratuitous arp job. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `consistency_check_timeout_period` - Period for consistency check job. Field introduced in 18.1.1. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `controller_resource_info_collection_period` - Periodically collect stats. Field introduced in 20.1.3. Unit is min. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `crashed_se_reboot` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dead_se_detection_timer` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `default_minimum_api_timeout` - Minimum api timeout value.if this value is not 60, it will be the default timeout for all apis that do not have a specific timeout.if an api has a specific timeout but is less than this value, this value will become the new timeout. Allowed values are 60-3600. Field introduced in 18.2.6. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `del_offline_se_after_reboot_delay` - The amount of time the controller will wait before deleting an offline se after it has been rebooted. For unresponsive ses, the total time will be  unresponsive_se_reboot + del_offline_se_after_reboot_delay. For crashed ses, the total time will be crashed_se_reboot + del_offline_se_after_reboot_delay. Field introduced in 20.1.5. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `detach_ip_retry_interval` - Amount of time to wait after last detach ip failure before attempting next detach ip retry. Field introduced in 21.1.3. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `detach_ip_retry_limit` - Maximum number of detach ip retries. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `detach_ip_timeout` - Time to wait before marking detach ip as failed. Field introduced in 21.1.3. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `dns_refresh_period` - Period for refresh pool and gslb dns job. Unit is min. Allowed in enterprise edition with any value, essentials edition(allowed values- 60), basic edition(allowed values- 60), enterprise with cloud services edition.
* `dummy` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `edit_system_limits` - Allow editing of system limits. Keep in mind that these system limits have been carefully selected based on rigorous testing in our testig environments. Modifying these limits could destabilize your cluster. Do this at your own risk!. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_api_sharding` - This setting enables the controller leader to shard api requests to the followers (if any). Field introduced in 18.1.5, 18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_memory_balancer` - Enable/disable memory balancer. Field introduced in 17.2.8. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_per_process_stop` - Enable stopping of individual processes if process cross the given threshold limit, even when the total controller memory usage is belowits threshold limit. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `enable_resmgr_log_cache_print` - Enable printing of cached logs inside resource manager. Used for debugging purposes only. Field introduced in 20.1.6. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `false_positive_learning_config` - False positive learning configuration. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `fatal_error_lease_time` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `federated_datastore_cleanup_duration` - Federated datastore will not cleanup diffs unless they are at least this duration in the past. Field introduced in 20.1.1. Unit is hours. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `file_object_cleanup_period` - Period for file object cleanup job. Field introduced in 20.1.1. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_dead_se_in_grp` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_pcap_per_tenant` - Maximum number of pcap files stored per tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_se_spawn_interval_delay` - Maximum delay possible to add to se_spawn_retry_interval after successive se spawn failure. Field introduced in 20.1.1. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_seq_attach_ip_failures` - Maximum number of consecutive attach ip failures that halts vs placement. Field introduced in 17.2.2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_seq_vnic_failures` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_threads_cc_vip_bg_worker` - Maximum number of threads in threadpool used by cloud connector ccvipbgworker. Allowed values are 1-100. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `permission_scoped_shared_admin_networks` - Network and vrfcontext objects from the admin tenant will not be shared to non-admin tenants unless admin permissions are granted. Field introduced in 18.2.7, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `persistence_key_rotate_period` - Period for rotate app persistence keys job. Allowed values are 1-1051200. Special values are 0 - disabled. Unit is min. Allowed in enterprise edition with any value, essentials edition(allowed values- 0), basic edition(allowed values- 0), enterprise with cloud services edition.
* `portal_request_burst_limit` - Burst limit on number of incoming requests. 0 to disable. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `portal_request_rate_limit` - Maximum average number of requests allowed per second. 0 to disable. Field introduced in 20.1.1. Unit is per_second. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `portal_token` - Token used for uploading tech-support to portal. Field introduced in 16.4.6,17.1.2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `process_locked_useraccounts_timeout_period` - Period for process locked user accounts job. Field introduced in 18.1.1. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `process_pki_profile_timeout_period` - Period for process pki profile job. Field introduced in 18.1.1. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `query_host_fail` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `resmgr_log_caching_period` - Period for each cycle of log caching in resource manager. At the end of each cycle, the in memory cached log history will be cleared. Field introduced in 20.1.5. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `restrict_cloud_read_access` - Restrict read access to cloud. Field introduced in 22.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `safenet_hsm_version` - Version of the safenet package installed on the controller. Field introduced in 16.5.2,17.2.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_create_timeout` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_failover_attempt_interval` - Interval between attempting failovers to an se. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_from_marketplace` - This setting decides whether se is to be deployed from the cloud marketplace or to be created by the controller. The setting is applicable only when byol license is selected. Enum options - MARKETPLACE, IMAGE_SE. Field introduced in 18.1.4, 18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_offline_del` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_spawn_retry_interval` - Default retry period before attempting another service engine spawn in se group. Field introduced in 20.1.1. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_upgrade_flow_cleanup_timeout` - Timeout for flows cleanup by serviceengine during upgrade.internal knob  to be exercised under the surveillance of vmware avi support team. Field introduced in 22.1.1. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_vnic_cooldown` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_vnic_gc_wait_time` - Duration to wait after last vnic addition before proceeding with vnic garbage collection. Used for testing purposes. Field introduced in 20.1.4. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `secure_channel_cleanup_timeout` - Period for secure channel cleanup job. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `secure_channel_controller_token_timeout` - Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `secure_channel_se_token_timeout` - Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `seupgrade_copy_pool_size` - This parameter defines the number of simultaneous se image downloads in a segroup. It is used to pace the se downloads so that controller network/cpu bandwidth is a bounded operation. A value of 0 will disable the pacing scheme and all the se(s) in the segroup will attempt to download the image. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `seupgrade_fabric_pool_size` - Pool size used for all fabric commands during se upgrade. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `seupgrade_segroup_min_dead_timeout` - Time to wait before marking segroup upgrade as stuck. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `shared_ssl_certificates` - Ssl certificates in the admin tenant can be used in non-admin tenants. Field introduced in 18.2.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ssl_certificate_expiry_warning_days` - Number of days for ssl certificate expiry warning. Unit is days. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `unresponsive_se_reboot` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `update_dns_entry_retry_limit` - Number of times to retry a dns entry update/delete operation. Field introduced in 21.1.4. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `update_dns_entry_timeout` - Timeout period for a dns entry update/delete operation. Field introduced in 21.1.4. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `upgrade_dns_ttl` - Time to account for dns ttl during upgrade. This is in addition to vs_scalein_timeout_for_upgrade in se_group. Field introduced in 17.1.1. Unit is sec. Allowed in enterprise edition with any value, essentials edition(allowed values- 5), basic edition(allowed values- 5), enterprise with cloud services edition.
* `upgrade_fat_se_lease_time` - Amount of time controller waits for a large-sized se (>=128gb memory) to reconnect after it is rebooted during upgrade. Field introduced in 18.2.10, 20.1.1. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `upgrade_lease_time` - Amount of time controller waits for a regular-sized se (<128gb memory) to reconnect after it is rebooted during upgrade. Starting 18.2.10/20.1.1, the default time has increased from 360 seconds to 600 seconds. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `upgrade_se_per_vs_scale_ops_txn_time` - This parameter defines the upper-bound value of the vs scale-in or vs scale-out operation executed in the sescalein and sescale context. User can tweak this parameter to a higher value if the segroup gets suspended due to sescalein or sescaleout timeout failure typically associated with high number of vs(es) scaled out. Field introduced in 18.2.10, 20.1.1. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `user_agent_cache_config` - Configuration for user-agent cache used in bot management. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vnic_op_fail_time` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_awaiting_se_timeout` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_key_rotate_period` - Period for rotate vs keys job. Allowed values are 1-1051200. Special values are 0 - disabled. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_scaleout_ready_check_interval` - Interval for checking scaleout_ready status while controller is waiting for scaleoutready rpc from the service engine. Field introduced in 18.2.2. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_se_attach_ip_fail` - Time to wait before marking attach ip operation on an se as failed. Field introduced in 17.2.2. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_se_bootup_fail` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_se_create_fail` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_se_ping_fail` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_se_vnic_fail` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vs_se_vnic_ip_fail` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vsphere_ha_detection_timeout` - Vsphere ha monitor detection timeout. If vsphere_ha_enabled is true and the controller is not able to reach the se, placement will wait for this duration for vsphere_ha_inprogress to be marked true before taking corrective action. Field introduced in 20.1.7, 21.1.3. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vsphere_ha_recovery_timeout` - Vsphere ha monitor recovery timeout. Once vsphere_ha_inprogress is set to true (meaning host failure detected and vsphere ha will recover the service engine), placement will wait for at least this duration for the se to reconnect to the controller before taking corrective action. Field introduced in 20.1.7, 21.1.3. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vsphere_ha_timer_interval` - Vsphere ha monitor timer interval for sending cc_check_se_status to cloud connector. Field introduced in 20.1.7, 21.1.3. Unit is sec. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `warmstart_se_reconnect_wait_time` - Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `warmstart_vs_resync_wait_time` - Timeout for warmstart vs resync. Field introduced in 18.1.4, 18.2.1. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

