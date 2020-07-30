package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceNetworkProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSNetworkProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_networkprofile.testNetworkProfile", "name", "test-System-TCP-Proxy-abc"),
				),
			},
		},
	})

}

const testAccAVIDSNetworkProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_networkprofile" "testNetworkProfile" {
	name = "test-System-TCP-Proxy-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	profile {
		tcp_proxy_profile {
			receive_window = "64"
			time_wait_delay = "2000"
			cc_algo = "CC_ALGO_NEW_RENO"
			nagles_algorithm = false
			max_syn_retransmissions = "8"
			ignore_time_wait = false
			idle_connection_type = "KEEP_ALIVE"
			aggressive_congestion_avoidance = false
			idle_connection_timeout = "600"
			max_retransmissions = "8"
			ip_dscp = "0"
			automatic = true
			use_interface_mtu = true
			reorder_threshold = "10"
			min_rexmt_timeout = "50"
		}
		type = "PROTOCOL_TYPE_TCP_PROXY"
	}
}

data "avi_networkprofile" "testNetworkProfile" {
    name= "${avi_networkprofile.testNetworkProfile.name}"
}
`
