package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVIAlertEmailConfigBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVIAlertEmailConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIAlertEmailConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIAlertEmailConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertEmailConfigExists("avi_alertemailconfig.testalertemailconfig"),
					resource.TestCheckResourceAttr(
						"avi_alertemailconfig.testalertemailconfig", "name", "aec-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIAlertEmailConfigExists("avi_alertemailconfig.testalertemailconfig"),
					resource.TestCheckResourceAttr(
						"avi_alertemailconfig.testalertemailconfig", "name", "aec-abc")),
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
			return fmt.Errorf("No Alert Email Config ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
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
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI Alert Email Config still exists")
		}
	}
	return nil
}

const testAccAVIAlertEmailConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_alertemailconfig" "testalertemailconfig" {
	name = "aec-%s"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
	cc_emails= "admin@avicontroller.net"
	description= "test alert email"
	to_emails= "admin@avicontroller.net"
}
`
