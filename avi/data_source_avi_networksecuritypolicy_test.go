package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceNetworkSecurityPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSNetworkSecurityPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_networksecuritypolicy.testNetworkSecurityPolicy", "name", "ns-abc-abc"),
				),
			},
		},
	})

}

const testAccAVIDSNetworkSecurityPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_networksecuritypolicy" "testNetworkSecurityPolicy" {
	name = "ns-abc-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	description = "test network policy"
}

data "avi_networksecuritypolicy" "testNetworkSecurityPolicy" {
    name= "${avi_networksecuritypolicy.testNetworkSecurityPolicy.name}"
}
`
