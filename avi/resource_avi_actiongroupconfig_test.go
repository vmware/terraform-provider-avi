package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
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
						"avi_actiongroupconfig.testActionGroupConfig", "name", "test-System-Alert-Level-High-abc"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "autoscale_trigger_notification", "false"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "external_only", "false"),
				),
			},
			{
				Config: testAccAVIActionGroupConfigupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIActionGroupConfigExists("avi_actiongroupconfig.testActionGroupConfig"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "name", "test-System-Alert-Level-High-updated"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "autoscale_trigger_notification", "true"),
					resource.TestCheckResourceAttr(
						"avi_actiongroupconfig.testActionGroupConfig", "external_only", "false"),
				),
			},
			{
				ResourceName:      "avi_actiongroupconfig.testActionGroupConfig",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIActionGroupConfigConfig,
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
	name = "test-System-Alert-Level-High-abc"
	tenant_ref = data.avi_tenant.default_tenant.id
	level = "ALERT_HIGH"
	autoscale_trigger_notification = false
	external_only = false
}
`

const testAccAVIActionGroupConfigupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_actiongroupconfig" "testActionGroupConfig" {
	name = "test-System-Alert-Level-High-updated"
	tenant_ref = data.avi_tenant.default_tenant.id
	level = "ALERT_HIGH"
	autoscale_trigger_notification = true
	external_only = false
}
`
