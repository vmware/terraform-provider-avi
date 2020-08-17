############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_serviceenginegroup"
sidebar_current: "docs-avi-resource-serviceenginegroup"
description: |-
  Creates and manages Avi ServiceEngineGroup.
---

# avi_serviceenginegroup

The ServiceEngineGroup resource allows the creation and management of Avi ServiceEngineGroup

## Example Usage

```hcl
resource "avi_serviceenginegroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `accelerated_networking` - (Optional) Enable accelerated networking option for azure se.
* `active_standby` - (Optional) Service engines in active/standby mode for ha failover.
* `aggressive_failure_detection` - (Optional) Enable aggressive failover configuration for ha.
* `algo` - (Optional) In compact placement, virtual services are placed on existing ses until max_vs_per_se limit is reached.
* `allow_burst` - (Optional) Allow ses to be created using burst license.
* `app_cache_percent` - (Optional) A percent value of total se memory reserved for applicationcaching.
* `app_learning_memory_percent` - (Optional) A percent value of total se memory reserved for application learning.
* `archive_shm_limit` - (Optional) Amount of se memory in gb until which shared memory is collected in core archive.
* `async_ssl` - (Optional) Ssl handshakes will be handled by dedicated ssl threads.requires se reboot.
* `async_ssl_threads` - (Optional) Number of async ssl threads per se_dp.requires se reboot.
* `auto_rebalance` - (Optional) If set, virtual services will be automatically migrated when load on an se is less than minimum or more than maximum thresholds.
* `auto_rebalance_capacity_per_se` - (Optional) Capacities of se for auto rebalance for each criteria.
* `auto_rebalance_criteria` - (Optional) Set of criteria for se auto rebalance.
* `auto_rebalance_interval` - (Optional) Frequency of rebalance, if 'auto rebalance' is enabled.
* `auto_redistribute_active_standby_load` - (Optional) Redistribution of virtual services from the takeover se to the replacement se can cause momentary traffic loss.
* `bgp_state_update_interval` - (Optional) Bgp peer state update interval.
* `buffer_se` - (Optional) Excess service engine capacity provisioned for ha failover.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `compress_ip_rules_for_each_ns_subnet` - (Optional) Compress ip rules into a single subnet based ip rule for each north-south ipam subnet configured in pcap mode in openshift/kubernetes node.
* `config_debugs_on_all_cores` - (Optional) Enable config debugs on all cores of se.
* `connection_memory_percentage` - (Optional) Percentage of memory for connection state.
* `core_shm_app_cache` - (Optional) Include shared memory for app cache in core file.requires se reboot.
* `core_shm_app_learning` - (Optional) Include shared memory for app learning in core file.requires se reboot.
* `cpu_reserve` - (Optional) Boolean flag to set cpu_reserve.
* `cpu_socket_affinity` - (Optional) Allocate all the cpu cores for the service engine virtual machines  on the same cpu socket.
* `custom_securitygroups_data` - (Optional) Custom security groups to be associated with data vnics for se instances in openstack and aws clouds.
* `custom_securitygroups_mgmt` - (Optional) Custom security groups to be associated with management vnic for se instances in openstack and aws clouds.
* `custom_tag` - (Optional) Custom tag will be used to create the tags for se instance in aws.
* `data_network_id` - (Optional) Subnet used to spin up the data nic for service engines, used only for azure cloud.
* `datascript_timeout` - (Optional) Number of instructions before datascript times out.
* `dedicated_dispatcher_core` - (Optional) Dedicate the core that handles packet receive/transmit from the network to just the dispatching function.
* `description` - (Optional) User defined description for the object.
* `disable_avi_securitygroups` - (Optional) By default, avi creates and manages security groups along with custom sg provided by user.
* `disable_csum_offloads` - (Optional) Stop using tcp/udp and ip checksum offload features of nics.
* `disable_gro` - (Optional) Disable generic receive offload (gro) in dpdk poll-mode driver packet receive path.
* `disable_se_memory_check` - (Optional) If set, disable the config memory check done in service engine.
* `disable_tso` - (Optional) Disable tcp segmentation offload (tso) in dpdk poll-mode driver packet transmit path.
* `disk_per_se` - (Optional) Amount of disk space for each of the service engine virtual machines.
* `distribute_load_active_standby` - (Optional) Use both the active and standby service engines for virtual service placement in the legacy active standby ha mode.
* `distribute_queues` - (Optional) Distributes queue ownership among cores so multiple cores handle dispatcher duties.
* `distribute_vnics` - (Optional) Distributes vnic ownership among cores so multiple cores handle dispatcher duties.requires se reboot.
* `enable_gratarp_permanent` - (Optional) Enable gratarp for vip_ip.
* `enable_hsm_priming` - (Optional) (this is a beta feature).
* `enable_multi_lb` - (Optional) Applicable only for azure cloud with basic sku lb.
* `enable_pcap_tx_ring` - (Optional) Enable tx ring support in pcap mode of operation.
* `ephemeral_portrange_end` - (Optional) End local ephemeral port number for outbound connections.
* `ephemeral_portrange_start` - (Optional) Start local ephemeral port number for outbound connections.
* `extra_config_multiplier` - (Optional) Multiplier for extra config to support large vs/pool config.
* `extra_shared_config_memory` - (Optional) Extra config memory to support large geo db configuration.
* `flow_table_new_syn_max_entries` - (Optional) Maximum number of flow table entries that have not completed tcp three-way handshake yet.
* `free_list_size` - (Optional) Number of entries in the free list.
* `gratarp_permanent_periodicity` - (Optional) Gratarp periodicity for vip-ip.
* `ha_mode` - (Optional) High availability mode for all the virtual services using this service engine group.
* `hardwaresecuritymodulegroup_ref` - (Optional) It is a reference to an object of type hardwaresecuritymodulegroup.
* `heap_minimum_config_memory` - (Optional) Minimum required heap memory to apply any configuration.
* `hm_on_standby` - (Optional) Enable active health monitoring from the standby se for all placed virtual services.
* `host_attribute_key` - (Optional) Key of a (key, value) pair identifying a label for a set of nodes usually in container clouds.
* `host_attribute_value` - (Optional) Value of a (key, value) pair identifying a label for a set of nodes usually in container clouds.
* `host_gateway_monitor` - (Optional) Enable the host gateway monitor when service engine is deployed as docker container.
* `hypervisor` - (Optional) Override default hypervisor.
* `ignore_rtt_threshold` - (Optional) Ignore rtt samples if it is above threshold.
* `ingress_access_data` - (Optional) Program se security group ingress rules to allow vip data access from remote cidr type.
* `ingress_access_mgmt` - (Optional) Program se security group ingress rules to allow ssh/icmp management access from remote cidr type.
* `instance_flavor` - (Optional) Instance/flavor name for se instance.
* `iptables` - (Optional) Iptable rules.
* `labels` - (Optional) Labels associated with this se group.
* `least_load_core_selection` - (Optional) Select core with least load for new flow.
* `license_tier` - (Optional) Specifies the license tier which would be used.
* `license_type` - (Optional) If no license type is specified then default license enforcement for the cloud type is chosen.
* `log_disksz` - (Optional) Maximum disk capacity (in mb) to be allocated to an se.
* `log_malloc_failure` - (Optional) Se will log memory allocation related failure to the se_trace file, wherever available.
* `max_concurrent_external_hm` - (Optional) Maximum number of external health monitors that can run concurrently in a service engine.
* `max_cpu_usage` - (Optional) When cpu usage on an se exceeds this threshold, virtual services hosted on this se may be rebalanced to other ses to reduce load.
* `max_memory_per_mempool` - (Optional) Max bytes that can be allocated in a single mempool.
* `max_public_ips_per_lb` - (Optional) Applicable to azure platform only.
* `max_queues_per_vnic` - (Optional) Maximum number of queues per vnic setting to '0' utilises all queues that are distributed across dispatcher cores.
* `max_rules_per_lb` - (Optional) Applicable to azure platform only.
* `max_scaleout_per_vs` - (Optional) Maximum number of active service engines for the virtual service.
* `max_se` - (Optional) Maximum number of services engines in this group.
* `max_vs_per_se` - (Optional) Maximum number of virtual services that can be placed on a single service engine.
* `mem_reserve` - (Optional) Boolean flag to set mem_reserve.
* `memory_for_config_update` - (Optional) Indicates the percent of memory reserved for config updates.
* `memory_per_se` - (Optional) Amount of memory for each of the service engine virtual machines.
* `mgmt_network_ref` - (Optional) Management network to use for avi service engines.
* `mgmt_subnet` - (Optional) Management subnet to use for avi service engines.
* `min_cpu_usage` - (Optional) When cpu usage on an se falls below the minimum threshold, virtual services hosted on the se may be consolidated onto other underutilized ses.
* `min_scaleout_per_vs` - (Optional) Minimum number of active service engines for the virtual service.
* `min_se` - (Optional) Minimum number of services engines in this group (relevant for se autorebalance only).
* `minimum_connection_memory` - (Optional) Indicates the percent of memory reserved for connections.
* `n_log_streaming_threads` - (Optional) Number of threads to use for log streaming.
* `non_significant_log_throttle` - (Optional) This setting limits the number of non-significant logs generated per second per core on this se.
* `num_dispatcher_cores` - (Optional) Number of dispatcher cores (0,1,2,4,8 or 16).
* `num_flow_cores_sum_changes_to_ignore` - (Optional) Number of changes in num flow cores sum to ignore.
* `openstack_availability_zones` - (Optional) Field introduced in 17.1.1.
* `openstack_mgmt_network_name` - (Optional) Avi management network name.
* `openstack_mgmt_network_uuid` - (Optional) Management network uuid.
* `os_reserved_memory` - (Optional) Amount of extra memory to be reserved for use by the operating system on a service engine.
* `pcap_tx_mode` - (Optional) Determines the pcap transmit mode of operation.
* `per_app` - (Optional) Per-app se mode is designed for deploying dedicated load balancers per app (vs).
* `placement_mode` - (Optional) If placement mode is 'auto', virtual services are automatically placed on service engines.
* `realtime_se_metrics` - (Optional) Enable or disable real time se metrics.
* `reboot_on_panic` - (Optional) Reboot the vm or host on kernel panic.
* `se_bandwidth_type` - (Optional) Select the se bandwidth for the bandwidth license.
* `se_delayed_flow_delete` - (Optional) Delay the cleanup of flowtable entry.
* `se_deprovision_delay` - (Optional) Duration to preserve unused service engine virtual machines before deleting them.
* `se_dos_profile` - (Optional) Dict settings for serviceenginegroup.
* `se_dp_vnic_queue_stall_event_sleep` - (Optional) Time (in seconds) service engine waits for after generating a vnic transmit queue stall event before resetting thenic.
* `se_dp_vnic_queue_stall_threshold` - (Optional) Number of consecutive transmit failures to look for before generating a vnic transmit queue stall event.
* `se_dp_vnic_queue_stall_timeout` - (Optional) Time (in milliseconds) to wait for network/nic recovery on detecting a transmit queue stall after which service engine resets the nic.
* `se_dp_vnic_restart_on_queue_stall_count` - (Optional) Number of consecutive transmit queue stall events in se_dp_vnic_stall_se_restart_window to look for before restarting se.
* `se_dp_vnic_stall_se_restart_window` - (Optional) Window of time (in seconds) during which se_dp_vnic_restart_on_queue_stall_count number of consecutive stalls results in a se restart.
* `se_dpdk_pmd` - (Optional) Determines if dpdk pool mode driver should be used or not   0  automatically determine based on hypervisor/nic type 1  unconditionally use dpdk poll mode driver 2  don't use dpdk poll mode driver.requires se reboot.
* `se_flow_probe_retries` - (Optional) Flow probe retry count if no replies are received.requires se reboot.
* `se_flow_probe_retry_timer` - (Optional) Timeout in milliseconds for flow probe retries.requires se reboot.
* `se_ipc_udp_port` - (Optional) Udp port for se_dp ipc in docker bridge mode.
* `se_kni_burst_factor` - (Optional) Knob to control burst size used in polling kni interfaces for traffic sent from kni towards dpdk application also controls burst size used by kni module to read pkts punted from dpdk application towards kni helps minimize drops in non-vip traffic in either pathfactor of (0-2) multiplies/divides burst size by 2^n.
* `se_lro` - (Optional) Enable or disable large receive optimization for vnics.
* `se_mp_ring_retry_count` - (Optional) The retry count for the multi-producer enqueue before yielding the cpu.
* `se_mtu` - (Optional) Mtu for the vnics of ses in the se group.
* `se_name_prefix` - (Optional) Prefix to use for virtual machine name of service engines.
* `se_pcap_lookahead` - (Optional) Enables lookahead mode of packet receive in pcap mode.
* `se_pcap_pkt_count` - (Optional) Max number of packets the pcap interface can hold and if the value is 0 the optimum value will be chosen.
* `se_pcap_pkt_sz` - (Optional) Max size of each packet in the pcap interface.
* `se_pcap_qdisc_bypass` - (Optional) Bypass the kernel's traffic control layer, to deliver packets directly to the driver.
* `se_pcap_reinit_frequency` - (Optional) Frequency in seconds at which periodically a pcap reinit check is triggered.
* `se_pcap_reinit_threshold` - (Optional) Threshold for input packet receive errors in pcap mode exceeding which a pcap reinit is triggered.
* `se_probe_port` - (Optional) Tcp port on se where echo service will be run.
* `se_remote_punt_udp_port` - (Optional) Udp port for punted packets in docker bridge mode.
* `se_rl_prop` - (Optional) Rate limiter properties.
* `se_rum_sampling_nav_interval` - (Optional) Minimum time to wait on server between taking sampleswhen sampling the navigation timing data from the end user client.
* `se_rum_sampling_nav_percent` - (Optional) Percentage of navigation timing data from the end user client, used for sampling to get client insights.
* `se_rum_sampling_res_interval` - (Optional) Minimum time to wait on server between taking sampleswhen sampling the resource timing data from the end user client.
* `se_rum_sampling_res_percent` - (Optional) Percentage of resource timing data from the end user client used for sampling to get client insight.
* `se_sb_dedicated_core` - (Optional) Sideband traffic will be handled by a dedicated core.requires se reboot.
* `se_sb_threads` - (Optional) Number of sideband threads per se.requires se reboot.
* `se_thread_multiplier` - (Optional) Multiplier for se threads based on vcpu.
* `se_tracert_port_range` - (Optional) Traceroute port range.
* `se_tunnel_mode` - (Optional) Determines if dsr from secondary se is active or not  0  automatically determine based on hypervisor type.
* `se_tunnel_udp_port` - (Optional) Udp port for tunneled packets from secondary to primary se in docker bridge mode.requires se reboot.
* `se_tx_batch_size` - (Optional) Number of packets to batch for transmit to the nic.
* `se_txq_threshold` - (Optional) Once the tx queue of the dispatcher reaches this threshold, hardware queues are not polled for further packets.
* `se_udp_encap_ipc` - (Optional) Determines if se-se ipc messages are encapsulated in a udp header  0  automatically determine based on hypervisor type.
* `se_use_dpdk` - (Optional) Determines if dpdk library should be used or not   0  automatically determine based on hypervisor type 1  use dpdk if pcap is not enabled 2  don't use dpdk.
* `se_vs_hb_max_pkts_in_batch` - (Optional) Maximum number of aggregated vs heartbeat packets to send in a batch.
* `se_vs_hb_max_vs_in_pkt` - (Optional) Maximum number of virtualservices for which heartbeat messages are aggregated in one packet.
* `self_se_election` - (Optional) Enable ses to elect a primary amongst themselves in the absence of a connectivity to controller.
* `service_ip6_subnets` - (Optional) Ipv6 subnets assigned to the se group.
* `service_ip_subnets` - (Optional) Subnets assigned to the se group.
* `shm_minimum_config_memory` - (Optional) Minimum required shared memory to apply any configuration.
* `significant_log_throttle` - (Optional) This setting limits the number of significant logs generated per second per core on this se.
* `ssl_preprocess_sni_hostname` - (Optional) (beta) preprocess ssl client hello for sni hostname extension.if set to true, this will apply sni child's ssl protocol(s), if they are different from sni parent's allowed ssl protocol(s).
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `udf_log_throttle` - (Optional) This setting limits the number of udf logs generated per second per core on this se.
* `use_standard_alb` - (Optional) Use standard sku azure load balancer.
* `vcenter_clusters` - (Optional) Dict settings for serviceenginegroup.
* `vcenter_datastore_mode` - (Optional) Enum options - vcenter_datastore_any, vcenter_datastore_local, vcenter_datastore_shared.
* `vcenter_datastores` - (Optional) List of list.
* `vcenter_datastores_include` - (Optional) Boolean flag to set vcenter_datastores_include.
* `vcenter_folder` - (Optional) Folder to place all the service engine virtual machines in vcenter.
* `vcenter_hosts` - (Optional) Dict settings for serviceenginegroup.
* `vcpus_per_se` - (Optional) Number of vcpus for each of the service engine virtual machines.
* `vip_asg` - (Optional) When vip_asg is set, vip configuration will be managed by avi.user will be able to configure vip_asg or vips individually at the time of create.
* `vs_host_redundancy` - (Optional) Ensure primary and secondary service engines are deployed on different physical hosts.
* `vs_scalein_timeout` - (Optional) Time to wait for the scaled in se to drain existing flows before marking the scalein done.
* `vs_scalein_timeout_for_upgrade` - (Optional) During se upgrade, time to wait for the scaled-in se to drain existing flows before marking the scalein done.
* `vs_scaleout_timeout` - (Optional) Time to wait for the scaled out se to become ready before marking the scaleout done.
* `vs_se_scaleout_additional_wait_time` - (Optional) Wait time for sending scaleout ready notification after virtual service is marked up.
* `vs_se_scaleout_ready_timeout` - (Optional) Timeout in seconds for service engine to sendscaleout ready notification of a virtual service.
* `vs_switchover_timeout` - (Optional) During se upgrade in a legacy active/standby segroup, time to wait for the new primary se to accept flows before marking the switchover done.
* `vss_placement` - (Optional) Parameters to place virtual services on only a subset of the cores of an se.
* `vss_placement_enabled` - (Optional) If set, virtual services will be placed on only a subset of the cores of an se.
* `waf_mempool` - (Optional) Enable memory pool for waf.requires se reboot.
* `waf_mempool_size` - (Optional) Memory pool size used for waf.requires se reboot.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

