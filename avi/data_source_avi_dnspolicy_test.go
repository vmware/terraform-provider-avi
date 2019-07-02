package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceDnsPolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSDnsPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_dnspolicy.testDnsPolicy", "name", "test-dp-abc"),
				),
			},
		},
	})

}

const testAccAVIDSDnsPolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_dnspolicy" "testDnsPolicy" {
	"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
	"name" = "test-dp-abc"
	"description" = "test dns policy"
}

data "avi_dnspolicy" "testDnsPolicy" {
    name= "${avi_dnspolicy.testDnsPolicy.name}"
}
`
