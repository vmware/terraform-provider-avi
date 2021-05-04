package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceActionGroupConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSActionGroupConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "name", "test-System-Alert-Level-High-abc"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "autoscale_trigger_notification", "false"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "external_only", "false"),
				),
			},
		},
	})

}

const testAccAVIDSActionGroupConfigConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_actiongroupconfig" "testActionGroupConfig" {
	autoscale_trigger_notification = false
	external_only = false
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-System-Alert-Level-High-abc"
	level = "ALERT_HIGH"
}

data "avi_actiongroupconfig" "testActionGroupConfig" {
    name= "${avi_actiongroupconfig.testActionGroupConfig.name}"
}
`
