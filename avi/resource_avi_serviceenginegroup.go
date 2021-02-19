/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"time"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceServiceEngineGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accelerated_networking": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"active_standby": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"aggressive_failure_detection": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"algo": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PLACEMENT_ALGO_PACKED",
		},
		"allow_burst": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"app_cache_percent": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"app_cache_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"app_learning_memory_percent": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"archive_shm_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  8,
		},
		"async_ssl": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"async_ssl_threads": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"auto_rebalance": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"auto_rebalance_capacity_per_se": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"auto_rebalance_criteria": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"auto_rebalance_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"auto_redistribute_active_standby_load": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"availability_zone_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"bgp_state_update_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"buffer_se": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"compress_ip_rules_for_each_ns_subnet": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"config_debugs_on_all_cores": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"connection_memory_percentage": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  50,
		},
		"core_shm_app_cache": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"core_shm_app_learning": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"cpu_reserve": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"cpu_socket_affinity": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"custom_securitygroups_data": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"custom_securitygroups_mgmt": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"custom_tag": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomTagSchema(),
		},
		"data_network_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"datascript_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1000000,
		},
		"dedicated_dispatcher_core": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"disable_avi_securitygroups": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disable_csum_offloads": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disable_flow_probes": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disable_gro": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"disable_se_memory_check": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disable_tso": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disk_per_se": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  15,
		},
		"distribute_load_active_standby": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"distribute_queues": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"distribute_vnics": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"dp_aggressive_deq_interval_msec": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"dp_aggressive_enq_interval_msec": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"dp_aggressive_hb_frequency": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"dp_aggressive_hb_timeout_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"dp_deq_interval_msec": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"dp_enq_interval_msec": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"dp_hb_frequency": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"dp_hb_timeout_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"enable_gratarp_permanent": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_hsm_priming": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_multi_lb": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_pcap_tx_ring": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"ephemeral_portrange_end": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"ephemeral_portrange_start": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"extra_config_multiplier": {
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "0.0",
		},
		"extra_shared_config_memory": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"flow_table_new_syn_max_entries": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"free_list_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1024,
		},
		"gcp_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGCPSeGroupConfigSchema(),
		},
		"gratarp_permanent_periodicity": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"ha_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "HA_MODE_SHARED",
		},
		"handle_per_pkt_attack": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"hardwaresecuritymodulegroup_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"heap_minimum_config_memory": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  8,
		},
		"hm_on_standby": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"host_attribute_key": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"host_attribute_value": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"host_gateway_monitor": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"hypervisor": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ignore_rtt_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5000,
		},
		"ingress_access_data": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SG_INGRESS_ACCESS_ALL",
		},
		"ingress_access_mgmt": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SG_INGRESS_ACCESS_ALL",
		},
		"instance_flavor": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"instance_flavor_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceCloudFlavorSchema(),
		},
		"iptables": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIptableRuleSetSchema(),
		},
		"labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"least_load_core_selection": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"license_tier": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"license_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"log_disksz": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10000,
		},
		"log_malloc_failure": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"max_concurrent_external_hm": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"max_cpu_usage": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  80,
		},
		"max_memory_per_mempool": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  64,
		},
		"max_num_se_dps": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"max_public_ips_per_lb": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"max_queues_per_vnic": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"max_rules_per_lb": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  150,
		},
		"max_scaleout_per_vs": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"max_se": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"max_vs_per_se": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"mem_reserve": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"memory_for_config_update": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  15,
		},
		"memory_per_se": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2048,
		},
		"mgmt_network_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"mgmt_subnet": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"min_cpu_usage": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"min_scaleout_per_vs": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"min_se": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"minimum_connection_memory": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"n_log_streaming_threads": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"netlink_poller_threads": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"netlink_sock_buf_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"non_significant_log_throttle": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"ns_helper_deq_interval_msec": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"num_dispatcher_cores": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"num_flow_cores_sum_changes_to_ignore": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  8,
		},
		"objsync_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceObjSyncConfigSchema(),
		},
		"objsync_port": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  9001,
		},
		"openstack_availability_zones": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"openstack_mgmt_network_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"openstack_mgmt_network_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"os_reserved_memory": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"pcap_tx_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PCAP_TX_AUTO",
		},
		"pcap_tx_ring_rd_balancing_factor": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"per_app": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"per_vs_admission_control": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"placement_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PLACEMENT_MODE_AUTO",
		},
		"realtime_se_metrics": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceMetricsRealTimeUpdateSchema(),
		},
		"reboot_on_panic": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"resync_time_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  65536,
		},
		"se_bandwidth_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_delayed_flow_delete": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"se_deprovision_delay": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"se_dos_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDosThresholdProfileSchema(),
		},
		"se_dp_hm_drops": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_dp_max_hb_version": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"se_dp_vnic_queue_stall_event_sleep": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_dp_vnic_queue_stall_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2000,
		},
		"se_dp_vnic_queue_stall_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10000,
		},
		"se_dp_vnic_restart_on_queue_stall_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"se_dp_vnic_stall_se_restart_window": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3600,
		},
		"se_dpdk_pmd": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_flow_probe_retries": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"se_flow_probe_retry_timer": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  40,
		},
		"se_group_analytics_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSeGroupAnalyticsPolicySchema(),
		},
		"se_hyperthreaded_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SE_CPU_HT_AUTO",
		},
		"se_ip_encap_ipc": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_kni_burst_factor": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_l3_encap_ipc": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_lro": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"se_mp_ring_retry_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  500,
		},
		"se_mtu": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"se_name_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Avi",
		},
		"se_pcap_lookahead": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"se_pcap_pkt_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_pcap_pkt_sz": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  69632,
		},
		"se_pcap_qdisc_bypass": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"se_pcap_reinit_frequency": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_pcap_reinit_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_probe_port": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  7,
		},
		"se_rl_prop": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRateLimiterPropertiesSchema(),
		},
		"se_rum_sampling_nav_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"se_rum_sampling_nav_percent": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"se_rum_sampling_res_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"se_rum_sampling_res_percent": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"se_sb_dedicated_core": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"se_sb_threads": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"se_thread_multiplier": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"se_tracert_port_range": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePortRangeSchema(),
		},
		"se_tunnel_mode": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_tunnel_udp_port": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1550,
		},
		"se_tx_batch_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  64,
		},
		"se_txq_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2048,
		},
		"se_udp_encap_ipc": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_use_dpdk": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_vnic_tx_sw_queue_flush_frequency": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"se_vnic_tx_sw_queue_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  256,
		},
		"se_vs_hb_max_pkts_in_batch": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  64,
		},
		"se_vs_hb_max_vs_in_pkt": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  256,
		},
		"self_se_election": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"send_se_ready_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"service_ip6_subnets": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"service_ip_subnets": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrPrefixSchema(),
		},
		"shm_minimum_config_memory": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"significant_log_throttle": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"ssl_preprocess_sni_hostname": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"transient_shared_memory_max": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"udf_log_throttle": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"use_hyperthreaded_cores": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"use_objsync": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"use_standard_alb": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vcenter_clusters": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVcenterClustersSchema(),
		},
		"vcenter_datastore_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VCENTER_DATASTORE_ANY",
		},
		"vcenter_datastores": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVcenterDatastoreSchema(),
		},
		"vcenter_datastores_include": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vcenter_folder": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "AviSeFolder",
		},
		"vcenter_hosts": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVcenterHostsSchema(),
		},
		"vcenters": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePlacementScopeConfigSchema(),
		},
		"vcpus_per_se": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"vip_asg": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVipAutoscaleGroupSchema(),
		},
		"vnic_dhcp_ip_check_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  6,
		},
		"vnic_dhcp_ip_max_retries": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"vnic_ip_delete_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"vnic_probe_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"vnic_rpc_retry_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"vnicdb_cmd_history_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  256,
		},
		"vs_host_redundancy": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"vs_scalein_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"vs_scalein_timeout_for_upgrade": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"vs_scaleout_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  600,
		},
		"vs_se_scaleout_additional_wait_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"vs_se_scaleout_ready_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"vs_switchover_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"vss_placement": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVssPlacementSchema(),
		},
		"vss_placement_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"waf_mempool": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"waf_mempool_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  64,
		},
	}
}

func resourceAviServiceEngineGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServiceEngineGroupCreate,
		Read:   ResourceAviServiceEngineGroupRead,
		Update: resourceAviServiceEngineGroupUpdate,
		Delete: resourceAviServiceEngineGroupDelete,
		Schema: ResourceServiceEngineGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceServiceEngineGroupImporter,
		},
	}
}

func ResourceServiceEngineGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceServiceEngineGroupSchema()
	return ResourceImporter(d, m, "serviceenginegroup", s)
}

func ResourceAviServiceEngineGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	err := APIRead(d, meta, "serviceenginegroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviServiceEngineGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	err := APICreateOrUpdate(d, meta, "serviceenginegroup", s)
	if err == nil {
		err = ResourceAviServiceEngineGroupRead(d, meta)
	}
	return err
}

func resourceAviServiceEngineGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceServiceEngineGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "serviceenginegroup", s)
	if err == nil {
		err = ResourceAviServiceEngineGroupRead(d, meta)
	}
	return err
}

func resourceAviServiceEngineGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "serviceenginegroup"
	client := meta.(*clients.AviClient)
	seDeprovisionExtraDelay := 2
	if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string), "api/cloud/") {
		cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
		cloudPath := "api/cloud/" + cloudUUID
		var robj interface{}
		if err := client.AviSession.Get(cloudPath, &robj); err == nil {
			if vcenterConfig, isVcenterConfig := robj.(map[string]interface{})["vcenter_configuration"]; isVcenterConfig {
				if privilege := vcenterConfig.(map[string]interface{})["privilege"].(string); privilege == "WRITE_ACCESS" {
					seGroupName := d.Get("name").(string)
					cloudName := robj.(map[string]interface{})["name"].(string)
					seDeprovisionDelay := d.Get("se_deprovision_delay").(int) + seDeprovisionExtraDelay
					log.Printf("Waiting for %v minutes to delete SE from SE Group %v of cloud %v", seDeprovisionDelay, seGroupName, cloudName)
					time.Sleep(time.Duration(seDeprovisionDelay) * time.Minute)
				}
			}
		}
	}
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviServiceEngineGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
