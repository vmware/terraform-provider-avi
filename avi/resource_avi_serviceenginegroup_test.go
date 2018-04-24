package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
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
						"avi_serviceenginegroup.testServiceEngineGroup", "name", "testDefault-Group")),
			},
			{
				Config: testAccAVIServiceEngineGroupupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIServiceEngineGroupExists("avi_serviceenginegroup.testServiceEngineGroup"),
					resource.TestCheckResourceAttr(
						"avi_serviceenginegroup.testServiceEngineGroup", "name", "testDefault-Group-abc")),
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
"archive_shm_limit" = "8"
"udf_log_throttle" = "100"
"se_ipc_udp_port" = "1500"
"se_vs_hb_max_vs_in_pkt" = "256"
"vcpus_per_se" = "1"
"hm_on_standby" = true
"disk_per_se" = "10"
"cpu_socket_affinity" = false
"license_tier" = "ENTERPRISE_18"
"se_name_prefix" = "Avi"
"auto_redistribute_active_standby_load" = false
"non_significant_log_throttle" = "100"
"auto_rebalance" = false
"aggressive_failure_detection" = false
"num_flow_cores_sum_changes_to_ignore" = "8"
"connection_memory_percentage" = "50"
"log_disksz" = "10000"
"se_tunnel_mode" = "0"
"se_sb_dedicated_core" = false
"cloud_ref" = "${data.avi_cloud.default_cloud.id}"
"extra_shared_config_memory" = "0"
"se_tunnel_udp_port" = "1550"
"auto_rebalance_interval" = "300"
"vs_scaleout_timeout" = "30"
"buffer_se" = "1"
"enable_routing" = false
"disable_tso" = true
"ha_mode" = "HA_MODE_SHARED"
"se_sb_threads" = "1"
"se_remote_punt_udp_port" = "1501"
"ignore_rtt_threshold" = "5000"
"active_standby" = false
"distribute_queues" = false
"disable_csum_offloads" = false
"enable_vip_on_all_interfaces" = true
"se_probe_port" = "7"
"se_udp_encap_ipc" = "0"
"extra_config_multiplier" = "0.0"
"max_vs_per_se" = "10"
"async_ssl_threads" = "1"
"min_cpu_usage" = "30"
"placement_mode" = "PLACEMENT_MODE_AUTO"
"max_scaleout_per_vs" = "4"
"ingress_access_mgmt" = "SG_INGRESS_ACCESS_ALL"
"vcenter_folder" = "AviSeFolder"
"os_reserved_memory" = "0"
"vs_scalein_timeout_for_upgrade" = "30"
"vs_scalein_timeout" = "30"
"dedicated_dispatcher_core" = false
"se_thread_multiplier" = "1"
"se_deprovision_delay" = "120"
"ingress_access_data" = "SG_INGRESS_ACCESS_ALL"
"disable_gro" = true
"vcenter_datastores_include" = false
"advertise_backend_networks" = false
"per_app" = false
"memory_per_se" = "2048"
"max_cpu_usage" = "80"
"async_ssl" = false
"se_vs_hb_max_pkts_in_batch" = "8"
"name" = "testDefault-Group"
"host_gateway_monitor" = false
"waf_mempool" = true
"min_scaleout_per_vs" = "1"
"flow_table_new_syn_max_entries" = "0"
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"vcenter_datastore_mode" = "VCENTER_DATASTORE_ANY"
"enable_hsm_priming" = false
"distribute_load_active_standby" = false
"least_load_core_selection" = true
"max_se" = "10"
"significant_log_throttle" = "100"
"waf_mempool_size" = "64"
"mem_reserve" = true
"cpu_reserve" = false
"algo" = "PLACEMENT_ALGO_DISTRIBUTED"
"se_bandwidth_type" = "SE_BANDWIDTH_UNLIMITED"
"enable_vmac" = false
"vs_host_redundancy" = true
"license_type" = "LIC_CORES"
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
"archive_shm_limit" = "8"
"udf_log_throttle" = "100"
"se_ipc_udp_port" = "1500"
"se_vs_hb_max_vs_in_pkt" = "256"
"vcpus_per_se" = "1"
"hm_on_standby" = true
"disk_per_se" = "10"
"cpu_socket_affinity" = false
"license_tier" = "ENTERPRISE_18"
"se_name_prefix" = "Avi"
"auto_redistribute_active_standby_load" = false
"non_significant_log_throttle" = "100"
"auto_rebalance" = false
"aggressive_failure_detection" = false
"num_flow_cores_sum_changes_to_ignore" = "8"
"connection_memory_percentage" = "50"
"log_disksz" = "10000"
"se_tunnel_mode" = "0"
"se_sb_dedicated_core" = false
"cloud_ref" = "${data.avi_cloud.default_cloud.id}"
"extra_shared_config_memory" = "0"
"se_tunnel_udp_port" = "1550"
"auto_rebalance_interval" = "300"
"vs_scaleout_timeout" = "30"
"buffer_se" = "1"
"enable_routing" = false
"disable_tso" = true
"ha_mode" = "HA_MODE_SHARED"
"se_sb_threads" = "1"
"se_remote_punt_udp_port" = "1501"
"ignore_rtt_threshold" = "5000"
"active_standby" = false
"distribute_queues" = false
"disable_csum_offloads" = false
"enable_vip_on_all_interfaces" = true
"se_probe_port" = "7"
"se_udp_encap_ipc" = "0"
"extra_config_multiplier" = "0.0"
"max_vs_per_se" = "10"
"async_ssl_threads" = "1"
"min_cpu_usage" = "30"
"placement_mode" = "PLACEMENT_MODE_AUTO"
"max_scaleout_per_vs" = "4"
"ingress_access_mgmt" = "SG_INGRESS_ACCESS_ALL"
"vcenter_folder" = "AviSeFolder"
"os_reserved_memory" = "0"
"vs_scalein_timeout_for_upgrade" = "30"
"vs_scalein_timeout" = "30"
"dedicated_dispatcher_core" = false
"se_thread_multiplier" = "1"
"se_deprovision_delay" = "120"
"ingress_access_data" = "SG_INGRESS_ACCESS_ALL"
"disable_gro" = true
"vcenter_datastores_include" = false
"advertise_backend_networks" = false
"per_app" = false
"memory_per_se" = "2048"
"max_cpu_usage" = "80"
"async_ssl" = false
"se_vs_hb_max_pkts_in_batch" = "8"
"name" = "testDefault-Group-abc"
"host_gateway_monitor" = false
"waf_mempool" = true
"min_scaleout_per_vs" = "1"
"flow_table_new_syn_max_entries" = "0"
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"vcenter_datastore_mode" = "VCENTER_DATASTORE_ANY"
"enable_hsm_priming" = false
"distribute_load_active_standby" = false
"least_load_core_selection" = true
"max_se" = "10"
"significant_log_throttle" = "100"
"waf_mempool_size" = "64"
"mem_reserve" = true
"cpu_reserve" = false
"algo" = "PLACEMENT_ALGO_DISTRIBUTED"
"se_bandwidth_type" = "SE_BANDWIDTH_UNLIMITED"
"enable_vmac" = false
"vs_host_redundancy" = true
"license_type" = "LIC_CORES"
}
`
