// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviControllerProperties() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviControllerPropertiesRead,
		Schema: map[string]*schema.Schema{
			"allow_admin_network_updates": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_ip_forwarding": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_unauthenticated_apis": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_unauthenticated_nodes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_idle_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_perf_logging_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"appviewx_compat_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"async_patch_merge_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"async_patch_request_cleanup_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attach_ip_retry_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attach_ip_retry_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bm_use_ansible": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"check_vsvip_fqdn_syntax": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cleanup_expired_authtoken_timeout_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cleanup_sessions_timeout_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_reconcile": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_ip_gratuitous_arp_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"consistency_check_timeout_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_resource_info_collection_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crashed_se_reboot": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dead_se_detection_timer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_minimum_api_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"del_offline_se_after_reboot_delay": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"detach_ip_retry_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"detach_ip_retry_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"detach_ip_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_refresh_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dummy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"edit_system_limits": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_api_sharding": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_memory_balancer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_per_process_stop": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_resmgr_log_cache_print": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"false_positive_learning_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceFalsePositiveLearningConfigSchema(),
			},
			"fatal_error_lease_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"federated_datastore_cleanup_duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_object_cleanup_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_reference_mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceFileReferenceMappingSchema(),
			},
			"gslb_purge_batch_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gslb_purge_sleep_time_ms": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_vrf_in_networksubnetlist": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_dead_se_in_grp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_pcap_per_tenant": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_se_spawn_interval_delay": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_seq_attach_ip_failures": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_seq_vnic_failures": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_threads_cc_vip_bg_worker": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permission_scoped_shared_admin_networks": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"persistence_key_rotate_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portal_request_burst_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portal_request_rate_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portal_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"postgres_vacuum_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"process_locked_useraccounts_timeout_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"process_pki_profile_timeout_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"query_host_fail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resmgr_log_caching_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"restrict_cloud_read_access": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"safenet_hsm_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_create_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_failover_attempt_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_from_marketplace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_offline_del": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_spawn_retry_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_upgrade_flow_cleanup_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_vnic_cooldown": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_vnic_gc_wait_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secure_channel_cleanup_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secure_channel_controller_token_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secure_channel_se_token_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seupgrade_copy_buffer_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seupgrade_copy_pool_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seupgrade_fabric_pool_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seupgrade_segroup_min_dead_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shared_ssl_certificates": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skopeo_retry_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skopeo_retry_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"soft_min_mem_per_se_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_certificate_expiry_warning_days": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"unresponsive_se_reboot": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_dns_entry_retry_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_dns_entry_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_dns_ttl": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_fat_se_lease_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_lease_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_se_per_vs_scale_ops_txn_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_agent_cache_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceUserAgentCacheConfigSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic_op_fail_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_awaiting_se_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_key_rotate_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_scaleout_ready_check_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_attach_ip_fail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_bootup_fail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_bootup_fail_patch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_create_fail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_ping_fail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_vnic_fail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_vnic_ip_fail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsphere_ha_detection_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsphere_ha_recovery_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsphere_ha_timer_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"warmstart_se_reconnect_wait_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"warmstart_vs_resync_wait_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
