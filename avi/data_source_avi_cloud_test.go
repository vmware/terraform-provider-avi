package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceCloudBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSCloudConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "name", "test-Default-Cloud-abc"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "enable_vip_static_routes", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "prefer_static_routes", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "mtu", "1500"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "apic_mode", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "dhcp_enabled", "false"),
					resource.TestCheckResourceAttr(
						"avi_cloud.testCloud", "state_based_dns_registration", "true"),
				),
			},
		},
	})

}

const testAccAVIDSCloudConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_cloud" "testCloud" {
	"vtype" = "CLOUD_NONE"
	"license_tier" = "ENTERPRISE_18"
	"name" = "test-Default-Cloud-abc"
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"enable_vip_static_routes" = false
	"prefer_static_routes" = false
	"license_type" = "LIC_CORES"
	"mtu" = "1500"
	"apic_mode" = false
	"dhcp_enabled" = false
	"state_based_dns_registration" = true
}

data "avi_cloud" "testCloud" {
    name= "${avi_cloud.testCloud.name}"
}
`
