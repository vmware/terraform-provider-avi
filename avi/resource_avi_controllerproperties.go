// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceControllerPropertiesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_admin_network_updates": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"allow_ip_forwarding": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"allow_unauthenticated_apis": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"allow_unauthenticated_nodes": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"api_idle_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "15",
			ValidateFunc: validateInteger,
		},
		"api_perf_logging_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10000",
			ValidateFunc: validateInteger,
		},
		"appviewx_compat_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"async_patch_merge_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"async_patch_request_cleanup_duration": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"attach_ip_retry_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "360",
			ValidateFunc: validateInteger,
		},
		"attach_ip_retry_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
		},
		"bm_use_ansible": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"check_vsvip_fqdn_syntax": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"cleanup_expired_authtoken_timeout_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"cleanup_sessions_timeout_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"cloud_discovery_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateInteger,
		},
		"cloud_reconcile": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"cloud_reconcile_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateInteger,
		},
		"cluster_ip_gratuitous_arp_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"consistency_check_timeout_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"controller_resource_info_collection_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "525600",
			ValidateFunc: validateInteger,
		},
		"crashed_se_reboot": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "900",
			ValidateFunc: validateInteger,
		},
		"dead_se_detection_timer": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "360",
			ValidateFunc: validateInteger,
		},
		"default_minimum_api_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"del_offline_se_after_reboot_delay": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"detach_ip_retry_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"detach_ip_retry_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
		},
		"detach_ip_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"dns_refresh_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"dummy": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"edit_system_limits": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_api_sharding": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"enable_memory_balancer": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"enable_per_process_stop": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_resmgr_log_cache_print": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"false_positive_learning_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceFalsePositiveLearningConfigSchema(),
		},
		"fatal_error_lease_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "120",
			ValidateFunc: validateInteger,
		},
		"federated_datastore_cleanup_duration": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "120",
			ValidateFunc: validateInteger,
		},
		"file_object_cleanup_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1440",
			ValidateFunc: validateInteger,
		},
		"fileobject_max_file_versions": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"gslb_purge_batch_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1000",
			ValidateFunc: validateInteger,
		},
		"gslb_purge_sleep_time_ms": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "50",
			ValidateFunc: validateInteger,
		},
		"ignore_vrf_in_networksubnetlist": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"max_dead_se_in_grp": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"max_pcap_per_tenant": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
		},
		"max_se_spawn_interval_delay": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1800",
			ValidateFunc: validateInteger,
		},
		"max_seq_attach_ip_failures": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"max_seq_vnic_failures": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"max_threads_cc_vip_bg_worker": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
		},
		"permission_scoped_shared_admin_networks": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"persistence_key_rotate_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"portal_request_burst_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"portal_request_rate_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"portal_token": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"postgres_vacuum_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20160",
			ValidateFunc: validateInteger,
		},
		"process_locked_useraccounts_timeout_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"process_pki_profile_timeout_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1440",
			ValidateFunc: validateInteger,
		},
		"query_host_fail": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "180",
			ValidateFunc: validateInteger,
		},
		"resmgr_log_caching_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "21600",
			ValidateFunc: validateInteger,
		},
		"restrict_cloud_read_access": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"safenet_hsm_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_create_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "900",
			ValidateFunc: validateInteger,
		},
		"se_failover_attempt_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"se_from_marketplace": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "IMAGE_SE",
		},
		"se_offline_del": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "172000",
			ValidateFunc: validateInteger,
		},
		"se_spawn_retry_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"se_upgrade_flow_cleanup_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "90",
			ValidateFunc: validateInteger,
		},
		"se_vnic_cooldown": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "120",
			ValidateFunc: validateInteger,
		},
		"se_vnic_gc_wait_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"secure_channel_cleanup_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"secure_channel_controller_token_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"secure_channel_se_token_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"seupgrade_copy_buffer_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "512",
			ValidateFunc: validateInteger,
		},
		"seupgrade_copy_pool_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateInteger,
		},
		"seupgrade_fabric_pool_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
		},
		"seupgrade_segroup_min_dead_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "360",
			ValidateFunc: validateInteger,
		},
		"shared_ssl_certificates": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"skopeo_retry_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateInteger,
		},
		"skopeo_retry_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"soft_min_mem_per_se_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1900",
			ValidateFunc: validateInteger,
		},
		"ssl_certificate_expiry_warning_days": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"system_report_cleanup_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"system_report_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"unresponsive_se_reboot": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"update_dns_entry_retry_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"update_dns_entry_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "120",
			ValidateFunc: validateInteger,
		},
		"upgrade_dns_ttl": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateInteger,
		},
		"upgrade_fat_se_lease_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1200",
			ValidateFunc: validateInteger,
		},
		"upgrade_lease_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "600",
			ValidateFunc: validateInteger,
		},
		"upgrade_se_per_vs_scale_ops_txn_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"user_agent_cache_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceUserAgentCacheConfigSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vnic_op_fail_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "180",
			ValidateFunc: validateInteger,
		},
		"vs_awaiting_se_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"vs_key_rotate_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "360",
			ValidateFunc: validateInteger,
		},
		"vs_scaleout_ready_check_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"vs_se_attach_ip_fail": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "600",
			ValidateFunc: validateInteger,
		},
		"vs_se_bootup_fail": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "480",
			ValidateFunc: validateInteger,
		},
		"vs_se_bootup_fail_patch": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "600",
			ValidateFunc: validateInteger,
		},
		"vs_se_create_fail": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1500",
			ValidateFunc: validateInteger,
		},
		"vs_se_ping_fail": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "60",
			ValidateFunc: validateInteger,
		},
		"vs_se_vnic_fail": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
		"vs_se_vnic_ip_fail": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "120",
			ValidateFunc: validateInteger,
		},
		"vsphere_ha_detection_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "120",
			ValidateFunc: validateInteger,
		},
		"vsphere_ha_recovery_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "480",
			ValidateFunc: validateInteger,
		},
		"vsphere_ha_timer_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
		},
		"warmstart_se_reconnect_wait_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "480",
			ValidateFunc: validateInteger,
		},
		"warmstart_vs_resync_wait_time": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "300",
			ValidateFunc: validateInteger,
		},
	}
}

func resourceAviControllerProperties() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviControllerPropertiesCreate,
		Read:   ResourceAviControllerPropertiesRead,
		Update: resourceAviControllerPropertiesUpdate,
		Delete: resourceAviControllerPropertiesDelete,
		Schema: ResourceControllerPropertiesSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceControllerPropertiesImporter,
		},
	}
}

func ResourceControllerPropertiesImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceControllerPropertiesSchema()
	return ResourceImporter(d, m, "controllerproperties", s)
}

func ResourceAviControllerPropertiesRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	err := APIRead(d, meta, "controllerproperties", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviControllerPropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	err := APICreateOrUpdate(d, meta, "controllerproperties", s)
	if err == nil {
		err = ResourceAviControllerPropertiesRead(d, meta)
	}
	return err
}

func resourceAviControllerPropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "controllerproperties", s)
	if err == nil {
		err = ResourceAviControllerPropertiesRead(d, meta)
	}
	return err
}

func resourceAviControllerPropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "controllerproperties")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
