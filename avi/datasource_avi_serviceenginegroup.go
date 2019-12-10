/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviServiceEngineGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServiceEngineGroupRead,
		Schema: map[string]*schema.Schema{
			"accelerated_networking": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"active_standby": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"aggressive_failure_detection": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"algo": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_burst": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"app_cache_percent": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"app_learning_memory_percent": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"archive_shm_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"async_ssl": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"async_ssl_threads": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auto_rebalance": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"auto_rebalance_capacity_per_se": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"auto_rebalance_criteria": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"auto_rebalance_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auto_redistribute_active_standby_load": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"bgp_state_update_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"buffer_se": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_debugs_on_all_cores": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"connection_memory_percentage": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpu_reserve": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cpu_socket_affinity": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"custom_securitygroups_data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"custom_securitygroups_mgmt": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"custom_tag": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCustomTagSchema(),
			},
			"data_network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datascript_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dedicated_dispatcher_core": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_avi_securitygroups": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disable_csum_offloads": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disable_gro": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disable_se_memory_check": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disable_tso": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disk_per_se": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distribute_load_active_standby": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"distribute_queues": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"distribute_vnics": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_gratarp_permanent": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_hsm_priming": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_multi_lb": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_pcap_tx_ring": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ephemeral_portrange_end": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ephemeral_portrange_start": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"extra_config_multiplier": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"extra_shared_config_memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"flow_table_new_syn_max_entries": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"free_list_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"gratarp_permanent_periodicity": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ha_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardwaresecuritymodulegroup_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"heap_minimum_config_memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hm_on_standby": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"host_attribute_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_attribute_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_gateway_monitor": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_rtt_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ingress_access_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingress_access_mgmt": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_flavor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iptables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIptableRuleSetSchema(),
			},
			"least_load_core_selection": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"license_tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_disksz": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_cpu_usage": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_memory_per_mempool": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_public_ips_per_lb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_rules_per_lb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_scaleout_per_vs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_se": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_vs_per_se": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mem_reserve": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"memory_for_config_update": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_per_se": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mgmt_network_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_subnet": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"min_cpu_usage": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_scaleout_per_vs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_se": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"minimum_connection_memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"n_log_streaming_threads": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"non_significant_log_throttle": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"num_dispatcher_cores": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"num_flow_cores_sum_changes_to_ignore": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"openstack_availability_zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"openstack_mgmt_network_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"openstack_mgmt_network_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_reserved_memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"per_app": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"placement_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"realtime_se_metrics": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceMetricsRealTimeUpdateSchema(),
			},
			"reboot_on_panic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"se_bandwidth_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_deprovision_delay": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_dos_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDosThresholdProfileSchema(),
			},
			"se_dp_vnic_queue_stall_event_sleep": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_dp_vnic_queue_stall_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_dp_vnic_queue_stall_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_dp_vnic_restart_on_queue_stall_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_dp_vnic_stall_se_restart_window": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_dpdk_pmd": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_flow_probe_retries": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_flow_probe_retry_timer": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_ipc_udp_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_kni_burst_factor": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_lro": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"se_name_prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_pcap_lookahead": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"se_pcap_pkt_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_pcap_pkt_sz": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_pcap_qdisc_bypass": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"se_pcap_reinit_frequency": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_pcap_reinit_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_probe_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_remote_punt_udp_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_rum_sampling_nav_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_rum_sampling_nav_percent": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_rum_sampling_res_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_rum_sampling_res_percent": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_sb_dedicated_core": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"se_sb_threads": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_thread_multiplier": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_tracert_port_range": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"se_tunnel_mode": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_tunnel_udp_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_tx_batch_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_udp_encap_ipc": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_use_dpdk": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_vs_hb_max_pkts_in_batch": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_vs_hb_max_vs_in_pkt": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"self_se_election": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"service_ip6_subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"service_ip_subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"shm_minimum_config_memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"significant_log_throttle": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssl_preprocess_sni_hostname": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udf_log_throttle": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"use_standard_alb": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_clusters": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceVcenterClustersSchema(),
			},
			"vcenter_datastore_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_datastores": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceVcenterDatastoreSchema(),
			},
			"vcenter_datastores_include": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vcenter_folder": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_hosts": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceVcenterHostsSchema(),
			},
			"vcpus_per_se": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vip_asg": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceVipAutoscaleGroupSchema(),
			},
			"vs_host_redundancy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vs_scalein_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_scalein_timeout_for_upgrade": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_scaleout_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_se_scaleout_additional_wait_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_se_scaleout_ready_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vs_switchover_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vss_placement": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceVssPlacementSchema(),
			},
			"vss_placement_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"waf_mempool": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"waf_mempool_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}
