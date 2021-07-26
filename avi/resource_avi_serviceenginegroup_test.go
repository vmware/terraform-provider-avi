package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIServiceEngineGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIServiceEngineGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIServiceEngineGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIServiceEngineGroupExists("avi_serviceenginegroup.testServiceEngineGroup"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "name", "test-Default-Group-abc"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "udf_log_throttle", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vcpus_per_se", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "active_standby", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "hm_on_standby", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disk_per_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "cpu_socket_affinity", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scalein_timeout_for_upgrade", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_redistribute_active_standby_load", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_vs_hb_max_pkts_in_batch", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_rebalance", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "enable_hsm_priming", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "log_disksz", "10000"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "memory_per_se", "2048"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_tunnel_mode", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_sb_dedicated_core", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "distribute_load_active_standby", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "extra_shared_config_memory", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_tunnel_udp_port", "1550"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_rebalance_interval", "300"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scaleout_timeout", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "buffer_se", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disable_tso", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_sb_threads", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "ignore_rtt_threshold", "5000"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "min_cpu_usage", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "distribute_queues", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "cpu_reserve", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disable_csum_offloads", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "async_ssl", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_probe_port", "7"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_udp_encap_ipc", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "archive_shm_limit", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_vs_per_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "async_ssl_threads", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "connection_memory_percentage", "50"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_scaleout_per_vs", "4"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "os_reserved_memory", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scalein_timeout", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "dedicated_dispatcher_core", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_deprovision_delay", "120"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_vs_hb_max_vs_in_pkt", "256"),
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
						"avi_serviceenginegroup.testServiceEngineGroup", "host_gateway_monitor", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "waf_mempool", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "mem_reserve", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "flow_table_new_syn_max_entries", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "num_flow_cores_sum_changes_to_ignore", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "least_load_core_selection", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "aggressive_failure_detection", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "significant_log_throttle", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "waf_mempool_size", "64"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_thread_multiplier", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_host_redundancy", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vnic_dhcp_ip_check_interval", "6"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vnic_dhcp_ip_max_retries", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vnic_ip_delete_interval", "5"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vnic_probe_interval", "5"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "sdb_flush_interval", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "sdb_pipeline_size", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "sdb_scan_count", "1000"),
				),
			},
			{
				Config: testAccAVIServiceEngineGroupupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIServiceEngineGroupExists("avi_serviceenginegroup.testServiceEngineGroup"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "name", "test-Default-Group-updated"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "udf_log_throttle", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vcpus_per_se", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "active_standby", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "hm_on_standby", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disk_per_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "cpu_socket_affinity", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scalein_timeout_for_upgrade", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_redistribute_active_standby_load", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_vs_hb_max_pkts_in_batch", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_rebalance", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "enable_hsm_priming", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "log_disksz", "10000"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "memory_per_se", "2048"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_tunnel_mode", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_sb_dedicated_core", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "distribute_load_active_standby", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "extra_shared_config_memory", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_tunnel_udp_port", "1550"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "auto_rebalance_interval", "300"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scaleout_timeout", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "buffer_se", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disable_tso", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_sb_threads", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "ignore_rtt_threshold", "5000"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "min_cpu_usage", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "distribute_queues", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "cpu_reserve", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "disable_csum_offloads", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "async_ssl", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_probe_port", "7"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_udp_encap_ipc", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "archive_shm_limit", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_vs_per_se", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "async_ssl_threads", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "connection_memory_percentage", "50"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "max_scaleout_per_vs", "4"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "os_reserved_memory", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_scalein_timeout", "30"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "dedicated_dispatcher_core", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_deprovision_delay", "120"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_vs_hb_max_vs_in_pkt", "256"),
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
						"avi_serviceenginegroup.testServiceEngineGroup", "host_gateway_monitor", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "waf_mempool", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "mem_reserve", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "flow_table_new_syn_max_entries", "0"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "num_flow_cores_sum_changes_to_ignore", "8"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "least_load_core_selection", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "aggressive_failure_detection", "false"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "significant_log_throttle", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "waf_mempool_size", "64"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "se_thread_multiplier", "1"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vs_host_redundancy", "true"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vnic_dhcp_ip_check_interval", "6"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vnic_dhcp_ip_max_retries", "10"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vnic_ip_delete_interval", "5"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "vnic_probe_interval", "5"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "sdb_flush_interval", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "sdb_pipeline_size", "100"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "sdb_scan_count", "1000"),
				),
			},
			{
				ResourceName:      "avi_serviceenginegroup.testServiceEngineGroup",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIServiceEngineGroupConfig,
			},
		},
	})

}

func testAccCheckAVIServiceEngineGroupExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ServiceEngineGroup ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIServiceEngineGroupDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_serviceenginegroup" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI ServiceEngineGroup still exists")
		}
	}
	return nil
}

const testAccAVIServiceEngineGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
resource "avi_serviceenginegroup" "testServiceEngineGroup" {
	name = "test-Default-Group-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
	extra_config_multiplier = "0.0"
	udf_log_throttle = "100"
	ingress_access_data = "SG_INGRESS_ACCESS_ALL"
	vcpus_per_se = "1"
	active_standby = false
	hm_on_standby = true
	disk_per_se = "10"
	cpu_socket_affinity = false
	vs_scalein_timeout_for_upgrade = "30"
	se_name_prefix = "Avi"
	auto_redistribute_active_standby_load = false
	se_vs_hb_max_pkts_in_batch = "8"
	auto_rebalance = false
	enable_hsm_priming = false
	log_disksz = "10000"
	memory_per_se = "2048"
	se_tunnel_mode = "0"
	se_sb_dedicated_core = false
	distribute_load_active_standby = false
	extra_shared_config_memory = "0"
	se_tunnel_udp_port = "1550"
	auto_rebalance_interval = "300"
	vs_scaleout_timeout = "30"
	buffer_se = "1"
	disable_tso = false
	ha_mode = "HA_MODE_SHARED"
	se_sb_threads = "1"
	ignore_rtt_threshold = "5000"
	min_cpu_usage = "30"
	distribute_queues = false
	cpu_reserve = false
	disable_csum_offloads = false
	max_se = "10"
	async_ssl = false
	se_probe_port = "7"
	se_udp_encap_ipc = "0"
	archive_shm_limit = "8"
	max_vs_per_se = "10"
	async_ssl_threads = "1"
	connection_memory_percentage = "50"
	placement_mode = "PLACEMENT_MODE_AUTO"
	max_scaleout_per_vs = "4"
	ingress_access_mgmt = "SG_INGRESS_ACCESS_ALL"
	vcenter_folder = "AviSeFolder"
	os_reserved_memory = "0"
	vs_scalein_timeout = "30"
	dedicated_dispatcher_core = false
	se_deprovision_delay = "120"
	se_vs_hb_max_vs_in_pkt = "256"
	disable_gro = true
	vcenter_datastores_include = false
	per_app = false
	non_significant_log_throttle = "100"
	max_cpu_usage = "80"
	min_scaleout_per_vs = "1"
	license_tier = "ENTERPRISE"
	host_gateway_monitor = false
	waf_mempool = true
	mem_reserve = true
	flow_table_new_syn_max_entries = "0"
	vcenter_datastore_mode = "VCENTER_DATASTORE_ANY"
	num_flow_cores_sum_changes_to_ignore = "8"
	least_load_core_selection = true
	aggressive_failure_detection = false
	significant_log_throttle = "100"
	waf_mempool_size = "64"
	se_bandwidth_type = "SE_BANDWIDTH_UNLIMITED"
	algo = "PLACEMENT_ALGO_DISTRIBUTED"
	se_thread_multiplier = "1"
	vs_host_redundancy = true
	license_type = "LIC_CORES"
	vnic_dhcp_ip_check_interval = "6"
	vnic_dhcp_ip_max_retries = "10"
	vnic_ip_delete_interval = "5"
	vnic_probe_interval = "5"
	sdb_flush_interval = "100"
	sdb_pipeline_size = "100"
	sdb_scan_count = "1000"
}
`

const testAccAVIServiceEngineGroupupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
resource "avi_serviceenginegroup" "testServiceEngineGroup" {
	name = "test-Default-Group-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	cloud_ref = data.avi_cloud.default_cloud.id
	extra_config_multiplier = "0.0"
	udf_log_throttle = "100"
	ingress_access_data = "SG_INGRESS_ACCESS_ALL"
	vcpus_per_se = "1"
	active_standby = false
	hm_on_standby = true
	disk_per_se = "10"
	cpu_socket_affinity = false
	vs_scalein_timeout_for_upgrade = "30"
	se_name_prefix = "Avi"
	auto_redistribute_active_standby_load = false
	se_vs_hb_max_pkts_in_batch = "8"
	auto_rebalance = false
	enable_hsm_priming = false
	log_disksz = "10000"
	memory_per_se = "2048"
	se_tunnel_mode = "0"
	se_sb_dedicated_core = false
	distribute_load_active_standby = false
	extra_shared_config_memory = "0"
	se_tunnel_udp_port = "1550"
	auto_rebalance_interval = "300"
	vs_scaleout_timeout = "30"
	buffer_se = "1"
	disable_tso = false
	ha_mode = "HA_MODE_SHARED"
	se_sb_threads = "1"
	ignore_rtt_threshold = "5000"
	min_cpu_usage = "30"
	distribute_queues = false
	cpu_reserve = false
	disable_csum_offloads = false
	max_se = "10"
	async_ssl = false
	se_probe_port = "7"
	se_udp_encap_ipc = "0"
	archive_shm_limit = "8"
	max_vs_per_se = "10"
	async_ssl_threads = "1"
	connection_memory_percentage = "50"
	placement_mode = "PLACEMENT_MODE_AUTO"
	max_scaleout_per_vs = "4"
	ingress_access_mgmt = "SG_INGRESS_ACCESS_ALL"
	vcenter_folder = "AviSeFolder"
	os_reserved_memory = "0"
	vs_scalein_timeout = "30"
	dedicated_dispatcher_core = false
	se_deprovision_delay = "120"
	se_vs_hb_max_vs_in_pkt = "256"
	disable_gro = true
	vcenter_datastores_include = false
	per_app = false
	non_significant_log_throttle = "100"
	max_cpu_usage = "80"
	min_scaleout_per_vs = "1"
	license_tier = "ENTERPRISE"
	host_gateway_monitor = false
	waf_mempool = true
	mem_reserve = true
	flow_table_new_syn_max_entries = "0"
	vcenter_datastore_mode = "VCENTER_DATASTORE_ANY"
	num_flow_cores_sum_changes_to_ignore = "8"
	least_load_core_selection = true
	aggressive_failure_detection = false
	significant_log_throttle = "100"
	waf_mempool_size = "64"
	se_bandwidth_type = "SE_BANDWIDTH_UNLIMITED"
	algo = "PLACEMENT_ALGO_DISTRIBUTED"
	se_thread_multiplier = "1"
	vs_host_redundancy = true
	license_type = "LIC_CORES"
	vnic_dhcp_ip_check_interval = "6"
	vnic_dhcp_ip_max_retries = "10"
	vnic_ip_delete_interval = "5"
	vnic_probe_interval = "5"
	sdb_flush_interval = "100"
	sdb_pipeline_size = "100"
	sdb_scan_count = "1000"
}
`
