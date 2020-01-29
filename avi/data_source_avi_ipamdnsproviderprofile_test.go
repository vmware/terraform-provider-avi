package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAVIDataSourceIpamDnsProviderProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSIpamDnsProviderProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "name", "test-ipam-abc-abc"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testIpamDnsProviderProfile", "allocate_ip_in_vrf", "false"),
				),
			},
		},
	})

}

const testAccAVIDSIpamDnsProviderProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_ipamdnsproviderprofile" "testIpamDnsProviderProfile" {
	tenant_ref = data.avi_tenant.default_tenant.id
	allocate_ip_in_vrf = false
	type = "IPAMDNS_TYPE_INTERNAL"
	name = "test-ipam-abc-abc"
	internal_profile {
		ttl = "31"
	}
}

data "avi_ipamdnsproviderprofile" "testIpamDnsProviderProfile" {
    name= "${avi_ipamdnsproviderprofile.testIpamDnsProviderProfile.name}"
}
`
