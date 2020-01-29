package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceIpAddrGroupBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSIpAddrGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_ipaddrgroup.testIpAddrGroup", "name", "test-Internal-abc"),
				),
			},
		},
	})

}

const testAccAVIDSIpAddrGroupConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipaddrgroup" "testIpAddrGroup" {
	prefixes {
	mask = "8"
	ip_addr {
		type = "V4"
		addr = "10.0.0.0"
	}
}
prefixes {
	mask = "16"
	ip_addr {
		type = "V4"
		addr = "192.168.0.0"
	}
}
prefixes {
	mask = "12"
	ip_addr {
		type = "V4"
		addr = "172.16.0.0"
	}
}
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-Internal-abc"
}

data "avi_ipaddrgroup" "testIpAddrGroup" {
    name= "${avi_ipaddrgroup.testIpAddrGroup.name}"
}
`
