package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAVISNMPTrapProfileBasic(t *testing.T) {
	updatedConfig := fmt.Sprintf(testAccAVISNMPTrapProfileConfig, "abc")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVISNMPTrapProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVISNMPTrapProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISNMPTrapProfileExists("avi_snmptrapprofile.testsnmptrapprofile"),
					resource.TestCheckResourceAttr(
						"avi_snmptrapprofile.testsnmptrapprofile", "name", "snmp-%s")),
			},
			{
				Config: updatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVISNMPTrapProfileExists("avi_snmptrapprofile.testsnmptrapprofile"),
					resource.TestCheckResourceAttr(
						"avi_snmptrapprofile.testsnmptrapprofile", "name", "snmp-abc")),
			},
		},
	})

}

func testAccCheckAVISNMPTrapProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No SNMP Trap Profile ID is set")
		}
		path := "api" + strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVISNMPTrapProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_snmptrapprofile" {
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
			return fmt.Errorf("AVI SNMP Trap Profile still exists")
		}
	}
	return nil
}

const testAccAVISNMPTrapProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}

resource "avi_snmptrapprofile" "testsnmptrapprofile" {
	name = "snmp-%s"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
