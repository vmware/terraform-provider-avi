/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviControllerProperties() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviControllerPropertiesRead,
		Schema: map[string]*schema.Schema{
			"allow_admin_network_updates": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_ip_forwarding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_unauthenticated_apis": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_unauthenticated_nodes": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"api_idle_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"api_perf_logging_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"appviewx_compat_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"async_patch_merge_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"async_patch_request_cleanup_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"attach_ip_retry_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"attach_ip_retry_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bm_use_ansible": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cleanup_expired_authtoken_timeout_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cleanup_sessions_timeout_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cloud_reconcile": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cluster_ip_gratuitous_arp_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"consistency_check_timeout_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"controller_resource_info_collection_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"crashed_se_reboot": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dead_se_detection_timer": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_minimum_api_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dns_refresh_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dummy": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"edit_system_limits": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_api_sharding": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_memory_balancer": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fatal_error_lease_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"federated_datastore_cleanup_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"file_object_cleanup_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_dead_se_in_grp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_pcap_per_tenant": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_se_spawn_interval_delay": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_seq_attach_ip_failures": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_seq_vnic_failures": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_threads_cc_vip_bg_worker": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"permission_scoped_shared_admin_networks": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"persistence_key_rotate_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"portal_request_burst_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"portal_request_rate_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"portal_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"process_locked_useraccounts_timeout_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"process_pki_profile_timeout_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"query_host_fail": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"resmgr_log_caching_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"safenet_hsm_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_create_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_failover_attempt_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_from_marketplace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_offline_del": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_spawn_retry_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_vnic_cooldown": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_vnic_gc_wait_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secure_channel_cleanup_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secure_channel_controller_token_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secure_channel_se_token_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"seupgrade_copy_pool_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"seupgrade_fabric_pool_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"seupgrade_segroup_min_dead_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"shared_ssl_certificates": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ssl_certificate_expiry_warning_days": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"unresponsive_se_reboot": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upgrade_dns_ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upgrade_fat_se_lease_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upgrade_lease_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upgrade_se_per_vs_scale_ops_txn_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vnic_op_fail_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_apic_scaleout_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_awaiting_se_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_key_rotate_period": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_scaleout_ready_check_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_se_attach_ip_fail": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_se_bootup_fail": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_se_create_fail": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_se_ping_fail": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_se_vnic_fail": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_se_vnic_ip_fail": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"warmstart_se_reconnect_wait_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"warmstart_vs_resync_wait_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}
