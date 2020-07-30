package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceTrafficCloneProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSTrafficCloneProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_trafficcloneprofile.testTrafficCloneProfile", "name", "test-tp-test-abc"),
				),
			},
		},
	})

}

const testAccAVIDSTrafficCloneProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_trafficcloneprofile" "testTrafficCloneProfile" {
	name = "test-tp-test-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
}

data "avi_trafficcloneprofile" "testTrafficCloneProfile" {
    name= "${avi_trafficcloneprofile.testTrafficCloneProfile.name}"
}
`
