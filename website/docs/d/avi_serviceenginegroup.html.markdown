<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `accelerated_networking` - Enable accelerated networking option for azure se. Accelerated networking enables single root i/o virtualization (sr-iov) to a se vm. This improves networking performance. Field introduced in 17.2.14,18.1.5,18.2.1.
* `active_standby` - Service engines in active/standby mode for ha failover.
* `aggressive_failure_detection` - Enable aggressive failover configuration for ha.
* `algo` - In compact placement, virtual services are placed on existing ses until max_vs_per_se limit is reached. Enum options - PLACEMENT_ALGO_PACKED, PLACEMENT_ALGO_DISTRIBUTED.
* `allow_burst` - Allow ses to be created using burst license. Field introduced in 17.2.5.
* `app_cache_percent` - A percent value of total se memory reserved for applicationcaching. This is an se bootup property and requires se restart.requires se reboot. Allowed values are 0 - 100. Special values are 0- 'disable'. Field introduced in 18.2.3.
* `app_learning_memory_percent` - A percent value of total se memory reserved for application learning. This is an se bootup property and requires se restart. Allowed values are 0 - 10. Field introduced in 18.2.3.
* `archive_shm_limit` - Amount of se memory in gb until which shared memory is collected in core archive. Field introduced in 17.1.3.
* `async_ssl` - Ssl handshakes will be handled by dedicated ssl threads.requires se reboot.
* `async_ssl_threads` - Number of async ssl threads per se_dp.requires se reboot. Allowed values are 1-16.
* `auto_rebalance` - If set, virtual services will be automatically migrated when load on an se is less than minimum or more than maximum thresholds. Only alerts are generated when the auto_rebalance is not set.
* `auto_rebalance_capacity_per_se` - Capacities of se for auto rebalance for each criteria. Field introduced in 17.2.4.
* `auto_rebalance_criteria` - Set of criteria for se auto rebalance. Enum options - SE_AUTO_REBALANCE_CPU, SE_AUTO_REBALANCE_PPS, SE_AUTO_REBALANCE_MBPS, SE_AUTO_REBALANCE_OPEN_CONNS, SE_AUTO_REBALANCE_CPS. Field introduced in 17.2.3.
* `auto_rebalance_interval` - Frequency of rebalance, if 'auto rebalance' is enabled.
* `auto_redistribute_active_standby_load` - Redistribution of virtual services from the takeover se to the replacement se can cause momentary traffic loss. If the auto-redistribute load option is left in its default off state, any desired rebalancing requires calls to rest api.
* `bgp_state_update_interval` - Bgp peer state update interval. Allowed values are 5-100. Field introduced in 17.2.14,18.1.5,18.2.1.
* `buffer_se` - Excess service engine capacity provisioned for ha failover.
* `cloud_ref` - It is a reference to an object of type cloud.
* `compress_ip_rules_for_each_ns_subnet` - Compress ip rules into a single subnet based ip rule for each north-south ipam subnet configured in pcap mode in openshift/kubernetes node. Field introduced in 18.2.9.
* `config_debugs_on_all_cores` - Enable config debugs on all cores of se. Field introduced in 17.2.13,18.1.5,18.2.1.
* `connection_memory_percentage` - Percentage of memory for connection state. This will come at the expense of memory used for http in-memory cache. Allowed values are 10-90.
* `core_shm_app_cache` - Include shared memory for app cache in core file.requires se reboot. Field introduced in 18.2.8.
* `core_shm_app_learning` - Include shared memory for app learning in core file.requires se reboot. Field introduced in 18.2.8.
* `cpu_reserve` - Boolean flag to set cpu_reserve.
* `cpu_socket_affinity` - Allocate all the cpu cores for the service engine virtual machines  on the same cpu socket. Applicable only for vcenter cloud.
* `custom_securitygroups_data` - Custom security groups to be associated with data vnics for se instances in openstack and aws clouds. Field introduced in 17.1.3.
* `custom_securitygroups_mgmt` - Custom security groups to be associated with management vnic for se instances in openstack and aws clouds. Field introduced in 17.1.3.
* `custom_tag` - Custom tag will be used to create the tags for se instance in aws. Note this is not the same as the prefix for se name.
* `data_network_id` - Subnet used to spin up the data nic for service engines, used only for azure cloud. Overrides the cloud level setting for service engine subnet. Field introduced in 18.2.3.
* `datascript_timeout` - Number of instructions before datascript times out. Allowed values are 0-100000000. Field introduced in 18.2.3.
* `dedicated_dispatcher_core` - Dedicate the core that handles packet receive/transmit from the network to just the dispatching function. Don't use it for tcp/ip and ssl functions.
* `description` - User defined description for the object.
* `disable_avi_securitygroups` - By default, avi creates and manages security groups along with custom sg provided by user. Set this to true to disallow avi to create and manage new security groups. Avi will only make use of custom security groups provided by user. This option is supported for aws and openstack cloud types. Field introduced in 17.2.13,18.1.4,18.2.1.
* `disable_csum_offloads` - Stop using tcp/udp and ip checksum offload features of nics. Field introduced in 17.1.14, 17.2.5, 18.1.1.
* `disable_gro` - Disable generic receive offload (gro) in dpdk poll-mode driver packet receive path. Gro is on by default on nics that do not support lro (large receive offload) or do not gain performance boost from lro. Field introduced in 17.2.5, 18.1.1.
* `disable_se_memory_check` - If set, disable the config memory check done in service engine. Field introduced in 18.1.2.
* `disable_tso` - Disable tcp segmentation offload (tso) in dpdk poll-mode driver packet transmit path. Tso is on by default on nics that support it. Field introduced in 17.2.5, 18.1.1.
* `disk_per_se` - Amount of disk space for each of the service engine virtual machines.
* `distribute_load_active_standby` - Use both the active and standby service engines for virtual service placement in the legacy active standby ha mode.
* `distribute_queues` - Distributes queue ownership among cores so multiple cores handle dispatcher duties. Requires se reboot. Deprecated from 18.2.8, instead use max_queues_per_vnic. Field introduced in 17.2.8.
* `distribute_vnics` - Distributes vnic ownership among cores so multiple cores handle dispatcher duties.requires se reboot. Field introduced in 18.2.5.
* `enable_gratarp_permanent` - Enable gratarp for vip_ip. Field introduced in 18.2.3.
* `enable_hsm_priming` - (this is a beta feature). Enable hsm key priming. If enabled, key handles on the hsm will be synced to se before processing client connections. Field introduced in 17.2.7, 18.1.1.
* `enable_multi_lb` - Applicable only for azure cloud with basic sku lb. If set, additional azure lbs will be automatically created if resources in existing lb are exhausted. Field introduced in 17.2.10, 18.1.2.
* `enable_pcap_tx_ring` - Enable tx ring support in pcap mode of operation. Tso feature is not supported with tx ring enabled. Deprecated from 18.2.8, instead use pcap_tx_mode. Requires se reboot. Field introduced in 18.2.5.
* `ephemeral_portrange_end` - End local ephemeral port number for outbound connections. Field introduced in 17.2.13, 18.1.5, 18.2.1.
* `ephemeral_portrange_start` - Start local ephemeral port number for outbound connections. Field introduced in 17.2.13, 18.1.5, 18.2.1.
* `extra_config_multiplier` - Multiplier for extra config to support large vs/pool config.
* `extra_shared_config_memory` - Extra config memory to support large geo db configuration. Field introduced in 17.1.1.
* `flow_table_new_syn_max_entries` - Maximum number of flow table entries that have not completed tcp three-way handshake yet. Field introduced in 17.2.5.
* `free_list_size` - Number of entries in the free list. Field introduced in 17.2.10, 18.1.2.
* `gratarp_permanent_periodicity` - Gratarp periodicity for vip-ip. Allowed values are 5-30. Field introduced in 18.2.3.
* `ha_mode` - High availability mode for all the virtual services using this service engine group. Enum options - HA_MODE_SHARED_PAIR, HA_MODE_SHARED, HA_MODE_LEGACY_ACTIVE_STANDBY.
* `hardwaresecuritymodulegroup_ref` - It is a reference to an object of type hardwaresecuritymodulegroup.
* `heap_minimum_config_memory` - Minimum required heap memory to apply any configuration. Allowed values are 0-100. Field introduced in 18.1.2.
* `hm_on_standby` - Enable active health monitoring from the standby se for all placed virtual services.
* `host_attribute_key` - Key of a (key, value) pair identifying a label for a set of nodes usually in container clouds. Needs to be specified together with host_attribute_value. Ses can be configured differently including ha modes across different se groups. May also be used for isolation between different classes of virtualservices. Virtualservices' se group may be specified via annotations/labels. A openshift/kubernetes namespace maybe annotated with a matching se group label as openshift.io/node-selector  apptype=prod. When multiple se groups are used in a cloud with host attributes specified,just a single se group can exist as a match-all se group without a host_attribute_key.
* `host_attribute_value` - Value of a (key, value) pair identifying a label for a set of nodes usually in container clouds. Needs to be specified together with host_attribute_key.
* `host_gateway_monitor` - Enable the host gateway monitor when service engine is deployed as docker container. Disabled by default. Field introduced in 17.2.4.
* `hypervisor` - Override default hypervisor. Enum options - DEFAULT, VMWARE_ESX, KVM, VMWARE_VSAN, XEN.
* `ignore_rtt_threshold` - Ignore rtt samples if it is above threshold. Field introduced in 17.1.6,17.2.2.
* `ingress_access_data` - Program se security group ingress rules to allow vip data access from remote cidr type. Enum options - SG_INGRESS_ACCESS_NONE, SG_INGRESS_ACCESS_ALL, SG_INGRESS_ACCESS_VPC. Field introduced in 17.1.5.
* `ingress_access_mgmt` - Program se security group ingress rules to allow ssh/icmp management access from remote cidr type. Enum options - SG_INGRESS_ACCESS_NONE, SG_INGRESS_ACCESS_ALL, SG_INGRESS_ACCESS_VPC. Field introduced in 17.1.5.
* `instance_flavor` - Instance/flavor name for se instance.
* `iptables` - Iptable rules.
* `labels` - Labels associated with this se group. Field introduced in 18.2.10.
* `least_load_core_selection` - Select core with least load for new flow.
* `license_tier` - Specifies the license tier which would be used. This field by default inherits the value from cloud. Enum options - ENTERPRISE_16, ENTERPRISE_18. Field introduced in 17.2.5.
* `license_type` - If no license type is specified then default license enforcement for the cloud type is chosen. Enum options - LIC_BACKEND_SERVERS, LIC_SOCKETS, LIC_CORES, LIC_HOSTS, LIC_SE_BANDWIDTH, LIC_METERED_SE_BANDWIDTH. Field introduced in 17.2.5.
* `log_disksz` - Maximum disk capacity (in mb) to be allocated to an se. This is exclusively used for debug and log data.
* `max_concurrent_external_hm` - Maximum number of external health monitors that can run concurrently in a service engine. This helps control the cpu and memory use by external health monitors. Special values are 0- 'value will be internally calculated based on cpu and memory'. Field introduced in 18.2.7.
* `max_cpu_usage` - When cpu usage on an se exceeds this threshold, virtual services hosted on this se may be rebalanced to other ses to reduce load. A new se may be created as part of this process. Allowed values are 40-90.
* `max_memory_per_mempool` - Max bytes that can be allocated in a single mempool. Field introduced in 18.1.5.
* `max_public_ips_per_lb` - Applicable to azure platform only. Maximum number of public ips per azure lb. Field introduced in 17.2.12, 18.1.2.
* `max_queues_per_vnic` - Maximum number of queues per vnic setting to '0' utilises all queues that are distributed across dispatcher cores. Allowed values are 0,1,2,4,8,16. Field introduced in 18.2.7.
* `max_rules_per_lb` - Applicable to azure platform only. Maximum number of rules per azure lb. Field introduced in 17.2.12, 18.1.2.
* `max_scaleout_per_vs` - Maximum number of active service engines for the virtual service. Allowed values are 1-64.
* `max_se` - Maximum number of services engines in this group. Allowed values are 0-1000.
* `max_vs_per_se` - Maximum number of virtual services that can be placed on a single service engine. Allowed values are 1-1000.
* `mem_reserve` - Boolean flag to set mem_reserve.
* `memory_for_config_update` - Indicates the percent of memory reserved for config updates. Allowed values are 0-100. Field introduced in 18.1.2.
* `memory_per_se` - Amount of memory for each of the service engine virtual machines. Changes to this setting do not affect existing ses.
* `mgmt_network_ref` - Management network to use for avi service engines. It is a reference to an object of type network.
* `mgmt_subnet` - Management subnet to use for avi service engines.
* `min_cpu_usage` - When cpu usage on an se falls below the minimum threshold, virtual services hosted on the se may be consolidated onto other underutilized ses. After consolidation, unused service engines may then be eligible for deletion. Allowed values are 20-60.
* `min_scaleout_per_vs` - Minimum number of active service engines for the virtual service. Allowed values are 1-64.
* `min_se` - Minimum number of services engines in this group (relevant for se autorebalance only). Allowed values are 0-1000. Field introduced in 17.2.13,18.1.3,18.2.1.
* `minimum_connection_memory` - Indicates the percent of memory reserved for connections. Allowed values are 0-100. Field introduced in 18.1.2.
* `n_log_streaming_threads` - Number of threads to use for log streaming. Allowed values are 1-100. Field introduced in 17.2.12, 18.1.2.
* `name` - Name of the object.
* `non_significant_log_throttle` - This setting limits the number of non-significant logs generated per second per core on this se. Default is 100 logs per second. Set it to zero (0) to disable throttling. Field introduced in 17.1.3.
* `num_dispatcher_cores` - Number of dispatcher cores (0,1,2,4,8 or 16). If set to 0, then number of dispatcher cores is deduced automatically.requires se reboot. Allowed values are 0,1,2,4,8,16. Field introduced in 17.2.12, 18.1.3, 18.2.1.
* `num_flow_cores_sum_changes_to_ignore` - Number of changes in num flow cores sum to ignore.
* `openstack_availability_zones` - Field introduced in 17.1.1.
* `openstack_mgmt_network_name` - Avi management network name.
* `openstack_mgmt_network_uuid` - Management network uuid.
* `os_reserved_memory` - Amount of extra memory to be reserved for use by the operating system on a service engine.
* `pcap_tx_mode` - Determines the pcap transmit mode of operation. Requires se reboot. Enum options - PCAP_TX_AUTO, PCAP_TX_SOCKET, PCAP_TX_RING. Field introduced in 18.2.8.
* `pcap_tx_ring_rd_balancing_factor` - In pcap mode, reserve a configured portion of tx ring resources for itself and  the remaining portion for the rx ring to achieve better balance in terms of queue depth. Requires se reboot. Allowed values are 10-100. Field introduced in 18.2.11.
* `per_app` - Per-app se mode is designed for deploying dedicated load balancers per app (vs). In this mode, each se is limited to a max of 2 vss. Vcpus in per-app ses count towards licensing usage at 25% rate.
* `per_vs_admission_control` - Enable/disable per vs level admission control.enabling this feature will cause the connection and packet throttling on a particular vs that has high packet buffer consumption. Field introduced in 18.2.12.
* `placement_mode` - If placement mode is 'auto', virtual services are automatically placed on service engines. Enum options - PLACEMENT_MODE_AUTO.
* `realtime_se_metrics` - Enable or disable real time se metrics.
* `reboot_on_panic` - Reboot the vm or host on kernel panic. Field introduced in 18.2.5.
* `se_bandwidth_type` - Select the se bandwidth for the bandwidth license. Enum options - SE_BANDWIDTH_UNLIMITED, SE_BANDWIDTH_25M, SE_BANDWIDTH_200M, SE_BANDWIDTH_1000M, SE_BANDWIDTH_10000M. Field introduced in 17.2.5.
* `se_delayed_flow_delete` - Delay the cleanup of flowtable entry. To be used under surveillance of avi support. Field introduced in 18.2.11.
* `se_deprovision_delay` - Duration to preserve unused service engine virtual machines before deleting them. If traffic to a virtual service were to spike up abruptly, this se would still be available to be utilized again rather than creating a new se. If this value is set to 0, controller will never delete any ses and administrator has to manually cleanup unused ses. Allowed values are 0-525600.
* `se_dos_profile` - Dict settings for serviceenginegroup.
* `se_dp_vnic_queue_stall_event_sleep` - Time (in seconds) service engine waits for after generating a vnic transmit queue stall event before resetting thenic. Field introduced in 18.2.5.
* `se_dp_vnic_queue_stall_threshold` - Number of consecutive transmit failures to look for before generating a vnic transmit queue stall event. Field introduced in 18.2.5.
* `se_dp_vnic_queue_stall_timeout` - Time (in milliseconds) to wait for network/nic recovery on detecting a transmit queue stall after which service engine resets the nic. Field introduced in 18.2.5.
* `se_dp_vnic_restart_on_queue_stall_count` - Number of consecutive transmit queue stall events in se_dp_vnic_stall_se_restart_window to look for before restarting se. Field introduced in 18.2.5.
* `se_dp_vnic_stall_se_restart_window` - Window of time (in seconds) during which se_dp_vnic_restart_on_queue_stall_count number of consecutive stalls results in a se restart. Field introduced in 18.2.5.
* `se_dpdk_pmd` - Determines if dpdk pool mode driver should be used or not   0  automatically determine based on hypervisor/nic type 1  unconditionally use dpdk poll mode driver 2  don't use dpdk poll mode driver.requires se reboot. Allowed values are 0-2. Field introduced in 18.1.3.
* `se_flow_probe_retries` - Flow probe retry count if no replies are received.requires se reboot. Allowed values are 0-5. Field introduced in 18.1.4, 18.2.1.
* `se_flow_probe_retry_timer` - Timeout in milliseconds for flow probe retries.requires se reboot. Allowed values are 20-50. Field introduced in 18.2.5.
* `se_ipc_udp_port` - Udp port for se_dp ipc in docker bridge mode. Field introduced in 17.1.2.
* `se_kni_burst_factor` - This knob controls the resource availability and burst size used between se datapath and kni. This helps in minimising packet drops when there is higher kni traffic (non-vip traffic from and to linux). The factor takes the following values      0-default. 1-doubles the burst size and kni resources. 2-quadruples the burst size and kni resources. Allowed values are 0-2. Field introduced in 18.2.6.
* `se_lro` - Enable or disable large receive optimization for vnics. Requires se reboot. Field introduced in 18.2.5.
* `se_mp_ring_retry_count` - The retry count for the multi-producer enqueue before yielding the cpu. To be used under surveillance of avi support. Field introduced in 18.2.11.
* `se_mtu` - Mtu for the vnics of ses in the se group. Allowed values are 512-9000. Field introduced in 18.2.8.
* `se_name_prefix` - Prefix to use for virtual machine name of service engines.
* `se_pcap_lookahead` - Enables lookahead mode of packet receive in pcap mode. Introduced to overcome an issue with hv_netvsc driver. Lookahead mode attempts to ensure that application and kernel's view of the receive rings are consistent. Field introduced in 18.2.3.
* `se_pcap_pkt_count` - Max number of packets the pcap interface can hold and if the value is 0 the optimum value will be chosen. The optimum value will be chosen based on se-memory, cloud type and number of interfaces.requires se reboot. Field introduced in 18.2.5.
* `se_pcap_pkt_sz` - Max size of each packet in the pcap interface. Requires se reboot. Field introduced in 18.2.5.
* `se_pcap_qdisc_bypass` - Bypass the kernel's traffic control layer, to deliver packets directly to the driver. Enabling this feature results in egress packets not being captured in host tcpdump. Note   brief packet reordering or loss may occur upon toggle. Field introduced in 18.2.6.
* `se_pcap_reinit_frequency` - Frequency in seconds at which periodically a pcap reinit check is triggered. May be used in conjunction with the configuration pcap_reinit_threshold. (valid range   15 mins - 12 hours, 0 - disables). Allowed values are 900-43200. Special values are 0- 'disable'. Field introduced in 17.2.13, 18.1.3, 18.2.1.
* `se_pcap_reinit_threshold` - Threshold for input packet receive errors in pcap mode exceeding which a pcap reinit is triggered. If not set, an unconditional reinit is performed. This value is checked every pcap_reinit_frequency interval. Field introduced in 17.2.13, 18.1.3, 18.2.1.
* `se_probe_port` - Tcp port on se where echo service will be run. Field introduced in 17.2.2.
* `se_remote_punt_udp_port` - Udp port for punted packets in docker bridge mode. Field introduced in 17.1.2.
* `se_rl_prop` - Rate limiter properties. Field introduced in 18.2.9.
* `se_rum_sampling_nav_interval` - Minimum time to wait on server between taking sampleswhen sampling the navigation timing data from the end user client. Field introduced in 18.2.6.
* `se_rum_sampling_nav_percent` - Percentage of navigation timing data from the end user client, used for sampling to get client insights. Field introduced in 18.2.6.
* `se_rum_sampling_res_interval` - Minimum time to wait on server between taking sampleswhen sampling the resource timing data from the end user client. Field introduced in 18.2.6.
* `se_rum_sampling_res_percent` - Percentage of resource timing data from the end user client used for sampling to get client insight. Field introduced in 18.2.6.
* `se_sb_dedicated_core` - Sideband traffic will be handled by a dedicated core.requires se reboot. Field introduced in 16.5.2, 17.1.9, 17.2.3.
* `se_sb_threads` - Number of sideband threads per se.requires se reboot. Allowed values are 1-128. Field introduced in 16.5.2, 17.1.9, 17.2.3.
* `se_thread_multiplier` - Multiplier for se threads based on vcpu. Allowed values are 1-10.
* `se_tracert_port_range` - Traceroute port range. Field introduced in 17.2.8.
* `se_tunnel_mode` - Determines if dsr from secondary se is active or not  0  automatically determine based on hypervisor type. 1  disable dsr unconditionally. 2  enable dsr unconditionally. Allowed values are 0-2. Field introduced in 17.1.1.
* `se_tunnel_udp_port` - Udp port for tunneled packets from secondary to primary se in docker bridge mode.requires se reboot. Field introduced in 17.1.3.
* `se_tx_batch_size` - Number of packets to batch for transmit to the nic. Requires se reboot. Field introduced in 18.2.5.
* `se_txq_threshold` - Once the tx queue of the dispatcher reaches this threshold, hardware queues are not polled for further packets. To be used under surveillance of avi support. Allowed values are 512-32768. Field introduced in 18.2.11.
* `se_udp_encap_ipc` - Determines if se-se ipc messages are encapsulated in a udp header  0  automatically determine based on hypervisor type. 1  use udp encap unconditionally.requires se reboot. Allowed values are 0-1. Field introduced in 17.1.2.
* `se_use_dpdk` - Determines if dpdk library should be used or not   0  automatically determine based on hypervisor type 1  use dpdk if pcap is not enabled 2  don't use dpdk. Allowed values are 0-2. Field introduced in 18.1.3.
* `se_vs_hb_max_pkts_in_batch` - Maximum number of aggregated vs heartbeat packets to send in a batch. Allowed values are 1-256. Field introduced in 17.1.1.
* `se_vs_hb_max_vs_in_pkt` - Maximum number of virtualservices for which heartbeat messages are aggregated in one packet. Allowed values are 1-1024. Field introduced in 17.1.1.
* `self_se_election` - Enable ses to elect a primary amongst themselves in the absence of a connectivity to controller. Field introduced in 18.1.2.
* `service_ip6_subnets` - Ipv6 subnets assigned to the se group. Required for vs group placement. Field introduced in 18.1.1.
* `service_ip_subnets` - Subnets assigned to the se group. Required for vs group placement. Field introduced in 17.1.1.
* `shm_minimum_config_memory` - Minimum required shared memory to apply any configuration. Allowed values are 0-100. Field introduced in 18.1.2.
* `significant_log_throttle` - This setting limits the number of significant logs generated per second per core on this se. Default is 100 logs per second. Set it to zero (0) to disable throttling. Field introduced in 17.1.3.
* `ssl_preprocess_sni_hostname` - (beta) preprocess ssl client hello for sni hostname extension.if set to true, this will apply sni child's ssl protocol(s), if they are different from sni parent's allowed ssl protocol(s). Field introduced in 17.2.12, 18.1.3.
* `tenant_ref` - It is a reference to an object of type tenant.
* `transient_shared_memory_max` - The threshold for the transient shared config memory in the se. Allowed values are 0-100. Field introduced in 18.2.10.
* `udf_log_throttle` - This setting limits the number of udf logs generated per second per core on this se. Udf logs are generated due to the configured client log filters or the rules with logging enabled. Default is 100 logs per second. Set it to zero (0) to disable throttling. Field introduced in 17.1.3.
* `use_standard_alb` - Use standard sku azure load balancer. By default cloud level flag is set. If not set, it inherits/uses the use_standard_alb flag from the cloud. Field introduced in 18.2.3.
* `uuid` - Unique object identifier of the object.
* `vcenter_clusters` - Dict settings for serviceenginegroup.
* `vcenter_datastore_mode` - Enum options - vcenter_datastore_any, vcenter_datastore_local, vcenter_datastore_shared.
* `vcenter_datastores` - List of list.
* `vcenter_datastores_include` - Boolean flag to set vcenter_datastores_include.
* `vcenter_folder` - Folder to place all the service engine virtual machines in vcenter.
* `vcenter_hosts` - Dict settings for serviceenginegroup.
* `vcpus_per_se` - Number of vcpus for each of the service engine virtual machines. Changes to this setting do not affect existing ses.
* `vip_asg` - When vip_asg is set, vip configuration will be managed by avi.user will be able to configure vip_asg or vips individually at the time of create. Field introduced in 17.2.12, 18.1.2.
* `vs_host_redundancy` - Ensure primary and secondary service engines are deployed on different physical hosts.
* `vs_scalein_timeout` - Time to wait for the scaled in se to drain existing flows before marking the scalein done.
* `vs_scalein_timeout_for_upgrade` - During se upgrade, time to wait for the scaled-in se to drain existing flows before marking the scalein done.
* `vs_scaleout_timeout` - Time to wait for the scaled out se to become ready before marking the scaleout done.
* `vs_se_scaleout_additional_wait_time` - Wait time for sending scaleout ready notification after virtual service is marked up. In certain deployments, there may be an additional delay to accept traffic. For example, for bgp, some time is needed for route advertisement. Allowed values are 0-20. Field introduced in 18.1.5,18.2.1.
* `vs_se_scaleout_ready_timeout` - Timeout in seconds for service engine to sendscaleout ready notification of a virtual service. Allowed values are 0-90. Field introduced in 18.1.5,18.2.1.
* `vs_switchover_timeout` - During se upgrade in a legacy active/standby segroup, time to wait for the new primary se to accept flows before marking the switchover done. Field introduced in 17.2.13,18.1.4,18.2.1.
* `vss_placement` - Parameters to place virtual services on only a subset of the cores of an se. Field introduced in 17.2.5.
* `vss_placement_enabled` - If set, virtual services will be placed on only a subset of the cores of an se. Field introduced in 18.1.1.
* `waf_mempool` - Enable memory pool for waf.requires se reboot. Field introduced in 17.2.3.
* `waf_mempool_size` - Memory pool size used for waf.requires se reboot. Field introduced in 17.2.3.

