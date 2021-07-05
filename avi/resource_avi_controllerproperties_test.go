package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIControllerPropertiesBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIControllerPropertiesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIControllerPropertiesConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIControllerPropertiesExists("avi_controllerproperties.testControllerProperties"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_ping_fail", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_create_fail", "1500"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "cluster_ip_gratuitous_arp_period", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "persistence_key_rotate_period", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "unresponsive_se_reboot", "300"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "attach_ip_retry_interval", "360"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_vnic_fail", "300"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "attach_ip_retry_limit", "4"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "se_vnic_cooldown", "120"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vnic_op_fail_time", "180"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "max_pcap_per_tenant", "4"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "enable_memory_balancer", "true"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "max_dead_se_in_grp", "1"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "seupgrade_fabric_pool_size", "20"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_key_rotate_period", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "seupgrade_segroup_min_dead_timeout", "360"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "upgrade_lease_time", "360"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "se_create_timeout", "900"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "query_host_fail", "180"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "bm_use_ansible", "true"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "se_offline_del", "172000"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_bootup_fail", "480"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "upgrade_dns_ttl", "5"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "fatal_error_lease_time", "120"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "allow_ip_forwarding", "false"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_attach_ip_fail", "3600"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "max_seq_vnic_failures", "3"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "allow_unauthenticated_nodes", "false"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "allow_unauthenticated_apis", "false"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_awaiting_se_timeout", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "warmstart_se_reconnect_wait_time", "300"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "dns_refresh_period", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "max_seq_attach_ip_failures", "3"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "secure_channel_cleanup_timeout", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_vnic_ip_fail", "120"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "secure_channel_se_token_timeout", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "secure_channel_controller_token_timeout", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "api_idle_timeout", "15"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "crashed_se_reboot", "900"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "appviewx_compat_mode", "false"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "se_failover_attempt_interval", "300"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "dead_se_detection_timer", "360"),
				),
			},
			{
				Config: testAccAVIControllerPropertiesupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIControllerPropertiesExists("avi_controllerproperties.testControllerProperties"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_ping_fail", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_create_fail", "1500"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "cluster_ip_gratuitous_arp_period", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "persistence_key_rotate_period", "70"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "unresponsive_se_reboot", "300"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "attach_ip_retry_interval", "360"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_vnic_fail", "300"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "attach_ip_retry_limit", "4"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "se_vnic_cooldown", "120"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vnic_op_fail_time", "180"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "max_pcap_per_tenant", "4"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "enable_memory_balancer", "true"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "max_dead_se_in_grp", "1"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "seupgrade_fabric_pool_size", "20"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_key_rotate_period", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "seupgrade_segroup_min_dead_timeout", "360"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "upgrade_lease_time", "360"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "se_create_timeout", "900"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "query_host_fail", "180"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "bm_use_ansible", "true"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "se_offline_del", "172000"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_bootup_fail", "480"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "upgrade_dns_ttl", "5"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "fatal_error_lease_time", "120"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "allow_ip_forwarding", "false"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_attach_ip_fail", "3600"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "max_seq_vnic_failures", "3"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "allow_unauthenticated_nodes", "false"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "allow_unauthenticated_apis", "false"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_awaiting_se_timeout", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "warmstart_se_reconnect_wait_time", "300"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "dns_refresh_period", "70"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "max_seq_attach_ip_failures", "3"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "secure_channel_cleanup_timeout", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "vs_se_vnic_ip_fail", "120"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "secure_channel_se_token_timeout", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "secure_channel_controller_token_timeout", "60"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "api_idle_timeout", "15"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "crashed_se_reboot", "900"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "appviewx_compat_mode", "false"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "se_failover_attempt_interval", "300"),
					resource.TestCheckResourceAttr(
						"avi_controllerproperties.testControllerProperties", "dead_se_detection_timer", "360"),
				),
			},
			{
				ResourceName:      "avi_controllerproperties.testControllerProperties",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIControllerPropertiesConfig,
			},
		},
	})

}

func testAccCheckAVIControllerPropertiesExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ControllerProperties ID is set")
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

func testAccCheckAVIControllerPropertiesDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_controllerproperties" {
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
			return fmt.Errorf("AVI ControllerProperties still exists")
		}
	}
	return nil
}

const testAccAVIControllerPropertiesConfig = `
resource "avi_controllerproperties" "testControllerProperties" {
	vs_se_ping_fail = "60"
	vs_se_create_fail = "1500"
	cluster_ip_gratuitous_arp_period = "60"
	persistence_key_rotate_period = "60"
	unresponsive_se_reboot = "300"
	attach_ip_retry_interval = "360"
	vs_se_vnic_fail = "300"
	attach_ip_retry_limit = "4"
	se_vnic_cooldown = "120"
	vnic_op_fail_time = "180"
	max_pcap_per_tenant = "4"
	enable_memory_balancer = true
	max_dead_se_in_grp = "1"
	seupgrade_fabric_pool_size = "20"
	vs_key_rotate_period = "60"
	seupgrade_segroup_min_dead_timeout = "360"
	upgrade_lease_time = "360"
	se_create_timeout = "900"
	query_host_fail = "180"
	bm_use_ansible = true
	se_offline_del = "172000"
	vs_se_bootup_fail = "480"
	upgrade_dns_ttl = "5"
	fatal_error_lease_time = "120"
	allow_ip_forwarding = false
	vs_se_attach_ip_fail = "3600"
	max_seq_vnic_failures = "3"
	allow_unauthenticated_nodes = false
	allow_unauthenticated_apis = false
	vs_awaiting_se_timeout = "60"
	warmstart_se_reconnect_wait_time = "300"
	dns_refresh_period = "60"
	max_seq_attach_ip_failures = "3"
	secure_channel_cleanup_timeout = "60"
	vs_se_vnic_ip_fail = "120"
	ssl_certificate_expiry_warning_days = ["30","7","1"]
	secure_channel_se_token_timeout = "60"
	secure_channel_controller_token_timeout = "60"
	api_idle_timeout = "15"
	crashed_se_reboot = "900"
	appviewx_compat_mode = false
	se_failover_attempt_interval = "300"
	dead_se_detection_timer = "360"
}
`

const testAccAVIControllerPropertiesupdatedConfig = `
resource "avi_controllerproperties" "testControllerProperties" {
	vs_se_ping_fail = "60"
	vs_se_create_fail = "1500"
	cluster_ip_gratuitous_arp_period = "60"
	persistence_key_rotate_period = "70"
	unresponsive_se_reboot = "300"
	attach_ip_retry_interval = "360"
	vs_se_vnic_fail = "300"
	attach_ip_retry_limit = "4"
	se_vnic_cooldown = "120"
	vnic_op_fail_time = "180"
	max_pcap_per_tenant = "4"
	enable_memory_balancer = true
	max_dead_se_in_grp = "1"
	seupgrade_fabric_pool_size = "20"
	vs_key_rotate_period = "60"
	seupgrade_segroup_min_dead_timeout = "360"
	upgrade_lease_time = "360"
	se_create_timeout = "900"
	query_host_fail = "180"
	bm_use_ansible = true
	se_offline_del = "172000"
	vs_se_bootup_fail = "480"
	upgrade_dns_ttl = "5"
	fatal_error_lease_time = "120"
	allow_ip_forwarding = false
	vs_se_attach_ip_fail = "3600"
	max_seq_vnic_failures = "3"
	allow_unauthenticated_nodes = false
	allow_unauthenticated_apis = false
	vs_awaiting_se_timeout = "60"
	warmstart_se_reconnect_wait_time = "300"
	dns_refresh_period = "70"
	max_seq_attach_ip_failures = "3"
	secure_channel_cleanup_timeout = "60"
	vs_se_vnic_ip_fail = "120"
	ssl_certificate_expiry_warning_days = ["30","7","1"]
	secure_channel_se_token_timeout = "60"
	secure_channel_controller_token_timeout = "60"
	api_idle_timeout = "15"
	crashed_se_reboot = "900"
	appviewx_compat_mode = false
	se_failover_attempt_interval = "300"
	dead_se_detection_timer = "360"
}
`
