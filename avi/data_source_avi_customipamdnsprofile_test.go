package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceCustomIpamDnsProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSCustomIpamDnsProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_customipamdnsprofile.testCustomIpamDnsProfile", "name", "test-ipam-abc"),
				),
			},
		},
	})

}

const testAccAVIDSCustomIpamDnsProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_customipamdnsprofile" "testCustomIpamDnsProfile" {
	name = "test-ipam-abc"
	script_uri = "/"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_customipamdnsprofile" "testCustomIpamDnsProfile" {
    name= "${avi_customipamdnsprofile.testCustomIpamDnsProfile.name}"
}
`
