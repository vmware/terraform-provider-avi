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
resource "ServiceEngineGroup" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `accelerated_networking` - (Optional ) argument_description.
        * `active_standby` - (Optional ) argument_description.
        * `advertise_backend_networks` - (Optional ) argument_description.
        * `aggressive_failure_detection` - (Optional ) argument_description.
        * `algo` - (Optional ) argument_description.
        * `allow_burst` - (Optional ) argument_description.
        * `archive_shm_limit` - (Optional ) argument_description.
        * `async_ssl` - (Optional ) argument_description.
        * `async_ssl_threads` - (Optional ) argument_description.
        * `auto_rebalance` - (Optional ) argument_description.
        * `auto_rebalance_capacity_per_se` - (Optional ) argument_description.
        * `auto_rebalance_criteria` - (Optional ) argument_description.
        * `auto_rebalance_interval` - (Optional ) argument_description.
        * `auto_redistribute_active_standby_load` - (Optional ) argument_description.
        * `bgp_state_update_interval` - (Optional ) argument_description.
        * `buffer_se` - (Optional ) argument_description.
        * `cloud_ref` - (Optional ) argument_description.
        * `config_debugs_on_all_cores` - (Optional ) argument_description.
        * `connection_memory_percentage` - (Optional ) argument_description.
        * `cpu_reserve` - (Optional ) argument_description.
        * `cpu_socket_affinity` - (Optional ) argument_description.
        * `custom_securitygroups_data` - (Optional ) argument_description.
        * `custom_securitygroups_mgmt` - (Optional ) argument_description.
        * `custom_tag` - (Optional ) argument_description.
        * `dedicated_dispatcher_core` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `disable_avi_securitygroups` - (Optional ) argument_description.
        * `disable_csum_offloads` - (Optional ) argument_description.
        * `disable_gro` - (Optional ) argument_description.
        * `disable_se_memory_check` - (Optional ) argument_description.
        * `disable_tso` - (Optional ) argument_description.
        * `disk_per_se` - (Optional ) argument_description.
        * `distribute_load_active_standby` - (Optional ) argument_description.
        * `distribute_queues` - (Optional ) argument_description.
        * `enable_hsm_priming` - (Optional ) argument_description.
        * `enable_multi_lb` - (Optional ) argument_description.
        * `enable_routing` - (Optional ) argument_description.
        * `enable_vip_on_all_interfaces` - (Optional ) argument_description.
        * `enable_vmac` - (Optional ) argument_description.
        * `ephemeral_portrange_end` - (Optional ) argument_description.
        * `ephemeral_portrange_start` - (Optional ) argument_description.
        * `extra_config_multiplier` - (Optional ) argument_description.
        * `extra_shared_config_memory` - (Optional ) argument_description.
        * `floating_intf_ip` - (Optional ) argument_description.
        * `floating_intf_ip_se_2` - (Optional ) argument_description.
        * `flow_table_new_syn_max_entries` - (Optional ) argument_description.
        * `free_list_size` - (Optional ) argument_description.
        * `ha_mode` - (Optional ) argument_description.
        * `hardwaresecuritymodulegroup_ref` - (Optional ) argument_description.
        * `heap_minimum_config_memory` - (Optional ) argument_description.
        * `hm_on_standby` - (Optional ) argument_description.
        * `host_attribute_key` - (Optional ) argument_description.
        * `host_attribute_value` - (Optional ) argument_description.
        * `host_gateway_monitor` - (Optional ) argument_description.
        * `hypervisor` - (Optional ) argument_description.
        * `ignore_rtt_threshold` - (Optional ) argument_description.
        * `ingress_access_data` - (Optional ) argument_description.
        * `ingress_access_mgmt` - (Optional ) argument_description.
        * `instance_flavor` - (Optional ) argument_description.
        * `iptables` - (Optional ) argument_description.
        * `least_load_core_selection` - (Optional ) argument_description.
        * `license_tier` - (Optional ) argument_description.
        * `license_type` - (Optional ) argument_description.
        * `log_disksz` - (Optional ) argument_description.
        * `max_cpu_usage` - (Optional ) argument_description.
        * `max_memory_per_mempool` - (Optional ) argument_description.
        * `max_public_ips_per_lb` - (Optional ) argument_description.
        * `max_rules_per_lb` - (Optional ) argument_description.
        * `max_scaleout_per_vs` - (Optional ) argument_description.
        * `max_se` - (Optional ) argument_description.
        * `max_vs_per_se` - (Optional ) argument_description.
        * `mem_reserve` - (Optional ) argument_description.
        * `memory_for_config_update` - (Optional ) argument_description.
        * `memory_per_se` - (Optional ) argument_description.
        * `mgmt_network_ref` - (Optional ) argument_description.
        * `mgmt_subnet` - (Optional ) argument_description.
        * `min_cpu_usage` - (Optional ) argument_description.
        * `min_scaleout_per_vs` - (Optional ) argument_description.
        * `min_se` - (Optional ) argument_description.
        * `minimum_connection_memory` - (Optional ) argument_description.
        * `n_log_streaming_threads` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `non_significant_log_throttle` - (Optional ) argument_description.
        * `num_dispatcher_cores` - (Optional ) argument_description.
        * `num_flow_cores_sum_changes_to_ignore` - (Optional ) argument_description.
        * `openstack_availability_zones` - (Optional ) argument_description.
        * `openstack_mgmt_network_name` - (Optional ) argument_description.
        * `openstack_mgmt_network_uuid` - (Optional ) argument_description.
        * `os_reserved_memory` - (Optional ) argument_description.
        * `per_app` - (Optional ) argument_description.
        * `placement_mode` - (Optional ) argument_description.
        * `realtime_se_metrics` - (Optional ) argument_description.
        * `se_bandwidth_type` - (Optional ) argument_description.
        * `se_deprovision_delay` - (Optional ) argument_description.
        * `se_dos_profile` - (Optional ) argument_description.
        * `se_dpdk_pmd` - (Optional ) argument_description.
        * `se_flow_probe_retries` - (Optional ) argument_description.
        * `se_flow_probe_timer` - (Optional ) argument_description.
        * `se_ipc_udp_port` - (Optional ) argument_description.
        * `se_name_prefix` - (Optional ) argument_description.
        * `se_pcap_reinit_frequency` - (Optional ) argument_description.
        * `se_pcap_reinit_threshold` - (Optional ) argument_description.
        * `se_probe_port` - (Optional ) argument_description.
        * `se_remote_punt_udp_port` - (Optional ) argument_description.
        * `se_sb_dedicated_core` - (Optional ) argument_description.
        * `se_sb_threads` - (Optional ) argument_description.
        * `se_thread_multiplier` - (Optional ) argument_description.
        * `se_tracert_port_range` - (Optional ) argument_description.
        * `se_tunnel_mode` - (Optional ) argument_description.
        * `se_tunnel_udp_port` - (Optional ) argument_description.
        * `se_udp_encap_ipc` - (Optional ) argument_description.
        * `se_use_dpdk` - (Optional ) argument_description.
        * `se_vs_hb_max_pkts_in_batch` - (Optional ) argument_description.
        * `se_vs_hb_max_vs_in_pkt` - (Optional ) argument_description.
        * `self_se_election` - (Optional ) argument_description.
        * `service_ip6_subnets` - (Optional ) argument_description.
        * `service_ip_subnets` - (Optional ) argument_description.
        * `shm_minimum_config_memory` - (Optional ) argument_description.
        * `significant_log_throttle` - (Optional ) argument_description.
        * `ssl_preprocess_sni_hostname` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `udf_log_throttle` - (Optional ) argument_description.
            * `vcenter_clusters` - (Optional ) argument_description.
        * `vcenter_datastore_mode` - (Optional ) argument_description.
        * `vcenter_datastores` - (Optional ) argument_description.
        * `vcenter_datastores_include` - (Optional ) argument_description.
        * `vcenter_folder` - (Optional ) argument_description.
        * `vcenter_hosts` - (Optional ) argument_description.
        * `vcpus_per_se` - (Optional ) argument_description.
        * `vip_asg` - (Optional ) argument_description.
        * `vs_host_redundancy` - (Optional ) argument_description.
        * `vs_scalein_timeout` - (Optional ) argument_description.
        * `vs_scalein_timeout_for_upgrade` - (Optional ) argument_description.
        * `vs_scaleout_timeout` - (Optional ) argument_description.
        * `vs_se_scaleout_additional_wait_time` - (Optional ) argument_description.
        * `vs_se_scaleout_ready_timeout` - (Optional ) argument_description.
        * `vs_switchover_timeout` - (Optional ) argument_description.
        * `vss_placement` - (Optional ) argument_description.
        * `vss_placement_enabled` - (Optional ) argument_description.
        * `waf_learning_interval` - (Optional ) argument_description.
        * `waf_learning_memory` - (Optional ) argument_description.
        * `waf_mempool` - (Optional ) argument_description.
        * `waf_mempool_size` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            * `uuid` - argument_description.
                                                                                        
