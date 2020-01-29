package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAVIDataSourceVrfContextBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSVrfContextConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_vrfcontext.testVrfContext", "name", "test-global-abc"),
					resource.TestCheckResourceAttr(
						"avi_vrfcontext.testVrfContext", "system_default", "false"),
				),
			},
		},
	})

}

const testAccAVIDSVrfContextConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_cloud" "default_cloud" {
    name= "Default-Cloud"
}
resource "avi_vrfcontext" "testVrfContext" {
	cloud_ref = data.avi_cloud.default_cloud.id
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-global-abc"
}

data "avi_vrfcontext" "testVrfContext" {
    name= "${avi_vrfcontext.testVrfContext.name}"
}
`
