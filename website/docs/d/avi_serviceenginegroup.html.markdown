############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_serviceenginegroup"
sidebar_current: "docs-avi-datasource-serviceenginegroup"
description: |-
  Get information of Avi ServiceEngineGroup.
---

# avi_serviceenginegroup

This data source is used to to get avi_serviceenginegroup objects.

## Example Usage

```hcl
data "avi_serviceenginegroup" "foo_serviceenginegroup" {
    uuid = "serviceenginegroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search ServiceEngineGroup by name.
* `uuid` - (Optional) Search ServiceEngineGroup by uuid.
* `cloud_ref` - (Optional) Search ServiceEngineGroup by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `accelerated_networking` - Enable accelerated networking option for azure se.
* `active_standby` - Service engines in active/standby mode for ha failover.
* `aggressive_failure_detection` - Enable aggressive failover configuration for ha.
* `algo` - In compact placement, virtual services are placed on existing ses until max_vs_per_se limit is reached.
* `allow_burst` - Allow ses to be created using burst license.
* `app_cache_percent` - A percent value of total se memory reserved for applicationcaching.
* `app_learning_memory_percent` - A percent value of total se memory reserved for application learning.
* `archive_shm_limit` - Amount of se memory in gb until which shared memory is collected in core archive.
* `async_ssl` - Ssl handshakes will be handled by dedicated ssl threads.requires se reboot.
* `async_ssl_threads` - Number of async ssl threads per se_dp.requires se reboot.
* `auto_rebalance` - If set, virtual services will be automatically migrated when load on an se is less than minimum or more than maximum thresholds.
* `auto_rebalance_capacity_per_se` - Capacities of se for auto rebalance for each criteria.
* `auto_rebalance_criteria` - Set of criteria for se auto rebalance.
* `auto_rebalance_interval` - Frequency of rebalance, if 'auto rebalance' is enabled.
* `auto_redistribute_active_standby_load` - Redistribution of virtual services from the takeover se to the replacement se can cause momentary traffic loss.
* `bgp_state_update_interval` - Bgp peer state update interval.
* `buffer_se` - Excess service engine capacity provisioned for ha failover.
* `cloud_ref` - It is a reference to an object of type cloud.
* `compress_ip_rules_for_each_ns_subnet` - Compress ip rules into a single subnet based ip rule for each north-south ipam subnet configured in pcap mode in openshift/kubernetes node.
* `config_debugs_on_all_cores` - Enable config debugs on all cores of se.
* `connection_memory_percentage` - Percentage of memory for connection state.
* `core_shm_app_cache` - Include shared memory for app cache in core file.requires se reboot.
* `core_shm_app_learning` - Include shared memory for app learning in core file.requires se reboot.
* `cpu_reserve` - Boolean flag to set cpu_reserve.
* `cpu_socket_affinity` - Allocate all the cpu cores for the service engine virtual machines  on the same cpu socket.
* `custom_securitygroups_data` - Custom security groups to be associated with data vnics for se instances in openstack and aws clouds.
* `custom_securitygroups_mgmt` - Custom security groups to be associated with management vnic for se instances in openstack and aws clouds.
* `custom_tag` - Custom tag will be used to create the tags for se instance in aws.
* `data_network_id` - Subnet used to spin up the data nic for service engines, used only for azure cloud.
* `datascript_timeout` - Number of instructions before datascript times out.
* `dedicated_dispatcher_core` - Dedicate the core that handles packet receive/transmit from the network to just the dispatching function.
* `description` - User defined description for the object.
* `disable_avi_securitygroups` - By default, avi creates and manages security groups along with custom sg provided by user.
* `disable_csum_offloads` - Stop using tcp/udp and ip checksum offload features of nics.
* `disable_gro` - Disable generic receive offload (gro) in dpdk poll-mode driver packet receive path.
* `disable_se_memory_check` - If set, disable the config memory check done in service engine.
* `disable_tso` - Disable tcp segmentation offload (tso) in dpdk poll-mode driver packet transmit path.
* `disk_per_se` - Amount of disk space for each of the service engine virtual machines.
* `distribute_load_active_standby` - Use both the active and standby service engines for virtual service placement in the legacy active standby ha mode.
* `distribute_queues` - Distributes queue ownership among cores so multiple cores handle dispatcher duties.
* `distribute_vnics` - Distributes vnic ownership among cores so multiple cores handle dispatcher duties.requires se reboot.
* `enable_gratarp_permanent` - Enable gratarp for vip_ip.
* `enable_hsm_priming` - (this is a beta feature).
* `enable_multi_lb` - Applicable only for azure cloud with basic sku lb.
* `enable_pcap_tx_ring` - Enable tx ring support in pcap mode of operation.
* `ephemeral_portrange_end` - End local ephemeral port number for outbound connections.
* `ephemeral_portrange_start` - Start local ephemeral port number for outbound connections.
* `extra_config_multiplier` - Multiplier for extra config to support large vs/pool config.
* `extra_shared_config_memory` - Extra config memory to support large geo db configuration.
* `flow_table_new_syn_max_entries` - Maximum number of flow table entries that have not completed tcp three-way handshake yet.
* `free_list_size` - Number of entries in the free list.
* `gratarp_permanent_periodicity` - Gratarp periodicity for vip-ip.
* `ha_mode` - High availability mode for all the virtual services using this service engine group.
* `hardwaresecuritymodulegroup_ref` - It is a reference to an object of type hardwaresecuritymodulegroup.
* `heap_minimum_config_memory` - Minimum required heap memory to apply any configuration.
* `hm_on_standby` - Enable active health monitoring from the standby se for all placed virtual services.
* `host_attribute_key` - Key of a (key, value) pair identifying a label for a set of nodes usually in container clouds.
* `host_attribute_value` - Value of a (key, value) pair identifying a label for a set of nodes usually in container clouds.
* `host_gateway_monitor` - Enable the host gateway monitor when service engine is deployed as docker container.
* `hypervisor` - Override default hypervisor.
* `ignore_rtt_threshold` - Ignore rtt samples if it is above threshold.
* `ingress_access_data` - Program se security group ingress rules to allow vip data access from remote cidr type.
* `ingress_access_mgmt` - Program se security group ingress rules to allow ssh/icmp management access from remote cidr type.
* `instance_flavor` - Instance/flavor name for se instance.
* `iptables` - Iptable rules.
* `labels` - Labels associated with this se group.
* `least_load_core_selection` - Select core with least load for new flow.
* `license_tier` - Specifies the license tier which would be used.
* `license_type` - If no license type is specified then default license enforcement for the cloud type is chosen.
* `log_disksz` - Maximum disk capacity (in mb) to be allocated to an se.
* `log_malloc_failure` - Se will log memory allocation related failure to the se_trace file, wherever available.
* `max_concurrent_external_hm` - Maximum number of external health monitors that can run concurrently in a service engine.
* `max_cpu_usage` - When cpu usage on an se exceeds this threshold, virtual services hosted on this se may be rebalanced to other ses to reduce load.
* `max_memory_per_mempool` - Max bytes that can be allocated in a single mempool.
* `max_public_ips_per_lb` - Applicable to azure platform only.
* `max_queues_per_vnic` - Maximum number of queues per vnic setting to '0' utilises all queues that are distributed across dispatcher cores.
* `max_rules_per_lb` - Applicable to azure platform only.
* `max_scaleout_per_vs` - Maximum number of active service engines for the virtual service.
* `max_se` - Maximum number of services engines in this group.
* `max_vs_per_se` - Maximum number of virtual services that can be placed on a single service engine.
* `mem_reserve` - Boolean flag to set mem_reserve.
* `memory_for_config_update` - Indicates the percent of memory reserved for config updates.
* `memory_per_se` - Amount of memory for each of the service engine virtual machines.
* `mgmt_network_ref` - Management network to use for avi service engines.
* `mgmt_subnet` - Management subnet to use for avi service engines.
* `min_cpu_usage` - When cpu usage on an se falls below the minimum threshold, virtual services hosted on the se may be consolidated onto other underutilized ses.
* `min_scaleout_per_vs` - Minimum number of active service engines for the virtual service.
* `min_se` - Minimum number of services engines in this group (relevant for se autorebalance only).
* `minimum_connection_memory` - Indicates the percent of memory reserved for connections.
* `n_log_streaming_threads` - Number of threads to use for log streaming.
* `name` - Name of the object.
* `non_significant_log_throttle` - This setting limits the number of non-significant logs generated per second per core on this se.
* `num_dispatcher_cores` - Number of dispatcher cores (0,1,2,4,8 or 16).
* `num_flow_cores_sum_changes_to_ignore` - Number of changes in num flow cores sum to ignore.
* `openstack_availability_zones` - Field introduced in 17.1.1.
* `openstack_mgmt_network_name` - Avi management network name.
* `openstack_mgmt_network_uuid` - Management network uuid.
* `os_reserved_memory` - Amount of extra memory to be reserved for use by the operating system on a service engine.
* `pcap_tx_mode` - Determines the pcap transmit mode of operation.
* `per_app` - Per-app se mode is designed for deploying dedicated load balancers per app (vs).
* `placement_mode` - If placement mode is 'auto', virtual services are automatically placed on service engines.
* `realtime_se_metrics` - Enable or disable real time se metrics.
* `reboot_on_panic` - Reboot the vm or host on kernel panic.
* `se_bandwidth_type` - Select the se bandwidth for the bandwidth license.
* `se_delayed_flow_delete` - Delay the cleanup of flowtable entry.
* `se_deprovision_delay` - Duration to preserve unused service engine virtual machines before deleting them.
* `se_dos_profile` - Dict settings for serviceenginegroup.
* `se_dp_vnic_queue_stall_event_sleep` - Time (in seconds) service engine waits for after generating a vnic transmit queue stall event before resetting thenic.
* `se_dp_vnic_queue_stall_threshold` - Number of consecutive transmit failures to look for before generating a vnic transmit queue stall event.
* `se_dp_vnic_queue_stall_timeout` - Time (in milliseconds) to wait for network/nic recovery on detecting a transmit queue stall after which service engine resets the nic.
* `se_dp_vnic_restart_on_queue_stall_count` - Number of consecutive transmit queue stall events in se_dp_vnic_stall_se_restart_window to look for before restarting se.
* `se_dp_vnic_stall_se_restart_window` - Window of time (in seconds) during which se_dp_vnic_restart_on_queue_stall_count number of consecutive stalls results in a se restart.
* `se_dpdk_pmd` - Determines if dpdk pool mode driver should be used or not   0  automatically determine based on hypervisor/nic type 1  unconditionally use dpdk poll mode driver 2  don't use dpdk poll mode driver.requires se reboot.
* `se_flow_probe_retries` - Flow probe retry count if no replies are received.requires se reboot.
* `se_flow_probe_retry_timer` - Timeout in milliseconds for flow probe retries.requires se reboot.
* `se_ipc_udp_port` - Udp port for se_dp ipc in docker bridge mode.
* `se_kni_burst_factor` - Knob to control burst size used in polling kni interfaces for traffic sent from kni towards dpdk application also controls burst size used by kni module to read pkts punted from dpdk application towards kni helps minimize drops in non-vip traffic in either pathfactor of (0-2) multiplies/divides burst size by 2^n.
* `se_lro` - Enable or disable large receive optimization for vnics.
* `se_mp_ring_retry_count` - The retry count for the multi-producer enqueue before yielding the cpu.
* `se_mtu` - Mtu for the vnics of ses in the se group.
* `se_name_prefix` - Prefix to use for virtual machine name of service engines.
* `se_pcap_lookahead` - Enables lookahead mode of packet receive in pcap mode.
* `se_pcap_pkt_count` - Max number of packets the pcap interface can hold and if the value is 0 the optimum value will be chosen.
* `se_pcap_pkt_sz` - Max size of each packet in the pcap interface.
* `se_pcap_qdisc_bypass` - Bypass the kernel's traffic control layer, to deliver packets directly to the driver.
* `se_pcap_reinit_frequency` - Frequency in seconds at which periodically a pcap reinit check is triggered.
* `se_pcap_reinit_threshold` - Threshold for input packet receive errors in pcap mode exceeding which a pcap reinit is triggered.
* `se_probe_port` - Tcp port on se where echo service will be run.
* `se_remote_punt_udp_port` - Udp port for punted packets in docker bridge mode.
* `se_rl_prop` - Rate limiter properties.
* `se_rum_sampling_nav_interval` - Minimum time to wait on server between taking sampleswhen sampling the navigation timing data from the end user client.
* `se_rum_sampling_nav_percent` - Percentage of navigation timing data from the end user client, used for sampling to get client insights.
* `se_rum_sampling_res_interval` - Minimum time to wait on server between taking sampleswhen sampling the resource timing data from the end user client.
* `se_rum_sampling_res_percent` - Percentage of resource timing data from the end user client used for sampling to get client insight.
* `se_sb_dedicated_core` - Sideband traffic will be handled by a dedicated core.requires se reboot.
* `se_sb_threads` - Number of sideband threads per se.requires se reboot.
* `se_thread_multiplier` - Multiplier for se threads based on vcpu.
* `se_tracert_port_range` - Traceroute port range.
* `se_tunnel_mode` - Determines if dsr from secondary se is active or not  0  automatically determine based on hypervisor type.
* `se_tunnel_udp_port` - Udp port for tunneled packets from secondary to primary se in docker bridge mode.requires se reboot.
* `se_tx_batch_size` - Number of packets to batch for transmit to the nic.
* `se_txq_threshold` - Once the tx queue of the dispatcher reaches this threshold, hardware queues are not polled for further packets.
* `se_udp_encap_ipc` - Determines if se-se ipc messages are encapsulated in a udp header  0  automatically determine based on hypervisor type.
* `se_use_dpdk` - Determines if dpdk library should be used or not   0  automatically determine based on hypervisor type 1  use dpdk if pcap is not enabled 2  don't use dpdk.
* `se_vs_hb_max_pkts_in_batch` - Maximum number of aggregated vs heartbeat packets to send in a batch.
* `se_vs_hb_max_vs_in_pkt` - Maximum number of virtualservices for which heartbeat messages are aggregated in one packet.
* `self_se_election` - Enable ses to elect a primary amongst themselves in the absence of a connectivity to controller.
* `service_ip6_subnets` - Ipv6 subnets assigned to the se group.
* `service_ip_subnets` - Subnets assigned to the se group.
* `shm_minimum_config_memory` - Minimum required shared memory to apply any configuration.
* `significant_log_throttle` - This setting limits the number of significant logs generated per second per core on this se.
* `ssl_preprocess_sni_hostname` - (beta) preprocess ssl client hello for sni hostname extension.if set to true, this will apply sni child's ssl protocol(s), if they are different from sni parent's allowed ssl protocol(s).
* `tenant_ref` - It is a reference to an object of type tenant.
* `transient_shared_memory_max` - The threshold for the transient shared config memory in the se.
* `udf_log_throttle` - This setting limits the number of udf logs generated per second per core on this se.
* `use_standard_alb` - Use standard sku azure load balancer.
* `uuid` - Unique object identifier of the object.
* `vcenter_clusters` - Dict settings for serviceenginegroup.
* `vcenter_datastore_mode` - Enum options - vcenter_datastore_any, vcenter_datastore_local, vcenter_datastore_shared.
* `vcenter_datastores` - List of list.
* `vcenter_datastores_include` - Boolean flag to set vcenter_datastores_include.
* `vcenter_folder` - Folder to place all the service engine virtual machines in vcenter.
* `vcenter_hosts` - Dict settings for serviceenginegroup.
* `vcpus_per_se` - Number of vcpus for each of the service engine virtual machines.
* `vip_asg` - When vip_asg is set, vip configuration will be managed by avi.user will be able to configure vip_asg or vips individually at the time of create.
* `vs_host_redundancy` - Ensure primary and secondary service engines are deployed on different physical hosts.
* `vs_scalein_timeout` - Time to wait for the scaled in se to drain existing flows before marking the scalein done.
* `vs_scalein_timeout_for_upgrade` - During se upgrade, time to wait for the scaled-in se to drain existing flows before marking the scalein done.
* `vs_scaleout_timeout` - Time to wait for the scaled out se to become ready before marking the scaleout done.
* `vs_se_scaleout_additional_wait_time` - Wait time for sending scaleout ready notification after virtual service is marked up.
* `vs_se_scaleout_ready_timeout` - Timeout in seconds for service engine to sendscaleout ready notification of a virtual service.
* `vs_switchover_timeout` - During se upgrade in a legacy active/standby segroup, time to wait for the new primary se to accept flows before marking the switchover done.
* `vss_placement` - Parameters to place virtual services on only a subset of the cores of an se.
* `vss_placement_enabled` - If set, virtual services will be placed on only a subset of the cores of an se.
* `waf_mempool` - Enable memory pool for waf.requires se reboot.
* `waf_mempool_size` - Memory pool size used for waf.requires se reboot.

