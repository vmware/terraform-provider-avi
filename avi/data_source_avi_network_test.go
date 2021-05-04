package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceNetworkBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSNetworkConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "name", "test-network-test-abc"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "exclude_discovered_subnets", "false"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "synced_from_se", "true"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "dhcp_enabled", "true"),
					resource.TestCheckResourceAttr(
						"avi_network.testNetwork", "vcenter_dvs", "true"),
				),
			},
		},
	})

}

const testAccAVIDSNetworkConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_network" "testNetwork" {
	name = "test-network-test-abc"
	exclude_discovered_subnets = false
	tenant_ref = data.avi_tenant.default_tenant.id
	synced_from_se = true
	dhcp_enabled = true
	vcenter_dvs = true
}

data "avi_network" "testNetwork" {
    name= "${avi_network.testNetwork.name}"
}
`
