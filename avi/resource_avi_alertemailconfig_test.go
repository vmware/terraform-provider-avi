package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
)

func TestAVIAlertEmailConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAlertEmailConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAlertEmailConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertEmailConfigExists("avi_alertemailconfig.testAlertEmailConfig"),
					resource.TestCheckResourceAttr(
						"avi_alertemailconfig.testAlertEmailConfig", "name", "test-aec-abc"),
				),
			},
			{
				Config: testAccAVIAlertEmailConfigupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertEmailConfigExists("avi_alertemailconfig.testAlertEmailConfig"),
					resource.TestCheckResourceAttr(
						"avi_alertemailconfig.testAlertEmailConfig", "name", "test-aec-updated"),
				),
			},
			{
				ResourceName:      "avi_alertemailconfig.testAlertEmailConfig",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIAlertEmailConfigConfig,
			},
		},
	})

}

func testAccCheckAVIAlertEmailConfigExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI AlertEmailConfig ID is set")
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

func testAccCheckAVIAlertEmailConfigDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_alertemailconfig" {
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
			return fmt.Errorf("AVI AlertEmailConfig still exists")
		}
	}
	return nil
}

const testAccAVIAlertEmailConfigConfig = `
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
`

const testAccAVIAlertEmailConfigupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_alertemailconfig" "testAlertEmailConfig" {
	to_emails = "admin@avicontroller.net"
	cc_emails = "admin@avicontroller.net"
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-aec-updated"
	description = "test alert email"
}
`
