package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceAlertEmailConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSAlertEmailConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_alertemailconfig.testAlertEmailConfig", "name", "test-aec-abc"),
				),
			},
		},
	})

}

const testAccAVIDSAlertEmailConfigConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_alertemailconfig" "testAlertEmailConfig" {
	to_emails = "admin@avicontroller.net"
	cc_emails = "admin@avicontroller.net"
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-aec-abc"
	description = "test alert email"

}

data "avi_alertemailconfig" "testAlertEmailConfig" {
    name= avi_alertemailconfig.testAlertEmailConfig.name
}
`
