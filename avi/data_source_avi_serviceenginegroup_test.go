package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAVIDataSourceServiceEngineGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSServiceEngineGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "name", "test-Default-Group-abc"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "archive_shm_limit", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "udf_log_throttle", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_ipc_udp_port", "1500"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_vs_hb_max_vs_in_pkt", "256"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vcpus_per_se", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "hm_on_standby", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disk_per_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "cpu_socket_affinity", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_redistribute_active_standby_load", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_rebalance", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "aggressive_failure_detection", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "num_flow_cores_sum_changes_to_ignore", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "connection_memory_percentage", "50"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "log_disksz", "10000"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_tunnel_mode", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_sb_dedicated_core", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "memory_per_se", "2048"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "extra_shared_config_memory", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_tunnel_udp_port", "1550"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_rebalance_interval", "300"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scaleout_timeout", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disable_tso", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_sb_threads", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_remote_punt_udp_port", "1501"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "ignore_rtt_threshold", "5000"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "active_standby", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "distribute_queues", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disable_csum_offloads", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "async_ssl", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_probe_port", "7"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_udp_encap_ipc", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "extra_config_multiplier", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_vs_per_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "async_ssl_threads", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "min_cpu_usage", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_scaleout_per_vs", "4"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "os_reserved_memory", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scalein_timeout_for_upgrade", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scalein_timeout", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "dedicated_dispatcher_core", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_thread_multiplier", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_deprovision_delay", "120"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disable_gro", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vcenter_datastores_include", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "per_app", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "non_significant_log_throttle", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_cpu_usage", "80"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "min_scaleout_per_vs", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_vs_hb_max_pkts_in_batch", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "host_gateway_monitor", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "waf_mempool", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "buffer_se", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "mem_reserve", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "flow_table_new_syn_max_entries", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "enable_hsm_priming", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "distribute_load_active_standby", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "least_load_core_selection", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "significant_log_throttle", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "waf_mempool_size", "64"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "cpu_reserve", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_host_redundancy", "true"),
				),
			},
		},
	})

}

const testAccAVIDSServiceEngineGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
resource "avi_serviceenginegroup" "testServiceEngineGroup" {
	archive_shm_limit = "8"
	udf_log_throttle = "100"
	se_ipc_udp_port = "1500"
	se_vs_hb_max_vs_in_pkt = "256"
	vcpus_per_se = "1"
	hm_on_standby = true
	disk_per_se = "10"
	cpu_socket_affinity = false
	license_tier = "ENTERPRISE_18"
	se_name_prefix = "Avi"
	auto_redistribute_active_standby_load = false
	auto_rebalance = false
	aggressive_failure_detection = false
	num_flow_cores_sum_changes_to_ignore = "8"
	connection_memory_percentage = "50"
	log_disksz = "10000"
	se_tunnel_mode = "0"
	se_sb_dedicated_core = false
	memory_per_se = "2048"
	cloud_ref = data.avi_cloud.default_cloud.id
	extra_shared_config_memory = "0"
	se_tunnel_udp_port = "1550"
	auto_rebalance_interval = "300"
	vs_scaleout_timeout = "30"
	disable_tso = false
	ha_mode = "HA_MODE_SHARED"
	se_sb_threads = "1"
	se_remote_punt_udp_port = "1501"
	ignore_rtt_threshold = "5000"
	active_standby = false
	distribute_queues = false
	disable_csum_offloads = false
	async_ssl = false
	se_probe_port = "7"
	se_udp_encap_ipc = "0"
	extra_config_multiplier = "0"
	max_vs_per_se = "10"
	async_ssl_threads = "1"
	min_cpu_usage = "30"
	placement_mode = "PLACEMENT_MODE_AUTO"
	max_scaleout_per_vs = "4"
	ingress_access_mgmt = "SG_INGRESS_ACCESS_ALL"
	vcenter_folder = "AviSeFolder"
	os_reserved_memory = "0"
	vs_scalein_timeout_for_upgrade = "30"
	vs_scalein_timeout = "30"
	dedicated_dispatcher_core = false
	se_thread_multiplier = "1"
	se_deprovision_delay = "120"
	ingress_access_data = "SG_INGRESS_ACCESS_ALL"
	disable_gro = true
	vcenter_datastores_include = false
	per_app = false
	non_significant_log_throttle = "100"
	max_cpu_usage = "80"
	min_scaleout_per_vs = "1"
	se_vs_hb_max_pkts_in_batch = "8"
	name = "test-Default-Group-abc"
	host_gateway_monitor = false
	waf_mempool = true
	buffer_se = "1"
	mem_reserve = true
	flow_table_new_syn_max_entries = "0"
	tenant_ref = data.avi_tenant.default_tenant.id
	vcenter_datastore_mode = "VCENTER_DATASTORE_ANY"
	enable_hsm_priming = false
	distribute_load_active_standby = false
	least_load_core_selection = true
	max_se = "10"
	significant_log_throttle = "100"
	waf_mempool_size = "64"
	cpu_reserve = false
	algo = "PLACEMENT_ALGO_DISTRIBUTED"
	se_bandwidth_type = "SE_BANDWIDTH_UNLIMITED"
	vs_host_redundancy = true
	license_type = "LIC_CORES"
}

data "avi_serviceenginegroup" "testServiceEngineGroup" {
    name= "${avi_serviceenginegroup.testServiceEngineGroup.name}"
}
`
