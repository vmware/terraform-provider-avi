package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceErrorPageProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSErrorPageProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_errorpageprofile.testErrorPageProfile", "name", "test-epp-abc"),
				),
			},
		},
	})

}

const testAccAVIDSErrorPageProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_errorpageprofile" "testErrorPageProfile" {
	name = "test-epp-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_errorpageprofile" "testErrorPageProfile" {
    name= "${avi_errorpageprofile.testErrorPageProfile.name}"
}
`
