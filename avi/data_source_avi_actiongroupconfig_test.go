package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
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
	name = "test-System-Alert-Level-High-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	level = "ALERT_HIGH"
	autoscale_trigger_notification = false
	external_only = false
}

data "avi_actiongroupconfig" "testActionGroupConfig" {
    name= "${avi_actiongroupconfig.testActionGroupConfig.name}"
}
`
