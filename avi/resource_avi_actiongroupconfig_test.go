package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIActionGroupConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIActionGroupConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIActionGroupConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIActionGroupConfigExists("avi_actiongroupconfig.testActionGroupConfig"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "name", "testSystem-Alert-Level-High")),
			},
			{
				Config: testAccAVIActionGroupConfigupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIActionGroupConfigExists("avi_actiongroupconfig.testActionGroupConfig"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "name", "testSystem-Alert-Level-High-abc")),
			},
		},
	})

}

func testAccCheckAVIActionGroupConfigExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ActionGroupConfig ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIActionGroupConfigDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_actiongroupconfig" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI ActionGroupConfig still exists")
		}
	}
	return nil
}

const testAccAVIActionGroupConfigConfig = `
data "avi_tenant" "default_tenant"{
        name= "admin"
}
resource "avi_actiongroupconfig" "testActionGroupConfig" {
"level" = "ALERT_HIGH"
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"autoscale_trigger_notification" = false
"external_only" = false
"name" = "testSystem-Alert-Level-High"
}
`

const testAccAVIActionGroupConfigupdatedConfig = `
data "avi_tenant" "default_tenant"{
        name= "admin"
}
resource "avi_actiongroupconfig" "testActionGroupConfig" {
"level" = "ALERT_HIGH"
"tenant_ref" = "${data.avi_tenant.default_tenant.id}"
"autoscale_trigger_notification" = false
"external_only" = false
"name" = "testSystem-Alert-Level-High-abc"
}
`
