package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceVsVipBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSVsVipConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_vsvip.testVsVip", "name", "test-vsvip-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSVsVipConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_vsvip" "testVsVip" {
	name = "test-vsvip-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	vip {
	vip_id = "1"
	avi_allocated_fip = false
	auto_allocate_ip = false
	enabled = false
	auto_allocate_floating_ip = false
	avi_allocated_vip = false
	ip_address {
		type = "V4"
		addr = "1.2.3.1"
	}
}
}

data "avi_vsvip" "testVsVip" {
    name= "${avi_vsvip.testVsVip.name}"
}
`
