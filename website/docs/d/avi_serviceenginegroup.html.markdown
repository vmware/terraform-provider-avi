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
data "ServiceEngineGroup" "foo_ServiceEngineGroup" {
    uuid = "ServiceEngineGroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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
* `advertise_backend_networks` - Advertise reach-ability of backend server networks via adc through bgp for default gateway feature.
* `aggressive_failure_detection` - Enable aggressive failover configuration for ha.
* `algo` - In compact placement, virtual services are placed on existing ses until max_vs_per_se limit is reached.
* `allow_burst` - Allow ses to be created using burst license.
* `app_cache_percent` - A percent value of total se memory reserved for application caching.
* `archive_shm_limit` - Amount of se memory in gb until which shared memory is collected in core archive.
* `async_ssl` - Ssl handshakes will be handled by dedicated ssl threads.
* `async_ssl_threads` - Number of async ssl threads per se_dp.
* `auto_rebalance` - If set, virtual services will be automatically migrated when load on an se is less than minimum or more than maximum thresholds.
* `auto_rebalance_capacity_per_se` - Capacities of se for auto rebalance for each criteria.
* `auto_rebalance_criteria` - Set of criteria for se auto rebalance.
* `auto_rebalance_interval` - Frequency of rebalance, if 'auto rebalance' is enabled.
* `auto_redistribute_active_standby_load` - Redistribution of virtual services from the takeover se to the replacement se can cause momentary traffic loss.
* `bgp_state_update_interval` - Bgp peer state update interval.
* `buffer_se` - Excess service engine capacity provisioned for ha failover.
* `cloud_ref` - It is a reference to an object of type cloud.
* `config_debugs_on_all_cores` - Enable config debugs on all cores of se.
* `connection_memory_percentage` - Percentage of memory for connection state.
* `cpu_reserve` - General description.
* `cpu_socket_affinity` - Allocate all the cpu cores for the service engine virtual machines  on the same cpu socket.
* `custom_securitygroups_data` - Custom security groups to be associated with data vnics for se instances in openstack and aws clouds.
* `custom_securitygroups_mgmt` - Custom security groups to be associated with management vnic for se instances in openstack and aws clouds.
* `custom_tag` - Custom tag will be used to create the tags for se instance in aws.
* `dedicated_dispatcher_core` - Dedicate the core that handles packet receive/transmit from the network to just the dispatching function.
* `description` - General description.
* `disable_avi_securitygroups` - By default, avi creates and manages security groups along with custom sg provided by user.
* `disable_csum_offloads` - Stop using tcp/udp and ip checksum offload features of nics.
* `disable_gro` - Disable generic receive offload (gro) in dpdk poll-mode driver packet receive path.
* `disable_se_memory_check` - If set, disable the config memory check done in service engine.
* `disable_tso` - Disable tcp segmentation offload (tso) in dpdk poll-mode driver packet transmit path.
* `disk_per_se` - Amount of disk space for each of the service engine virtual machines.
* `distribute_load_active_standby` - Use both the active and standby service engines for virtual service placement in the legacy active standby ha mode.
* `distribute_queues` - Distributes queue ownership among cores so multiple cores handle dispatcher duties.
* `enable_hsm_priming` - (this is a beta feature).
* `enable_multi_lb` - Applicable only for azure cloud with basic sku lb.
* `enable_routing` - Enable routing for this serviceenginegroup .
* `enable_vip_on_all_interfaces` - Enable vip on all interfaces of se.
* `enable_vmac` - Use virtual mac address for interfaces on which floating interface ips are placed.
* `ephemeral_portrange_end` - End local ephemeral port number for outbound connections.
* `ephemeral_portrange_start` - Start local ephemeral port number for outbound connections.
* `extra_config_multiplier` - Multiplier for extra config to support large vs/pool config.
* `extra_shared_config_memory` - Extra config memory to support large geo db configuration.
* `floating_intf_ip` - If serviceenginegroup is configured for legacy 1+1 active standby ha mode, floating ip's will be advertised only by the active se in the pair.
* `floating_intf_ip_se_2` - If serviceenginegroup is configured for legacy 1+1 active standby ha mode, floating ip's will be advertised only by the active se in the pair.
* `flow_table_new_syn_max_entries` - Maximum number of flow table entries that have not completed tcp three-way handshake yet.
* `free_list_size` - Number of entries in the free list.
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
* `instance_flavor` - Instance/flavor type for se instance.
* `iptables` - Iptable rules.
* `least_load_core_selection` - Select core with least load for new flow.
* `license_tier` - Specifies the license tier which would be used.
* `license_type` - If no license type is specified then default license enforcement for the cloud type is chosen.
* `log_disksz` - Maximum disk capacity (in mb) to be allocated to an se.
* `max_cpu_usage` - When cpu usage on an se exceeds this threshold, virtual services hosted on this se may be rebalanced to other ses to reduce load.
* `max_memory_per_mempool` - Max bytes that can be allocated in a single mempool.
* `max_public_ips_per_lb` - Applicable to azure platform only.
* `max_rules_per_lb` - Applicable to azure platform only.
* `max_scaleout_per_vs` - Maximum number of active service engines for the virtual service.
* `max_se` - Maximum number of services engines in this group.
* `max_vs_per_se` - Maximum number of virtual services that can be placed on a single service engine.
* `mem_reserve` - General description.
* `memory_for_config_update` - Indicates the percent of memory reserved for config updates.
* `memory_per_se` - Amount of memory for each of the service engine virtual machines.
* `mgmt_network_ref` - Management network to use for avi service engines.
* `mgmt_subnet` - Management subnet to use for avi service engines.
* `min_cpu_usage` - When cpu usage on an se falls below the minimum threshold, virtual services hosted on the se may be consolidated onto other underutilized ses.
* `min_scaleout_per_vs` - Minimum number of active service engines for the virtual service.
* `min_se` - Minimum number of services engines in this group (relevant for se autorebalance only).
* `minimum_connection_memory` - Indicates the percent of memory reserved for connections.
* `n_log_streaming_threads` - Number of threads to use for log streaming.
* `name` - General description.
* `non_significant_log_throttle` - This setting limits the number of non-significant logs generated per second per core on this se.
* `num_dispatcher_cores` - Number of dispatcher cores (0,1,2,4,8 or 16).
* `num_flow_cores_sum_changes_to_ignore` - Number of changes in num flow cores sum to ignore.
* `openstack_availability_zones` - Field introduced in 17.1.1.
* `openstack_mgmt_network_name` - Avi management network name.
* `openstack_mgmt_network_uuid` - Management network uuid.
* `os_reserved_memory` - Amount of extra memory to be reserved for use by the operating system on a service engine.
* `per_app` - Per-app se mode is designed for deploying dedicated load balancers per app (vs).
* `placement_mode` - If placement mode is 'auto', virtual services are automatically placed on service engines.
* `realtime_se_metrics` - Enable or disable real time se metrics.
* `se_bandwidth_type` - Select the se bandwidth for the bandwidth license.
* `se_deprovision_delay` - Duration to preserve unused service engine virtual machines before deleting them.
* `se_dos_profile` - General description.
* `se_dpdk_pmd` - Determines if dpdk pool mode driver should be used or not   0  automatically determine based on hypervisor/nic type 1  unconditionally use dpdk poll mode driver 2  don't use dpdk poll mode driver.
* `se_flow_probe_retries` - Flow probe retry count if no replies are received.
* `se_flow_probe_timer` - Timeout in milliseconds for flow probe entries.
* `se_ipc_udp_port` - Udp port for se_dp ipc in docker bridge mode.
* `se_name_prefix` - Prefix to use for virtual machine name of service engines.
* `se_pcap_reinit_frequency` - Frequency in seconds at which periodically a pcap reinit check is triggered.
* `se_pcap_reinit_threshold` - Threshold for input packet receive errors in pcap mode exceeding which a pcap reinit is triggered.
* `se_probe_port` - Tcp port on se where echo service will be run.
* `se_remote_punt_udp_port` - Udp port for punted packets in docker bridge mode.
* `se_sb_dedicated_core` - Sideband traffic will be handled by a dedicated core.
* `se_sb_threads` - Number of sideband threads per se.
* `se_thread_multiplier` - Multiplier for se threads based on vcpu.
* `se_tracert_port_range` - Traceroute port range.
* `se_tunnel_mode` - Determines if dsr from secondary se is active or not  0  automatically determine based on hypervisor type.
* `se_tunnel_udp_port` - Udp port for tunneled packets from secondary to primary se in docker bridge mode.
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
* `udf_log_throttle` - This setting limits the number of udf logs generated per second per core on this se.
* `uuid` - General description.
* `vcenter_clusters` - General description.
* `vcenter_datastore_mode` - Enum options - vcenter_datastore_any, vcenter_datastore_local, vcenter_datastore_shared.
* `vcenter_datastores` - General description.
* `vcenter_datastores_include` - General description.
* `vcenter_folder` - Folder to place all the service engine virtual machines in vcenter.
* `vcenter_hosts` - General description.
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
* `waf_learning_interval` - Frequency with which se publishes waf learning.
* `waf_learning_memory` - Amount of memory reserved on se for waf learning.
* `waf_mempool` - Enable memory pool for waf.
* `waf_mempool_size` - Memory pool size used for waf.

