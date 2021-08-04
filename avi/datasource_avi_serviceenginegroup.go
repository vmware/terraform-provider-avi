// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviServiceEngineGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServiceEngineGroupRead,
		Schema: map[string]*schema.Schema{
			"accelerated_networking": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active_standby": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aggressive_failure_detection": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"algo": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_burst": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_cache_percent": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_learning_memory_percent": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"archive_shm_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"async_ssl": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"async_ssl_threads": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_rebalance": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_redistribute_active_standby_load": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bgp_state_update_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"buffer_se": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compress_ip_rules_for_each_ns_subnet": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_debugs_on_all_cores": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_memory_percentage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"core_shm_app_cache": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"core_shm_app_learning": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_reserve": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_socket_affinity": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"dedicated_dispatcher_core": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_avi_securitygroups": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_csum_offloads": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_gro": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_se_memory_check": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_tso": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disk_per_se": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_load_active_standby": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_queues": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_vnics": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_gratarp_permanent": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_hsm_priming": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_multi_lb": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_pcap_tx_ring": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ephemeral_portrange_end": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ephemeral_portrange_start": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"extra_config_multiplier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"extra_shared_config_memory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_table_new_syn_max_entries": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"free_list_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gratarp_permanent_periodicity": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"hm_on_standby": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_rtt_threshold": {
				Type:     schema.TypeString,
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
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceKeyValueSchema(),
			},
			"least_load_core_selection": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_concurrent_external_hm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_cpu_usage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_memory_per_mempool": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_public_ips_per_lb": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_queues_per_vnic": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_rules_per_lb": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_scaleout_per_vs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_se": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_vs_per_se": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mem_reserve": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_for_config_update": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_per_se": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_scaleout_per_vs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_se": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"minimum_connection_memory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"n_log_streaming_threads": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"non_significant_log_throttle": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"num_dispatcher_cores": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"num_flow_cores_sum_changes_to_ignore": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"pcap_tx_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pcap_tx_ring_rd_balancing_factor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"per_app": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"per_vs_admission_control": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_bandwidth_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_delayed_flow_delete": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_deprovision_delay": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_dos_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDosThresholdProfileSchema(),
			},
			"se_dp_vnic_queue_stall_event_sleep": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_dp_vnic_queue_stall_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_dp_vnic_queue_stall_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_dp_vnic_restart_on_queue_stall_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_dp_vnic_stall_se_restart_window": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_dpdk_pmd": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_flow_probe_retries": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_flow_probe_retry_timer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_ipc_udp_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_kni_burst_factor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_lro": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_mp_ring_retry_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_mtu": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_name_prefix": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_pcap_lookahead": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_pcap_pkt_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_pcap_pkt_sz": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_pcap_qdisc_bypass": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_pcap_reinit_frequency": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_pcap_reinit_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_probe_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_remote_punt_udp_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_rl_prop": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRateLimiterPropertiesSchema(),
			},
			"se_rum_sampling_nav_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_rum_sampling_nav_percent": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_rum_sampling_res_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_rum_sampling_res_percent": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_sb_dedicated_core": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_sb_threads": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_thread_multiplier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_tracert_port_range": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePortRangeSchema(),
			},
			"se_tunnel_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_tunnel_udp_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_tx_batch_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_txq_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_udp_encap_ipc": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_use_dpdk": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_vs_hb_max_pkts_in_batch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_vs_hb_max_vs_in_pkt": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"self_se_election": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"significant_log_throttle": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_preprocess_sni_hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transient_shared_memory_max": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"udf_log_throttle": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_standard_alb": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"vip_asg": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceVipAutoscaleGroupSchema(),
			},
			"vs_host_redundancy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_scalein_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_scalein_timeout_for_upgrade": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_scaleout_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_scaleout_additional_wait_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_se_scaleout_ready_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_switchover_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vss_placement": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceVssPlacementSchema(),
			},
			"vss_placement_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"waf_mempool": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"waf_mempool_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
