package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceServerAutoScalePolicyBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSServerAutoScalePolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_serverautoscalepolicy.testServerAutoScalePolicy", "name", "test-ssp-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSServerAutoScalePolicyConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_serverautoscalepolicy" "testServerAutoScalePolicy" {
	scalein_alertconfig_refs = []
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-ssp-test-abc"
	scaleout_alertconfig_refs = []

}

data "avi_serverautoscalepolicy" "testServerAutoScalePolicy" {
    name= avi_serverautoscalepolicy.testServerAutoScalePolicy.name
}
`
