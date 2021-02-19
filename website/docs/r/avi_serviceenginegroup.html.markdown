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
* `accelerated_networking` - (Optional) Enable accelerated networking option for azure se. Accelerated networking enables single root i/o virtualization (sr-iov) to a se vm. This improves networking performance. Field introduced in 17.2.14,18.1.5,18.2.1.
* `active_standby` - (Optional) Service engines in active/standby mode for ha failover.
* `aggressive_failure_detection` - (Optional) Enable aggressive failover configuration for ha. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `algo` - (Optional) In compact placement, virtual services are placed on existing ses until max_vs_per_se limit is reached. Enum options - PLACEMENT_ALGO_PACKED, PLACEMENT_ALGO_DISTRIBUTED.
* `allow_burst` - (Optional) Allow ses to be created using burst license. Field introduced in 17.2.5.
* `app_cache_percent` - (Optional) A percent value of total se memory reserved for applicationcaching. This is an se bootup property and requires se restart.requires se reboot. Allowed values are 0 - 100. Special values are 0- 'disable'. Field introduced in 18.2.3. Unit is percent. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition. Special default for basic edition is 0, essentials edition is 0, enterprise is 10.
* `app_cache_threshold` - (Optional) The max memory that can be allocated for the app cache. This value will act as an upper bound on the cache size specified in app_cache_percent. Special values are 0- 'disable'. Field introduced in 20.1.1. Unit is gb.
* `app_learning_memory_percent` - (Optional) A percent value of total se memory reserved for application learning. This is an se bootup property and requires se restart. Allowed values are 0 - 10. Field introduced in 18.2.3. Unit is percent.
* `archive_shm_limit` - (Optional) Amount of se memory in gb until which shared memory is collected in core archive. Field introduced in 17.1.3. Unit is gb.
* `async_ssl` - (Optional) Ssl handshakes will be handled by dedicated ssl threads.requires se reboot. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `async_ssl_threads` - (Optional) Number of async ssl threads per se_dp.requires se reboot. Allowed values are 1-16.
* `auto_rebalance` - (Optional) If set, virtual services will be automatically migrated when load on an se is less than minimum or more than maximum thresholds. Only alerts are generated when the auto_rebalance is not set. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `auto_rebalance_capacity_per_se` - (Optional) Capacities of se for auto rebalance for each criteria. Field introduced in 17.2.4.
* `auto_rebalance_criteria` - (Optional) Set of criteria for se auto rebalance. Enum options - SE_AUTO_REBALANCE_CPU, SE_AUTO_REBALANCE_PPS, SE_AUTO_REBALANCE_MBPS, SE_AUTO_REBALANCE_OPEN_CONNS, SE_AUTO_REBALANCE_CPS. Field introduced in 17.2.3.
* `auto_rebalance_interval` - (Optional) Frequency of rebalance, if 'auto rebalance' is enabled. Unit is sec.
* `auto_redistribute_active_standby_load` - (Optional) Redistribution of virtual services from the takeover se to the replacement se can cause momentary traffic loss. If the auto-redistribute load option is left in its default off state, any desired rebalancing requires calls to rest api. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `availability_zone_refs` - (Optional) Availability zones for virtual service high availability. It is a reference to an object of type availabilityzone. Field introduced in 20.1.1.
* `bgp_state_update_interval` - (Optional) Bgp peer state update interval. Allowed values are 5-100. Field introduced in 17.2.14,18.1.5,18.2.1. Unit is sec.
* `buffer_se` - (Optional) Excess service engine capacity provisioned for ha failover.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `compress_ip_rules_for_each_ns_subnet` - (Optional) Compress ip rules into a single subnet based ip rule for each north-south ipam subnet configured in pcap mode in openshift/kubernetes node. Field introduced in 18.2.9, 20.1.1.
* `config_debugs_on_all_cores` - (Optional) Enable config debugs on all cores of se. Field introduced in 17.2.13,18.1.5,18.2.1.
* `connection_memory_percentage` - (Optional) Percentage of memory for connection state. This will come at the expense of memory used for http in-memory cache. Allowed values are 10-90. Unit is percent.
* `core_shm_app_cache` - (Optional) Include shared memory for app cache in core file.requires se reboot. Field introduced in 18.2.8, 20.1.1.
* `core_shm_app_learning` - (Optional) Include shared memory for app learning in core file.requires se reboot. Field introduced in 18.2.8, 20.1.1.
* `cpu_reserve` - (Optional) Boolean flag to set cpu_reserve.
* `cpu_socket_affinity` - (Optional) Allocate all the cpu cores for the service engine virtual machines  on the same cpu socket. Applicable only for vcenter cloud.
* `custom_securitygroups_data` - (Optional) Custom security groups to be associated with data vnics for se instances in openstack and aws clouds. Field introduced in 17.1.3.
* `custom_securitygroups_mgmt` - (Optional) Custom security groups to be associated with management vnic for se instances in openstack and aws clouds. Field introduced in 17.1.3.
* `custom_tag` - (Optional) Custom tag will be used to create the tags for se instance in aws. Note this is not the same as the prefix for se name.
* `data_network_id` - (Optional) Subnet used to spin up the data nic for service engines, used only for azure cloud. Overrides the cloud level setting for service engine subnet. Field introduced in 18.2.3.
* `datascript_timeout` - (Optional) Number of instructions before datascript times out. Allowed values are 0-100000000. Field introduced in 18.2.3.
* `dedicated_dispatcher_core` - (Optional) Dedicate the core that handles packet receive/transmit from the network to just the dispatching function. Don't use it for tcp/ip and ssl functions.
* `description` - (Optional) User defined description for the object.
* `disable_avi_securitygroups` - (Optional) By default, avi creates and manages security groups along with custom sg provided by user. Set this to true to disallow avi to create and manage new security groups. Avi will only make use of custom security groups provided by user. This option is supported for aws and openstack cloud types. Field introduced in 17.2.13,18.1.4,18.2.1.
* `disable_csum_offloads` - (Optional) Stop using tcp/udp and ip checksum offload features of nics. Field introduced in 17.1.14, 17.2.5, 18.1.1.
* `disable_flow_probes` - (Optional) Disable flow probes for scaled out vs'es. Field introduced in 20.1.3.
* `disable_gro` - (Optional) Disable generic receive offload (gro) in dpdk poll-mode driver packet receive path. Gro is on by default on nics that do not support lro (large receive offload) or do not gain performance boost from lro. Field introduced in 17.2.5, 18.1.1.
* `disable_se_memory_check` - (Optional) If set, disable the config memory check done in service engine. Field introduced in 18.1.2.
* `disable_tso` - (Optional) Disable tcp segmentation offload (tso) in dpdk poll-mode driver packet transmit path. Tso is on by default on nics that support it. Field introduced in 17.2.5, 18.1.1.
* `disk_per_se` - (Optional) Amount of disk space for each of the service engine virtual machines. Unit is gb.
* `distribute_load_active_standby` - (Optional) Use both the active and standby service engines for virtual service placement in the legacy active standby ha mode. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `distribute_queues` - (Optional) Distributes queue ownership among cores so multiple cores handle dispatcher duties. Requires se reboot. Deprecated from 18.2.8, instead use max_queues_per_vnic. Field introduced in 17.2.8. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `distribute_vnics` - (Optional) Distributes vnic ownership among cores so multiple cores handle dispatcher duties.requires se reboot. Field introduced in 18.2.5.
* `dp_aggressive_deq_interval_msec` - (Optional) Dequeue interval for receive queue from se_dp in aggressive mode. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds.
* `dp_aggressive_enq_interval_msec` - (Optional) Enqueue interval for request queue to se_dp in aggressive mode. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds.
* `dp_aggressive_hb_frequency` - (Optional) Frequency of se - se hb messages when aggressive failure mode detection is enabled. Field introduced in 20.1.3. Unit is milliseconds.
* `dp_aggressive_hb_timeout_count` - (Optional) Consecutive hb failures after which failure is reported to controller,when aggressive failure mode detection is enabled. Field introduced in 20.1.3.
* `dp_deq_interval_msec` - (Optional) Dequeue interval for receive queue from se_dp. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds.
* `dp_enq_interval_msec` - (Optional) Enqueue interval for request queue to se_dp. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds.
* `dp_hb_frequency` - (Optional) Frequency of se - se hb messages when aggressive failure mode detection is not enabled. Field introduced in 20.1.3. Unit is milliseconds.
* `dp_hb_timeout_count` - (Optional) Consecutive hb failures after which failure is reported to controller, when aggressive failure mode detection is not enabled. Field introduced in 20.1.3.
* `enable_gratarp_permanent` - (Optional) Enable gratarp for vip_ip. Field introduced in 18.2.3.
* `enable_hsm_priming` - (Optional) (this is a beta feature). Enable hsm key priming. If enabled, key handles on the hsm will be synced to se before processing client connections. Field introduced in 17.2.7, 18.1.1.
* `enable_multi_lb` - (Optional) Applicable only for azure cloud with basic sku lb. If set, additional azure lbs will be automatically created if resources in existing lb are exhausted. Field introduced in 17.2.10, 18.1.2.
* `enable_pcap_tx_ring` - (Optional) Enable tx ring support in pcap mode of operation. Tso feature is not supported with tx ring enabled. Deprecated from 18.2.8, instead use pcap_tx_mode. Requires se reboot. Field introduced in 18.2.5.
* `ephemeral_portrange_end` - (Optional) End local ephemeral port number for outbound connections. Field introduced in 17.2.13, 18.1.5, 18.2.1.
* `ephemeral_portrange_start` - (Optional) Start local ephemeral port number for outbound connections. Field introduced in 17.2.13, 18.1.5, 18.2.1.
* `extra_config_multiplier` - (Optional) Multiplier for extra config to support large vs/pool config.
* `extra_shared_config_memory` - (Optional) Extra config memory to support large geo db configuration. Field introduced in 17.1.1. Unit is mb.
* `flow_table_new_syn_max_entries` - (Optional) Maximum number of flow table entries that have not completed tcp three-way handshake yet. Field introduced in 17.2.5.
* `free_list_size` - (Optional) Number of entries in the free list. Field introduced in 17.2.10, 18.1.2.
* `gcp_config` - (Optional) Google cloud platform, service engine group configuration. Field introduced in 20.1.3.
* `gratarp_permanent_periodicity` - (Optional) Gratarp periodicity for vip-ip. Allowed values are 5-30. Field introduced in 18.2.3. Unit is min.
* `ha_mode` - (Optional) High availability mode for all the virtual services using this service engine group. Enum options - HA_MODE_SHARED_PAIR, HA_MODE_SHARED, HA_MODE_LEGACY_ACTIVE_STANDBY. Allowed in basic(allowed values- ha_mode_legacy_active_standby) edition, essentials(allowed values- ha_mode_legacy_active_standby) edition, enterprise edition. Special default for basic edition is ha_mode_legacy_active_standby, essentials edition is ha_mode_legacy_active_standby, enterprise is ha_mode_shared.
* `handle_per_pkt_attack` - (Optional) Configuration to handle per packet attack handling.for example, dns reflection attack is a type of attack where a response packet is sent to the dns vs.this configuration tells if such packets should be dropped without further processing. Field introduced in 20.1.3.
* `hardwaresecuritymodulegroup_ref` - (Optional) It is a reference to an object of type hardwaresecuritymodulegroup.
* `heap_minimum_config_memory` - (Optional) Minimum required heap memory to apply any configuration. Allowed values are 0-100. Field introduced in 18.1.2. Unit is mb.
* `hm_on_standby` - (Optional) Enable active health monitoring from the standby se for all placed virtual services. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition. Special default for basic edition is false, essentials edition is false, enterprise is true.
* `host_attribute_key` - (Optional) Key of a (key, value) pair identifying a label for a set of nodes usually in container clouds. Needs to be specified together with host_attribute_value. Ses can be configured differently including ha modes across different se groups. May also be used for isolation between different classes of virtualservices. Virtualservices' se group may be specified via annotations/labels. A openshift/kubernetes namespace maybe annotated with a matching se group label as openshift.io/node-selector  apptype=prod. When multiple se groups are used in a cloud with host attributes specified,just a single se group can exist as a match-all se group without a host_attribute_key.
* `host_attribute_value` - (Optional) Value of a (key, value) pair identifying a label for a set of nodes usually in container clouds. Needs to be specified together with host_attribute_key.
* `host_gateway_monitor` - (Optional) Enable the host gateway monitor when service engine is deployed as docker container. Disabled by default. Field introduced in 17.2.4.
* `hypervisor` - (Optional) Override default hypervisor. Enum options - DEFAULT, VMWARE_ESX, KVM, VMWARE_VSAN, XEN.
* `ignore_rtt_threshold` - (Optional) Ignore rtt samples if it is above threshold. Field introduced in 17.1.6,17.2.2. Unit is milliseconds.
* `ingress_access_data` - (Optional) Program se security group ingress rules to allow vip data access from remote cidr type. Enum options - SG_INGRESS_ACCESS_NONE, SG_INGRESS_ACCESS_ALL, SG_INGRESS_ACCESS_VPC. Field introduced in 17.1.5.
* `ingress_access_mgmt` - (Optional) Program se security group ingress rules to allow ssh/icmp management access from remote cidr type. Enum options - SG_INGRESS_ACCESS_NONE, SG_INGRESS_ACCESS_ALL, SG_INGRESS_ACCESS_VPC. Field introduced in 17.1.5.
* `instance_flavor` - (Optional) Instance/flavor name for se instance.
* `instance_flavor_info` - (Optional) Additional information associated with instance_flavor. Field introduced in 20.1.1.
* `iptables` - (Optional) Iptable rules. Maximum of 128 items allowed.
* `labels` - (Optional) Labels associated with this se group. Field introduced in 20.1.1. Maximum of 1 items allowed.
* `least_load_core_selection` - (Optional) Select core with least load for new flow.
* `license_tier` - (Optional) Specifies the license tier which would be used. This field by default inherits the value from cloud. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS. Field introduced in 17.2.5.
* `license_type` - (Optional) If no license type is specified then default license enforcement for the cloud type is chosen. Enum options - LIC_BACKEND_SERVERS, LIC_SOCKETS, LIC_CORES, LIC_HOSTS, LIC_SE_BANDWIDTH, LIC_METERED_SE_BANDWIDTH. Field introduced in 17.2.5.
* `log_disksz` - (Optional) Maximum disk capacity (in mb) to be allocated to an se. This is exclusively used for debug and log data. Unit is mb.
* `log_malloc_failure` - (Optional) Se will log memory allocation related failure to the se_trace file, wherever available. Field introduced in 20.1.2. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `max_concurrent_external_hm` - (Optional) Maximum number of external health monitors that can run concurrently in a service engine. This helps control the cpu and memory use by external health monitors. Special values are 0- 'value will be internally calculated based on cpu and memory'. Field introduced in 18.2.7.
* `max_cpu_usage` - (Optional) When cpu usage on an se exceeds this threshold, virtual services hosted on this se may be rebalanced to other ses to reduce load. A new se may be created as part of this process. Allowed values are 40-90. Unit is percent.
* `max_memory_per_mempool` - (Optional) Max bytes that can be allocated in a single mempool. Field introduced in 18.1.5. Unit is mb.
* `max_num_se_dps` - (Optional) Configures the maximum number of se_dp processes created on the se, requires se reboot. If not configured, defaults to the number of cpus on the se. This should only be used if user wants to limit the number of se_dps to less than the available cpus on the se. Allowed values are 1-128. Field introduced in 20.1.1. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `max_public_ips_per_lb` - (Optional) Applicable to azure platform only. Maximum number of public ips per azure lb. Field introduced in 17.2.12, 18.1.2.
* `max_queues_per_vnic` - (Optional) Maximum number of queues per vnic setting to '0' utilises all queues that are distributed across dispatcher cores. Allowed values are 0,1,2,4,8,16. Field introduced in 18.2.7, 20.1.1.
* `max_rules_per_lb` - (Optional) Applicable to azure platform only. Maximum number of rules per azure lb. Field introduced in 17.2.12, 18.1.2.
* `max_scaleout_per_vs` - (Optional) Maximum number of active service engines for the virtual service. Allowed values are 1-64.
* `max_se` - (Optional) Maximum number of services engines in this group. Allowed values are 0-1000.
* `max_vs_per_se` - (Optional) Maximum number of virtual services that can be placed on a single service engine. Allowed values are 1-1000.
* `mem_reserve` - (Optional) Boolean flag to set mem_reserve.
* `memory_for_config_update` - (Optional) Indicates the percent of memory reserved for config updates. Allowed values are 0-100. Field introduced in 18.1.2. Unit is percent.
* `memory_per_se` - (Optional) Amount of memory for each of the service engine virtual machines. Changes to this setting do not affect existing ses.
* `mgmt_network_ref` - (Optional) Management network to use for avi service engines. It is a reference to an object of type network.
* `mgmt_subnet` - (Optional) Management subnet to use for avi service engines.
* `min_cpu_usage` - (Optional) When cpu usage on an se falls below the minimum threshold, virtual services hosted on the se may be consolidated onto other underutilized ses. After consolidation, unused service engines may then be eligible for deletion. Allowed values are 20-60. Unit is percent.
* `min_scaleout_per_vs` - (Optional) Minimum number of active service engines for the virtual service. Allowed values are 1-64.
* `min_se` - (Optional) Minimum number of services engines in this group (relevant for se autorebalance only). Allowed values are 0-1000. Field introduced in 17.2.13,18.1.3,18.2.1.
* `minimum_connection_memory` - (Optional) Indicates the percent of memory reserved for connections. Allowed values are 0-100. Field introduced in 18.1.2. Unit is percent.
* `n_log_streaming_threads` - (Optional) Number of threads to use for log streaming. Allowed values are 1-100. Field introduced in 17.2.12, 18.1.2.
* `netlink_poller_threads` - (Optional) Number of threads to poll for netlink messages excluding the thread for default namespace. Requires se reboot. Allowed values are 1-32. Field introduced in 20.1.3.
* `netlink_sock_buf_size` - (Optional) Socket buffer size for the netlink sockets. Requires se reboot. Allowed values are 1-128. Field introduced in 20.1.3. Unit is mega_bytes.
* `non_significant_log_throttle` - (Optional) This setting limits the number of non-significant logs generated per second per core on this se. Default is 100 logs per second. Set it to zero (0) to deactivate throttling. Field introduced in 17.1.3. Unit is per_second.
* `ns_helper_deq_interval_msec` - (Optional) Dequeue interval for receive queue from ns helper. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is milliseconds.
* `num_dispatcher_cores` - (Optional) Number of dispatcher cores (0,1,2,4,8 or 16). If set to 0, then number of dispatcher cores is deduced automatically.requires se reboot. Allowed values are 0,1,2,4,8,16. Field introduced in 17.2.12, 18.1.3, 18.2.1.
* `num_flow_cores_sum_changes_to_ignore` - (Optional) Number of changes in num flow cores sum to ignore.
* `objsync_config` - (Optional) Configuration knobs for interse object distribution. Field introduced in 20.1.3.
* `objsync_port` - (Optional) Tcp port on se management interface for interse object distribution. Supported only for externally managed security groups. Not supported on full access deployments. Requires se reboot. Field introduced in 20.1.3.
* `openstack_availability_zones` - (Optional) Field introduced in 17.1.1. Maximum of 5 items allowed.
* `openstack_mgmt_network_name` - (Optional) Avi management network name.
* `openstack_mgmt_network_uuid` - (Optional) Management network uuid.
* `os_reserved_memory` - (Optional) Amount of extra memory to be reserved for use by the operating system on a service engine. Unit is mb.
* `pcap_tx_mode` - (Optional) Determines the pcap transmit mode of operation. Requires se reboot. Enum options - PCAP_TX_AUTO, PCAP_TX_SOCKET, PCAP_TX_RING. Field introduced in 18.2.8, 20.1.1.
* `pcap_tx_ring_rd_balancing_factor` - (Optional) In pcap mode, reserve a configured portion of tx ring resources for itself and the remaining portion for the rx ring to achieve better balance in terms of queue depth. Requires se reboot. Allowed values are 10-100. Field introduced in 20.1.3. Unit is percent.
* `per_app` - (Optional) Per-app se mode is designed for deploying dedicated load balancers per app (vs). In this mode, each se is limited to a max of 2 vss. Vcpus in per-app ses count towards licensing usage at 25% rate. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `per_vs_admission_control` - (Optional) Enable/disable per vs level admission control.enabling this feature will cause the connection and packet throttling on a particular vs that has high packet buffer consumption. Field introduced in 20.1.3.
* `placement_mode` - (Optional) If placement mode is 'auto', virtual services are automatically placed on service engines. Enum options - PLACEMENT_MODE_AUTO.
* `realtime_se_metrics` - (Optional) Enable or deactivate real time se metrics.
* `reboot_on_panic` - (Optional) Reboot the vm or host on kernel panic. Field introduced in 18.2.5.
* `resync_time_interval` - (Optional) Time interval to re-sync se's time with wall clock time. Allowed values are 8-600000. Field introduced in 20.1.1. Unit is milliseconds.
* `se_bandwidth_type` - (Optional) Select the se bandwidth for the bandwidth license. Enum options - SE_BANDWIDTH_UNLIMITED, SE_BANDWIDTH_25M, SE_BANDWIDTH_200M, SE_BANDWIDTH_1000M, SE_BANDWIDTH_10000M. Field introduced in 17.2.5. Allowed in basic(allowed values- se_bandwidth_unlimited) edition, essentials(allowed values- se_bandwidth_unlimited) edition, enterprise edition.
* `se_delayed_flow_delete` - (Optional) Delay the cleanup of flowtable entry. To be used under surveillance of avi support. Field introduced in 20.1.2. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `se_deprovision_delay` - (Optional) Duration to preserve unused service engine virtual machines before deleting them. If traffic to a virtual service were to spike up abruptly, this se would still be available to be utilized again rather than creating a new se. If this value is set to 0, controller will never delete any ses and administrator has to manually cleanup unused ses. Allowed values are 0-525600. Unit is min.
* `se_dos_profile` - (Optional) Dict settings for serviceenginegroup.
* `se_dp_hm_drops` - (Optional) Internal only. Used to simulate se - se hb failure. Field introduced in 20.1.3.
* `se_dp_max_hb_version` - (Optional) The highest supported se-se heartbeat protocol version. This version is reported by secondary se to primary se in heartbeat response messages. Allowed values are 1-3. Field introduced in 20.1.1.
* `se_dp_vnic_queue_stall_event_sleep` - (Optional) Time (in seconds) service engine waits for after generating a vnic transmit queue stall event before resetting thenic. Field introduced in 18.2.5.
* `se_dp_vnic_queue_stall_threshold` - (Optional) Number of consecutive transmit failures to look for before generating a vnic transmit queue stall event. Field introduced in 18.2.5.
* `se_dp_vnic_queue_stall_timeout` - (Optional) Time (in milliseconds) to wait for network/nic recovery on detecting a transmit queue stall after which service engine resets the nic. Field introduced in 18.2.5.
* `se_dp_vnic_restart_on_queue_stall_count` - (Optional) Number of consecutive transmit queue stall events in se_dp_vnic_stall_se_restart_window to look for before restarting se. Field introduced in 18.2.5.
* `se_dp_vnic_stall_se_restart_window` - (Optional) Window of time (in seconds) during which se_dp_vnic_restart_on_queue_stall_count number of consecutive stalls results in a se restart. Field introduced in 18.2.5.
* `se_dpdk_pmd` - (Optional) Determines if dpdk pool mode driver should be used or not   0  automatically determine based on hypervisor/nic type 1  unconditionally use dpdk poll mode driver 2  don't use dpdk poll mode driver.requires se reboot. Allowed values are 0-2. Field introduced in 18.1.3.
* `se_flow_probe_retries` - (Optional) Flow probe retry count if no replies are received.requires se reboot. Allowed values are 0-5. Field introduced in 18.1.4, 18.2.1.
* `se_flow_probe_retry_timer` - (Optional) Timeout in milliseconds for flow probe retries.requires se reboot. Allowed values are 20-50. Field introduced in 18.2.5. Unit is milliseconds.
* `se_group_analytics_policy` - (Optional) Analytics policy for serviceenginegroup. Field introduced in 20.1.3.
* `se_hyperthreaded_mode` - (Optional) Controls the distribution of se data path processes on cpus which support hyper-threading. Requires hyper-threading to be enabled at host level. Requires se reboot. For more details please refer to se placement kb. Enum options - SE_CPU_HT_AUTO, SE_CPU_HT_SPARSE_DISPATCHER_PRIORITY, SE_CPU_HT_SPARSE_PROXY_PRIORITY, SE_CPU_HT_PACKED_CORES. Field introduced in 20.1.1.
* `se_ip_encap_ipc` - (Optional) Determines if se-se ipc messages are encapsulated in an ip header       0        automatically determine based on hypervisor type    1        use ip encap unconditionally    ~[0,1]   don't use ip encaprequires se reboot. Field introduced in 20.1.3.
* `se_kni_burst_factor` - (Optional) This knob controls the resource availability and burst size used between se datapath and kni. This helps in minimising packet drops when there is higher kni traffic (non-vip traffic from and to linux). The factor takes the following values      0-default. 1-doubles the burst size and kni resources. 2-quadruples the burst size and kni resources. Allowed values are 0-2. Field introduced in 18.2.6.
* `se_l3_encap_ipc` - (Optional) Determines if se-se ipc messages use se interface ip instead of vip        0        automatically determine based on hypervisor type    1        use se interface ip unconditionally    ~[0,1]   don't use se interface iprequires se reboot. Field introduced in 20.1.3.
* `se_lro` - (Optional) Enable or disable large receive optimization for vnics. Requires se reboot. Field introduced in 18.2.5.
* `se_mp_ring_retry_count` - (Optional) The retry count for the multi-producer enqueue before yielding the cpu. To be used under surveillance of avi support. Field introduced in 20.1.3. Allowed in basic(allowed values- 500) edition, essentials(allowed values- 500) edition, enterprise edition.
* `se_mtu` - (Optional) Mtu for the vnics of ses in the se group. Allowed values are 512-9000. Field introduced in 18.2.8, 20.1.1.
* `se_name_prefix` - (Optional) Prefix to use for virtual machine name of service engines.
* `se_pcap_lookahead` - (Optional) Enables lookahead mode of packet receive in pcap mode. Introduced to overcome an issue with hv_netvsc driver. Lookahead mode attempts to ensure that application and kernel's view of the receive rings are consistent. Field introduced in 18.2.3.
* `se_pcap_pkt_count` - (Optional) Max number of packets the pcap interface can hold and if the value is 0 the optimum value will be chosen. The optimum value will be chosen based on se-memory, cloud type and number of interfaces.requires se reboot. Field introduced in 18.2.5.
* `se_pcap_pkt_sz` - (Optional) Max size of each packet in the pcap interface. Requires se reboot. Field introduced in 18.2.5.
* `se_pcap_qdisc_bypass` - (Optional) Bypass the kernel's traffic control layer, to deliver packets directly to the driver. Enabling this feature results in egress packets not being captured in host tcpdump. Note   brief packet reordering or loss may occur upon toggle. Field introduced in 18.2.6.
* `se_pcap_reinit_frequency` - (Optional) Frequency in seconds at which periodically a pcap reinit check is triggered. May be used in conjunction with the configuration pcap_reinit_threshold. (valid range   15 mins - 12 hours, 0 - disables). Allowed values are 900-43200. Special values are 0- 'disable'. Field introduced in 17.2.13, 18.1.3, 18.2.1. Unit is sec.
* `se_pcap_reinit_threshold` - (Optional) Threshold for input packet receive errors in pcap mode exceeding which a pcap reinit is triggered. If not set, an unconditional reinit is performed. This value is checked every pcap_reinit_frequency interval. Field introduced in 17.2.13, 18.1.3, 18.2.1. Unit is metric_count.
* `se_probe_port` - (Optional) Tcp port on se where echo service will be run. Field introduced in 17.2.2.
* `se_rl_prop` - (Optional) Rate limiter properties. Field introduced in 20.1.1.
* `se_rum_sampling_nav_interval` - (Optional) Minimum time to wait on server between taking sampleswhen sampling the navigation timing data from the end user client. Field introduced in 18.2.6. Unit is sec.
* `se_rum_sampling_nav_percent` - (Optional) Percentage of navigation timing data from the end user client, used for sampling to get client insights. Field introduced in 18.2.6.
* `se_rum_sampling_res_interval` - (Optional) Minimum time to wait on server between taking sampleswhen sampling the resource timing data from the end user client. Field introduced in 18.2.6. Unit is sec.
* `se_rum_sampling_res_percent` - (Optional) Percentage of resource timing data from the end user client used for sampling to get client insight. Field introduced in 18.2.6.
* `se_sb_dedicated_core` - (Optional) Sideband traffic will be handled by a dedicated core.requires se reboot. Field introduced in 16.5.2, 17.1.9, 17.2.3.
* `se_sb_threads` - (Optional) Number of sideband threads per se.requires se reboot. Allowed values are 1-128. Field introduced in 16.5.2, 17.1.9, 17.2.3.
* `se_thread_multiplier` - (Optional) Multiplier for se threads based on vcpu. Allowed values are 1-10. Allowed in basic(allowed values- 1) edition, essentials(allowed values- 1) edition, enterprise edition.
* `se_tracert_port_range` - (Optional) Traceroute port range. Field introduced in 17.2.8.
* `se_tunnel_mode` - (Optional) Determines if direct secondary return (dsr) from secondary se is active or not  0  automatically determine based on hypervisor type. 1  enable tunnel mode - dsr is unconditionally disabled. 2  disable tunnel mode - dsr is unconditionally enabled. Tunnel mode can be enabled or disabled at run-time. Allowed values are 0-2. Field introduced in 17.1.1. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `se_tunnel_udp_port` - (Optional) Udp port for tunneled packets from secondary to primary se in docker bridge mode.requires se reboot. Field introduced in 17.1.3.
* `se_tx_batch_size` - (Optional) Number of packets to batch for transmit to the nic. Requires se reboot. Field introduced in 18.2.5.
* `se_txq_threshold` - (Optional) Once the tx queue of the dispatcher reaches this threshold, hardware queues are not polled for further packets. To be used under surveillance of avi support. Allowed values are 512-32768. Field introduced in 20.1.2. Allowed in basic(allowed values- 2048) edition, essentials(allowed values- 2048) edition, enterprise edition.
* `se_udp_encap_ipc` - (Optional) Determines if se-se ipc messages are encapsulated in a udp header  0  automatically determine based on hypervisor type. 1  use udp encap unconditionally.requires se reboot. Allowed values are 0-1. Field introduced in 17.1.2.
* `se_use_dpdk` - (Optional) Determines if dpdk library should be used or not   0  automatically determine based on hypervisor type 1  use dpdk if pcap is not enabled 2  don't use dpdk. Allowed values are 0-2. Field introduced in 18.1.3.
* `se_vnic_tx_sw_queue_flush_frequency` - (Optional) Configure the frequency in milliseconds of software transmit spillover queue flush when enabled. This is necessary to flush any packets in the spillover queue in the absence of a packet transmit in the normal course of operation. Allowed values are 50-500. Special values are 0- 'disable'. Field introduced in 20.1.1. Unit is milliseconds.
* `se_vnic_tx_sw_queue_size` - (Optional) Configure the size of software transmit spillover queue when enabled. Requires se reboot. Allowed values are 128-2048. Field introduced in 20.1.1.
* `se_vs_hb_max_pkts_in_batch` - (Optional) Maximum number of aggregated vs heartbeat packets to send in a batch. Allowed values are 1-256. Field introduced in 17.1.1.
* `se_vs_hb_max_vs_in_pkt` - (Optional) Maximum number of virtualservices for which heartbeat messages are aggregated in one packet. Allowed values are 1-1024. Field introduced in 17.1.1.
* `self_se_election` - (Optional) Enable ses to elect a primary amongst themselves in the absence of a connectivity to controller. Field introduced in 18.1.2. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `send_se_ready_timeout` - (Optional) Timeout for sending se_ready without ns helper registration completion. Allowed values are 10-600. Field introduced in 21.1.1. Unit is seconds.
* `service_ip6_subnets` - (Optional) Ipv6 subnets assigned to the se group. Required for vs group placement. Field introduced in 18.1.1. Maximum of 128 items allowed.
* `service_ip_subnets` - (Optional) Subnets assigned to the se group. Required for vs group placement. Field introduced in 17.1.1. Maximum of 128 items allowed.
* `shm_minimum_config_memory` - (Optional) Minimum required shared memory to apply any configuration. Allowed values are 0-100. Field introduced in 18.1.2. Unit is mb.
* `significant_log_throttle` - (Optional) This setting limits the number of significant logs generated per second per core on this se. Default is 100 logs per second. Set it to zero (0) to deactivate throttling. Field introduced in 17.1.3. Unit is per_second.
* `ssl_preprocess_sni_hostname` - (Optional) (beta) preprocess ssl client hello for sni hostname extension.if set to true, this will apply sni child's ssl protocol(s), if they are different from sni parent's allowed ssl protocol(s). Field introduced in 17.2.12, 18.1.3.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `transient_shared_memory_max` - (Optional) The threshold for the transient shared config memory in the se. Allowed values are 0-100. Field introduced in 20.1.1. Unit is percent.
* `udf_log_throttle` - (Optional) This setting limits the number of udf logs generated per second per core on this se. Udf logs are generated due to the configured client log filters or the rules with logging enabled. Default is 100 logs per second. Set it to zero (0) to deactivate throttling. Field introduced in 17.1.3. Unit is per_second.
* `use_hyperthreaded_cores` - (Optional) Enables the use of hyper-threaded cores on se. Requires se reboot. Field introduced in 20.1.1.
* `use_objsync` - (Optional) Enable interse objsyc distribution framework. Field introduced in 20.1.3. Allowed in basic edition, essentials edition, enterprise edition.
* `use_standard_alb` - (Optional) Use standard sku azure load balancer. By default cloud level flag is set. If not set, it inherits/uses the use_standard_alb flag from the cloud. Field introduced in 18.2.3.
* `vcenter_clusters` - (Optional) Dict settings for serviceenginegroup.
* `vcenter_datastore_mode` - (Optional) Enum options - vcenter_datastore_any, vcenter_datastore_local, vcenter_datastore_shared.
* `vcenter_datastores` - (Optional) List of list.
* `vcenter_datastores_include` - (Optional) Boolean flag to set vcenter_datastores_include.
* `vcenter_folder` - (Optional) Folder to place all the service engine virtual machines in vcenter.
* `vcenter_hosts` - (Optional) Dict settings for serviceenginegroup.
* `vcenters` - (Optional) Vcenter information for scoping at host/cluster level. Field introduced in 20.1.1.
* `vcpus_per_se` - (Optional) Number of vcpus for each of the service engine virtual machines. Changes to this setting do not affect existing ses.
* `vip_asg` - (Optional) When vip_asg is set, vip configuration will be managed by avi.user will be able to configure vip_asg or vips individually at the time of create. Field introduced in 17.2.12, 18.1.2.
* `vnic_dhcp_ip_check_interval` - (Optional) Dhcp ip check interval. Allowed values are 1-1000. Field introduced in 21.1.1. Unit is sec.
* `vnic_dhcp_ip_max_retries` - (Optional) Dhcp ip max retries. Field introduced in 21.1.1.
* `vnic_ip_delete_interval` - (Optional) Wait interval before deleting ip. Field introduced in 21.1.1. Unit is sec.
* `vnic_probe_interval` - (Optional) Probe vnic interval. Field introduced in 21.1.1. Unit is sec.
* `vnic_rpc_retry_interval` - (Optional) Time interval for retrying the failed vnic rpc requests. Field introduced in 21.1.1. Unit is sec.
* `vnicdb_cmd_history_size` - (Optional) Size of vnicdb command history. Allowed values are 0-65535. Field introduced in 21.1.1.
* `vs_host_redundancy` - (Optional) Ensure primary and secondary service engines are deployed on different physical hosts. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition. Special default for basic edition is true, essentials edition is true, enterprise is true.
* `vs_scalein_timeout` - (Optional) Time to wait for the scaled in se to drain existing flows before marking the scalein done. Unit is sec.
* `vs_scalein_timeout_for_upgrade` - (Optional) During se upgrade, time to wait for the scaled-in se to drain existing flows before marking the scalein done. Unit is sec.
* `vs_scaleout_timeout` - (Optional) Time to wait for the scaled out se to become ready before marking the scaleout done. Unit is sec.
* `vs_se_scaleout_additional_wait_time` - (Optional) Wait time for sending scaleout ready notification after virtual service is marked up. In certain deployments, there may be an additional delay to accept traffic. For example, for bgp, some time is needed for route advertisement. Allowed values are 0-20. Field introduced in 18.1.5,18.2.1. Unit is sec.
* `vs_se_scaleout_ready_timeout` - (Optional) Timeout in seconds for service engine to sendscaleout ready notification of a virtual service. Allowed values are 0-90. Field introduced in 18.1.5,18.2.1. Unit is sec.
* `vs_switchover_timeout` - (Optional) During se upgrade in a legacy active/standby segroup, time to wait for the new primary se to accept flows before marking the switchover done. Field introduced in 17.2.13,18.1.4,18.2.1. Unit is sec.
* `vss_placement` - (Optional) Parameters to place virtual services on only a subset of the cores of an se. Field introduced in 17.2.5.
* `vss_placement_enabled` - (Optional) If set, virtual services will be placed on only a subset of the cores of an se. Field introduced in 18.1.1.
* `waf_mempool` - (Optional) Enable memory pool for waf.requires se reboot. Field introduced in 17.2.3.
* `waf_mempool_size` - (Optional) Memory pool size used for waf.requires se reboot. Field introduced in 17.2.3. Unit is kb.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

